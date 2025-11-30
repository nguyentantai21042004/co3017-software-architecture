package consumer

import (
	"context"
	"encoding/json"

	"learner-model-service/internal/learner"
	"learner-model-service/internal/model"
	"learner-model-service/pkg/log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// AMQPChannel is an interface for amqp.Channel to enable testing
type AMQPChannel interface {
	ExchangeDeclare(name, kind string, durable, autoDelete, internal, noWait bool, args amqp.Table) error
	QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) (amqp.Queue, error)
	QueueBind(name, key, exchange string, noWait bool, args amqp.Table) error
	Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error)
	Close() error
}

// AMQPConnection is an interface for amqp.Connection to enable testing
type AMQPConnection interface {
	Channel() (AMQPChannel, error)
	Close() error
	IsClosed() bool
}

// AMQPChannelWrapper wraps *amqp.Channel to implement AMQPChannel interface
type AMQPChannelWrapper struct {
	ch *amqp.Channel
}

func (w *AMQPChannelWrapper) ExchangeDeclare(name, kind string, durable, autoDelete, internal, noWait bool, args amqp.Table) error {
	return w.ch.ExchangeDeclare(name, kind, durable, autoDelete, internal, noWait, args)
}

func (w *AMQPChannelWrapper) QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) (amqp.Queue, error) {
	return w.ch.QueueDeclare(name, durable, autoDelete, exclusive, noWait, args)
}

func (w *AMQPChannelWrapper) QueueBind(name, key, exchange string, noWait bool, args amqp.Table) error {
	return w.ch.QueueBind(name, key, exchange, noWait, args)
}

func (w *AMQPChannelWrapper) Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error) {
	return w.ch.Consume(queue, consumer, autoAck, exclusive, noLocal, noWait, args)
}

func (w *AMQPChannelWrapper) Close() error {
	return w.ch.Close()
}

// AMQPConnectionWrapper wraps *amqp.Connection to implement AMQPConnection interface
type AMQPConnectionWrapper struct {
	conn *amqp.Connection
}

func (w *AMQPConnectionWrapper) Channel() (AMQPChannel, error) {
	ch, err := w.conn.Channel()
	if err != nil {
		return nil, err
	}
	return &AMQPChannelWrapper{ch: ch}, nil
}

func (w *AMQPConnectionWrapper) Close() error {
	return w.conn.Close()
}

func (w *AMQPConnectionWrapper) IsClosed() bool {
	return w.conn.IsClosed()
}

type EventConsumer interface {
	Start() error
	Close() error
}

type rabbitmqConsumer struct {
	conn    AMQPConnection
	channel AMQPChannel
	uc      learner.UseCase
	queue   amqp.Queue
	l       log.Logger
}

// NewRabbitMQConsumer creates a new RabbitMQ consumer
func NewRabbitMQConsumer(url string, uc learner.UseCase, l log.Logger) (EventConsumer, error) {
	ctx := context.Background()

	// Connect to RabbitMQ
	rawConn, err := amqpDialer(url)
	if err != nil {
		l.Errorf(ctx, "consumer.NewRabbitMQConsumer: failed to connect | error=%v", err)
		return nil, err
	}

	// Wrap connection for interface compatibility
	conn := &AMQPConnectionWrapper{conn: rawConn}

	return newRabbitMQConsumerWithConnection(conn, uc, l)
}

// newRabbitMQConsumerWithConnection creates a consumer with an existing connection (for testing)
func newRabbitMQConsumerWithConnection(conn AMQPConnection, uc learner.UseCase, l log.Logger) (EventConsumer, error) {
	ctx := context.Background()

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

	l.Infof(ctx, "consumer.newRabbitMQConsumerWithConnection: RabbitMQ Consumer connected successfully")

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
	var err error
	// Check if channel is not nil by attempting to close it
	// Since amqp.Channel is an interface, we can't directly compare to nil
	if c.channel != nil {
		if closeErr := c.channel.Close(); closeErr != nil {
			err = closeErr
		}
	}
	if c.conn != nil {
		if closeErr := c.conn.Close(); closeErr != nil {
			err = closeErr
		}
	}
	return err
}
