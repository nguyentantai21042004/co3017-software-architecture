package http

import (
	"context"
	"learner-model-service/internal/learner"
	"learner-model-service/pkg/log" // Import the log package

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

// MockUseCase is a mock implementation of learner.UseCase
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

// CORSMiddleware handles CORS configuration for testing purposes
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// SetupIntegrationTestRouter configures a gin router for integration testing,
// mapping only the specific API routes needed for these tests.
func SetupIntegrationTestRouter(mockUc *MockUseCase, mockLogger *MockLogger) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	r.Use(CORSMiddleware()) // Use the extracted CORS middleware

	handler := New(mockLogger, mockUc)

	// Health check (global)
	r.GET("/health", handler.Health)

	// API routes
	internal := r.Group("/internal/learner")
	internal.GET("/mastery/:user_id", handler.GetMastery)
	
	return r
}

// SetupTestRouter configures a gin router for testing (wrapper for integration test setup)
func SetupTestRouter(mockUc learner.UseCase, mockLogger log.Logger) *gin.Engine {
	// Cast interfaces to concrete mock types for SetupIntegrationTestRouter
	concreteMockUc, ok := mockUc.(*MockUseCase)
	if !ok {
		panic("mockUc is not of type *MockUseCase")
	}
	concreteMockLogger, ok := mockLogger.(*MockLogger)
	if !ok {
		panic("mockLogger is not of type *MockLogger")
	}
	return SetupIntegrationTestRouter(concreteMockUc, concreteMockLogger)
}
