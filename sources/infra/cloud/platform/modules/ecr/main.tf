module "ecr" {
  source = "terraform-aws-modules/ecr/aws"

  for_each = var.services

  repository_name                   = "${var.base_repository}/${each.key}"
  repository_force_delete           = true // NOTE: Not for production!
  repository_read_write_access_arns = [aws_iam_role.github.arn]
  repository_image_tag_mutability   = "MUTABLE"
  repository_lifecycle_policy = jsonencode({
    rules = [
      {
        rulePriority = 1,
        description  = "Keep last 3 images",
        selection = {
          tagStatus   = "any",
          countType   = "imageCountMoreThan",
          countNumber = 3
        },
        action = {
          type = "expire"
        }
      }
    ]
  })

  tags = var.tags
}

resource "aws_iam_openid_connect_provider" "github" {
  url = "https://token.actions.githubusercontent.com"

  client_id_list  = ["sts.amazonaws.com"]
  thumbprint_list = []

  tags = var.tags
}

data "aws_iam_policy_document" "github_oidc" {
  statement {
    actions = ["sts:AssumeRoleWithWebIdentity"]

    principals {
      type        = "Federated"
      identifiers = [aws_iam_openid_connect_provider.github.arn]
    }

    condition {
      test     = "StringEquals"
      values   = ["sts.amazonaws.com"]
      variable = "token.actions.githubusercontent.com:aud"
    }

    condition {
      test     = "StringLike"
      values   = ["repo:${var.github_repository_filter}"]
      variable = "token.actions.githubusercontent.com:sub"
    }
  }
}

resource "aws_iam_role" "github" {
  name               = "github_oidc_role"
  assume_role_policy = data.aws_iam_policy_document.github_oidc.json

  tags = var.tags
}

data "aws_iam_policy_document" "ci" {
  statement {
    effect = "Allow"
    actions = [
      "ecr:*",
    ]
    resources = ["*"]
  }
}

resource "aws_iam_policy" "ci" {
  name        = "${var.base_repository}-ci-policy"
  description = "Policy used for deployments on CI"
  policy      = data.aws_iam_policy_document.ci.json

  tags = var.tags
}

resource "aws_iam_role_policy_attachment" "attach_ci" {
  role       = aws_iam_role.github.name
  policy_arn = aws_iam_policy.ci.arn
}
