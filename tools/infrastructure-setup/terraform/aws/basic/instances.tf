provider "aws" {
  profile    = "default"
  region     = "${var.region}"
}

resource "aws_instance" "elk-stack" {
    count           = "${var.instance_count}"
    ami             = "${var.amazon_linux_2}"
    instance_type   = "m5a.2xlarge"
    security_groups = ["${aws_security_group.setup.name}", "${aws_security_group.vpc_internal.name}"]
    key_name        = "${var.key}"
    ebs_block_device {
        device_name = "/dev/xvda"
        volume_type = "gp2"
        volume_size = 20
    }
    ebs_block_device {
        device_name = "/dev/sdb"
        volume_type = "st1"
        volume_size = 503
    }
}

output "public_dns" {
    value = "${aws_instance.elk-stack.*.public_dns}"
}