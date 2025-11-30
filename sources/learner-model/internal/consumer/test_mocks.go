package consumer

import (
	"context"
	"time"

	"learner-model-service/internal/learner"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/mock"
)

// MockUseCase for testing the consumer
type MockUseCase struct {
	mock.Mock
}

func (m *MockUseCase) GetMastery(ctx context.Context, input learner.GetMasteryInput) (learner.MasteryOutput, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		return learner.MasteryOutput{}, args.Error(1)
	}
	return args.Get(0).(learner.MasteryOutput), args.Error(1)
}

func (m *MockUseCase) UpdateMasteryFromEvent(ctx context.Context, input learner.UpdateMasteryInput) error {
	args := m.Called(ctx, input)
	return args.Error(0)
}

// MockLogger for testing
type MockLogger struct {
	mock.Mock
}

func (m *MockLogger) Debug(ctx context.Context, arg ...any) { m.Called(ctx, arg) }
func (m *MockLogger) Debugf(ctx context.Context, format string, arg ...any) {
	m.Called(ctx, format, arg)
}
func (m *MockLogger) Info(ctx context.Context, arg ...any) { m.Called(ctx, arg) }
func (m *MockLogger) Infof(ctx context.Context, format string, arg ...any) {
	m.Called(ctx, format, arg)
}
func (m *MockLogger) Warn(ctx context.Context, arg ...any) { m.Called(ctx, arg) }
func (m *MockLogger) Warnf(ctx context.Context, format string, arg ...any) {
	m.Called(ctx, format, arg)
}
func (m *MockLogger) Error(ctx context.Context, arg ...any) { m.Called(ctx, arg) }
func (m *MockLogger) Errorf(ctx context.Context, format string, arg ...any) {
	m.Called(ctx, format, arg)
}
func (m *MockLogger) Fatal(ctx context.Context, arg ...any) { m.Called(ctx, arg) }
func (m *MockLogger) Fatalf(ctx context.Context, format string, arg ...any) {
	m.Called(ctx, format, arg)
}

func NewMockLogger() *MockLogger {
	mockLogger := new(MockLogger)
	mockLogger.On("Infof", mock.Anything, mock.Anything, mock.Anything).Return()
	mockLogger.On("Errorf", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return()
	return mockLogger
}

// Concrete type for amqp.Delivery to pass to the consumer
type mockAMQPDelivery struct {
	Body          []byte
	ConsumerTag   string
	DeliveryTag   uint64
	Redelivered   bool
	Exchange      string
	RoutingKey    string
	Acknowledger  amqp.Acknowledger
	Headers       amqp.Table
	ContentType   string
	ContentEncoding string
	Timestamp     time.Time
	Type          string
	// ... other fields as needed for the consumer
}

func (m *mockAMQPDelivery) Ack(multiple bool) error {
	// Not used in auto-ack mode, but good to have a mock
	return nil
}

func (m *mockAMQPDelivery) Nack(multiple bool, requeue bool) error {
	return nil
}

func (m *mockAMQPDelivery) Reject(requeue bool) error {
	return nil
}

// Mock of amqp.Channel for testing
type MockAMQPChannel struct {
	mock.Mock
	queue chan amqp.Delivery
}

func NewMockAMQPChannel() *MockAMQPChannel {
	return &MockAMQPChannel{
		queue: make(chan amqp.Delivery, 100), // Buffer messages
	}
}

// Implement methods of amqp.Channel interface explicitly
// Only include methods actually called by the consumer, or mock them if needed.
// For a comprehensive mock, all methods would need to be implemented.
func (m *MockAMQPChannel) Close() error {
	args := m.Called()
	return args.Error(0)
}
func (m *MockAMQPChannel) Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error) {
	arg := m.Called(queue, consumer, autoAck, exclusive, noLocal, noWait, args)
	if arg.Get(1) != nil {
		return nil, arg.Error(1)
	}
	return m.queue, nil
}
func (m *MockAMQPChannel) ExchangeDeclare(name, kind string, durable, autoDelete, internal, noWait bool, args amqp.Table) error {
	mockArgs := m.Called(name, kind, durable, autoDelete, internal, noWait, args)
	return mockArgs.Error(0)
}
func (m *MockAMQPChannel) QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) (amqp.Queue, error) {
	mockArgs := m.Called(name, durable, autoDelete, exclusive, noWait, args)
	if mockArgs.Get(1) != nil { // Check for error first
		return amqp.Queue{}, mockArgs.Error(1)
	}
	return mockArgs.Get(0).(amqp.Queue), nil
}
func (m *MockAMQPChannel) QueueBind(name, key, exchange string, noWait bool, args amqp.Table) error {
	mockArgs := m.Called(name, key, exchange, noWait, args)
	return mockArgs.Error(0)
}
func (m *MockAMQPChannel) Qos(prefetchCount, prefetchSize int, global bool) error {
	args := m.Called(prefetchCount, prefetchSize, global)
	return args.Error(0)
}
func (m *MockAMQPChannel) Flow(active bool) error {
	args := m.Called(active)
	return args.Error(0)
}
func (m *MockAMQPChannel) Confirm(noWait bool) chan amqp.Confirmation {
	mockArgs := m.Called(noWait)
	if mockArgs.Get(0) == nil {
		return nil
	}
	return mockArgs.Get(0).(chan amqp.Confirmation)
}
func (m *MockAMQPChannel) Nack(tag uint64, multiple, requeue bool) error {
	args := m.Called(tag, multiple, requeue)
	return args.Error(0)
}
func (m *MockAMQPChannel) Reject(tag uint64, requeue bool) error {
	args := m.Called(tag, requeue)
	return args.Error(0)
}
func (m *MockAMQPChannel) Get(queue string, autoAck bool) (amqp.Delivery, bool, error) {
	args := m.Called(queue, autoAck)
	return args.Get(0).(amqp.Delivery), args.Bool(1), args.Error(2)
}
func (m *MockAMQPChannel) Publish(exchange, routingKey string, mandatory, immediate bool, msg amqp.Publishing) error {
	args := m.Called(exchange, routingKey, mandatory, immediate, msg)
	return args.Error(0)
}
func (m *MockAMQPChannel) Tx(noWait bool) error {
	args := m.Called(noWait)
	return args.Error(0)
}
func (m *MockAMQPChannel) TxCommit() error {
	args := m.Called()
	return args.Error(0)
}
func (m *MockAMQPChannel) TxRollback() error {
	args := m.Called()
	return args.Error(0)
}
func (m *MockAMQPChannel) NotifyClose(c chan *amqp.Error) chan *amqp.Error {
	args := m.Called(c)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(chan *amqp.Error)
}
func (m *MockAMQPChannel) NotifyConfirm(ack, nack chan uint64) (chan uint64, chan uint64) {
	args := m.Called(ack, nack)
	if args.Get(0) == nil || args.Get(1) == nil {
		return nil, nil
	}
	return args.Get(0).(chan uint64), args.Get(1).(chan uint64)
}
func (m *MockAMQPChannel) NotifyReturn(c chan amqp.Return) chan amqp.Return {
	args := m.Called(c)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(chan amqp.Return)
}
func (m *MockAMQPChannel) NotifyFlow(c chan bool) chan bool {
	args := m.Called(c)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(chan bool)
}
func (m *MockAMQPChannel) NotifyCancel(c chan string) chan string {
	args := m.Called(c)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(chan string)
}
func (m *MockAMQPChannel) ExchangeDelete(name string, ifUnused, noWait bool) error {
	args := m.Called(name, ifUnused, noWait)
	return args.Error(0)
}
func (m *MockAMQPChannel) QueueDelete(name string, ifUnused, ifEmpty, noWait bool) (int, error) {
	args := m.Called(name, ifUnused, ifEmpty, noWait)
	return args.Int(0), args.Error(1)
}
func (m *MockAMQPChannel) QueuePurge(name string, noWait bool) (int, error) {
	args := m.Called(name, noWait)
	return args.Int(0), args.Error(1)
}
func (m *MockAMQPChannel) ConsumerCount(queue string) (int, error) {
	args := m.Called(queue)
	return args.Int(0), args.Error(1)
}
func (m *MockAMQPChannel) MessageCount(queue string) (int, error) {
	args := m.Called(queue)
	return args.Int(0), args.Error(1)
}
func (m *MockAMQPChannel) ExchangeBind(destination, routingKey, source string, noWait bool, args amqp.Table) error {
	mockArgs := m.Called(destination, routingKey, source, noWait, args)
	return mockArgs.Error(0)
}
func (m *MockAMQPChannel) ExchangeUnbind(destination, routingKey, source string, noWait bool, args amqp.Table) error {
	mockArgs := m.Called(destination, routingKey, source, noWait, args)
	return mockArgs.Error(0)
}
func (m *MockAMQPChannel) QueueUnbind(name, routingKey, exchange string, args amqp.Table) error {
	mockArgs := m.Called(name, routingKey, exchange, args)
	return mockArgs.Error(0)
}
func (m *MockAMQPChannel) BasicPublish(exchange, routingKey string, mandatory, immediate bool, msg amqp.Publishing) error {
	args := m.Called(exchange, routingKey, mandatory, immediate, msg)
	return args.Error(0)
}
func (m *MockAMQPChannel) ExchangeInspect(name string) error {
	args := m.Called(name)
	return args.Error(0)
}
func (m *MockAMQPChannel) ExchangeDeclarePassive(name, kind string, durable, autoDelete, internal, noWait bool, args amqp.Table) error {
	mockArgs := m.Called(name, kind, durable, autoDelete, internal, noWait, args)
	return mockArgs.Error(0)
}
func (m *MockAMQPChannel) QueueDeclarePassive(name string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) (amqp.Queue, error) {
	mockArgs := m.Called(name, durable, autoDelete, exclusive, noWait, args)
	return mockArgs.Get(0).(amqp.Queue), mockArgs.Error(1)
}
func (m *MockAMQPChannel) ExchangeDeleteNoWait(name string, ifUnused bool) error {
	args := m.Called(name, ifUnused)
	return args.Error(0)
}
func (m *MockAMQPChannel) ExchangeBindNoWait(destination, routingKey, source string, args amqp.Table) error {
	mockArgs := m.Called(destination, routingKey, source, args)
	return mockArgs.Error(0)
}
func (m *MockAMQPChannel) ExchangeUnbindNoWait(destination, routingKey, source string, args amqp.Table) error {
	mockArgs := m.Called(destination, routingKey, source, args)
	return mockArgs.Error(0)
}
func (m *MockAMQPChannel) QueueDeleteNoWait(name string, ifUnused, ifEmpty bool) error {
	args := m.Called(name, ifUnused, ifEmpty)
	return args.Error(0)
}
func (m *MockAMQPChannel) QueueUnbindNoWait(name, routingKey, exchange string, args amqp.Table) error {
	mockArgs := m.Called(name, routingKey, exchange, args)
	return mockArgs.Error(0)
}


// Mock of amqp.Connection for testing
type MockAMQPConnection struct {
	mock.Mock
	channel *MockAMQPChannel
}

func NewMockAMQPConnection(ch *MockAMQPChannel) *MockAMQPConnection {
	return &MockAMQPConnection{channel: ch}
}

// Implement methods of amqp.Connection interface
func (m *MockAMQPConnection) Close() error {
	args := m.Called()
	return args.Error(0)
}
func (m *MockAMQPConnection) Channel() (AMQPChannel, error) {
	args := m.Called()
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	if args.Get(0) == nil {
		return nil, nil
	}
	// Return the mock channel which implements AMQPChannel interface
	return args.Get(0).(AMQPChannel), nil
}
func (m *MockAMQPConnection) LocalAddr() string {
	args := m.Called()
	return args.String(0)
}
func (m *MockAMQPConnection) RemoteAddr() string {
	args := m.Called()
	return args.String(0)
}
func (m *MockAMQPConnection) ConnectionState() interface{} {
	mockArgs := m.Called()
	if mockArgs.Get(0) == nil {
		return nil
	}
	return mockArgs.Get(0)
}
func (m *MockAMQPConnection) NotifyClose(c chan *amqp.Error) chan *amqp.Error {
	args := m.Called(c)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(chan *amqp.Error)
}
func (m *MockAMQPConnection) NotifyBlocked(c chan amqp.Blocking) chan amqp.Blocking {
	args := m.Called(c)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(chan amqp.Blocking)
}
func (m *MockAMQPConnection) NotifyFlow(c chan bool) chan bool {
	args := m.Called(c)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(chan bool)
}
func (m *MockAMQPConnection) IsClosed() bool {
	args := m.Called()
	return args.Bool(0)
}
func (m *MockAMQPConnection) TLSConnectionState() (interface{}, bool) {
	mockArgs := m.Called()
	if mockArgs.Get(0) == nil {
		return nil, mockArgs.Bool(1)
	}
	return mockArgs.Get(0), mockArgs.Bool(1)
}


// amqpDialer is a function variable that can be overridden for testing
// It defaults to the real amqp.Dial function
var amqpDialer = func(url string) (*amqp.Connection, error) {
	return amqp.Dial(url)
}