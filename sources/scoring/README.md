# Scoring Service

**Microservice đánh giá câu trả lời và xuất bản sự kiện**

---

## Mục lục

- [Tổng quan](#tổng-quan)
- [Công nghệ](#công-nghệ)
- [Kiến trúc](#kiến-trúc)
- [Cơ sở dữ liệu](#cơ-sở-dữ-liệu)
- [API Endpoints](#api-endpoints)
- [Event Publishing (RabbitMQ)](#event-publishing-rabbitmq)
- [Cấu hình](#cấu-hình)
- [Phát triển Local](#phát-triển-local)
- [Testing](#testing)
- [Hexagonal Architecture](#hexagonal-architecture)

---

## Tổng quan

Scoring Service là microservice chịu trách nhiệm đánh giá câu trả lời của học sinh và xuất bản sự kiện trong hệ thống ITS (Intelligent Tutoring System). Service này cung cấp các chức năng:

- **Đánh giá câu trả lời**: So sánh câu trả lời của học sinh với đáp án đúng
- **Tính điểm tự động**: Tính toán điểm số dựa trên độ chính xác
- **Lưu trữ submission**: Ghi lại lịch sử trả lời của học sinh
- **Xuất bản sự kiện**: Gửi sự kiện đến RabbitMQ để các service khác xử lý
- **Theo dõi tiến độ**: Cung cấp API để truy vấn câu hỏi đã trả lời

Scoring Service hoạt động như một **trung gian xử lý** giữa học sinh và hệ thống đánh giá, đồng thời kích hoạt các quy trình cập nhật mô hình học sinh thông qua event-driven architecture.

---

## Công nghệ

### Core Technologies

- **Go 1.23**: Ngôn ngữ lập trình chính
- **Gin**: HTTP web framework
- **PostgreSQL**: Cơ sở dữ liệu quan hệ
- **RabbitMQ**: Message broker cho event publishing
- **SQLBoiler**: Type-safe ORM và code generation

### Libraries & Tools

- **amqp091-go**: RabbitMQ client library
- **Zap**: Structured logging
- **Swagger/Swaggo**: API documentation tự động
- **UUID**: Unique identifier generation
- **go-i18n**: Internationalization support

### Architecture Pattern

- **Hexagonal Architecture (Ports & Adapters)**: Tách biệt business logic khỏi infrastructure
- **Module-First Architecture**: Tổ chức code theo feature/domain
- **Event-Driven Architecture**: Sử dụng RabbitMQ cho async communication
- **Clean Architecture**: Dependency inversion và separation of concerns

---

## Kiến trúc

Scoring Service được thiết kế theo **Hexagonal Architecture** (Ports & Adapters) với **Module-First** approach:

```
┌─────────────────────────────────────────────────────────┐
│                  HTTP Delivery Layer                    │
│  (delivery/http/) - Handlers, Routes, DTOs              │
└─────────────────────────────────────────────────────────┘
                        ↓
┌─────────────────────────────────────────────────────────┐
│                   UseCase Layer                         │
│  (usecase/) - Business Logic & Orchestration            │
└─────────────────────────────────────────────────────────┘
                        ↓
┌──────────────────┬──────────────────┬───────────────────┐
│   Repository     │  Event Publisher │  HTTP Client      │
│   (PostgreSQL)   │   (RabbitMQ)     │  (Content Svc)    │
└──────────────────┴──────────────────┴───────────────────┘
```

### Hexagonal Architecture Layers

**1. Core (Business Logic)**

- UseCase layer chứa business rules
- Không phụ thuộc vào infrastructure
- Định nghĩa interfaces (ports) cho external dependencies

**2. Adapters (Infrastructure)**

- **Primary Adapters** (Inbound): HTTP handlers nhận requests
- **Secondary Adapters** (Outbound): Repository, Event Publisher, HTTP Client

**3. Ports (Interfaces)**

- Repository interface: Data access port
- EventPublisher interface: Messaging port
- ContentClient interface: External service port

### Request Flow

```
HTTP Request (POST /api/scoring/submit)
    ↓
Handler (delivery/http/handler.go)
    ↓ [Validate & Transform Request]
UseCase (usecase/scoring.go)
    ↓
    ├─→ ContentClient.GetQuestion() [HTTP Call to Content Service]
    │   ↓ [Fetch correct answer & skill tag]
    ├─→ Score Answer [Business Logic]
    │   ↓ [Compare submitted answer with correct answer]
    ├─→ Repository.Create() [Save to PostgreSQL]
    │   ↓ [Store submission record]
    └─→ EventPublisher.PublishSubmissionEvent() [Async]
        ↓ [Send event to RabbitMQ]
        ↓
    RabbitMQ Exchange (its.events)
        ↓ [Routing key: event.submission]
    Learner Model Consumer
        ↓ [Update student mastery]
```

### Directory Structure

```
sources/scoring/
├── cmd/
│   └── api/
│       ├── main.go                      # Application entry point
│       └── Dockerfile                   # Container image
│
├── internal/
│   ├── scoring/                         # MODULE (feature-based)
│   │   ├── delivery/                    # Primary Adapters
│   │   │   └── http/
│   │   │       ├── handler.go           # HTTP handlers
│   │   │       ├── routes.go            # Route mapping
│   │   │       ├── presenter.go         # Response transformation
│   │   │       └── process_request.go   # Request transformation
│   │   ├── usecase/                     # Core Business Logic
│   │   │   ├── scoring.go               # Scoring logic
│   │   │   ├── new.go                   # UseCase constructor
│   │   │   └── errors.go                # Error messages
│   │   ├── repository/                  # Secondary Adapter (Data)
│   │   │   ├── interface.go             # Repository port
│   │   │   └── postgre/
│   │   │       └── submission.go        # PostgreSQL implementation
│   │   ├── interface.go                 # UseCase interface (port)
│   │   ├── type.go                      # Input/Output types
│   │   └── error.go                     # Module errors
│   │
│   ├── publisher/                       # Secondary Adapter (Events)
│   │   └── publisher.go                 # RabbitMQ publisher
│   │
│   └── model/
│       └── submission.go                # Domain models
│
├── pkg/                                 # Shared packages
│   ├── curl/                            # HTTP client for external services
│   │   └── content_service.go           # Content Service client
│   ├── log/                             # Structured logging
│   ├── response/                        # Standard response format
│   └── errors/                          # Error handling utilities
│
├── config/
│   └── config.go                        # Configuration management
│
├── migration/
│   └── 01_add_project.sql               # Database migrations
│
└── docs/
    ├── swagger.json                     # OpenAPI specification
    └── swagger.yaml
```

---

## Cơ sở dữ liệu

### Database Schema

Scoring Service sử dụng PostgreSQL database với schema sau:

```mermaid
erDiagram
    SUBMISSIONS {
        BIGINT id PK "Primary Key, Auto-increment"
        VARCHAR user_id "ID của học sinh"
        BIGINT question_id "ID câu hỏi"
        VARCHAR submitted_answer "Câu trả lời đã nộp"
        INTEGER score_awarded "Điểm số đạt được"
        BOOLEAN is_passed "Đạt hay không (>= 50)"
        TIMESTAMP created_at "Thời gian nộp bài"
    }
```

### Table: submissions

| Column             | Type         | Constraints                 | Description                     |
| ------------------ | ------------ | --------------------------- | ------------------------------- |
| `id`               | BIGINT       | PRIMARY KEY, AUTO_INCREMENT | ID duy nhất của submission      |
| `user_id`          | VARCHAR(255) | NOT NULL                    | ID của học sinh                 |
| `question_id`      | BIGINT       | NOT NULL                    | ID câu hỏi (từ Content Service) |
| `submitted_answer` | VARCHAR(255) | NOT NULL                    | Câu trả lời học sinh đã nộp     |
| `score_awarded`    | INTEGER      | NOT NULL                    | Điểm số (0-100)                 |
| `is_passed`        | BOOLEAN      | NOT NULL                    | True nếu điểm >= 50             |
| `created_at`       | TIMESTAMP    | NOT NULL                    | Thời gian nộp bài               |

### Indexes

```sql
CREATE INDEX idx_submissions_user_id ON submissions(user_id);
CREATE INDEX idx_submissions_question_id ON submissions(question_id);
CREATE INDEX idx_submissions_created_at ON submissions(created_at);
```

### Scoring Logic

```
IF submitted_answer == correct_answer:
    score = 100
    is_passed = true
ELSE:
    score = 0
    is_passed = false
```

---

## API Endpoints

### Base URL

```
http://localhost:8082
```

### Swagger UI

Truy cập API documentation tại:

```
http://localhost:8082/scoring/swagger/index.html
```

### Endpoints

#### 1. Health Check

**GET** `/health`

Kiểm tra trạng thái service.

**Response:**

```json
{
  "error_code": 0,
  "message": "Healthy",
  "data": {
    "status": "healthy",
    "service": "scoring"
  }
}
```

**Curl Example:**

```bash
curl http://localhost:8082/health
```

---

#### 2. Submit Answer

**POST** `/api/scoring/submit`

Nộp câu trả lời để đánh giá và nhận feedback ngay lập tức.

**Request Body:**

```json
{
  "user_id": "user_01",
  "question_id": 1,
  "answer": "A"
}
```

**Response:**

```json
{
  "error_code": 0,
  "message": "Success",
  "data": {
    "correct": true,
    "score": 100,
    "feedback": "Correct! Well done."
  }
}
```

**Error Response (Wrong Answer):**

```json
{
  "error_code": 0,
  "message": "Success",
  "data": {
    "correct": false,
    "score": 0,
    "feedback": "Incorrect answer. Please try again!"
  }
}
```

**Curl Example:**

```bash
curl -X POST http://localhost:8082/api/scoring/submit \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "user_01",
    "question_id": 1,
    "answer": "A"
  }'
```

**Flow:**

1. Validate request (user_id, question_id, answer không được rỗng)
2. Gọi Content Service để lấy correct_answer và skill_tag
3. So sánh submitted_answer với correct_answer
4. Tính điểm: 100 nếu đúng, 0 nếu sai
5. Lưu submission vào database
6. Xuất bản event đến RabbitMQ (async)
7. Trả về kết quả ngay lập tức

---

#### 3. Get Answered Questions

**GET** `/api/scoring/answered-questions`

Lấy danh sách ID các câu hỏi mà học sinh đã trả lời cho một kỹ năng cụ thể.

**Query Parameters:**

- `user_id` (string, required): ID của học sinh
- `skill` (string, required): Skill tag (vd: "math_arithmetic")

**Response:**

```json
{
  "error_code": 0,
  "message": "Success",
  "data": [1, 3, 5, 7, 9]
}
```

**Curl Example:**

```bash
curl "http://localhost:8082/api/scoring/answered-questions?user_id=user_01&skill=math_arithmetic"
```

**Use Case:**

- Content Service sử dụng endpoint này để lọc câu hỏi đã trả lời
- Adaptive Engine sử dụng để tránh đề xuất câu hỏi trùng lặp

---

## Event Publishing (RabbitMQ)

### Overview

Scoring Service sử dụng **RabbitMQ** để xuất bản sự kiện submission đến các service khác trong hệ thống. Đây là một phần quan trọng của **Event-Driven Architecture**.

### RabbitMQ Configuration

```
Exchange: its.events
Type: topic
Routing Key: event.submission
Durable: true
```

### Event Flow

```
Scoring Service
    ↓ [After saving submission]
EventPublisher.PublishSubmissionEvent()
    ↓ [Marshal to JSON]
RabbitMQ Exchange (its.events)
    ↓ [Route by key: event.submission]
Queue: learner_queue
    ↓ [Bind to exchange]
Learner Model Consumer
    ↓ [Process event]
Update Student Mastery (BKT Algorithm)
```

### Event Structure

**SubmissionEvent:**

```json
{
  "event": "SubmissionCompleted",
  "user_id": "user_01",
  "skill_tag": "math_arithmetic",
  "score_obtained": 100,
  "timestamp": "2024-12-07T10:30:00Z"
}
```

**Fields:**

| Field            | Type   | Description                               |
| ---------------- | ------ | ----------------------------------------- |
| `event`          | string | Event type (always "SubmissionCompleted") |
| `user_id`        | string | ID của học sinh                           |
| `skill_tag`      | string | Kỹ năng được đánh giá                     |
| `score_obtained` | int    | Điểm số đạt được (0-100)                  |
| `timestamp`      | string | Thời gian event (RFC3339 format)          |

### Publisher Implementation

**File:** `internal/publisher/publisher.go`

```go
type EventPublisher interface {
    PublishSubmissionEvent(event interface{}) error
    Close() error
}

type rabbitmqPublisher struct {
    conn    *amqp.Connection
    channel *amqp.Channel
    l       log.Logger
}
```

**Key Features:**

1. **Connection Management**: Tự động kết nối và tạo channel
2. **Exchange Declaration**: Tự động tạo exchange nếu chưa tồn tại
3. **Timeout Handling**: Context timeout 5 giây cho mỗi publish
4. **Error Logging**: Log chi tiết khi publish thất bại
5. **Graceful Shutdown**: Close connection khi service shutdown

### Async Publishing

Event publishing được thực hiện **asynchronously** để không block HTTP response:

```go
// Step 4: Publish event to RabbitMQ (async, don't block response)
go func() {
    bgCtx := context.Background()
    event := model.SubmissionEvent{
        Event:         "SubmissionCompleted",
        UserID:        input.UserID,
        SkillTag:      skillTag,
        ScoreObtained: score,
        Timestamp:     time.Now().Format(time.RFC3339),
    }

    err := uc.publisher.PublishSubmissionEvent(event)
    if err != nil {
        uc.l.Errorf(bgCtx, "Failed to publish event | error=%v", err)
    }
}()
```

**Benefits:**

- HTTP response trả về ngay lập tức
- Không bị delay bởi RabbitMQ latency
- Nếu publish thất bại, không ảnh hưởng đến user experience

### Monitoring Events

**View RabbitMQ Management UI:**

```bash
# Access RabbitMQ Management
open http://localhost:15672

# Default credentials
Username: admintest
Password: adminTest2025
```

**Check Exchange:**

- Navigate to "Exchanges" tab
- Find "its.events" exchange
- View bindings and message rates

**Check Queue:**

- Navigate to "Queues" tab
- Find "learner_queue"
- View message count and consumers

---

## Cấu hình

### Environment Variables

Scoring Service sử dụng các biến môi trường sau:

| Variable               | Description              | Default                                          | Required |
| ---------------------- | ------------------------ | ------------------------------------------------ | -------- |
| `APP_HOST`             | Server host              | `0.0.0.0`                                        | No       |
| `APP_PORT`             | Server port              | `8082`                                           | No       |
| `POSTGRES_HOST`        | PostgreSQL host          | `localhost`                                      | Yes      |
| `POSTGRES_PORT`        | PostgreSQL port          | `5432`                                           | Yes      |
| `POSTGRES_USER`        | Database username        | `postgres`                                       | Yes      |
| `POSTGRES_PASSWORD`    | Database password        | `postgres`                                       | Yes      |
| `POSTGRES_DB`          | Database name            | `scoring_db`                                     | Yes      |
| `POSTGRES_SSLMODE`     | SSL mode                 | `disable`                                        | No       |
| `RABBITMQ_URL`         | RabbitMQ connection URL  | `amqp://admintest:adminTest2025@localhost:5672/` | Yes      |
| `RABBITMQ_EXCHANGE`    | RabbitMQ exchange name   | `its.events`                                     | No       |
| `RABBITMQ_QUEUE`       | RabbitMQ queue name      | `scoring_queue`                                  | No       |
| `RABBITMQ_ROUTING_KEY` | RabbitMQ routing key     | `event.submission`                               | No       |
| `CONTENT_SERVICE_URL`  | Content Service base URL | `http://localhost:8081/api/content`              | Yes      |
| `LOG_LEVEL`            | Logging level            | `debug`                                          | No       |

### Configuration File

**File:** `config/config.go`

```go
type Config struct {
    HTTPServer struct {
        Host string
        Port int
        Mode string // "debug" or "release"
    }
    Postgres struct {
        Host     string
        Port     int
        User     string
        Password string
        DBName   string
        SSLMode  string
    }
    RabbitMQ struct {
        URL        string
        Exchange   string
        Queue      string
        RoutingKey string
    }
    ContentServiceURL string
    Logger struct {
        Level    string
        Mode     string
        Encoding string
    }
}
```

### Docker Compose Configuration

**File:** `sources/docker-compose.yml`

```yaml
scoring-service:
  build:
    context: ./scoring
    dockerfile: cmd/api/Dockerfile
  environment:
    APP_HOST: 0.0.0.0
    APP_PORT: 8082
    POSTGRES_HOST: postgres
    POSTGRES_PORT: 5432
    POSTGRES_USER: postgres
    POSTGRES_PASSWORD: postgres
    POSTGRES_DB: scoring_db
    POSTGRES_SSLMODE: disable
    RABBITMQ_URL: amqp://admintest:adminTest2025@rabbitmq:5672/
    CONTENT_SERVICE_URL: http://content-service:8081/api/content
    LOG_LEVEL: debug
  ports:
    - "8082:8082"
  networks:
    - its-network
```

---

## Phát triển Local

### Prerequisites

- **Go 1.23** hoặc cao hơn
- **PostgreSQL 14+** (hoặc sử dụng Docker)
- **RabbitMQ 3.12+** (hoặc sử dụng Docker)
- **Content Service** đang chạy (port 8081)
- **IDE**: VS Code với Go extension, GoLand, hoặc Vim

### Setup Infrastructure

#### Option 1: Sử dụng Docker Compose (Recommended)

```bash
# Từ thư mục sources/
cd sources

# Start infrastructure services (PostgreSQL, RabbitMQ)
docker-compose -f docker-compose.infra.yml up -d

# Verify services are running
docker ps | grep -E "postgres|rabbitmq"
```

#### Option 2: Local Installation

**PostgreSQL:**

```bash
# macOS
brew install postgresql@14
brew services start postgresql@14

# Create database
createdb scoring_db
```

**RabbitMQ:**

```bash
# macOS
brew install rabbitmq
brew services start rabbitmq

# Enable management plugin
rabbitmq-plugins enable rabbitmq_management
```

### Setup Environment Variables

**Create `.env` file:**

```bash
cd sources/scoring

# Copy template
cp template.env .env

# Edit .env with your local configuration
```

**Example `.env` for local development:**

```env
# Server Configuration
APP_HOST=0.0.0.0
APP_PORT=8082

# Database Configuration (Local)
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=scoring_db
POSTGRES_SSLMODE=disable

# RabbitMQ Configuration (Local)
RABBITMQ_URL=amqp://guest:guest@localhost:5672/

# External Services
CONTENT_SERVICE_URL=http://localhost:8081/api/content

# Logging
LOG_LEVEL=debug
```

### Build Project

```bash
cd sources/scoring

# Download dependencies
go mod download

# Tidy dependencies
go mod tidy

# Build binary
go build -o bin/scoring-api cmd/api/main.go
```

### Run Application

#### Option 1: Go Run (Development)

```bash
# Ensure Content Service is running first
cd sources/scoring

# Run with go run
go run cmd/api/main.go
```

#### Option 2: Compiled Binary

```bash
# Build first
go build -o bin/scoring-api cmd/api/main.go

# Run binary
./bin/scoring-api
```

#### Option 3: Using Makefile

```bash
cd sources/scoring

# Run service
make run

# Or build and run
make build
make start
```

### Verify Service

```bash
# Check health
curl http://localhost:8082/health

# Expected output:
# {"error_code":0,"message":"Healthy","data":{"status":"healthy","service":"scoring"}}

# Access Swagger UI
open http://localhost:8082/scoring/swagger/index.html

# Test submit endpoint
curl -X POST http://localhost:8082/api/scoring/submit \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "test_user",
    "question_id": 1,
    "answer": "A"
  }'
```

### Hot Reload (Development)

Sử dụng **Air** cho hot reload:

```bash
# Install Air
go install github.com/cosmtrek/air@latest

# Run with Air
cd sources/scoring
air

# Air sẽ tự động rebuild khi code thay đổi
```

**Create `.air.toml`:**

```toml
root = "."
tmp_dir = "tmp"

[build]
  cmd = "go build -o ./tmp/main cmd/api/main.go"
  bin = "tmp/main"
  include_ext = ["go", "tpl", "tmpl", "html"]
  exclude_dir = ["tmp", "vendor", "bin"]
  delay = 1000
```

---

## Testing

### Test Structure

```
sources/scoring/
├── internal/
│   └── scoring/
│       ├── delivery/http/
│       │   └── handler_test.go          # HTTP handler tests
│       ├── usecase/
│       │   └── scoring_test.go          # Business logic tests
│       └── repository/postgre/
│           └── submission_test.go       # Repository tests
│
└── pkg/
    └── curl/
        └── content_service_test.go      # HTTP client tests
```

### Running Tests

```bash
cd sources/scoring

# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html

# View coverage report
open coverage.html

# Run specific test
go test -v ./internal/scoring/usecase -run TestSubmitAnswer

# Run with race detector
go test -race ./...
```

### Test Categories

#### 1. Unit Tests (UseCase Layer)

Test business logic với mocked dependencies:

```go
func TestSubmitAnswer_CorrectAnswer(t *testing.T) {
    // Arrange
    mockRepo := new(MockRepository)
    mockPublisher := new(MockEventPublisher)
    mockContentClient := new(MockContentClient)

    uc := usecase.New(logger, mockRepo, mockPublisher, mockContentClient)

    // Mock Content Service response
    mockContentClient.On("GetQuestion", mock.Anything, int64(1)).
        Return(&ContentResponse{
            Data: Question{
                CorrectAnswer: "A",
                SkillTag: "math",
            },
        }, nil)

    // Mock Repository
    mockRepo.On("Create", mock.Anything, mock.Anything).Return(nil)

    // Act
    output, err := uc.SubmitAnswer(ctx, scoring.SubmitInput{
        UserID: "user_01",
        QuestionID: 1,
        Answer: "A",
    })

    // Assert
    assert.NoError(t, err)
    assert.True(t, output.Correct)
    assert.Equal(t, 100, output.Score)
}
```

#### 2. Integration Tests (HTTP Layer)

Test HTTP handlers với mocked use case:

```go
func TestHandler_SubmitAnswer(t *testing.T) {
    // Setup
    gin.SetMode(gin.TestMode)
    mockUC := new(MockUseCase)
    handler := http.New(logger, mockUC)

    router := gin.New()
    router.POST("/api/scoring/submit", handler.SubmitAnswer)

    // Mock use case
    mockUC.On("SubmitAnswer", mock.Anything, mock.Anything).
        Return(scoring.SubmitOutput{
            Correct: true,
            Score: 100,
            Feedback: "Correct!",
        }, nil)

    // Request
    body := `{"user_id":"user_01","question_id":1,"answer":"A"}`
    req := httptest.NewRequest("POST", "/api/scoring/submit", strings.NewReader(body))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()

    // Execute
    router.ServeHTTP(w, req)

    // Assert
    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "Correct!")
}
```

#### 3. Repository Tests

Test database operations với test database:

```go
func TestRepository_Create(t *testing.T) {
    // Setup test database
    db := setupTestDB(t)
    defer db.Close()

    repo := postgre.New(db, logger)

    // Test data
    submission := &model.Submission{
        UserID: "user_01",
        QuestionID: 1,
        SubmittedAnswer: "A",
        ScoreAwarded: 100,
        IsPassed: true,
    }

    // Execute
    err := repo.Create(context.Background(), submission)

    // Assert
    assert.NoError(t, err)
    assert.NotZero(t, submission.ID)
}
```

### Test Coverage Goals

- **Unit Tests**: > 80% coverage
- **Integration Tests**: All API endpoints
- **Repository Tests**: All CRUD operations

### Mocking

Sử dụng **testify/mock** cho mocking:

```bash
# Install mockery
go install github.com/vektra/mockery/v2@latest

# Generate mocks
mockery --name=UseCase --dir=internal/scoring --output=internal/scoring/mocks
mockery --name=Repository --dir=internal/scoring/repository --output=internal/scoring/repository/mocks
```

---

## Hexagonal Architecture

### Overview

Scoring Service áp dụng **Hexagonal Architecture** (còn gọi là Ports & Adapters) để đạt được:

- **Testability**: Dễ dàng test business logic độc lập
- **Flexibility**: Dễ dàng thay đổi infrastructure (database, message broker)
- **Maintainability**: Code rõ ràng, dễ maintain
- **Independence**: Business logic không phụ thuộc vào framework

### Core Concepts

```
┌─────────────────────────────────────────────────────────┐
│                    PRIMARY ADAPTERS                     │
│              (Inbound - Drive the app)                  │
│                                                         │
│  ┌─────────────────────────────────────────────────┐    │
│  │         HTTP Handler (Gin)                      │    │
│  │  - Receive HTTP requests                        │    │
│  │  - Validate input                               │    │
│  │  - Call use case                                │    │
│  │  - Return HTTP response                         │    │
│  └─────────────────────────────────────────────────┘    │
└─────────────────────────────────────────────────────────┘
                        ↓
┌─────────────────────────────────────────────────────────┐
│                    CORE (HEXAGON)                       │
│                  Business Logic                         │
│                                                         │
│  ┌─────────────────────────────────────────────────┐    │
│  │         UseCase (Scoring Logic)                 │    │
│  │  - Fetch question from Content Service          │    │
│  │  - Compare answers                              │    │
│  │  - Calculate score                              │    │
│  │  - Save submission                              │    │
│  │  - Publish event                                │    │
│  └─────────────────────────────────────────────────┘    │
│                                                         │
│  Ports (Interfaces):                                    │
│  - Repository (data access)                             │
│  - EventPublisher (messaging)                           │
│  - ContentClient (external service)                     │
└─────────────────────────────────────────────────────────┘
                        ↓
┌─────────────────────────────────────────────────────────┐
│                  SECONDARY ADAPTERS                     │
│            (Outbound - Driven by the app)               │
│                                                         │
│  ┌──────────────┬──────────────┬──────────────────┐     │
│  │  PostgreSQL  │   RabbitMQ   │  HTTP Client     │     │
│  │  Repository  │   Publisher  │  (Content Svc)   │     │
│  └──────────────┴──────────────┴──────────────────┘     │
└─────────────────────────────────────────────────────────┘
```

### Ports (Interfaces)

**1. Repository Port (Data Access)**

```go
// internal/scoring/repository/interface.go
type Repository interface {
    Create(ctx context.Context, submission *model.Submission) error
    FindAnsweredQuestionIDs(ctx context.Context, userID, skillTag string) ([]int64, error)
}
```

**2. EventPublisher Port (Messaging)**

```go
// internal/publisher/publisher.go
type EventPublisher interface {
    PublishSubmissionEvent(event interface{}) error
    Close() error
}
```

**3. ContentClient Port (External Service)**

```go
// pkg/curl/interfaces.go
type ContentServiceClient interface {
    GetQuestion(ctx context.Context, questionID int64) (*ContentResponse, error)
}
```

### Adapters (Implementations)

**Primary Adapter: HTTP Handler**

```go
// internal/scoring/delivery/http/handler.go
type handler struct {
    l  log.Logger
    uc scoring.UseCase  // Depends on port (interface)
}

func (h *handler) SubmitAnswer(c *gin.Context) {
    // 1. Parse HTTP request
    var req SubmitRequest
    c.ShouldBindJSON(&req)

    // 2. Convert to domain input
    input := req.ToSubmitInput()

    // 3. Call use case (core business logic)
    output, err := h.uc.SubmitAnswer(ctx, input)

    // 4. Convert to HTTP response
    response.OK(c, ToSubmitResponse(output))
}
```

**Secondary Adapter: PostgreSQL Repository**

```go
// internal/scoring/repository/postgre/submission.go
type submissionRepository struct {
    db *sql.DB
    l  log.Logger
}

func (r *submissionRepository) Create(ctx context.Context, submission *model.Submission) error {
    query := `INSERT INTO submissions (...) VALUES (...)`
    return r.db.QueryRowContext(ctx, query, ...).Scan(&submission.ID)
}
```

**Secondary Adapter: RabbitMQ Publisher**

```go
// internal/publisher/publisher.go
type rabbitmqPublisher struct {
    conn    *amqp.Connection
    channel *amqp.Channel
    l       log.Logger
}

func (p *rabbitmqPublisher) PublishSubmissionEvent(event interface{}) error {
    body, _ := json.Marshal(event)
    return p.channel.PublishWithContext(ctx, "its.events", "event.submission", ...)
}
```

### Dependency Injection

**main.go orchestrates dependency injection:**

```go
func main() {
    // Load config
    cfg, _ := config.Load()

    // Initialize logger
    log := pkglog.Init(...)

    // Initialize infrastructure (Secondary Adapters)
    db, _ := connectDB(cfg.Postgres)
    eventPublisher, _ := publisher.NewRabbitMQPublisher(cfg.RabbitMQ.URL, log)
    contentClient := curl.NewContentServiceClient(cfg.ContentServiceURL)

    // Initialize layers (Dependency Injection)
    submissionRepo := scoringrepo.New(db, log)
    scoringUC := scoringusecase.New(log, submissionRepo, eventPublisher, contentClient)
    scoringHandler := scoringhttp.New(log, scoringUC)

    // Setup router (Primary Adapter)
    router := gin.Default()
    scoringhttp.MapScoringRoutes(router.Group("/api/scoring"), scoringHandler)

    // Start server
    router.Run(":8082")
}
```

### Benefits of Hexagonal Architecture

**1. Testability**

Business logic có thể test độc lập với infrastructure:

```go
// Test use case với mocked dependencies
func TestSubmitAnswer(t *testing.T) {
    mockRepo := new(MockRepository)
    mockPublisher := new(MockEventPublisher)
    mockClient := new(MockContentClient)

    uc := usecase.New(logger, mockRepo, mockPublisher, mockClient)

    // Test business logic without real database/RabbitMQ/HTTP
    output, err := uc.SubmitAnswer(ctx, input)

    assert.NoError(t, err)
    assert.True(t, output.Correct)
}
```

**2. Flexibility**

Dễ dàng thay đổi infrastructure:

```go
// Thay PostgreSQL bằng MongoDB
mongoRepo := mongorepo.New(mongoClient, log)
scoringUC := scoringusecase.New(log, mongoRepo, eventPublisher, contentClient)

// Thay RabbitMQ bằng Kafka
kafkaPublisher := kafka.NewPublisher(kafkaConfig, log)
scoringUC := scoringusecase.New(log, submissionRepo, kafkaPublisher, contentClient)
```

**3. Independence**

Business logic không phụ thuộc vào:

- Web framework (Gin)
- Database (PostgreSQL)
- Message broker (RabbitMQ)
- External services

**4. Clear Boundaries**

Mỗi layer có trách nhiệm rõ ràng:

- **HTTP Handler**: HTTP concerns only
- **UseCase**: Business logic only
- **Repository**: Data access only
- **Publisher**: Messaging only

### Anti-Patterns to Avoid

**Business Logic in Handler:**

```go
// BAD: Business logic in HTTP handler
func (h *handler) SubmitAnswer(c *gin.Context) {
    // Fetching from Content Service (should be in use case)
    resp, _ := http.Get("http://content-service/questions/1")

    // Scoring logic (should be in use case)
    if answer == correctAnswer {
        score = 100
    }

    // Direct database access (should be in repository)
    db.Exec("INSERT INTO submissions ...")
}
```

**Correct: Thin Handler, Fat Use Case:**

```go
// GOOD: Handler only handles HTTP concerns
func (h *handler) SubmitAnswer(c *gin.Context) {
    var req SubmitRequest
    c.ShouldBindJSON(&req)

    input := req.ToSubmitInput()
    output, err := h.uc.SubmitAnswer(ctx, input)  // Delegate to use case

    response.OK(c, ToSubmitResponse(output))
}
```

---

## Service Dependencies

### External Services

Scoring Service phụ thuộc vào các services sau:

#### 1. Content Service (Required)

**Purpose**: Lấy thông tin câu hỏi và đáp án đúng

**Endpoint Used:**

```
GET http://content-service:8081/api/content/{question_id}
```

**Response:**

```json
{
  "error_code": 0,
  "message": "Success",
  "data": {
    "id": 1,
    "content": "What is 2 + 2?",
    "correctAnswer": "4",
    "skillTag": "math_arithmetic",
    "difficultyLevel": 1
  }
}
```

**Integration:**

- Sử dụng `pkg/curl/content_service.go` client
- Timeout: 10 seconds
- Retry: None (fail fast)

#### 2. PostgreSQL (Required)

**Purpose**: Lưu trữ submission records

**Connection:**

```
postgresql://postgres:postgres@localhost:5432/scoring_db?sslmode=disable
```

**Tables:**

- `submissions`: Lưu trữ lịch sử trả lời

#### 3. RabbitMQ (Required)

**Purpose**: Xuất bản sự kiện submission

**Connection:**

```
amqp://admintest:adminTest2025@localhost:5672/
```

**Exchange:**

- Name: `its.events`
- Type: `topic`
- Routing Key: `event.submission`

### Service Communication Pattern

```
┌─────────────────┐
│  Client (Web)   │
└────────┬────────┘
         │ HTTP POST /api/scoring/submit
         ↓
┌─────────────────┐
│ Scoring Service │
└────────┬────────┘
         │
         ├─→ [1] GET /api/content/{id}
         │   Content Service (Synchronous)
         │   ↓ Returns: correct_answer, skill_tag
         │
         ├─→ [2] INSERT INTO submissions
         │   PostgreSQL (Synchronous)
         │   ↓ Returns: submission_id
         │
         └─→ [3] PUBLISH event.submission
             RabbitMQ (Asynchronous)
             ↓
         ┌───────────────────┐
         │ Learner Model     │
         │ Consumer          │
         └───────────────────┘
```

---

## Troubleshooting

### Common Issues

#### 1. Cannot Connect to PostgreSQL

**Error:**

```
Failed to connect to database | error=dial tcp [::1]:5432: connect: connection refused
```

**Solutions:**

```bash
# Check if PostgreSQL is running
docker ps | grep postgres

# Start PostgreSQL if not running
docker-compose -f docker-compose.infra.yml up -d postgres

# Check connection
psql -h localhost -U postgres -d scoring_db

# Verify environment variables
echo $POSTGRES_HOST
echo $POSTGRES_PORT
```

#### 2. Cannot Connect to RabbitMQ

**Error:**

```
Failed to connect to RabbitMQ | error=dial tcp [::1]:5672: connect: connection refused
```

**Solutions:**

```bash
# Check if RabbitMQ is running
docker ps | grep rabbitmq

# Start RabbitMQ if not running
docker-compose -f docker-compose.infra.yml up -d rabbitmq

# Check RabbitMQ logs
docker logs rabbitmq

# Access management UI
open http://localhost:15672
```

#### 3. Content Service Not Available

**Error:**

```
failed to fetch question from content service | error=Get "http://localhost:8081/api/content/1": dial tcp [::1]:8081: connect: connection refused
```

**Solutions:**

```bash
# Check if Content Service is running
curl http://localhost:8081/health

# Start Content Service
cd sources/content
mvn spring-boot:run

# Or use Docker
docker-compose up content-service
```

#### 4. Port Already in Use

**Error:**

```
Failed to start server | error=listen tcp :8082: bind: address already in use
```

**Solutions:**

```bash
# Find process using port 8082
lsof -i :8082

# Kill the process
kill -9 <PID>

# Or change port in .env
APP_PORT=8083
```

#### 5. Database Migration Issues

**Error:**

```
ERROR: relation "submissions" does not exist
```

**Solutions:**

```bash
# Run migrations manually
psql -h localhost -U postgres -d scoring_db -f migration/01_add_project.sql

# Or recreate database
dropdb scoring_db
createdb scoring_db
psql -h localhost -U postgres -d scoring_db -f migration/01_add_project.sql
```

### Debugging Tips

**1. Enable Debug Logging:**

```bash
# Set log level to debug
export LOG_LEVEL=debug

# Run service
go run cmd/api/main.go
```

**2. Check Service Logs:**

```bash
# Docker logs
docker logs its-scoring-service -f

# Local logs (if configured)
tail -f scoring.log
```

**3. Test RabbitMQ Connection:**

```bash
# Test publish manually
curl -X POST http://localhost:15672/api/exchanges/%2F/its.events/publish \
  -u admintest:adminTest2025 \
  -H "Content-Type: application/json" \
  -d '{
    "properties":{},
    "routing_key":"event.submission",
    "payload":"{\"event\":\"test\"}",
    "payload_encoding":"string"
  }'
```

**4. Verify Database State:**

```bash
# Connect to database
psql -h localhost -U postgres -d scoring_db

# Check submissions
SELECT * FROM submissions ORDER BY created_at DESC LIMIT 10;

# Count submissions by user
SELECT user_id, COUNT(*) FROM submissions GROUP BY user_id;
```

**5. Test API Endpoints:**

```bash
# Health check
curl http://localhost:8082/health

# Submit answer (should fail if Content Service is down)
curl -X POST http://localhost:8082/api/scoring/submit \
  -H "Content-Type: application/json" \
  -d '{"user_id":"test","question_id":1,"answer":"A"}'

# Get answered questions
curl "http://localhost:8082/api/scoring/answered-questions?user_id=test&skill=math"
```

---

## Performance Considerations

### Async Event Publishing

Event publishing được thực hiện asynchronously để:

- Không block HTTP response
- Tăng throughput
- Giảm latency cho user

### Database Indexes

Indexes được tạo cho:

- `user_id`: Truy vấn submissions theo user
- `question_id`: Truy vấn submissions theo question
- `created_at`: Sắp xếp theo thời gian

### Connection Pooling

Go's `database/sql` tự động quản lý connection pool:

- Default max open connections: unlimited
- Default max idle connections: 2
- Connection lifetime: unlimited

**Tuning (nếu cần):**

```go
db.SetMaxOpenConns(25)
db.SetMaxIdleConns(5)
db.SetConnMaxLifetime(5 * time.Minute)
```

---

## Tài liệu tham khảo

- [Architecture Documentation](./document/architecture.md) - Chi tiết về Hexagonal Architecture
- [API Documentation](./document/api.md) - Sequence diagrams và API flows
- [Swagger API Docs](http://localhost:8082/scoring/swagger/index.html) - Interactive API documentation
- [Hexagonal Architecture](https://alistair.cockburn.us/hexagonal-architecture/) - Original article by Alistair Cockburn
- [Go Documentation](https://go.dev/doc/)
- [Gin Framework](https://gin-gonic.com/docs/)
- [RabbitMQ Tutorials](https://www.rabbitmq.com/getstarted.html)

---

## Liên hệ

Để biết thêm thông tin về Scoring Service hoặc hệ thống ITS, vui lòng tham khảo:

- **Root README**: `../../README.md`
- **Sources README**: `../README.md`
- **Architecture Report**: `../../report/main.pdf`
- **Content Service**: `../content/README.md`
- **Learner Model Service**: `../learner-model/README.md`

---

**Scoring Service** - Phần của Intelligent Tutoring System (ITS)  
CO3017 - Kỹ Thuật Phần Mềm - HCMUT
