variable "base_repository" {
  description = "Base repository on ECR"
  type        = string
}

variable "services" {
  description = "List of services that require a repository"
  type        = set(string)
}

variable "github_repository_filter" {
  description = "Github repository filter for OIDC"
  type        = string
}

variable "tags" {
  description = "Map of tags for all resources"
  type        = map(string)
  default     = {}
}
