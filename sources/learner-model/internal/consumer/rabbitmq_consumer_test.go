package consumer

import (
	"context"
	"testing"

	"learner-model-service/internal/learner"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUseCase is a mock implementation of learner.UseCase
type MockUseCase struct {
	mock.Mock
}

func (m *MockUseCase) GetMastery(ctx context.Context, input learner.GetMasteryInput) (learner.MasteryOutput, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(learner.MasteryOutput), args.Error(1)
}

func (m *MockUseCase) UpdateMasteryFromEvent(ctx context.Context, input learner.UpdateMasteryInput) error {
	args := m.Called(ctx, input)
	return args.Error(0)
}

// TestNewRabbitMQConsumer_InvalidURL tests consumer creation with invalid URL
func TestNewRabbitMQConsumer_InvalidURL(t *testing.T) {
	mockUC := new(MockUseCase)
	mockLogger := &mockLogger{}

	// Try to create consumer with invalid URL
	consumer, err := NewRabbitMQConsumer("invalid://url", mockUC, mockLogger)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, consumer)
}

// TestEventConsumer_Interface tests that rabbitmqConsumer implements EventConsumer
func TestEventConsumer_Interface(t *testing.T) {
	mockUC := new(MockUseCase)
	mockLogger := &mockLogger{}

	// This test verifies the interface at compile time
	var _ EventConsumer = (*rabbitmqConsumer)(nil)

	// Create a consumer struct (without actual connection)
	consumer := &rabbitmqConsumer{
		uc: mockUC,
		l:  mockLogger,
	}

	// Verify it's not nil
	assert.NotNil(t, consumer)
	assert.NotNil(t, consumer.uc)
	assert.NotNil(t, consumer.l)
}

// TestClose_NilConnection tests Close with nil connection
func TestClose_NilConnection(t *testing.T) {
	consumer := &rabbitmqConsumer{
		conn:    nil,
		channel: nil,
	}

	err := consumer.Close()

	// Should not error with nil connection
	assert.NoError(t, err)
}

// mockLogger is a simple mock logger for testing
type mockLogger struct{}

func (m *mockLogger) Infof(ctx context.Context, format string, args ...interface{})  {}
func (m *mockLogger) Warnf(ctx context.Context, format string, args ...interface{})  {}
func (m *mockLogger) Errorf(ctx context.Context, format string, args ...interface{}) {}
func (m *mockLogger) Fatalf(ctx context.Context, format string, args ...interface{}) {}
func (m *mockLogger) Debugf(ctx context.Context, format string, args ...interface{}) {}
func (m *mockLogger) Debug(ctx context.Context, args ...interface{})                 {}
func (m *mockLogger) Info(ctx context.Context, args ...interface{})                  {}
func (m *mockLogger) Warn(ctx context.Context, args ...interface{})                  {}
func (m *mockLogger) Error(ctx context.Context, args ...interface{})                 {}
func (m *mockLogger) Fatal(ctx context.Context, args ...interface{})                 {}
