variable "name" {
  description = "Name of VPC"
  type        = string
}

variable "cidr" {
  description = "CIDR for the VPC"
  type        = string
}

variable "azs" {
  description = "List of availability zones"
  type        = list(string)
}

variable "public_subnets" {
  description = "List of public subnet CIDRs"
  type        = list(string)
}

variable "private_subnets" {
  description = "List of private subnet CIDRs"
  type        = list(string)
}

variable "enable_bastion_host" {
  description = "Flag to enable bastion host"
  type        = bool
  default     = false
}

variable "bastion_ami" {
  description = "AMI for the bastion host"
  type        = string
  default     = ""

  validation {
    condition     = !var.enable_bastion_host || var.bastion_ami != ""
    error_message = "Requred if bastion host is enabled"
  }
}

variable "bastion_ingress_cidrs" {
  description = "List of CIDRs for bastion ingress"
  type        = list(string)
  default     = []

  validation {
    condition     = !var.enable_bastion_host || length(var.bastion_ingress_cidrs) > 0
    error_message = "Requred if bastion host is enabled"
  }
}

variable "bastion_ssh_key" {
  description = "Name of the SSH key to access bastion"
  type        = string
  default     = ""

  validation {
    condition     = !var.enable_bastion_host || var.bastion_ssh_key != ""
    error_message = "Requred if bastion host is enabled"
  }
}

variable "tags" {
  description = "Map of tags for all resources"
  type        = map(string)
  default     = {}
}
