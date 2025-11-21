package consumer

import (
	"encoding/json"
	"fmt"
	"log"

	"learner-model-service/internal/model"
	"learner-model-service/internal/service"

	amqp "github.com/rabbitmq/amqp091-go"
)

type EventConsumer interface {
	Start() error
	Close() error
}

type rabbitmqConsumer struct {
	conn            *amqp.Connection
	channel         *amqp.Channel
	learnerService  service.LearnerService
	queue           amqp.Queue
}

func NewRabbitMQConsumer(url string, learnerService service.LearnerService) (EventConsumer, error) {
	// Connect to RabbitMQ
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to open channel: %w", err)
	}

	// Declare exchange (make sure it exists)
	err = ch.ExchangeDeclare(
		"its.events", // exchange name
		"topic",      // exchange type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, fmt.Errorf("failed to declare exchange: %w", err)
	}

	// Declare queue
	queue, err := ch.QueueDeclare(
		"learner.updates", // queue name
		true,              // durable
		false,             // auto-delete
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, fmt.Errorf("failed to declare queue: %w", err)
	}

	// Bind queue to exchange
	err = ch.QueueBind(
		queue.Name,         // queue name
		"event.submission", // routing key
		"its.events",       // exchange
		false,              // no-wait
		nil,                // arguments
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, fmt.Errorf("failed to bind queue: %w", err)
	}

	log.Println("✅ RabbitMQ Consumer connected, queue bound to exchange")

	return &rabbitmqConsumer{
		conn:           conn,
		channel:        ch,
		learnerService: learnerService,
		queue:          queue,
	}, nil
}

func (c *rabbitmqConsumer) Start() error {
	// Start consuming messages
	msgs, err := c.channel.Consume(
		c.queue.Name, // queue
		"",           // consumer tag
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	if err != nil {
		return fmt.Errorf("failed to register consumer: %w", err)
	}

	log.Println("🎧 Listening for events on queue: learner.updates")

	// Process messages in a goroutine
	go func() {
		for msg := range msgs {
			c.handleMessage(msg)
		}
	}()

	return nil
}

func (c *rabbitmqConsumer) handleMessage(msg amqp.Delivery) {
	log.Printf("📥 Received message: %s", string(msg.Body))

	// Parse event
	var event model.SubmissionEvent
	err := json.Unmarshal(msg.Body, &event)
	if err != nil {
		log.Printf("❌ Failed to parse event: %v", err)
		return
	}

	// Process event
	err = c.learnerService.UpdateMasteryFromEvent(&event)
	if err != nil {
		log.Printf("❌ Failed to process event: %v", err)
		return
	}

	log.Printf("✅ Successfully processed event for user: %s, skill: %s",
		event.UserID, event.SkillTag)
}

func (c *rabbitmqConsumer) Close() error {
	if c.channel != nil {
		c.channel.Close()
	}
	if c.conn != nil {
		c.conn.Close()
	}
	return nil
}
