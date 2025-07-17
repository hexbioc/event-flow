variable "region" {
  description = "AWS region"
  type        = string
  default     = "us-east-2"
}

variable "project" {
  description = "Name of the project"
  type        = string
}

// NOTE: Length of `azs`, `public_subnets` and `private_subnets`
// must be equal, and a minimum of 2 for EKS
variable "azs" {
  description = "List of availability zones for the project"
  type        = list(string)
  default     = ["us-east-2a", "us-east-2b"]
}

variable "cidr" {
  description = "CIDR for the VPC"
  type        = string
  default     = "10.10.0.0/16"
}

variable "public_subnets" {
  description = "List of public subnet CIDRs"
  type        = list(string)
  default     = ["10.10.1.0/24", "10.10.3.0/24"]
}

variable "private_subnets" {
  description = "List of private subnet CIDRs"
  type        = list(string)
  default     = ["10.10.2.0/24", "10.10.4.0/24"]
}

variable "enable_bastion_host" {
  description = "Flag to enable bastion host"
  type        = bool
  default     = false
}

variable "bastion_ami" {
  description = "AMI for the bastion host"
  type        = string
}

variable "bastion_ingress_cidrs" {
  description = "List of CIDRs for bastion ingress"
  type        = list(string)
}

variable "bastion_ssh_key" {
  description = "Name of the SSH key to access bastion"
  type        = string
}

variable "eks_min_nodes" {
  description = "Minimum number of nodes in EKS node pool"
  type        = number
  default     = 1
}

variable "eks_max_nodes" {
  description = "Maximum number of nodes in EKS node pool"
  type        = number
  default     = 1
}

variable "ecr_services" {
  description = "List of services that require a repository"
  type        = set(string)
}

variable "github_repository_filter" {
  description = "Github repository filter for OIDC with Github Actions"
  type        = string
}

variable "rmq_username" {
  description = "Username for RabbitMQ"
  type        = string
}

variable "rmq_password" {
  description = "Password for RabbitMQ"
  type        = string
}

variable "common_tags" {
  description = "Map of common tags for all created resources"
  type        = map(string)
  default     = {}
}
