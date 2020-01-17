provider "aws" {
  profile = "default"
  region  = var.region
}

resource "aws_instance" "elastic-masters" {
  count           = var.master_node_count
  ami             = var.amazon_linux_2
  instance_type   = "t3a.large"
  security_groups = [aws_security_group.setup.name, aws_security_group.vpc_internal.name]
  key_name        = var.key
  ebs_block_device {
    device_name = "/dev/xvda"
    volume_type = "gp2"
    volume_size = 20
  }
  
  tags = {
    Name = "lsar-demo-cluster-master"
  }
}

resource "aws_instance" "elastic-hot-data-nodes" {
  count           = var.hot_data_node_count
  ami             = var.amazon_linux_2
  instance_type   = "m5a.2xlarge"
  security_groups = [aws_security_group.setup.name, aws_security_group.vpc_internal.name]
  key_name        = var.key
  ebs_block_device {
    device_name = "/dev/xvda"
    volume_type = "io1"
    iops        = 3000
    volume_size = 4000
  }

  tags = {
    Name = "lsar-demo-cluster-hot"
  }
}

resource "aws_instance" "elastic-warm-data-nodes" {
  count           = var.warm_data_node_count
  ami             = var.amazon_linux_2
  instance_type   = "t3a.large"
  security_groups = [aws_security_group.setup.name, aws_security_group.vpc_internal.name]
  key_name        = var.key
  ebs_block_device {
    device_name = "/dev/xvda"
    volume_type = "gp2"
    volume_size = 20
  }
  ebs_block_device {
    device_name = "/dev/sdb"
    volume_type = "st1"
    volume_size = 5000
  }

  tags = {
    Name = "lsar-demo-cluster-warm"
  }
}

resource "aws_instance" "elastic-cold-data-nodes" {
  count           = var.cold_data_node_count
  ami             = var.amazon_linux_2
  instance_type   = "t3a.large"
  security_groups = [aws_security_group.setup.name, aws_security_group.vpc_internal.name]
  key_name        = var.key
  ebs_block_device {
    device_name = "/dev/xvda"
    volume_type = "gp2"
    volume_size = 20
  }
  ebs_block_device {
    device_name = "/dev/sdb"
    volume_type = "sc1"
    volume_size = 7000
  }

  tags = {
    Name = "lsar-demo-cluster-cold"
  }
}

resource "aws_instance" "elastic-coordinating-nodes" {
  count           = var.coordinating_node_count
  ami             = var.amazon_linux_2
  instance_type   = "r5a.2xlarge"
  security_groups = [aws_security_group.setup.name, aws_security_group.vpc_internal.name]
  key_name        = var.key
  ebs_block_device {
    device_name = "/dev/xvda"
    volume_type = "gp2"
    volume_size = 20
  }
  
  tags = {
    Name = "lsar-demo-cluster-coordinating"
  }
}

resource "aws_instance" "elastic-logstash-nodes" {
  count           = var.logstash_node_count
  ami             = var.amazon_linux_2
  instance_type   = "c5.2xlarge"
  security_groups = [aws_security_group.setup.name, aws_security_group.vpc_internal.name]
  key_name        = var.key
  ebs_block_device {
    device_name = "/dev/xvda"
    volume_type = "io1"
    iops        = 3000
    volume_size = 100
  }

  tags = {
    Name = "lsar-demo-cluster-logstash"
  }
}


output "public_dns_master" {
  value = aws_instance.elastic-masters.*.public_dns
}

output "public_dns_hot" {
  value = aws_instance.elastic-hot-data-nodes.*.public_dns
}

output "public_dns_warm" {
  value = aws_instance.elastic-warm-data-nodes.*.public_dns
}

output "public_dns_cold" {
  value = aws_instance.elastic-cold-data-nodes.*.public_dns
}

output "public_dns_coordinating" {
  value = aws_instance.elastic-coordinating-nodes.*.public_dns
}

output "public_dns_logstash" {
  value = aws_instance.elastic-logstash-nodes.*.public_dns
}
