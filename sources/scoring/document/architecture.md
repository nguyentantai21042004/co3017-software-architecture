# Scoring Service Architecture

## Overview

The **Scoring Service** handles answer submission, automatic scoring, feedback generation, and event publishing in the Intelligent Tutoring System. It follows **Module-First** architecture with clean separation of concerns.

## Architecture Style

**Module-First (Feature-Based) Architecture** with:
- **Clean Architecture** principles
- **Hexagonal Architecture** (Ports & Adapters)
- **Dependency Injection**
- **Interface-based design**

## Module Structure

```
internal/
└── scoring/                    # MODULE (feature-based)
    ├── delivery/               # Delivery Layer (Adapters)
    │   └── http/
    │       ├── handler.go      # HTTP handlers
    │       ├── routes.go       # Route mapping
    │       ├── presenter.go    # Response transformation
    │       ├── process_request.go  # Request transformation
    │       ├── errors.go       # HTTP error messages
    │       └── new.go          # Handler constructor
    ├── usecase/                # Use Case Layer (Business Logic)
    │   ├── scoring.go          # Core business logic
    │   ├── new.go              # UseCase constructor
    │   └── errors.go           # UseCase error messages
    ├── repository/             # Repository Layer (Data Access)
    │   ├── interface.go        # Repository interface
    │   ├── errors.go           # Repository errors
    │   └── postgre/
    │       └── submission.go   # PostgreSQL implementation
    ├── interface.go            # Module UseCase interface
    ├── type.go                 # Input/Output types
    └── error.go                # Module-level errors
```

## Layer Responsibilities

### 1. Delivery Layer (`delivery/http/`)

**Purpose**: Handle HTTP requests/responses, validation, and transformation.

**Files**:
- `handler.go`: HTTP request handlers
- `routes.go`: Route mapping
- `presenter.go`: Convert domain output to HTTP response
- `process_request.go`: Convert HTTP request to domain input
- `errors.go`: HTTP-specific error messages

**Rules**:
- ✅ Parse and validate HTTP requests
- ✅ Call UseCase layer
- ✅ Transform responses using `pkg/response`
- ✅ Use structured logging with context
- ❌ NO business logic
- ❌ NO direct database access
- ❌ NO external HTTP calls

**Example**:
```go
func (h *handler) SubmitAnswer(c *gin.Context) {
    ctx := c.Request.Context()
    var req SubmitRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        h.l.Errorf(ctx, "scoring.delivery.http.handler.SubmitAnswer: %s | error=%v",
            ErrMsgBindRequestFailed, err)
        response.Error(c, errors.NewHTTPError(http.StatusBadRequest, ErrMsgBindRequestFailed), nil)
        return
    }

    input := req.ToSubmitInput()
    output, err := h.uc.SubmitAnswer(ctx, input)
    if err != nil {
        response.Error(c, errors.NewHTTPError(http.StatusInternalServerError, ErrMsgSubmitAnswerFailed), nil)
        return
    }

    response.OK(c, ToSubmitResponse(output))
}
```

### 2. UseCase Layer (`usecase/`)

**Purpose**: Implement business logic, orchestrate data flow.

**Files**:
- `scoring.go`: Core scoring business logic
- `new.go`: UseCase constructor with dependencies
- `errors.go`: UseCase error messages and context keys

**Rules**:
- ✅ Implement business logic
- ✅ Call `pkg/curl` for external services
- ✅ Call repository for database operations
- ✅ Publish events via EventPublisher
- ✅ Use structured logging with detailed context
- ❌ NO HTTP-specific code
- ❌ NO direct SQL queries

**Example**:
```go
func (uc *usecase) SubmitAnswer(ctx context.Context, input scoring.SubmitInput) (scoring.SubmitOutput, error) {
    uc.l.Infof(ctx, "scoring.usecase.SubmitAnswer: starting | user_id=%s | question_id=%d",
        input.UserID, input.QuestionID)

    // Fetch question from Content Service
    questionResp, err := uc.contentClient.GetQuestion(ctx, input.QuestionID)
    if err != nil {
        uc.l.Errorf(ctx, "scoring.usecase.SubmitAnswer: %s | error=%v",
            ErrMsgFetchQuestionFailed, err)
        return scoring.SubmitOutput{}, fmt.Errorf("%s: %w", ErrMsgFetchQuestionFailed, err)
    }

    // Score the answer
    isCorrect := input.Answer == questionResp.Data.CorrectAnswer
    score := 0
    if isCorrect {
        score = 100
    }

    // Save submission
    submission := &model.Submission{
        UserID:          input.UserID,
        QuestionID:      input.QuestionID,
        SubmittedAnswer: input.Answer,
        ScoreAwarded:    score,
        IsPassed:        score >= 50,
    }

    if err := uc.repo.Create(submission); err != nil {
        return scoring.SubmitOutput{}, err
    }

    // Publish event (async)
    go uc.publishSubmissionEvent(input.UserID, questionResp.Data.SkillTag, score)

    return scoring.SubmitOutput{
        Correct:  isCorrect,
        Score:    score,
        Feedback: generateFeedback(isCorrect),
    }, nil
}
```

### 3. Repository Layer (`repository/`)

**Purpose**: Data access abstraction for database operations.

**Files**:
- `interface.go`: Repository interface
- `errors.go`: Repository-specific errors
- `postgre/submission.go`: PostgreSQL implementation

**Rules**:
- ✅ ONLY database operations
- ✅ Implement repository interface
- ✅ Handle database errors
- ❌ NO business logic
- ❌ NO HTTP calls to external services
- ❌ NO event publishing

**Example**:
```go
type Repository interface {
    Create(submission *model.Submission) error
}

type submissionRepository struct {
    db *sql.DB
}

func New(db *sql.DB) repository.Repository {
    return &submissionRepository{db: db}
}

func (r *submissionRepository) Create(submission *model.Submission) error {
    query := `
        INSERT INTO submissions (user_id, question_id, submitted_answer, score_awarded, is_passed, created_at)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id
    `
    submission.CreatedAt = time.Now()
    return r.db.QueryRow(query, submission.UserID, submission.QuestionID,
        submission.SubmittedAnswer, submission.ScoreAwarded, submission.IsPassed,
        submission.CreatedAt).Scan(&submission.ID)
}
```

## External Service Communication

### pkg/curl Package

All HTTP calls to external services (Content Service, Learner Service) are handled via `pkg/curl`:

```
pkg/
└── curl/
    ├── client.go              # Base HTTP client
    ├── errors.go              # Error types
    ├── error_messages.go      # Error messages & context keys
    ├── types.go               # Request/Response types
    ├── content_service.go     # Content Service client
    ├── content_service_test.go
    └── interfaces.go          # Client interfaces (for mocking)
```

**Rules**:
- ✅ UseCase layer calls `pkg/curl` clients
- ✅ Repository layer is ONLY for database
- ✅ Detailed error messages with context
- ✅ Timeout handling
- ✅ Unit tests with mocked HTTP responses

## Error Handling

### Error Message Convention

Define error constants at the point of error:

```go
// usecase/errors.go
const (
    ErrMsgFetchQuestionFailed     = "failed to fetch question from content service"
    ErrMsgSaveSubmissionFailed    = "failed to save submission to database"
    ErrMsgPublishEventFailed      = "failed to publish submission event"
)

const (
    ErrCtxUserID       = "user_id"
    ErrCtxQuestionID   = "question_id"
    ErrCtxScore        = "score"
)
```

### Error Logging Pattern

```go
uc.l.Errorf(ctx, "scoring.usecase.SubmitAnswer: %s | %s=%s | %s=%d | error=%v",
    ErrMsgFetchQuestionFailed, ErrCtxUserID, input.UserID, 
    ErrCtxQuestionID, input.QuestionID, err)
```

## Logging Convention

### Logger Initialization

```go
// main.go
logger := pkglog.Init(pkglog.ZapConfig{
    Level:    cfg.Logger.Level,
    Mode:     cfg.Logger.Mode,
    Encoding: cfg.Logger.Encoding,
})
```

### Logging Pattern

```go
// Info logs
uc.l.Infof(ctx, "scoring.usecase.SubmitAnswer: starting | user_id=%s | question_id=%d",
    input.UserID, input.QuestionID)

// Error logs
uc.l.Errorf(ctx, "scoring.usecase.SubmitAnswer: %s | user_id=%s | error=%v",
    ErrMsgFetchQuestionFailed, input.UserID, err)
```

**Rules**:
- ✅ All logs use initialized logger from `main.go`
- ✅ Pass `context.Context` to all log calls
- ✅ Use structured logging format: `module.layer.function: message | key=value`
- ❌ NO standard `log` package

## Response Format

All HTTP responses follow `pkg/response.Resp`:

```go
type Resp struct {
    ErrorCode int         `json:"error_code"`
    Message   string      `json:"message"`
    Data      interface{} `json:"data,omitempty"`
}
```

**Success**:
```go
response.OK(c, SubmitResponse{
    Correct:  true,
    Score:    100,
    Feedback: "Correct! Well done.",
})
```

**Error**:
```go
httpErr := errors.NewHTTPError(http.StatusBadRequest, ErrMsgBindRequestFailed)
response.Error(c, httpErr, nil)
```

## Dependency Flow

```
main.go
  ↓
  ├─→ config.Load()
  ├─→ pkglog.Init()
  ├─→ Database Connection
  ├─→ RabbitMQ Publisher
  ├─→ pkg/curl Clients
  ↓
  ├─→ repository/postgre.New(db)
  ├─→ usecase.New(logger, repo, publisher, contentClient)
  ├─→ delivery/http.New(logger, usecase)
  ↓
  └─→ Gin Router + Routes
```

## Testing Strategy

### Unit Tests

```
internal/scoring/
├── delivery/http/
│   └── handler_test.go        # Mock usecase
├── usecase/
│   └── scoring_test.go        # Mock repo + clients
└── repository/postgre/
    └── submission_test.go     # Mock DB or use testcontainers

pkg/curl/
├── content_service_test.go    # Mock HTTP server
└── learner_service_test.go    # Mock HTTP server
```

### Test Commands

```bash
make test              # Run all tests
make test-coverage     # Generate coverage report
make test-short        # Run short tests
```

## Key Principles

1. **Module-First**: Organize by feature/domain, not by technical layer
2. **Separation of Concerns**: Each layer has clear responsibilities
3. **Dependency Inversion**: Depend on interfaces, not implementations
4. **External Service Separation**: Use `pkg/curl` for HTTP calls, not repository
5. **Detailed Error Handling**: Define errors at point of occurrence
6. **Structured Logging**: Use Zap logger with context throughout
7. **Testability**: Use interfaces and dependency injection for easy mocking
8. **Consistency**: Follow established patterns across all modules

## Anti-Patterns to Avoid

❌ **Layer-First Architecture**:
```
internal/
├── handler/        # BAD: Groups by layer
├── service/
└── repository/
```

❌ **HTTP Calls in Repository Layer**:
```go
// BAD: Repository should ONLY access database
func (r *repo) FetchQuestion(id int64) (*Question, error) {
    resp, _ := http.Get("http://content-service/questions/" + id)
    // ...
}
```

❌ **Generic Error Messages**:
```go
// BAD: Too generic
return fmt.Errorf("error fetching data")

// GOOD: Specific with context
return fmt.Errorf("%s: %w | user_id=%s | question_id=%d", 
    ErrMsgFetchQuestionFailed, err, userID, questionID)
```

❌ **Using Standard log Package**:
```go
// BAD
log.Printf("Error: %v", err)

// GOOD
uc.l.Errorf(ctx, "scoring.usecase.SubmitAnswer: %s | error=%v", 
    ErrMsgFetchQuestionFailed, err)
```

## References

- [Clean Architecture by Robert C. Martin](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Hexagonal Architecture](https://alistair.cockburn.us/hexagonal-architecture/)
- [Go Project Layout](https://github.com/golang-standards/project-layout)
