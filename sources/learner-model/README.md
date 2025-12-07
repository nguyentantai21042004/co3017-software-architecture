# Learner Model Service - Dịch vụ Mô hình Người học

> Microservice theo dõi mức độ thành thạo kỹ năng (Skill Mastery) của học sinh trong Hệ thống Gia sư Thông minh (ITS)

---

## Mục lục

- [Tổng quan](#tổng-quan)
- [Công nghệ](#công-nghệ)
- [Kiến trúc](#kiến-trúc)
- [Database Schema](#database-schema)
- [API Endpoints](#api-endpoints)
- [Event Consumption](#event-consumption)
- [Thuật toán BKT](#thuật-toán-bkt)
- [Cấu hình](#cấu-hình)
- [Phát triển Cục bộ](#phát-triển-cục-bộ)
- [Testing](#testing)
- [Tích hợp MinIO](#tích-hợp-minio)
- [Cấu trúc Thư mục](#cấu-trúc-thư-mục)

---

## Tổng quan

**Learner Model Service** là microservice chịu trách nhiệm theo dõi và quản lý mức độ thành thạo kỹ năng của học sinh. Service này bao gồm hai thành phần chính:

### 1. API Service (Port 8083)

REST API cung cấp các endpoint để:

- Truy vấn mức độ thành thạo (Mastery Level) của học sinh cho từng kỹ năng
- Kiểm tra sức khỏe service (Health Check)
- Cung cấp tài liệu API qua Swagger

### 2. Consumer Service (Background Worker)

Background worker lắng nghe và xử lý events từ RabbitMQ:

- Nhận events từ Scoring Service khi học sinh hoàn thành bài kiểm tra
- Cập nhật mức độ thành thạo dựa trên điểm số đạt được
- Sử dụng thuật toán Moving Average để tính toán mastery score mới

### Vai trò trong Hệ thống ITS

```mermaid
flowchart LR
    subgraph Frontend
        Client[Client App]
    end

    subgraph Backend
        Scoring[Scoring Service]
        Learner[Learner Model Service]
        Adaptive[Adaptive Engine]
    end

    subgraph Infrastructure
        RabbitMQ[(RabbitMQ)]
        PostgreSQL[(PostgreSQL)]
    end

    Client -->|Truy vấn mastery| Adaptive
    Adaptive -->|Lấy mastery| Learner
    Scoring -->|Publish event| RabbitMQ
    RabbitMQ -->|Consume event| Learner
    Learner -->|Lưu/Đọc| PostgreSQL
```

**Luồng dữ liệu:**

1. Học sinh làm bài kiểm tra → Scoring Service chấm điểm
2. Scoring Service publish event `submission.scored` đến RabbitMQ
3. Learner Model Consumer nhận event và cập nhật mastery score
4. Adaptive Engine truy vấn mastery để đề xuất nội dung phù hợp

---

## Công nghệ

| Công nghệ      | Phiên bản | Mục đích                             |
| -------------- | --------- | ------------------------------------ |
| **Go**         | 1.23+     | Ngôn ngữ lập trình chính             |
| **Gin**        | Latest    | HTTP web framework                   |
| **PostgreSQL** | 15        | Cơ sở dữ liệu chính                  |
| **RabbitMQ**   | 3.x       | Message broker cho event consumption |
| **MinIO**      | Latest    | Object storage (tùy chọn)            |
| **SQLBoiler**  | v4.19.5   | ORM type-safe                        |
| **Swagger**    | v1.16.6   | API documentation                    |
| **Zap**        | Latest    | Structured logging                   |
| **Docker**     | >= 20.10  | Containerization                     |

---

## Kiến trúc

### Kiến trúc Tổng quan (API + Consumer)

Learner Model Service sử dụng **Clean Architecture** với pattern **Module-First**, tách biệt rõ ràng giữa các layers:

```
┌────────────────────────────────────────────────────────────────────┐
│                    Learner Model Service                           │
├────────────────────────────────────────────────────────────────────┤
│                                                                    │
│  ┌─────────────────────────┐    ┌─────────────────────────────┐    │
│  │     API Service         │    │     Consumer Service        │    │
│  │     (Port 8083)         │    │     (Background Worker)     │    │
│  │                         │    │                             │    │
│  │  ┌─────────────────┐    │    │  ┌─────────────────────┐    │    │
│  │  │ HTTP Handlers   │    │    │  │ RabbitMQ Consumer   │    │    │
│  │  │ (Gin Framework) │    │    │  │ (Event Listener)    │    │    │
│  │  └────────┬────────┘    │    │  └──────────┬──────────┘    │    │
│  │           │             │    │             │               │    │
│  └───────────┼─────────────┘    └─────────────┼───────────────┘    │
│              │                                │                    │
│              └────────────────┬───────────────┘                    │
│                               │                                    │
│                    ┌──────────▼──────────┐                         │
│                    │    UseCase Layer    │                         │
│                    │  (Business Logic)   │                         │
│                    │  - GetMastery       │                         │
│                    │  - UpdateMastery    │                         │
│                    └──────────┬──────────┘                         │
│                               │                                    │
│                    ┌──────────▼──────────┐                         │
│                    │  Repository Layer   │                         │
│                    │  (Data Access)      │                         │
│                    └──────────┬──────────┘                         │
│                               │                                    │
└───────────────────────────────┼────────────────────────────────────┘
                                │
                     ┌──────────▼──────────┐
                     │     PostgreSQL      │
                     │    (learner_db)     │
                     └─────────────────────┘
```

### Clean Architecture Layers

```

┌──────────────────────────────────────────────────────────────────┐
│ 1. Delivery Layer                                                │
│ - HTTP Handlers (API endpoints)                                  │
│ - RabbitMQ Consumer (Event processing)                           │
│ - Request/Response DTOs                                          │
│ - Input validation                                               │
└─────────────────────────────┬────────────────────────────────────┘
                              │ Calls
                              ▼
┌──────────────────────────────────────────────────────────────────┐
│ 2. UseCase Layer                                                 │
│ - GetMastery: Truy vấn mastery score                             │
│ - UpdateMasteryFromEvent: Cập nhật từ submission event           │
│ - Business logic và validation                                   │
│ - Tính toán mastery score mới                                    │
└─────────────────────────────┬────────────────────────────────────┘
                              │ Calls
                              ▼
┌──────────────────────────────────────────────────────────────────┐
│ 3. Repository Layer                                              │
│ - GetByUserAndSkill: Lấy mastery theo user và skill              │
│ - CreateOrUpdate: Tạo mới hoặc cập nhật mastery                  │
│ - PostgreSQL implementation                                      │
└──────────────────────────────────────────────────────────────────┘

```

### Nguyên tắc Thiết kế

1. **Dependency Inversion**: UseCase không phụ thuộc vào Repository implementation
2. **Interface Segregation**: Mỗi layer chỉ biết interface của layer dưới
3. **Single Responsibility**: Mỗi component có một trách nhiệm duy nhất
4. **Stateless API**: API service không lưu state giữa các requests

---

## Database Schema

### Bảng `skill_mastery`

Lưu trữ mức độ thành thạo của học sinh cho từng kỹ năng:

```sql
CREATE TABLE IF NOT EXISTS skill_mastery (
    user_id VARCHAR(255) NOT NULL,
    skill_tag VARCHAR(255) NOT NULL,
    current_score INTEGER NOT NULL DEFAULT 0
        CHECK (current_score >= 0 AND current_score <= 100),
    last_updated TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (user_id, skill_tag)
);

-- Indexes cho hiệu năng truy vấn
CREATE INDEX IF NOT EXISTS idx_skill_mastery_user_id ON skill_mastery(user_id);
CREATE INDEX IF NOT EXISTS idx_skill_mastery_skill_tag ON skill_mastery(skill_tag);
CREATE INDEX IF NOT EXISTS idx_skill_mastery_last_updated ON skill_mastery(last_updated);
```

### ERD (Entity Relationship Diagram)

```mermaid
erDiagram
    SKILL_MASTERY {
        varchar user_id PK "ID người dùng"
        varchar skill_tag PK "Tag kỹ năng"
        integer current_score "Điểm mastery (0-100)"
        timestamptz last_updated "Thời gian cập nhật"
    }
```

### Mô tả Các Cột

| Cột             | Kiểu         | Mô tả                               |
| --------------- | ------------ | ----------------------------------- |
| `user_id`       | VARCHAR(255) | ID của học sinh (Primary Key)       |
| `skill_tag`     | VARCHAR(255) | Tag định danh kỹ năng (Primary Key) |
| `current_score` | INTEGER      | Điểm thành thạo từ 0-100            |
| `last_updated`  | TIMESTAMPTZ  | Thời điểm cập nhật gần nhất         |

**Lưu ý:**

- Composite Primary Key: `(user_id, skill_tag)` - mỗi học sinh có một record cho mỗi kỹ năng
- `current_score` có constraint `CHECK (current_score >= 0 AND current_score <= 100)`

---

## API Endpoints

### Base URL

```
http://localhost:8083/internal/learner
```

### Swagger Documentation

```
http://localhost:8083/learner-model/swagger/index.html
```

### 1. Health Check

Kiểm tra trạng thái hoạt động của service.

**Request:**

```bash
curl http://localhost:8083/health
```

**Response (200 OK):**

```json
{
  "error_code": 0,
  "message": "Healthy",
  "data": {
    "status": "healthy",
    "service": "learner-model"
  }
}
```

### 2. Get Mastery - Lấy Mức độ Thành thạo

Truy vấn điểm mastery của học sinh cho một kỹ năng cụ thể.

**Endpoint:**

```
GET /internal/learner/:user_id/mastery?skill=:skill_tag
```

**Parameters:**

| Tham số   | Vị trí | Bắt buộc | Mô tả                        |
| --------- | ------ | -------- | ---------------------------- |
| `user_id` | Path   | Có       | ID của học sinh              |
| `skill`   | Query  | Có       | Tag của kỹ năng cần truy vấn |

**Request:**

```bash
# Lấy mastery của user_01 cho kỹ năng math_algebra
curl "http://localhost:8083/internal/learner/user_01/mastery?skill=math_algebra"
```

**Response (200 OK):**

```json
{
  "error_code": 0,
  "message": "Success",
  "data": {
    "user_id": "user_01",
    "skill_tag": "math_algebra",
    "mastery_score": 75,
    "last_updated": "2024-12-07T10:30:00Z"
  }
}
```

**Response khi chưa có dữ liệu (200 OK):**

```json
{
  "error_code": 0,
  "message": "Success",
  "data": {
    "user_id": "user_01",
    "skill_tag": "new_skill",
    "mastery_score": 0,
    "last_updated": "2024-12-07T10:30:00Z"
  }
}
```

**Error Responses:**

| Status Code | Mô tả                                  |
| ----------- | -------------------------------------- |
| 400         | Thiếu `user_id` hoặc `skill` parameter |
| 500         | Lỗi server khi truy vấn database       |

```json
// 400 Bad Request - Thiếu skill parameter
{
  "error_code": 400,
  "message": "skill tag is required"
}
```

### Ví dụ Sử dụng với Các Kỹ năng Khác nhau

```bash
# Kỹ năng DevOps
curl "http://localhost:8083/internal/learner/user_01/mastery?skill=devops_basics"

# Kỹ năng Kubernetes
curl "http://localhost:8083/internal/learner/user_01/mastery?skill=kubernetes_fundamentals"

# Kỹ năng Python
curl "http://localhost:8083/internal/learner/user_01/mastery?skill=python_loops"
```

---

## Event Consumption

### RabbitMQ Configuration

Consumer lắng nghe events từ Scoring Service thông qua RabbitMQ:

| Cấu hình          | Giá trị            |
| ----------------- | ------------------ |
| **Exchange**      | `its.events`       |
| **Exchange Type** | `topic`            |
| **Queue**         | `learner.updates`  |
| **Routing Key**   | `event.submission` |

### Event Format

**Submission Event** - Được publish bởi Scoring Service:

```json
{
  "event": "submission.scored",
  "user_id": "user_01",
  "skill_tag": "math_algebra",
  "score_obtained": 80,
  "timestamp": "2024-12-07T10:30:00Z"
}
```

### Event Processing Flow

```mermaid
sequenceDiagram
    participant Scoring as Scoring Service
    participant RabbitMQ as RabbitMQ
    participant Consumer as Learner Consumer
    participant UseCase as UseCase Layer
    participant DB as PostgreSQL

    Scoring->>RabbitMQ: Publish event.submission
    RabbitMQ->>Consumer: Deliver message
    Consumer->>Consumer: Parse JSON event
    Consumer->>UseCase: UpdateMasteryFromEvent(input)
    UseCase->>DB: GetByUserAndSkill(user_id, skill_tag)
    DB-->>UseCase: Current mastery (or null)
    UseCase->>UseCase: Calculate new score
    UseCase->>DB: CreateOrUpdate(mastery)
    DB-->>UseCase: Success
    UseCase-->>Consumer: Success
    Consumer->>Consumer: Log success
```

### Consumer Code Structure

```go
// Event được parse từ RabbitMQ message
type SubmissionEvent struct {
    Event         string `json:"event"`
    UserID        string `json:"user_id"`
    SkillTag      string `json:"skill_tag"`
    ScoreObtained int    `json:"score_obtained"`
    Timestamp     string `json:"timestamp"`
}

// Consumer xử lý message
func (c *rabbitmqConsumer) handleMessage(msg amqp.Delivery) {
    // 1. Parse event từ JSON
    var event model.SubmissionEvent
    json.Unmarshal(msg.Body, &event)

    // 2. Chuyển đổi thành input cho UseCase
    input := learner.UpdateMasteryInput{
        UserID:        event.UserID,
        SkillTag:      event.SkillTag,
        ScoreObtained: event.ScoreObtained,
    }

    // 3. Gọi UseCase để cập nhật mastery
    c.uc.UpdateMasteryFromEvent(ctx, input)
}
```

---

## Thuật toán BKT

### Bayesian Knowledge Tracing (BKT) - Đơn giản hóa

Service sử dụng phiên bản đơn giản hóa của thuật toán **Bayesian Knowledge Tracing** để tính toán mức độ thành thạo. Thay vì sử dụng full BKT với các tham số phức tạp, chúng tôi áp dụng **Moving Average** để cập nhật mastery score:

### Công thức Tính toán

```
NewScore = (OldScore + ScoreObtained) / 2
```

**Trong đó:**

- `OldScore`: Điểm mastery hiện tại (0 nếu chưa có)
- `ScoreObtained`: Điểm đạt được từ bài kiểm tra mới (0-100)
- `NewScore`: Điểm mastery mới (được giới hạn trong khoảng 0-100)

### Ví dụ Tính toán

| Lần | OldScore | ScoreObtained | NewScore | Giải thích                      |
| --- | -------- | ------------- | -------- | ------------------------------- |
| 1   | 0        | 60            | 30       | Lần đầu làm bài, đạt 60 điểm    |
| 2   | 30       | 80            | 55       | Làm tốt hơn, mastery tăng       |
| 3   | 55       | 90            | 72       | Tiếp tục cải thiện              |
| 4   | 72       | 100           | 86       | Gần đạt mastery cao             |
| 5   | 86       | 70            | 78       | Điểm thấp hơn, mastery giảm nhẹ |

### Implementation Code

```go
// Tính toán mastery score mới
func (uc *usecase) UpdateMasteryFromEvent(ctx context.Context, input learner.UpdateMasteryInput) error {
    // Lấy mastery hiện tại
    currentMastery, err := uc.repo.GetByUserAndSkill(ctx, input.UserID, input.SkillTag)

    oldScore := 0
    if currentMastery != nil {
        oldScore = currentMastery.CurrentScore
    }

    // Công thức Moving Average
    newScore := (oldScore + input.ScoreObtained) / 2

    // Đảm bảo score trong khoảng [0, 100]
    if newScore < 0 {
        newScore = 0
    }
    if newScore > 100 {
        newScore = 100
    }

    // Lưu mastery mới
    mastery := &model.SkillMastery{
        UserID:       input.UserID,
        SkillTag:     input.SkillTag,
        CurrentScore: newScore,
    }

    return uc.repo.CreateOrUpdate(ctx, mastery)
}
```

### Ưu điểm của Phương pháp

1. **Đơn giản**: Dễ hiểu và implement
2. **Ổn định**: Không thay đổi đột ngột khi có một bài kiểm tra tốt/xấu
3. **Phản ánh xu hướng**: Mastery tăng dần khi học sinh cải thiện
4. **Có trọng số**: Kết quả gần đây ảnh hưởng 50% đến score mới

### Ngưỡng Mastery (Mastery Threshold)

```go
const MASTERY_THRESHOLD = 50
```

Học sinh được coi là "đã thành thạo" một kỹ năng khi `current_score >= 50`.

---

## Cấu hình

### Biến Môi trường

| Biến                 | Mô tả                    | Giá trị Mặc định      |
| -------------------- | ------------------------ | --------------------- |
| **Server**           |                          |                       |
| `HOST`               | Host address             | `""` (all interfaces) |
| `APP_PORT`           | Port của API service     | `8083`                |
| `API_MODE`           | Gin mode (debug/release) | `debug`               |
| **Database**         |                          |                       |
| `POSTGRES_HOST`      | PostgreSQL host          | `localhost`           |
| `POSTGRES_PORT`      | PostgreSQL port          | `5432`                |
| `POSTGRES_USER`      | Database user            | `postgres`            |
| `POSTGRES_PASSWORD`  | Database password        | `postgres`            |
| `POSTGRES_DB`        | Database name            | `learner_db`          |
| `POSTGRES_SSLMODE`   | SSL mode                 | `disable`             |
| **RabbitMQ**         |                          |                       |
| `RABBITMQ_URL`       | RabbitMQ connection URL  | -                     |
| **Logging**          |                          |                       |
| `LOGGER_LEVEL`       | Log level                | `debug`               |
| `LOGGER_MODE`        | Logger mode              | `debug`               |
| `LOGGER_ENCODING`    | Log encoding             | `console`             |
| **MinIO (Optional)** |                          |                       |
| `MINIO_ENDPOINT`     | MinIO endpoint           | `localhost:9000`      |
| `MINIO_ACCESS_KEY`   | Access key               | `minioadmin`          |
| `MINIO_SECRET_KEY`   | Secret key               | `minioadmin`          |
| `MINIO_USE_SSL`      | Use SSL                  | `false`               |
| `MINIO_BUCKET`       | Bucket name              | -                     |

### Cấu hình Docker Compose

Trong `docker-compose.yml` (từ thư mục `sources/`):

```yaml
# API Service
learner-model-api:
  build:
    context: ./learner-model
    dockerfile: cmd/api/Dockerfile
  environment:
    APP_HOST: 0.0.0.0
    APP_PORT: 8083
    POSTGRES_HOST: postgres
    POSTGRES_PORT: 5432
    POSTGRES_USER: postgres
    POSTGRES_PASSWORD: postgres
    POSTGRES_DB: learner_db
    POSTGRES_SSLMODE: disable
    LOG_LEVEL: debug
  ports:
    - "8083:8083"

# Consumer Service
learner-model-consumer:
  build:
    context: ./learner-model
    dockerfile: cmd/consumer/Dockerfile
  environment:
    POSTGRES_HOST: postgres
    POSTGRES_PORT: 5432
    POSTGRES_USER: postgres
    POSTGRES_PASSWORD: postgres
    POSTGRES_DB: learner_db
    POSTGRES_SSLMODE: disable
    RABBITMQ_URL: amqp://admintest:adminTest2025@rabbitmq:5672/
    LOG_LEVEL: debug
```

### File `.env` Mẫu

Tạo file `.env` từ `template.env`:

```bash
cp template.env .env
```

Nội dung `.env` cho development:

```env
# Server
HOST=
APP_PORT=8083
API_MODE=debug

# Logger
LOGGER_LEVEL=debug
LOGGER_MODE=debug
LOGGER_ENCODING=console

# PostgreSQL
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=learner_db
POSTGRES_SSLMODE=disable

# RabbitMQ
RABBITMQ_URL=amqp://admintest:adminTest2025@localhost:5672/

# MinIO (Optional)
MINIO_ENDPOINT=localhost:9000
MINIO_ACCESS_KEY=minioadmin
MINIO_SECRET_KEY=minioadmin
MINIO_USE_SSL=false
```

---

## Phát triển Cục bộ

### Yêu cầu Hệ thống

- **Go**: >= 1.23
- **PostgreSQL**: >= 15
- **RabbitMQ**: >= 3.x (cho Consumer)
- **Make**: Để chạy Makefile commands

### Bước 1: Khởi động Infrastructure

Từ thư mục `sources/`:

```bash
# Khởi động PostgreSQL, RabbitMQ, MinIO
docker-compose -f docker-compose.infra.yml up -d
```

Hoặc sử dụng Makefile:

```bash
make dev
```

### Bước 2: Cấu hình Database

Chạy migration để tạo bảng:

```bash
# Kết nối vào PostgreSQL
psql -h localhost -U postgres -d learner_db

# Chạy migration script
\i migration/01_create_skill_mastery_table.sql
```

### Bước 3: Chạy API Service

```bash
# Từ thư mục learner-model/
cd sources/learner-model

# Cài đặt dependencies
go mod download

# Chạy API service
make run-api
# Hoặc
go run cmd/api/main.go
```

API sẽ chạy tại: `http://localhost:8083`

### Bước 4: Chạy Consumer Service

Mở terminal mới:

```bash
# Từ thư mục learner-model/
cd sources/learner-model

# Chạy Consumer
make run-consumer
# Hoặc
go run cmd/consumer/main.go
```

Consumer sẽ bắt đầu lắng nghe events từ RabbitMQ.

### Kiểm tra Service

```bash
# Health check
curl http://localhost:8083/health

# Test API endpoint
curl "http://localhost:8083/internal/learner/test_user/mastery?skill=test_skill"
```

### Build Binary

```bash
# Build API
make build-api
# Output: bin/api

# Build Consumer
make build-consumer
# Output: bin/consumer

# Build cả hai
make build-all
```

---

## Testing

### Chạy Unit Tests

```bash
# Chạy tất cả tests
make test

# Chạy tests với verbose output
go test -v -race ./...

# Chạy short tests (bỏ qua integration tests)
make test-short
```

### Test Coverage

```bash
# Tạo coverage report
make test-coverage

# Xem coverage trong browser
go tool cover -html=coverage.out
```

### Test Files

| File                                             | Mô tả                            |
| ------------------------------------------------ | -------------------------------- |
| `internal/learner/usecase/learner_test.go`       | Unit tests cho UseCase layer     |
| `internal/consumer/rabbitmq_consumer_test.go`    | Unit tests cho RabbitMQ Consumer |
| `internal/consumer/consumer_integration_test.go` | Integration tests cho Consumer   |

### Ví dụ Test Case

```go
// Test GetMastery khi chưa có dữ liệu
func TestGetMastery_NotFound_ReturnsDefault(t *testing.T) {
    // Setup mock repository
    mockRepo := new(mocks.Repository)
    mockRepo.On("GetByUserAndSkill", mock.Anything, "user_01", "skill_01").
        Return(nil, repository.ErrNotFound)

    uc := usecase.New(logger, mockRepo)

    // Execute
    output, err := uc.GetMastery(ctx, learner.GetMasteryInput{
        UserID:   "user_01",
        SkillTag: "skill_01",
    })

    // Assert
    assert.NoError(t, err)
    assert.Equal(t, 0, output.MasteryScore)
}
```

---

## Tích hợp MinIO

### Mục đích

MinIO được cấu hình trong service để lưu trữ các file bổ sung như:

- Learning analytics data
- Exported reports
- Backup data

### Cấu hình MinIO

```go
type MinIOConfig struct {
    Endpoint  string `env:"MINIO_ENDPOINT" envDefault:"localhost:9000"`
    AccessKey string `env:"MINIO_ACCESS_KEY" envDefault:"minioadmin"`
    SecretKey string `env:"MINIO_SECRET_KEY" envDefault:"minioadmin"`
    UseSSL    bool   `env:"MINIO_USE_SSL" envDefault:"false"`
    Region    string `env:"MINIO_REGION" envDefault:"us-east-1"`
    Bucket    string `env:"MINIO_BUCKET"`
}
```

### Truy cập MinIO Console

Khi chạy với Docker Compose:

```
URL: http://localhost:9001
Username: minioadmin
Password: minioadmin
```

**Lưu ý:** MinIO là tùy chọn và không bắt buộc cho chức năng core của Learner Model Service.

---

## Cấu trúc Thư mục

```
learner-model/
├── cmd/                          # Entry points
│   ├── api/                      # API service
│   │   ├── main.go              # API main entry
│   │   └── Dockerfile           # Docker build cho API
│   └── consumer/                 # Consumer service
│       ├── main.go              # Consumer main entry
│       └── Dockerfile           # Docker build cho Consumer
│
├── config/                       # Configuration
│   ├── config.go                # Config loader
│   ├── minio/                   # MinIO config
│   └── postgre/                 # PostgreSQL config
│
├── docs/                         # Swagger documentation
│   ├── docs.go                  # Generated docs
│   ├── swagger.json             # OpenAPI spec (JSON)
│   └── swagger.yaml             # OpenAPI spec (YAML)
│
├── document/                     # Technical documentation
│   ├── api.md                   # API documentation
│   ├── architecture.md          # Architecture details
│   └── overview.md              # Service overview
│
├── internal/                     # Internal packages
│   ├── consumer/                # RabbitMQ consumer
│   │   ├── rabbitmq_consumer.go # Consumer implementation
│   │   └── *_test.go            # Consumer tests
│   │
│   ├── learner/                 # Learner module (Module-First)
│   │   ├── delivery/http/       # HTTP handlers
│   │   │   ├── handler.go       # API handlers
│   │   │   ├── routes.go        # Route definitions
│   │   │   ├── presenter.go     # Response formatting
│   │   │   └── new.go           # Handler constructor
│   │   │
│   │   ├── repository/          # Data access layer
│   │   │   ├── interface.go     # Repository interface
│   │   │   ├── errors.go        # Repository errors
│   │   │   └── postgre/         # PostgreSQL implementation
│   │   │
│   │   ├── usecase/             # Business logic
│   │   │   ├── new.go           # UseCase constructor
│   │   │   ├── learner.go       # Business logic
│   │   │   └── learner_test.go  # Unit tests
│   │   │
│   │   ├── interface.go         # UseCase interface
│   │   ├── type.go              # Input/Output types
│   │   └── error.go             # Module errors
│   │
│   ├── model/                   # Domain models
│   │   └── skill_mastery.go     # SkillMastery model
│   │
│   └── sqlboiler/               # Generated DB models
│
├── migration/                    # Database migrations
│   └── 01_create_skill_mastery_table.sql
│
├── pkg/                          # Shared packages
│   ├── errors/                  # Error types
│   ├── log/                     # Logging
│   ├── minio/                   # MinIO client
│   ├── postgre/                 # PostgreSQL utilities
│   ├── rabbitmq/                # RabbitMQ client
│   └── response/                # HTTP response helpers
│
├── scripts/                      # Build scripts
│   ├── build.sh                 # Build API script
│   └── build-consumer.sh        # Build Consumer script
│
├── .env                         # Environment variables (local)
├── template.env                 # Environment template
├── go.mod                       # Go modules
├── go.sum                       # Go dependencies checksum
├── Makefile                     # Build automation
├── sqlboiler.toml               # SQLBoiler config
└── README.md                    # This file
```

---

## Makefile Commands

| Command               | Mô tả                           |
| --------------------- | ------------------------------- |
| `make run-api`        | Chạy API service locally        |
| `make run-consumer`   | Chạy Consumer service locally   |
| `make build-api`      | Build API binary                |
| `make build-consumer` | Build Consumer binary           |
| `make build-all`      | Build cả API và Consumer        |
| `make test`           | Chạy tất cả tests               |
| `make test-coverage`  | Chạy tests với coverage report  |
| `make test-short`     | Chạy short tests                |
| `make swagger`        | Generate Swagger documentation  |
| `make models`         | Generate SQLBoiler models       |
| `make docker-build`   | Build Docker image cho API      |
| `make consumer-build` | Build Docker image cho Consumer |
| `make help`           | Hiển thị tất cả commands        |

---

## Troubleshooting

### Lỗi Kết nối Database

```
Failed to connect to database
```

**Giải pháp:**

1. Kiểm tra PostgreSQL đang chạy: `docker ps | grep postgres`
2. Kiểm tra biến môi trường `POSTGRES_*`
3. Kiểm tra database `learner_db` đã được tạo

### Lỗi Kết nối RabbitMQ

```
Failed to connect to RabbitMQ
```

**Giải pháp:**

1. Kiểm tra RabbitMQ đang chạy: `docker ps | grep rabbitmq`
2. Kiểm tra `RABBITMQ_URL` đúng format
3. Truy cập RabbitMQ Management: `http://localhost:15672`

### Consumer Không Nhận Events

**Giải pháp:**

1. Kiểm tra exchange `its.events` đã được tạo
2. Kiểm tra queue `learner.updates` đã bind với routing key `event.submission`
3. Kiểm tra Scoring Service đang publish events

### API Trả về 500 Error

**Giải pháp:**

1. Kiểm tra logs: `docker logs its-learner-model-api`
2. Kiểm tra database connection
3. Kiểm tra migration đã chạy

---

## Liên kết Liên quan

### Tài liệu Dự án

| Tài liệu           | Đường dẫn                                        | Mô tả                                |
| ------------------ | ------------------------------------------------ | ------------------------------------ |
| **Root README**    | [../../README.md](../../README.md)               | Tổng quan dự án, cấu trúc repository |
| **Sources README** | [../README.md](../README.md)                     | Hướng dẫn microservices và Docker    |
| **Report README**  | [../../report/README.md](../../report/README.md) | Hướng dẫn build báo cáo LaTeX        |
| **Báo cáo PDF**    | [../../report/main.pdf](../../report/main.pdf)   | Báo cáo kiến trúc phần mềm           |

### Service READMEs Khác

| Service             | Đường dẫn                                                    | Mô tả                     |
| ------------------- | ------------------------------------------------------------ | ------------------------- |
| **Content Service** | [../content/README.md](../content/README.md)                 | Quản lý nội dung học tập  |
| **Scoring Service** | [../scoring/README.md](../scoring/README.md)                 | Chấm điểm, publish events |
| **Adaptive Engine** | [../adaptive-engine/README.md](../adaptive-engine/README.md) | Consume mastery data      |
| **Client**          | [../client/README.md](../client/README.md)                   | Frontend UI/UX            |

### Tài liệu Kiến trúc

| Tài liệu                   | Đường dẫn                                                                                              | Nội dung              |
| -------------------------- | ------------------------------------------------------------------------------------------------------ | --------------------- |
| **SOLID Principles**       | [../../markdown/report/6-SOLID-principles.md](../../markdown/report/6-SOLID-principles.md)             | Ví dụ áp dụng SOLID   |
| **Architecture Decisions** | [../../markdown/report/5-architecture-decisions.md](../../markdown/report/5-architecture-decisions.md) | ADRs                  |
| **Microservices Analysis** | [../../markdown/microservices.md](../../markdown/microservices.md)                                     | Chi tiết domain model |

### API Documentation

- **Swagger UI**: [http://localhost:8083/learner-model/swagger/index.html](http://localhost:8083/learner-model/swagger/index.html)
- **OpenAPI Spec**: [docs/swagger.json](./docs/swagger.json)

### Internal Documentation

- **Architecture**: [document/architecture.md](./document/architecture.md) - Chi tiết Clean Architecture
- **API Flows**: [document/api.md](./document/api.md) - API documentation
- **Overview**: [document/overview.md](./document/overview.md) - Service overview

### Configuration Files

- **Docker Compose**: [../docker-compose.yml](../docker-compose.yml) - Cấu hình deployment
- **Infrastructure**: [../docker-compose.infra.yml](../docker-compose.infra.yml) - PostgreSQL, RabbitMQ, MinIO

---

**Cập nhật lần cuối:** 2024-12-07  
**Phiên bản:** 1.0.0
