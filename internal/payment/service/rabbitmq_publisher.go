package service

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
	"gomicro/internal/payment/model"
)

type IRabbitMQPublisher interface {
	SendStockUpdateEvent(event *model.StockUpdateEvent) error
	Close()
}

type RabbitMQPublisher struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	topic   string
}

func NewRabbitMQPublisher(url, topic string) (*RabbitMQPublisher, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to open a channel: %v", err)
	}

	err = ch.ExchangeDeclare(
		topic,   // name
		"topic", // type
		true,    // durable
		false,   // auto-deleted
		false,   // internal
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, fmt.Errorf("failed to declare an exchange: %v", err)
	}

	return &RabbitMQPublisher{
		conn:    conn,
		channel: ch,
		topic:   topic,
	}, nil
}

func (p *RabbitMQPublisher) SendStockUpdateEvent(event *model.StockUpdateEvent) error {
	body, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %v", err)
	}

	err = p.channel.Publish(
		p.topic,           // exchange
		"stock.update",    // routing key
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message: %v", err)
	}

	log.Printf("Published stock update event: %+v", event)
	return nil
}

func (p *RabbitMQPublisher) Close() {
	if p.channel != nil {
		p.channel.Close()
	}
	if p.conn != nil {
		p.conn.Close()
	}
} 