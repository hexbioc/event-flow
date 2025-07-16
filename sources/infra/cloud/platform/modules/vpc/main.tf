module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "5.8.1"

  name = var.name

  cidr            = var.cidr
  azs             = var.azs
  private_subnets = var.private_subnets
  public_subnets  = var.public_subnets

  enable_nat_gateway   = true
  single_nat_gateway   = true
  enable_dns_hostnames = true

  public_subnet_tags = {
    "kubernetes.io/role/elb" = 1
  }
  private_subnet_tags = {
    "kubernetes.io/role/internal-elb" = 1
  }

  tags = var.tags
}

resource "aws_security_group" "bastion" {
  count = var.enable_bastion_host ? 1 : 0

  name        = "${var.name}-bastion-sg"
  description = "Security group for bastion host"
  vpc_id      = module.vpc.vpc_id

  ingress {
    description = "Allow SSH from specified CIDRs"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = var.bastion_ingress_cidrs
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = var.tags
}

resource "aws_instance" "bastion" {
  count = var.enable_bastion_host ? 1 : 0

  ami                         = var.bastion_ami
  instance_type               = "t3.micro"
  key_name                    = var.bastion_ssh_key
  subnet_id                   = module.vpc.public_subnets[0]
  vpc_security_group_ids      = [aws_security_group.bastion[0].id]
  associate_public_ip_address = true

  tags = merge(var.tags, { name = "${var.name}-bastion" })

  lifecycle {
    create_before_destroy = true
  }
}
