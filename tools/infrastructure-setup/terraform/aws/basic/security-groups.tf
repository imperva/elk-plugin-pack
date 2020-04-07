
resource "aws_security_group" "ssh-sg" {

  name = "elk-basic-ssh-sg"
  description = "Default security group that allows inbound and outbound ssh traffic"

  ingress {
    from_port = 22
    to_port   = 22
    protocol  = "tcp"
    cidr_blocks = ["${data.external.control_cidr.result["ip"]}/32"]
  }

  egress {
    from_port   = "0"
    to_port     = "0"
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
    self        = true
  }
}

resource "aws_security_group" "vpc-sg" {
  name = "elk-basic-vpc-sg"
  description = "Default VPC security group"

  ingress {
    protocol  = -1
    self      = true
    from_port = 0
    to_port   = 0
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group" "kibana-sg" {
  name = "elk-basic-kibana-sg"
  description = "Kibana HTTP security group"

  ingress {
    from_port = 5601
    to_port   = 5601
    protocol  = "tcp"
    cidr_blocks = ["${data.external.control_cidr.result["ip"]}/32"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group" "logstash-sg" {
  name = "elk-basic-logstash-sg"
  description = "Logstash security group"

  ingress {
    from_port = 5514
    to_port   = 5514
    protocol  = "tcp"
    cidr_blocks = ["${data.external.control_cidr.result["ip"]}/32"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}
