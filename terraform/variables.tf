variable "project" {
  type        = string
  description = "project name"
  default = "DU-Capstone"
}

variable "region" {
  type        = string
  description = "AWS region"
  default     = "us-east-1"
}

variable "lambda_src_path" {
  type = string
  description = "path to the source code for the lambda function"
  default = "blue/lambda-src"
}

variable "glue_src_path" {
  type = string
  description = "path to the source code for the glue job"
  default = "blue/glue-src"
}

variable "enable_dns_hostnames" {
  type        = bool
  description = "Enable DNS hostnames in VPC"
  default     = true
}

variable "vpc_cidr_block" {
  type        = string
  description = "Base CIDR Block for VPC"
  default     = "10.0.0.0/16"
}

variable "vpc_subnets_cidr_block" {
  type        = list(string)
  description = "CIDR Block for Subnets in VPC"
  default     = ["10.0.0.0/24", "10.0.1.0/24"]
}

variable "map_public_ip_on_launch" {
  type        = bool
  description = "Map a public IP address for Subnet instances"
  default     = true
}