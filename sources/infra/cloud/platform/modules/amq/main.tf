resource "aws_mq_broker" "rabbitmq" {
  broker_name         = var.name
  engine_type         = "RabbitMQ"
  engine_version      = "3.13"
  host_instance_type  = "mq.t3.micro"
  deployment_mode     = "SINGLE_INSTANCE"
  subnet_ids          = var.subnets
  security_groups     = [aws_security_group.allow_amqp.id]
  publicly_accessible = false

  user {
    username = var.username
    password = var.password
  }

  auto_minor_version_upgrade = true
  apply_immediately          = true
  maintenance_window_start_time {
    day_of_week = "MONDAY"
    time_of_day = "18:00"
    time_zone   = "UTC"
  }

  tags = var.tags
}


data "aws_vpc" "provided" {
  id = var.vpc
}

resource "aws_security_group" "allow_amqp" {
  name        = "allow_amqp"
  description = "Allow AMQP and console inbound traffic"
  vpc_id      = var.vpc

  tags = merge(var.tags, { Name = "${var.name}-allow-amqp" })
}

resource "aws_vpc_security_group_ingress_rule" "allow_rmq_ports" {
  for_each = toset(["443", "5671", "15671"])

  security_group_id = aws_security_group.allow_amqp.id
  cidr_ipv4         = data.aws_vpc.provided.cidr_block
  from_port         = tonumber(each.value)
  ip_protocol       = "tcp"
  to_port           = tonumber(each.value)
}
