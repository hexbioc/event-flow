provider "aws" {
  region = var.region
}

locals {
  common_tags = merge({ "project" = var.project }, var.common_tags)
}

module "vpc" {
  source = "./modules/vpc"

  name = "${var.project}-vpc"

  cidr            = var.cidr
  azs             = var.azs
  public_subnets  = var.public_subnets
  private_subnets = var.private_subnets

  enable_bastion_host   = var.enable_bastion_host
  bastion_ami           = var.bastion_ami
  bastion_ingress_cidrs = var.bastion_ingress_cidrs
  bastion_ssh_key       = var.bastion_ssh_key

  tags = local.common_tags
}

module "eks" {
  source = "./modules/eks"

  name = "${var.project}-eks"

  vpc_id     = module.vpc.id
  subnet_ids = module.vpc.private_subnets

  // Configuring node pool to be a maximum of 3 nodes
  ng_min_size     = var.eks_min_nodes
  ng_desired_size = var.eks_min_nodes
  ng_max_size     = var.eks_max_nodes

  tags = local.common_tags
}

module "ecr" {
  source = "./modules/ecr"

  base_repository          = var.project
  services                 = var.ecr_services
  github_repository_filter = var.github_repository_filter
}

module "amq" {
  source = "./modules/amq"

  name    = "${var.project}-mq"
  vpc     = module.vpc.id
  subnets = [module.vpc.private_subnets[0]]

  username = var.rmq_username
  password = var.rmq_password
}
