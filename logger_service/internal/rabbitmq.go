package internal

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn *amqp.Connection
}

func NewRabbitMQ(uri string) *RabbitMQ {
	conn, err := amqp.Dial(uri)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	return &RabbitMQ{conn: conn}
}

func (r *RabbitMQ) ListenAndSaveLogs(mongoClient *MongoClient) error {
	ch, err := r.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"logs", true, false, false, false, nil,
	)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		q.Name, "", true, false, false, false, nil,
	)
	if err != nil {
		return err
	}

	for msg := range msgs {
		var logEntry map[string]interface{}
		if err := json.Unmarshal(msg.Body, &logEntry); err != nil {
			log.Printf("Failed to unmarshal log: %v", err)
			continue
		}
		mongoClient.SaveLog(logEntry)
	}
	return nil
}
