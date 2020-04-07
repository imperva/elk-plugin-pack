

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

> Note Well : If you intend to use annotations for `pod` IAM roles, this will not work unless KIAM is deployed. You can deploy KIAM using its community developed Helm chart <https://github.com/helm/charts/tree/master/stable/kiam>. See appendix `Deploy KIAM - Kubernetes IAM` at the end of this document.

### Elasticsearch Helm Chart

See <https://hub.helm.sh/charts/stable/elastic-stack> 

~~~bash
$ helm install stable/elastic-stack --version 1.8.0

....

~~~

The elasticsearch cluster and associated extras have been installed.

~~~bash
$ kubectl get pods --namespace=default -l app=elasticsearch
NAME                                                 READY   STATUS    RESTARTS   AGE
tinseled-gnat-elasticsearch-client-bf8f4b848-77d7r   1/1     Running   0          15m
tinseled-gnat-elasticsearch-client-bf8f4b848-ggg47   1/1     Running   0          15m
tinseled-gnat-elasticsearch-data-0                   1/1     Running   0          15m
tinseled-gnat-elasticsearch-data-1                   1/1     Running   0          14m
tinseled-gnat-elasticsearch-master-0                 1/1     Running   0          15m
tinseled-gnat-elasticsearch-master-1                 1/1     Running   0          13m
tinseled-gnat-elasticsearch-master-2                 1/1     Running   0          13m

$ kubectl get pods --namespace=default -l app=kibana
NAME                                    READY   STATUS    RESTARTS   AGE
tinseled-gnat-kibana-85df678bf9-h6wx2   1/1     Running   0          15m

$ kubectl get pods --namespace=default -l app=logstash
NAME                       READY   STATUS    RESTARTS   AGE
tinseled-gnat-logstash-0   1/1     Running   0          17m
~~~

Kibana can be accessed:

- Within your cluster, at the DNS name `tinseled-gnat-elastic-stack.default.svc.cluster.local` at port `9200`
- From outside the cluster, run these commands in the same shell:

~~~bash
$ export POD_NAME=$(kubectl get pods --namespace default \
        -l "app=kibana" -o jsonpath="{.items[0].metadata.name}")

$ kubectl port-forward --namespace default $POD_NAME 5601:5601
~~~

To connect to Kibana, access <http://127.0.0.1:5601> via a browser.

### Expose the Kibana Service

~~~bash
$ kubectl get svc tinseled-gnat-kibana
NAME                   TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)   AGE
tinseled-gnat-kibana   ClusterIP   10.100.154.240   <none>        443/TCP   23m
~~~

Currently the Service does not have an External IP, so let’s now recreate the Service to use a cloud load balancer, just change the Type of `tinseled-gnat-kibana` Service from `ClusterIP` to `LoadBalancer`:

~~~bash
$ kubectl edit svc tinseled-gnat-kibana
~~~ 

~~~bash
$ kubectl get svc tinseled-gnat-kibana
NAME                   TYPE           CLUSTER-IP       EXTERNAL-IP                                                             PORT(S)         AGE
tinseled-gnat-kibana   LoadBalancer   10.100.154.240   afc2c47970d1011eaac7d0235f2b7372-19733194.us-east-2.elb.amazonaws.com   443:31629/TCP   25m
~~~

The `Kibana` service is now reachable externally from `afc2c47970d1011eaac7d0235f2b7372-19733194.us-east-2.elb.amazonaws.com:443`. 

Similarly for `logstash`

~~~bash
$ kubectl get  svc tinseled-gnat-logstash
NAME                     TYPE           CLUSTER-IP     EXTERNAL-IP                                                               PORT(S)          AGE
tinseled-gnat-logstash   LoadBalancer   10.100.44.47   afc2e77140d1011eaac7d0235f2b7372-1537181351.us-east-2.elb.amazonaws.com   5044:30220/TCP   29m
~~~

The `LogStash` service is now reachable externally from `afc2e77140d1011eaac7d0235f2b7372-1537181351.us-east-2.elb.amazonaws.com:5044 `. 

### Appendices

#### Deploy KIAM - Kubernetes IAM

If you intend to use annotations for `pod` IAM roles, this will not work unless KIAM is deployed. You can deploy KIAM using its community developed Helm chart <https://github.com/helm/charts/tree/master/stable/kiam>.

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

##### Using KIAM 

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
