package http

import (
	"context"
	"scoring/internal/scoring"
	"scoring/pkg/log" // Import the log package

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

// Mock UseCase for integration tests
type MockUseCase struct {
	mock.Mock
}

func (m *MockUseCase) SubmitAnswer(ctx context.Context, input scoring.SubmitInput) (scoring.SubmitOutput, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		return scoring.SubmitOutput{}, args.Error(1)
	}
	return args.Get(0).(scoring.SubmitOutput), args.Error(1)
}

func (m *MockUseCase) GetAnsweredQuestions(ctx context.Context, userID, skillTag string) ([]int64, error) {
	args := m.Called(ctx, userID, skillTag)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]int64), args.Error(1)
}

// Mock Logger for integration tests
type MockLogger struct {
	mock.Mock
}

func (m *MockLogger) Debug(ctx context.Context, arg ...any) {
	m.Called(ctx, arg)
}

func (m *MockLogger) Debugf(ctx context.Context, template string, arg ...any) {
	m.Called(ctx, template, arg)
}

func (m *MockLogger) Info(ctx context.Context, arg ...any) {
	m.Called(ctx, arg)
}

func (m *MockLogger) Infof(ctx context.Context, template string, arg ...any) {
	m.Called(ctx, template, arg)
}

func (m *MockLogger) Warn(ctx context.Context, arg ...any) {
	m.Called(ctx, arg)
}

func (m *MockLogger) Warnf(ctx context.Context, template string, arg ...any) {
	m.Called(ctx, template, arg)
}

func (m *MockLogger) Error(ctx context.Context, arg ...any) {
	m.Called(ctx, arg)
}

func (m *MockLogger) Errorf(ctx context.Context, template string, arg ...any) {
	m.Called(ctx, template, arg)
}

func (m *MockLogger) Fatal(ctx context.Context, arg ...any) {
	m.Called(ctx, arg)
}

func (m *MockLogger) Fatalf(ctx context.Context, template string, arg ...any) {
	m.Called(ctx, template, arg)
}

func NewMockLogger() *MockLogger {
	mockLogger := new(MockLogger)
	mockLogger.On("Infof", mock.Anything, mock.Anything, mock.Anything).Return()
	mockLogger.On("Errorf", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return()
	return mockLogger
}


// setupRouter configures a gin router for testing
func SetupTestRouter(mockUc scoring.UseCase, mockLogger log.Logger) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	// CORS Middleware - copy from main.go
	r.Use(func(c *gin.Context) {
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
	})

	handler := New(mockLogger, mockUc)
	api := r.Group("/api/scoring")
	MapScoringRoutes(api, handler)
	// Health check (global)
	r.GET("/health", handler.Health)

	return r
}
