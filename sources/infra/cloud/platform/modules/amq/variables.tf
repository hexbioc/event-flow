variable "name" {
  description = "Name for the broker instance"
  type        = string
}

variable "vpc" {
  description = "VPC ID in which RabbitMQ is to be deployed"
  type        = string
}

variable "subnets" {
  description = "Subnet IDs for RabbitMQ"
  type        = list(string)
}

variable "username" {
  description = "Username for RabbitMQ"
  type        = string
}

variable "password" {
  description = "Password for RabbitMQ"
  type        = string
}

variable "tags" {
  description = "Map of tags for all resources"
  type        = map(string)
  default     = {}
}
