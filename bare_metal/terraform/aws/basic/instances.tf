
terraform {
  required_version = ">= 0.12, < 0.13"
}

provider "aws" {
  profile = "default"
  region  = var.aws_region
  version = "~> 2.44"
}

data "external" "control_cidr" {
  program = ["bash", "-c", "curl -s 'https://api.ipify.org?format=json'"]
}

resource "aws_instance" "elk-master" {
    instance_type   = var.instance_type 
    ami             = data.aws_ami.amazon_linux_2.id
    security_groups = [ "${aws_security_group.ssh-sg.name}",
                        "${aws_security_group.vpc-sg.name}",
                        "${aws_security_group.kibana-sg.name}" ]
    key_name        = var.keypair
    ebs_block_device {
        device_name = "/dev/xvda"
        volume_type = "io1"
        iops        = 3000
        volume_size = 500
    }
}

resource "aws_instance" "elk-worker" {
    instance_type   = var.instance_type 
    count           = var.worker_count
    ami             = data.aws_ami.amazon_linux_2.id
    security_groups = [ "${aws_security_group.ssh-sg.name}",
                        "${aws_security_group.vpc-sg.name}",
                        "${aws_security_group.logstash-sg.name}" ]
    key_name        = var.keypair
    ebs_block_device {
        device_name = "/dev/xvda"
        volume_type = "io1"
        iops        = 3000
        volume_size = 500
    }
}

output "master_public_dns" {
  value = aws_instance.elk-master.public_dns
}

output "worker_public_dns" {
  value = join( ",", aws_instance.elk-worker.*.public_dns )
}

# output "worker_public_ips" {
#   value = join( ",", aws_instance.elk-worker.*.public_ip )
# }

