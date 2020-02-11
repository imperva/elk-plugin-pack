
variable "aws_region" {
  default = "us-east-1"
}

variable "worker_count" {
  default = 3
}

variable "instance_type" {
  default = "m5a.2xlarge" 
}

variable "keypair" {
  default = "elk-basic"
}

