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
  default = "blue"
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

variable "subnet_public_cidr_block" {
  type        = string
  description = "CIDR Block for Public Subnets in VPC"
  default     = "10.0.0.0/24"
}

variable "subnet_private_cidr_block" {
  type        = string
  description = "CIDR Block for Pirvate Subnets in VPC"
  default     = "10.0.8.0/24"
}
