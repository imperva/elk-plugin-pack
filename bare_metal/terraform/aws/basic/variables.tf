
variable "aws_region" {
  default = "us-east-1"
}

variable "node_count" {
  default = 6
}

variable "instance_type" {
  default = "m5a.2xlarge" 
}

variable "keypair" {
  default = "elk-basic"
}

