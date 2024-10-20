package internal

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
)

func ConnectRabbitMQ(uri string) (*amqp.Connection, error) {
	conn, err := amqp.Dial(uri)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}
	return conn, nil
}

func PublishToRabbitMQ(conn *amqp.Connection, logData interface{}) error {
	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open RabbitMQ channel: %w", err)
	}
	defer ch.Close()

	body, err := json.Marshal(logData)
	if err != nil {
		return fmt.Errorf("failed to marshal log message: %w", err)
	}

	q, err := ch.QueueDeclare("logs", true, false, false, false, nil)
	if err != nil {
		return fmt.Errorf("failed to declare RabbitMQ queue: %w", err)
	}

	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	})
	if err != nil {
		return fmt.Errorf("failed to publish log message: %w", err)
	}
	return nil
}
