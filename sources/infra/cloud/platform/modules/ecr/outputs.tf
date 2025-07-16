output "repositories" {
  description = "List of created repositories"
  value       = values(module.ecr).*.repository_arn
}

output "aws_role_for_github" {
  description = "AWS Role for Github Actions"
  value       = aws_iam_role.github.arn
}
