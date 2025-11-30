package consumer

import (
	"encoding/json"
	"errors"
	"testing"
	"time"

	"learner-model-service/internal/learner"
	"learner-model-service/internal/model"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// MockUseCase, MockLogger, NewMockLogger, MockAMQPChannel, NewMockAMQPChannel,
// MockAMQPConnection, NewMockAMQPConnection, and amqpDialer are now defined in test_mocks.go

func TestRabbitMQConsumer_ProcessSubmissionEvent(t *testing.T) {
	// Arrange
	mockUC := new(MockUseCase)
	mockLogger := NewMockLogger()
	mockChannel := NewMockAMQPChannel()
	mockConnection := NewMockAMQPConnection(mockChannel)

	// Mock logger expectations
	mockLogger.On("Infof", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return()
	mockLogger.On("Errorf", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return()

	// Setup mock expectations
	mockConnection.On("Channel").Return(mockChannel, nil)
	mockConnection.On("IsClosed").Return(false)
	mockConnection.On("Close").Return(nil)
	mockChannel.On("ExchangeDeclare", "its.events", "topic", true, false, false, false, amqp.Table(nil)).Return(nil)
	mockChannel.On("QueueDeclare", "learner.updates", true, false, false, false, amqp.Table(nil)).Return(amqp.Queue{Name: "learner.updates"}, nil)
	mockChannel.On("QueueBind", "learner.updates", "event.submission", "its.events", false, amqp.Table(nil)).Return(nil)
	mockChannel.On("Consume", "learner.updates", "", true, false, false, false, amqp.Table(nil)).Return(mockChannel.queue, nil)
	mockChannel.On("Close").Return(nil)

	// Expected event data
	eventData := model.SubmissionEvent{
		Event:         "SubmissionCompleted",
		UserID:        "user123",
		SkillTag:      "math",
		ScoreObtained: 75,
		Timestamp:     time.Now().Format(time.RFC3339),
	}
	eventBody, _ := json.Marshal(eventData)

	// Mock UseCase expectation
	expectedUpdateInput := learner.UpdateMasteryInput{
		UserID:        eventData.UserID,
		SkillTag:      eventData.SkillTag,
		ScoreObtained: eventData.ScoreObtained,
	}
	mockUC.On("UpdateMasteryFromEvent", mock.Anything, expectedUpdateInput).Return(nil)

	// Create consumer using test helper with mock connection
	consumer, err := newRabbitMQConsumerWithConnection(mockConnection, mockUC, mockLogger)
	require.NoError(t, err)
	err = consumer.Start()
	require.NoError(t, err)

	// Simulate message delivery
	mockChannel.queue <- amqp.Delivery{
		Body: eventBody,
		// Other fields can be left default as they are not used by handleMessage
	}

	// Give the goroutine some time to process the message
	time.Sleep(50 * time.Millisecond)

	// Assert UseCase expectations
	mockUC.AssertExpectations(t)
	// The logger is called with variadic arguments, so we check for the call with the format string
	mockLogger.AssertCalled(t, "Infof", mock.Anything, "consumer.handleMessage: successfully processed event | user_id=%s | skill_tag=%s", mock.Anything)

	// Clean up
	consumer.Close()
	mockChannel.AssertCalled(t, "Close")
	mockConnection.AssertCalled(t, "Close")
}

func TestRabbitMQConsumer_UnmarshalError(t *testing.T) {
	// Arrange
	mockUC := new(MockUseCase)
	mockLogger := NewMockLogger()
	mockChannel := NewMockAMQPChannel()
	mockConnection := NewMockAMQPConnection(mockChannel)

	// Mock logger expectations
	mockLogger.On("Infof", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return()
	mockLogger.On("Errorf", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return()

	// Setup mock expectations
	mockConnection.On("Channel").Return(mockChannel, nil)
	mockConnection.On("IsClosed").Return(false)
	mockConnection.On("Close").Return(nil)
	mockChannel.On("ExchangeDeclare", "its.events", "topic", true, false, false, false, amqp.Table(nil)).Return(nil)
	mockChannel.On("QueueDeclare", "learner.updates", true, false, false, false, amqp.Table(nil)).Return(amqp.Queue{Name: "learner.updates"}, nil)
	mockChannel.On("QueueBind", "learner.updates", "event.submission", "its.events", false, amqp.Table(nil)).Return(nil)
	mockChannel.On("Consume", "learner.updates", "", true, false, false, false, amqp.Table(nil)).Return(mockChannel.queue, nil)
	mockChannel.On("Close").Return(nil)

	// Create consumer using test helper with mock connection
	consumer, err := newRabbitMQConsumerWithConnection(mockConnection, mockUC, mockLogger)
	require.NoError(t, err)
	err = consumer.Start()
	require.NoError(t, err)

	// Simulate invalid message delivery
	invalidBody := []byte("{invalid json")
	mockChannel.queue <- amqp.Delivery{
		Body: invalidBody,
	}

	// Give the goroutine some time to process the message
	time.Sleep(50 * time.Millisecond)

	// Assert that UpdateMasteryFromEvent was NOT called
	mockUC.AssertNotCalled(t, "UpdateMasteryFromEvent", mock.Anything, mock.Anything)
	mockLogger.AssertCalled(t, "Errorf", mock.Anything, "consumer.handleMessage: failed to parse event | error=%v", mock.Anything)

	// Clean up
	consumer.Close()
	mockChannel.AssertCalled(t, "Close")
	mockConnection.AssertCalled(t, "Close")
}

func TestRabbitMQConsumer_UseCaseError(t *testing.T) {
	// Arrange
	mockUC := new(MockUseCase)
	mockLogger := NewMockLogger()
	mockChannel := NewMockAMQPChannel()
	mockConnection := NewMockAMQPConnection(mockChannel)

	// Mock logger expectations
	mockLogger.On("Infof", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return()
	mockLogger.On("Errorf", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return()

	// Setup mock expectations
	mockConnection.On("Channel").Return(mockChannel, nil)
	mockConnection.On("IsClosed").Return(false)
	mockConnection.On("Close").Return(nil) 
	mockChannel.On("ExchangeDeclare", "its.events", "topic", true, false, false, false, amqp.Table(nil)).Return(nil)
	mockChannel.On("QueueDeclare", "learner.updates", true, false, false, false, amqp.Table(nil)).Return(amqp.Queue{Name: "learner.updates"}, nil)
	mockChannel.On("QueueBind", "learner.updates", "event.submission", "its.events", false, amqp.Table(nil)).Return(nil)
	mockChannel.On("Consume", "learner.updates", "", true, false, false, false, amqp.Table(nil)).Return(mockChannel.queue, nil)
	mockConnection.On("Close").Return(nil)
	mockChannel.On("Close").Return(nil)

	// Expected event data
	eventData := model.SubmissionEvent{
		Event:         "SubmissionCompleted",
		UserID:        "user123",
		SkillTag:      "math",
		ScoreObtained: 75,
		Timestamp:     time.Now().Format(time.RFC3339),
	}
	eventBody, _ := json.Marshal(eventData)

	// Mock UseCase expectation - simulate an error
	expectedUseCaseError := errors.New("usecase processing failed")
	mockUC.On("UpdateMasteryFromEvent", mock.Anything, mock.Anything).Return(expectedUseCaseError)

	// Create consumer using test helper with mock connection
	consumer, err := newRabbitMQConsumerWithConnection(mockConnection, mockUC, mockLogger)
	require.NoError(t, err)
	err = consumer.Start()
	require.NoError(t, err)

	// Simulate message delivery
	mockChannel.queue <- amqp.Delivery{
		Body: eventBody,
	}

	// Give the goroutine some time to process the message
	time.Sleep(50 * time.Millisecond)

	// Assert UseCase expectations (called once, returned error)
	mockUC.AssertExpectations(t)
	mockLogger.AssertCalled(t, "Errorf", mock.Anything, "consumer.handleMessage: failed to process event | user_id=%s | skill_tag=%s | error=%v", mock.Anything)

	// Clean up
	consumer.Close()
	mockChannel.AssertCalled(t, "Close")
	mockConnection.AssertCalled(t, "Close")
}
