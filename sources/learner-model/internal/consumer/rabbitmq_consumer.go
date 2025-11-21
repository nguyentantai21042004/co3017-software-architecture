package consumer

import (
	"context"
	"encoding/json"

	"learner-model-service/internal/learner"
	"learner-model-service/internal/model"
	"learner-model-service/pkg/log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type EventConsumer interface {
	Start() error
	Close() error
}

type rabbitmqConsumer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	uc      learner.UseCase
	queue   amqp.Queue
	l       log.Logger
}

func NewRabbitMQConsumer(url string, uc learner.UseCase, l log.Logger) (EventConsumer, error) {
	ctx := context.Background()

	// Connect to RabbitMQ
	conn, err := amqp.Dial(url)
	if err != nil {
		l.Errorf(ctx, "consumer.NewRabbitMQConsumer: failed to connect | error=%v", err)
		return nil, err
	}

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		l.Errorf(ctx, "consumer.NewRabbitMQConsumer: failed to open channel | error=%v", err)
		return nil, err
	}

	// Declare exchange
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
		l.Errorf(ctx, "consumer.NewRabbitMQConsumer: failed to declare exchange | error=%v", err)
		return nil, err
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
		l.Errorf(ctx, "consumer.NewRabbitMQConsumer: failed to declare queue | error=%v", err)
		return nil, err
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
		l.Errorf(ctx, "consumer.NewRabbitMQConsumer: failed to bind queue | error=%v", err)
		return nil, err
	}

	l.Infof(ctx, "consumer.NewRabbitMQConsumer: RabbitMQ Consumer connected successfully")

	return &rabbitmqConsumer{
		conn:    conn,
		channel: ch,
		uc:      uc,
		queue:   queue,
		l:       l,
	}, nil
}

func (c *rabbitmqConsumer) Start() error {
	ctx := context.Background()

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
		c.l.Errorf(ctx, "consumer.Start: failed to register consumer | error=%v", err)
		return err
	}

	c.l.Infof(ctx, "consumer.Start: Listening for events on queue: learner.updates")

	// Process messages in a goroutine
	go func() {
		for msg := range msgs {
			c.handleMessage(msg)
		}
	}()

	return nil
}

func (c *rabbitmqConsumer) handleMessage(msg amqp.Delivery) {
	ctx := context.Background()
	c.l.Infof(ctx, "consumer.handleMessage: received message | body=%s", string(msg.Body))

	// Parse event
	var event model.SubmissionEvent
	err := json.Unmarshal(msg.Body, &event)
	if err != nil {
		c.l.Errorf(ctx, "consumer.handleMessage: failed to parse event | error=%v", err)
		return
	}

	// Process event through use case
	input := learner.UpdateMasteryInput{
		UserID:        event.UserID,
		SkillTag:      event.SkillTag,
		ScoreObtained: event.ScoreObtained,
	}

	err = c.uc.UpdateMasteryFromEvent(ctx, input)
	if err != nil {
		c.l.Errorf(ctx, "consumer.handleMessage: failed to process event | user_id=%s | skill_tag=%s | error=%v",
			event.UserID, event.SkillTag, err)
		return
	}

	c.l.Infof(ctx, "consumer.handleMessage: successfully processed event | user_id=%s | skill_tag=%s",
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
