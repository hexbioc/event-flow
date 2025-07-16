package messaging

import (
	"collector/config"
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

const eventsQueue = "eventq"

type Handler struct {
	rmqUri     string
	connection *amqp.Connection
	channel    *amqp.Channel
}

func New(cfg *config.Config) *Handler {
	proto := "amqp"
	if cfg.RMQTLS == "true" {
		proto = "amqps"
	}

	return &Handler{
		rmqUri: fmt.Sprintf(
			"%s://%s:%s@%s/%s",
			proto,
			cfg.RMQUser,
			cfg.RMQPassword,
			cfg.RMQHostname,
			cfg.RMQVhost,
		),
	}
}

func (h *Handler) Connect() (*amqp.Connection, *amqp.Channel) {
	// Establish connection to RMQ
	conn, err := amqp.Dial(h.rmqUri)
	if err != nil {
		panic("Unable to establish connection to RabbitMQ")
	}
	h.connection = conn

	// Create channel for communication
	channel, err := conn.Channel()
	if err != nil {
		panic("Unable to create channel on RabbitMQ")
	}
	h.channel = channel

	// Setup queue
	_, err = h.channel.QueueDeclare(
		eventsQueue,
		true,  // durable
		false, // auto-delete
		false, // exclusive
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		panic("Unable to create queue on RabbitMQ")
	}

	return h.connection, h.channel
}

func (h *Handler) PublishJSON(payload any) error {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	err = h.channel.Publish(
		"",
		eventsQueue,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         payloadBytes,
			DeliveryMode: amqp.Persistent,
		},
	)

	return err
}
