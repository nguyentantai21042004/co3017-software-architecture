package publisher

import (
	"context"
	"encoding/json"
	"time"

	"scoring/pkg/log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type EventPublisher interface {
	PublishSubmissionEvent(event interface{}) error
	Close() error
}

type rabbitmqPublisher struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	l       log.Logger
}

func NewRabbitMQPublisher(url string, logger log.Logger) (EventPublisher, error) {
	ctx := context.Background()

	// Connect to RabbitMQ
	conn, err := amqp.Dial(url)
	if err != nil {
		logger.Errorf(ctx, "publisher.NewRabbitMQPublisher: failed to connect to RabbitMQ | error=%v", err)
		return nil, err
	}

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		logger.Errorf(ctx, "publisher.NewRabbitMQPublisher: failed to open channel | error=%v", err)
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
		logger.Errorf(ctx, "publisher.NewRabbitMQPublisher: failed to declare exchange | error=%v", err)
		return nil, err
	}

	logger.Infof(ctx, "publisher.NewRabbitMQPublisher: RabbitMQ Publisher connected and exchange declared")

	return &rabbitmqPublisher{
		conn:    conn,
		channel: ch,
		l:       logger,
	}, nil
}

func (p *rabbitmqPublisher) PublishSubmissionEvent(event interface{}) error {
	ctx := context.Background()

	// Marshal event to JSON
	body, err := json.Marshal(event)
	if err != nil {
		p.l.Errorf(ctx, "publisher.rabbitmqPublisher.PublishSubmissionEvent: failed to marshal event | error=%v", err)
		return err
	}

	// Publish with context timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = p.channel.PublishWithContext(
		ctx,
		"its.events",       // exchange
		"event.submission", // routing key
		false,              // mandatory
		false,              // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
			Timestamp:   time.Now(),
		},
	)

	if err != nil {
		p.l.Errorf(ctx, "publisher.rabbitmqPublisher.PublishSubmissionEvent: failed to publish event | error=%v", err)
		return err
	}

	p.l.Infof(ctx, "publisher.rabbitmqPublisher.PublishSubmissionEvent: Published event | event=%s", string(body))
	return nil
}

func (p *rabbitmqPublisher) Close() error {
	if p.channel != nil {
		p.channel.Close()
	}
	if p.conn != nil {
		p.conn.Close()
	}
	return nil
}
