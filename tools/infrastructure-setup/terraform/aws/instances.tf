provider "aws" {
  profile    = "default"
  region     = "${var.region}"
}

resource "aws_instance" "elk-stack" {
    count           = "${var.instance_count}"
    ami             = "ami-00c03f7f7f2ec15c3"
    instance_type   = "m5ad.2xlarge"
    security_groups = ["${aws_security_group.setup.name}", "${aws_security_group.vpc_internal.name}"]
    key_name        = "${var.key}"
    ebs_block_device {
        device_name = "/dev/xvda"
        volume_type = "io1"
        iops        = 10000
        volume_size = 500
    }
}

output "public_dns" {
    value = "${aws_instance.elk-stack.*.public_dns}"
}