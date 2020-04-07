
################################################################################
#
# Ubunti AMI

data "aws_ami" "ubuntu_server_18_04" {
    most_recent = true
    owners = ["${var.ubuntu_account_number}"]
    filter {
        name   = "name"
        values = ["ubuntu/images/hvm-ssd/ubuntu-bionic-18.04-amd64-server-*"]
    }
    filter {
        name   = "virtualization-type"
        values = ["hvm"]
    }
    filter {
        name   = "image-type"
        values = ["machine"]
    }
    filter {
        name   = "architecture"
        values = ["x86_64"]
    }
}

variable "ubuntu_account_number" {
  default = "099720109477"
}


################################################################################
#
# Amazon Linux AMI

data "aws_ami" "amazon_linux_2" {
    most_recent = true
    owners = ["${var.amazon_account_number}"]
    filter {
        name   = "name"
        values = ["amzn2-ami-hvm-2.0.*"]
    }
    filter {
        name   = "virtualization-type"
        values = ["hvm"]
    }
    filter {
        name   = "image-type"
        values = ["machine"]
    }
    filter {
        name   = "architecture"
        values = ["x86_64"]
    }
}

variable "amazon_account_number" {
  default = "137112412989"
}


