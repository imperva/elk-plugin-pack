# Infrastructure tooling

## Terraform
The Terraform scripts are set up to create x number of instances in a selected region (defined in variables.tf). Instance details are configured in instances.tf and security groups in security_groups.tf. Two security groups are created to allow ssh access and inter VPC communication.

### Prerequsites 

- Have a key pair created on AWS. Change the 'key' variable in variables.tf to the name of your key.
- Have the [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-install.html) installed and [configured](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html) with your security keys.
- Have Terraform installed.
- Change AMI in instances.tf to **var.amazon_linux_2** if you plan to use our Ansible plays to create a swarm or **var.ubuntu_server_18_04** if you plan to create a Kubernetes cluster

### Creating infrastructure
Run the following commands from the './terraform/aws' directory

Initalise terraform and install providers

    terraform init

Plan and apply your desired setup

    terraform plan
    terraform apply

### Destroying infrastructure
To permanently destroy all the infrastructure you have created run:

    terraform destroy

## Ansible
To quickly create a Docker swarm or Kubernetes Cluster, use the Ansible plays we have created. 

### Docker
This play will set up your cluster with swarm managers and workers. **Nothing will be run on the manager nodes**.

Add in IP addresses to the different blocks in the hosts file './ansible/inventory/production/swarm_hosts'.

| Group                          | Explaination                                                                                                                                     |
| -------------------------------| ------------------------------------------------------------------------------------------------------------------------------------------------ |
| swarm_manager                  | Single host. Creates and administers the swarm                                                                                                   |
| additional_swarm_managers      | Additional swarm managers are listed here. We recommend 2 (2 here plus the swarm-manager)                                                        |
| swarm_workers                  | Add remaining hosts here. They will run the Elastic cluster                                                                                      |
| master_node                    | Nodes listed here will be reserved for Elasticsearch nodes configured as master nodes. Check the Elastic website for the latest recommendations. |
| data_nodes                     | Nodes listed here will be reserved for Elasticsearch nodes configured as data nodes                                                              |

#### Running
Once you have divvied up your machines, run the following command to create your swarm:

    ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook ./playbooks/setup_swarm.yml -i ./inventory/production/swarm_hosts --private-key=./{{ KEY_NAME }}.pem -u ec2-user --fork 10 -e serial_number=10

### Kubernetes
This play will set up your cluster with a master and workers.

Add in IP addresses to the different blocks in the hosts file './ansible/inventory/production/kube_hosts'.

| Group         | Explaination                                                                                                                                     |
| --------------| ------------------------------------------------------------------------------------------------------------------------------------------------ |
| all:vars      | Tell Ansible to use Python3                                                                                                    |
| masters       | Master nodes go here                                                                                                           |
| workers       | Add remaining hosts here                                                                                                       |
| kube-master   | Used to initialise the Kubernetes cluster                                                                                      |

#### Running
Once you have divvied up your machines, run the following command to create your Kubernetes cluster:

    ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook ./playbooks/setup_kube_cluster.yml -i ./inventory/production/kube_hosts --private-key=./{{ KEY_NAME }}.pem -u ubuntu --fork 10 -e serial_number=10
