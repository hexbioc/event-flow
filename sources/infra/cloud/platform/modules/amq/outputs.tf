output "instances" {
  description = "Instance information of RabbitMQ"
  value       = aws_mq_broker.rabbitmq.instances
}
