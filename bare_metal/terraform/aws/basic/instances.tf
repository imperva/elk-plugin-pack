
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

resource "aws_instance" "node" {
    instance_type   = var.instance_type 
    count           = var.node_count
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

output "node_public_dns" {
  value = aws_instance.node.*.public_dns
}
