output "region" {
  description = "AWS region"
  value       = var.region
}

output "cluster_endpoint" {
  description = "Endpoint for EKS control plane"
  value       = module.eks.cluster_endpoint
}

output "cluster_name" {
  description = "Kubernetes cluster name"
  value       = module.eks.cluster_name
}

output "cluster_security_group_id" {
  description = "Security group IDs attached to the cluster control plane"
  value       = module.eks.cluster_security_group_id
}

output "ecr_repositories" {
  description = "List of created ECR repositories"
  value       = module.ecr.repositories
}

output "aws_role_for_github" {
  description = "AWS Role for Github Actions"
  value       = module.ecr.aws_role_for_github
}

output "bastion_public_ip" {
  description = "Public IP of the bastion host, if created"
  value       = var.enable_bastion_host ? module.vpc.bastion_public_ip : "Not enabled"
}
