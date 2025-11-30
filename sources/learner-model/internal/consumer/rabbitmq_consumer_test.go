package consumer

import (
	"testing"

	// "learner-model-service/internal/learner" // Not directly used here, rely on shared mocks

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUseCase is a mock implementation of learner.UseCase (defined in test_mocks.go)
// MockLogger and NewMockLogger are defined in test_mocks.go

// TestNewRabbitMQConsumer_InvalidURL tests consumer creation with invalid URL
func TestNewRabbitMQConsumer_InvalidURL(t *testing.T) {
	mockUC := new(MockUseCase)
	mockLogger := NewMockLogger() // Use the factory function

	// Mock logger expectations - need to account for all log calls
	mockLogger.On("Infof", mock.Anything, mock.Anything, mock.Anything).Return()
	mockLogger.On("Errorf", mock.Anything, mock.Anything, mock.Anything).Return() 

	// Try to create consumer with invalid URL
	consumer, err := NewRabbitMQConsumer("invalid://url", mockUC, mockLogger)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, consumer)
	mockLogger.AssertExpectations(t)
}

// TestEventConsumer_Interface tests that rabbitmqConsumer implements EventConsumer
func TestEventConsumer_Interface(t *testing.T) {
	mockUC := new(MockUseCase)
	mockLogger := NewMockLogger() // Use the factory function

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
