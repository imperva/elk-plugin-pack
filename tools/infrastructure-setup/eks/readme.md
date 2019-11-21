

### Create an AWS EKS Cluster 

Aside: we will be doing the following as AWS user `k8s_cluster_admin`

~~~bash
$ aws iam get-user
{
    "User": {
        "Path": "/",
        "UserName": "k8s_cluster_admin",
        "UserId": "AIDAYCSDSTVY66LYESC7X",
        "Arn": "arn:aws:iam::555266252145:user/k8s_cluster_admin",
  ....
~~~

First create the cluster using `eksctl`

~~~bash
$ eksctl create cluster \
        --name lsar-eks-cluster \
        --version 1.14 \
        --region us-east-2 \
        --nodegroup-name lsar-cluster-workers \
        --node-type m5ad.2xlarge \
        --nodes     4 \
        --nodes-min 3 \
        --nodes-max 8 \
        --alb-ingress-access \
        --node-ami auto
~~~

This will create the following

- Subnets for each availability zone in the region
- Nodegroup `lsar-cluster-workers` using `ami-053250833d1030033` [AmazonLinux2/1.14]
- 4 EC2 `m5ad.2xlarge` instances 
- 2 separate CloudFormation stacks for the cluster itself and the initial nodegroup

It will also configure `kubectl` to use this cluster as the current context, so all `kubectl` command will be routed to it. This also includes tools, like `helm` and `kops`, that use `kubectl` and it's configuration.

~~~
$ kubectl config current-context
k8s_cluster_admin@lsar-eks-cluster.us-east-2.eksctl.io

$ kubectl cluster-info
Kubernetes master is running at https://5710B0F1A9497BDDEB1FA136A7531C87.yl4.us-east-2.eks.amazonaws.com
CoreDNS is running at https://5710B0F1A9497BDDEB1FA136A7531C87.yl4.us-east-2.eks.amazonaws.com/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy


~~~

You can check the cluster details by typing

~~~bash
$ eksctl utils describe-stacks \
        --region=us-east-2 \
        --cluster=lsar-eks-cluster
....
  StackId: "arn:aws:cloudformation:us-east-2:555266252145:stack/eksctl-lsar-eks-cluster-nodegroup-lsar-cluster-workers/2e668ff0-0c4d-11ea-a919-0af17cbfa87c",
  StackName: "eksctl-lsar-eks-cluster-nodegroup-lsar-cluster-workers",
  StackStatus: "CREATE_COMPLETE",
....
~~~

and enable CloudWatch logging as follows

~~~bash
$ eksctl utils update-cluster-logging \
        --region=us-east-2 \
        --cluster=lsar-eks-cluster \
        --enable-types=api --approve
~~~


### Configure Helm

We will be deploying a number of `K8S` resources using `helm`. `helm` (the command line tool) works by talking to a server-side component that runs in your cluster called `tiller`. `Tiller` by default runs as the `tiller-deploy` Deployment in `kube-system` namespace. Much of the power and flexibility of Helm comes from the fact that charts can contain just about any Kubernetes resource.

#### Configure Helm access with RBAC

Helm relies on a service called `tiller` that requires special permission on the kubernetes cluster, so we need to build a `Service Account` for tiller to use. Then apply this to the cluster.

Create a new service account manifest `tiller-rbac.yaml` with the following content.

~~~yaml
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: tiller
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: tiller
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
  - kind: ServiceAccount
    name: tiller
    namespace: kube-system
~~~

Apply the config

~~~bash
$ kubectl apply -f  tiller-rbac.yaml
serviceaccount/tiller created
clusterrolebinding.rbac.authorization.k8s.io/tiller created
~~~

Then we can install `tiller` using the `helm` tooling

~~~bash
$ helm init --service-account tiller

$ kubectl get pods -n kube-system | grep tiller 
tiller-deploy-7f4d76c4b6-n2mg6   1/1     Running   0          31m
~~~

This sets up the `helm` configuration at `$HOME/.helm.` and installs Tiller (the Helm server-side component) into your Kubernetes Cluster.

Please note: by default, Tiller is deployed with an insecure `allow unauthenticated users` policy.
To prevent this, run `helm init` with the `--tiller-tls-verify` flag.
For more information on securing your installation see: `https://docs.helm.sh/using_helm/#securing-your-helm-installation`

#### Deploy `ExternalDNS`

Deploy the `external-dns` service. `ExternalDNS` <https://github.com/kubernetes-sigs/external-dns> is a Kubernetes addon that configures public DNS servers with information about exposed Kubernetes services to make them discoverable. You can deploy it using the community-developed `Helm` chart <https://github.com/helm/charts/tree/master/stable/external-dns>.

~~~bash
$ helm install --name lsar-release  \
        --set provider=aws \
        --set aws.zoneType=public \
        --set aws.region=us-east-2 \
        stable/external-dns
        
....
==> v1/Pod(related)
NAME                                        READY  STATUS             RESTARTS  AGE
lsar-release-external-dns-76ccc995d8-r9bbl  0/1    ContainerCreating  0         1s

==> v1/Service
NAME                       TYPE       CLUSTER-IP      EXTERNAL-IP  PORT(S)   AGE
lsar-release-external-dns  ClusterIP  10.100.210.240  <none>       7979/TCP  1s
....
~~~

It will take some time to deploy the `helm` chart. To verify that `external-dns` has started, run:

~~~bash
$ kubectl --namespace=default get pods \
        -l "app.kubernetes.io/name=external-dns,app.kubernetes.io/instance=lsar-release"

NAME                                         READY   STATUS    RESTARTS   AGE
lsar-release-external-dns-76ccc995d8-r9bbl   1/1     Running   0          3m12s
~~~

To uninstall/delete the `lsar-release` deployment

~~~bash
$ helm delete lsar-release
~~~

#### Deploy KIAM - Kubernetes IAM

Annotations for pod IAM roles will not work if KIAM is not deployed. You can deploy KIAM using its community developed Helm chart <https://github.com/helm/charts/tree/master/stable/kiam>.

~~~bash
$ helm install stable/kiam --name lsar-kiam-release

....
==> v1/Service
NAME                      TYPE       CLUSTER-IP  EXTERNAL-IP  PORT(S)           AGE
lsar-kiam-release-agent   ClusterIP  None        <none>       9620/TCP          1s
lsar-kiam-release-server  ClusterIP  None        <none>       9620/TCP,443/TCP  1s
....
~~~

To verify that kiam has started, run:

~~~bash
$  kubectl --namespace=default get pods -l "app=kiam,release=lsar-kiam-release"
NAME                             READY   STATUS    RESTARTS   AGE
lsar-kiam-release-agent-97tzp    1/1     Running   2          2m29s
....
lsar-kiam-release-server-k9b6q   1/1     Running   0          2m29s
....
~~~

Using KIAM 

- Add an annotation to your namespace as below:

~~~yaml
kind: Namespace
metadata:
  name: iam-example
  annotations:
    iam.amazonaws.com/permitted: “<Role ARN or a Regex matching role ARN(s)>”
~~~

- Add an `iam.amazonaws.com/role` annotation to your pods with the role you want them to assume.
Use `curl` to verify the pod's role from within: `curl http://169.254.169.254/latest/meta-data/iam/security-credentials/`

### Elasticsearch Helm Chart

See <https://github.com/elastic/helm-charts/blob/master/elasticsearch/README.md>

#### Install ELK via Helm

##### Using Helm repository

Add the elastic helm charts repository

~~~bash
$ helm repo add elastic https://helm.elastic.co

$ helm install --name elasticsearch elastic/elasticsearch
...
==> v1/Service
NAME                           TYPE       CLUSTER-IP    EXTERNAL-IP  PORT(S)            AGE
elasticsearch-master           ClusterIP  10.100.54.90  <none>       9200/TCP,9300/TCP  1s
elasticsearch-master-headless  ClusterIP  None          <none>       9200/TCP,9300/TCP  1s
...
~~~

- Watch all cluster members come up.
 
~~~bash
$ kubectl get pods --namespace=default -l app=elasticsearch-master -w
NAME                     READY   STATUS    RESTARTS   AGE
elasticsearch-master-0   1/1     Running   0          2m46s
elasticsearch-master-1   1/1     Running   0          2m46s
elasticsearch-master-2   1/1     Running   0          2m46s
~~~

- Test cluster health using Helm test.

~~~bash
$ helm test elasticsearch
RUNNING: elasticsearch-ayulj-test
PASSED: elasticsearch-ayulj-test
~~~

##### Using master `git` branch

- Clone the git repo

~~~bash
$ git clone git@github.com:elastic/helm-charts.git
~~~

- Install it

~~~bash
$ helm install --name elasticsearch ./helm-charts/elasticsearch
~~~


