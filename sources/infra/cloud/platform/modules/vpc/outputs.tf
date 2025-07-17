output "id" {
  description = "VPC ID"
  value       = module.vpc.vpc_id
}

output "private_subnets" {
  description = "List of private subnets of the VPC"
  value       = module.vpc.private_subnets
}

output "public_subnets" {
  description = "List of public subnets of the VPC"
  value       = module.vpc.public_subnets
}

output "bastion_public_ip" {
  description = "Public IP of the bastion host, if created"
  value       = var.enable_bastion_host ? aws_instance.bastion[0].public_ip : "Not enabled"
}
