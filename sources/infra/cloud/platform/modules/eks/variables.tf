variable "name" {
  description = "Name of the EKS cluster"
  type        = string
}

variable "vpc_id" {
  description = "VPC ID for deployment"
  type        = string
}

variable "subnet_ids" {
  description = "List of subnet IDs"
  type        = list(string)
}

variable "ng_min_size" {
  description = "Minimum size of node group"
  type        = number
  default     = 1
}

variable "ng_max_size" {
  description = "Minimum size of node group"
  type        = number
  default     = 1
}

variable "ng_desired_size" {
  description = "Desired size of node group"
  type        = number
  default     = 1
}

variable "tags" {
  description = "Map of tags for all resources"
  type        = map(string)
  default     = {}
}
