variable "region" {
  default = "us-east-2"
}

variable "master_node_count" {
  default = 1
}

variable "hot_data_node_count" {
  default = 2
}

variable "warm_data_node_count" {
  default = 3
}

variable "cold_data_node_count" {
  default = 3
}

variable "key" {
  default = "audit-deploy"
}

variable "amazon_linux_2" {
  default = "ami-00c03f7f7f2ec15c3"
}

variable "ubuntu_server_18_04" {
  default = "ami-0d5d9d301c853a04a"
}

