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
