# Terraform
The Terraform scripts are set up to create x number of instances in a selected region (defined in variables.tf). Instance details are configured in instances.tf and security groups in security_groups.tf. Two security groups are created to allow ssh access and inter VPC communication.

## Prerequsites 

- Have a key pair created on AWS. Change the 'key' variable in variables.tf to the name of your key.
- Have the [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-install.html) installed and [configured](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html) with your security keys.
- Have Terraform installed.

## Creating infrastructure
Run the following commands from the './terraform/aws' directory

Initalise terraform and install providers

    terraform init

Plan and apply your desired setup

    terraform plan
    terraform apply

## Destroying infrastructure
To permanently destroy all the infrastructure you have created run:

    terraform destroy
