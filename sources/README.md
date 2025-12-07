# Hệ thống Gia sư Thông minh - Kiến trúc Microservices

[![Docker](https://img.shields.io/badge/Docker-Ready-blue)](https://www.docker.com/)
[![Microservices](https://img.shields.io/badge/Architecture-Microservices-green)](https://microservices.io/)
[![Java](https://img.shields.io/badge/Java-17-orange)](https://openjdk.org/)
[![Go](https://img.shields.io/badge/Go-1.23-00ADD8)](https://golang.org/)
[![Next.js](https://img.shields.io/badge/Next.js-15-black)](https://nextjs.org/)

## Mục lục

- [Tổng quan Hệ thống](#tổng-quan-hệ-thống)
- [Kiến trúc Microservices](#kiến-trúc-microservices)
- [Technology Stack](#technology-stack)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Running Services](#running-services)
- [Health Checks](#health-checks)
- [Testing](#testing)
- [Troubleshooting](#troubleshooting)
- [Development Workflow](#development-workflow)
- [API Gateway](#api-gateway)

## Tổng quan Hệ thống

Hệ thống Gia sư Thông minh (Intelligent Tutoring System - ITS) là một nền tảng học tập thích ứng được xây dựng theo kiến trúc Microservices. Hệ thống cung cấp trải nghiệm học tập được cá nhân hóa bằng cách:

- **Đánh giá kiến thức**: Sử dụng bài kiểm tra đầu vào (diagnostic test) để xác định trình độ học sinh
- **Chấm điểm tự động**: Đánh giá câu trả lời và cung cấp phản hồi tức thì với gợi ý (hints)
- **Theo dõi tiến độ**: Sử dụng thuật toán Bayesian Knowledge Tracing (BKT) để tính toán mức độ thành thạo kỹ năng
- **Đề xuất thông minh**: Gợi ý nội dung học tập phù hợp dựa trên mô hình AI
- **Gamification**: Tạo động lực học tập với XP, badges, và streaks

### Đặc điểm Kiến trúc

- **Database-per-Service**: Mỗi service có database PostgreSQL riêng biệt
- **Event-Driven Communication**: Sử dụng RabbitMQ cho async messaging
- **API Gateway Pattern**: Client giao tiếp thông qua Adaptive Engine
- **Containerization**: Tất cả services chạy trong Docker containers
- **Independent Deployment**: Mỗi service có thể deploy và scale độc lập
- **Polyglot Architecture**: Java (Content), Go (Scoring, Learner Model, Adaptive Engine), Node.js (Client)

## Kiến trúc Microservices

### Sơ đồ Kiến trúc Tổng quan

```mermaid
graph TB
    subgraph "Client Layer"
        Client[Client - Next.js<br/>Port 3001<br/>Giao diện người dùng]
    end

    subgraph "Application Services Layer"
        AE[Adaptive Engine<br/>Go - Port 8084<br/>Orchestration & Recommendation<br/>Stateless]
        CS[Content Service<br/>Java/Spring Boot - Port 8081<br/>Quản lý Nội dung Học tập<br/>Database: content_db]
        SS[Scoring Service<br/>Go - Port 8082<br/>Chấm điểm & Phản hồi<br/>Database: scoring_db]
        LM[Learner Model Service<br/>Go - Port 8083<br/>Theo dõi Skill Mastery<br/>Database: learner_db]
        LMC[Learner Model Consumer<br/>Go - Background Process<br/>RabbitMQ Event Consumer]
    end

    subgraph "Infrastructure Layer"
        PG[(PostgreSQL<br/>3 Databases<br/>content_db, scoring_db, learner_db)]
        RMQ[RabbitMQ<br/>Message Queue<br/>Port 5672, 15672<br/>Exchange: scoring_exchange]
        MinIO[MinIO<br/>Object Storage<br/>Port 9000, 9001<br/>Lưu trữ files]
    end

    Client -->|HTTP/REST| AE
    Client -->|HTTP/REST| CS

    AE -->|Query Mastery| LM
    AE -->|Query Content| CS

    SS -->|Publish Events<br/>submission.scored| RMQ
    RMQ -->|Consume Events| LMC
    LMC -->|Update Mastery| LM

    CS --> PG
    SS --> PG
    LM --> PG
    LM --> MinIO

    style Client fill:#e1f5ff
    style AE fill:#fff4e1
    style CS fill:#e8f5e9
    style SS fill:#fce4ec
    style LM fill:#f3e5f5
    style LMC fill:#f3e5f5
    style PG fill:#e0e0e0
    style RMQ fill:#ffe0b2
    style MinIO fill:#e0e0e0
```

### Mô tả Chi tiết các Microservices

#### 1. Content Service (Java/Spring Boot)

**Trách nhiệm**: Quản lý nội dung học tập (courses, chapters, lessons, quizzes)

- **Port**: 8081
- **Database**: `content_db` (PostgreSQL)
- **Technology**: Java 17, Spring Boot 3.5.6, JPA/Hibernate
- **Architecture Pattern**: Clean Architecture (Layered)

**Chức năng chính**:

- CRUD operations cho courses và content units
- Metadata tagging (skills, difficulty levels)
- Content versioning (hỗ trợ rollback)
- Path conditions (unlock logic cho adaptive learning)
- API documentation với Swagger

**README chi tiết**: [content/README.md](./content/README.md)

---

#### 2. Scoring Service (Go)

**Trách nhiệm**: Chấm điểm bài tập và phát sự kiện (events)

- **Port**: 8082
- **Database**: `scoring_db` (PostgreSQL)
- **Technology**: Go 1.23, Gin Framework
- **Architecture Pattern**: Hexagonal Architecture

**Chức năng chính**:

- Auto-grading cho MCQ (Multiple Choice Questions)
- Auto-grading cho coding tasks
- Manual review workflow cho bài tập phức tạp
- RabbitMQ event publishing (`submission.scored`)
- Feedback generation với hints và giải thích

**README chi tiết**: [scoring/README.md](./scoring/README.md)

---

#### 3. Learner Model Service (Go)

**Trách nhiệm**: Theo dõi mức độ thành thạo kỹ năng (skill mastery)

- **Port**: 8083 (API), Consumer (background process)
- **Database**: `learner_db` (PostgreSQL)
- **Technology**: Go 1.23, Gin Framework
- **Architecture Pattern**: API + Consumer (Event-Driven)

**Chức năng chính**:

- Bayesian Knowledge Tracing (BKT) algorithm
- RabbitMQ event consumption (từ Scoring Service)
- Skill mastery score calculation
- Learner profile management
- Gamification (XP, badges, streaks)
- MinIO integration cho file storage

**README chi tiết**: [learner-model/README.md](./learner-model/README.md)

---

#### 4. Adaptive Engine (Go)

**Trách nhiệm**: Đề xuất nội dung học tập cá nhân hóa

- **Port**: 8084
- **Database**: Không có (stateless service)
- **Technology**: Go 1.23, Gin Framework
- **Architecture Pattern**: Orchestration Service

**Chức năng chính**:

- Recommendation algorithm (mastery-based)
- Spaced repetition logic
- Live model swapping (Blue/Green deployment)
- Aggregates data từ Content và Learner Model services
- API Gateway cho Client

**README chi tiết**: [adaptive-engine/README.md](./adaptive-engine/README.md)

---

#### 5. Client (Next.js)

**Trách nhiệm**: Giao diện người dùng (UI/UX)

- **Port**: 3001
- **Technology**: Next.js 15, React 19, Tailwind CSS, Zustand
- **Architecture Pattern**: App Router (Next.js 15)

**Chức năng chính**:

- Dashboard (progress tracking, analytics)
- Learning interface (lessons, quizzes)
- Profile management
- Real-time feedback display
- E2E tests với Playwright

**README chi tiết**: [client/README.md](./client/README.md)

## Technology Stack

### Backend Services

| Service                   | Language | Framework         | Database      | Message Queue        | Storage |
| ------------------------- | -------- | ----------------- | ------------- | -------------------- | ------- |
| **Content Service**       | Java 17  | Spring Boot 3.5.6 | PostgreSQL 15 | -                    | -       |
| **Scoring Service**       | Go 1.23  | Gin               | PostgreSQL 15 | RabbitMQ (Publisher) | -       |
| **Learner Model Service** | Go 1.23  | Gin               | PostgreSQL 15 | RabbitMQ (Consumer)  | MinIO   |
| **Adaptive Engine**       | Go 1.23  | Gin               | -             | -                    | -       |

### Frontend

| Component            | Technology   | Version |
| -------------------- | ------------ | ------- |
| **Framework**        | Next.js      | 15.1.3  |
| **UI Library**       | React        | 19.0.0  |
| **Styling**          | Tailwind CSS | 3.4.1   |
| **State Management** | Zustand      | 5.0.2   |
| **Testing**          | Playwright   | 1.49.1  |

### Infrastructure

| Component          | Version                | Purpose                               |
| ------------------ | ---------------------- | ------------------------------------- |
| **PostgreSQL**     | 15-alpine              | Relational database cho mỗi service   |
| **RabbitMQ**       | 3.13-management-alpine | Message queue cho async communication |
| **MinIO**          | Latest                 | Object storage cho files              |
| **Docker**         | 20.10+                 | Containerization                      |
| **Docker Compose** | 2.0+                   | Multi-container orchestration         |

### Development Tools

- **API Documentation**: Swagger/OpenAPI 3.0
- **Database Migration**: Flyway (Java), golang-migrate (Go)
- **Testing**: JUnit 5 (Java), Go testing framework, Playwright (E2E)
- **Logging**: Logback (Java), Zap (Go)
- **Monitoring**: Health check endpoints

## Prerequisites

### Yêu cầu Bắt buộc

Trước khi bắt đầu, đảm bảo máy tính của bạn đã cài đặt:

| Phần mềm           | Phiên bản Tối thiểu | Mục đích                    | Kiểm tra Cài đặt           |
| ------------------ | ------------------- | --------------------------- | -------------------------- |
| **Docker Desktop** | 4.0+                | Container runtime           | `docker --version`         |
| **Docker Compose** | 2.0+                | Orchestration               | `docker-compose --version` |
| **Git**            | 2.30+               | Version control             | `git --version`            |
| **Make**           | 3.81+               | Build automation (tùy chọn) | `make --version`           |

### Yêu cầu Tùy chọn (Local Development)

Nếu bạn muốn chạy services locally (không dùng Docker):

| Phần mềm       | Phiên bản | Service                                 |
| -------------- | --------- | --------------------------------------- |
| **Java JDK**   | 17+       | Content Service                         |
| **Maven**      | 3.8+      | Content Service build                   |
| **Go**         | 1.23+     | Scoring, Learner Model, Adaptive Engine |
| **Node.js**    | 20+       | Client                                  |
| **npm**        | 10+       | Client package manager                  |
| **PostgreSQL** | 15+       | Databases                               |
| **RabbitMQ**   | 3.13+     | Message queue                           |

### Hệ điều hành Hỗ trợ

- **macOS**: Intel & Apple Silicon (M1/M2/M3)
- **Linux**: Ubuntu 20.04+, Debian 11+, Fedora 35+
- **Windows**: 10/11 với WSL2

### Yêu cầu Phần cứng

- **RAM**: Tối thiểu 8GB (khuyến nghị 16GB)
- **Disk**: Tối thiểu 10GB trống
- **CPU**: 4 cores (khuyến nghị)
- **Network**: Kết nối internet (cho Docker image pulls)

### Kiểm tra Prerequisites

```bash
# Kiểm tra Docker
docker --version
docker-compose --version

# Kiểm tra Docker đang chạy
docker ps

# Kiểm tra disk space
df -h

# Kiểm tra memory
free -h  # Linux
vm_stat  # macOS
```

## Installation

### Bước 1: Clone Repository

```bash
git clone https://github.com/your-username/co3017-software-architecture.git
cd co3017-software-architecture/sources
```

### Bước 2: Cấu hình Environment Variables (Tùy chọn)

Hệ thống sử dụng environment variables mặc định trong `docker-compose.yml`. Nếu cần tùy chỉnh:

```bash
# Content Service
cd content
cp template.env .env
# Chỉnh sửa .env nếu cần

# Scoring Service
cd ../scoring
cp template.env .env

# Learner Model Service
cd ../learner-model
cp template.env .env

# Adaptive Engine
cd ../adaptive-engine
cp template.env .env

# Client
cd ../client
cp .env.local.example .env.local
```

### Bước 3: Tạo Docker Network

```bash
# Tạo network cho tất cả services
docker network create its-network
```

## Running Services

### Phương pháp 1: Sử dụng Makefile (Khuyến nghị)

Makefile cung cấp các commands tiện lợi để quản lý hệ thống:

```bash
# Xem tất cả commands có sẵn
make help

# Khởi động infrastructure (PostgreSQL, RabbitMQ)
make infra

# Đợi infrastructure khởi động hoàn toàn
sleep 10

# Khởi tạo databases
make db-init

# Build và khởi động tất cả application services
make services

# Kiểm tra health của tất cả services
make health
```

### Phương pháp 2: Sử dụng Docker Compose Trực tiếp

#### Bước 1: Khởi động Infrastructure

```bash
# Khởi động PostgreSQL, RabbitMQ, MinIO
docker-compose -f docker-compose.infra.yml up -d

# Kiểm tra infrastructure đang chạy
docker-compose -f docker-compose.infra.yml ps
```

#### Bước 2: Khởi tạo Databases

```bash
# Chạy database initialization scripts
docker exec -it its-postgres psql -U postgres -d content_db -f /docker-entrypoint-initdb.d/01-init-content-db.sql
docker exec -it its-postgres psql -U postgres -d scoring_db -f /docker-entrypoint-initdb.d/02-init-scoring-db.sql
docker exec -it its-postgres psql -U postgres -d learner_db -f /docker-entrypoint-initdb.d/03-init-learner-db.sql
```

#### Bước 3: Khởi động Application Services

```bash
# Build và khởi động tất cả services
docker-compose up -d --build

# Hoặc khởi động từng service riêng lẻ
docker-compose up -d content-service
docker-compose up -d scoring-service
docker-compose up -d learner-model-api
docker-compose up -d learner-model-consumer
docker-compose up -d adaptive-engine
docker-compose up -d client
```

#### Bước 4: Kiểm tra Services

```bash
# Xem logs của tất cả services
docker-compose logs -f

# Xem logs của một service cụ thể
docker-compose logs -f content-service

# Kiểm tra status
docker-compose ps
```

### Phương pháp 3: Local Development (Không dùng Docker)

Nếu bạn muốn chạy services locally để development:

```bash
# Khởi động chỉ infrastructure
make infra

# Chạy Content Service
cd content
mvn spring-boot:run

# Chạy Scoring Service (terminal mới)
cd scoring
go run cmd/api/main.go

# Chạy Learner Model API (terminal mới)
cd learner-model
go run cmd/api/main.go

# Chạy Learner Model Consumer (terminal mới)
cd learner-model
go run cmd/consumer/main.go

# Chạy Adaptive Engine (terminal mới)
cd adaptive-engine
go run cmd/api/main.go

# Chạy Client (terminal mới)
cd client
npm install
npm run dev
```

### Truy cập Ứng dụng

Sau khi tất cả services đã khởi động:

| Service                 | URL                                      | Mô tả                                                       |
| ----------------------- | ---------------------------------------- | ----------------------------------------------------------- |
| **Client (Frontend)**   | http://localhost:3001                    | Giao diện người dùng                                        |
| **Content Service API** | http://localhost:8081/swagger-ui.html    | API documentation                                           |
| **Scoring Service API** | http://localhost:8082/swagger/index.html | API documentation                                           |
| **Learner Model API**   | http://localhost:8083/swagger/index.html | API documentation                                           |
| **Adaptive Engine API** | http://localhost:8084/swagger/index.html | API documentation                                           |
| **RabbitMQ Management** | http://localhost:15672                   | Queue monitoring (user: `admintest`, pass: `adminTest2025`) |
| **MinIO Console**       | http://localhost:9001                    | Object storage (user: `minioadmin`, pass: `minioadmin123`)  |

## Docker Deployment Strategy

### Docker Compose Architecture

Hệ thống sử dụng **hai file Docker Compose** để tách biệt infrastructure và application services:

1. **docker-compose.infra.yml** - Infrastructure layer (databases, message queues, storage)
2. **docker-compose.yml** - Application layer (microservices và frontend)

Kiến trúc này cho phép:

- **Independent lifecycle management**: Infrastructure có thể chạy độc lập
- **Faster development**: Rebuild application services mà không ảnh hưởng infrastructure
- **Resource optimization**: Infrastructure chỉ cần khởi động một lần
- **Clear separation of concerns**: Dễ quản lý và troubleshoot

### Infrastructure Layer (docker-compose.infra.yml)

#### PostgreSQL - Unified Database Service

```yaml
postgres:
  image: postgres:15-alpine
  container_name: its-postgres
  environment:
    POSTGRES_USER: postgres
    POSTGRES_PASSWORD: postgres
    POSTGRES_MULTIPLE_DATABASES: content_db,scoring_db,learner_db
  ports:
    - "5432:5432"
  volumes:
    - postgres-data:/var/lib/postgresql/data
    - ./scripts/init-multiple-postgresql-databases.sh:/docker-entrypoint-initdb.d/
```

**Đặc điểm**:

- **Single PostgreSQL instance** với 3 databases riêng biệt (database-per-service pattern)
- **Persistent volume** `postgres-data` để lưu trữ data
- **Initialization scripts** tự động tạo databases khi container khởi động lần đầu
- **Health check** để đảm bảo database sẵn sàng trước khi services kết nối

**Databases**:

- `content_db` - Content Service (courses, lessons, quizzes)
- `scoring_db` - Scoring Service (submissions, results)
- `learner_db` - Learner Model Service (skill mastery, profiles)

#### RabbitMQ - Message Queue Service

```yaml
rabbitmq:
  image: rabbitmq:3.13-management-alpine
  container_name: its-rabbitmq
  environment:
    RABBITMQ_DEFAULT_USER: admintest
    RABBITMQ_DEFAULT_PASS: adminTest2025
  ports:
    - "5672:5672" # AMQP protocol
    - "15672:15672" # Management UI
  volumes:
    - rabbitmq-data:/var/lib/rabbitmq
```

**Đặc điểm**:

- **Management plugin** enabled cho monitoring và debugging
- **Persistent volume** để lưu trữ queues và messages
- **Default vhost** `/` cho tất cả services
- **Health check** để verify RabbitMQ sẵn sàng

**Exchanges và Queues**:

- Exchange: `scoring_exchange` (topic type)
- Queue: `learner_queue` (bound to exchange)
- Routing key: `submission.scored`

### Application Layer (docker-compose.yml)

#### Service Dependencies và Startup Order

Hệ thống có dependency chain rõ ràng:

```
Infrastructure (PostgreSQL, RabbitMQ)
    ↓
Base Services (Content, Scoring, Learner Model API)
    ↓
Consumer (Learner Model Consumer) - depends on Scoring
    ↓
Orchestration (Adaptive Engine) - depends on Content + Learner Model
    ↓
Frontend (Client) - depends on Adaptive Engine + Content
```

**Startup sequence**:

1. **Infrastructure** khởi động trước (manual start với `docker-compose.infra.yml`)
2. **Content Service** - Không có dependencies, có thể start ngay
3. **Scoring Service** - Không có dependencies, có thể start ngay
4. **Learner Model API** - Không có dependencies, có thể start ngay
5. **Learner Model Consumer** - Đợi Scoring Service (để đảm bảo events được publish)
6. **Adaptive Engine** - Đợi Content Service và Learner Model API
7. **Client** - Đợi Adaptive Engine và Content Service

**Docker Compose dependencies**:

```yaml
learner-model-consumer:
  depends_on:
    scoring-service:
      condition: service_started

adaptive-engine:
  depends_on:
    content-service:
      condition: service_started
    learner-model-api:
      condition: service_started

client:
  depends_on:
    adaptive-engine:
      condition: service_started
    content-service:
      condition: service_started
```

**Lưu ý**: `depends_on` chỉ đảm bảo startup order, không đảm bảo service đã sẵn sàng. Sử dụng health checks và retry logic trong application code.

#### Service Configuration Patterns

Mỗi service được configure với:

**1. Build Context**:

```yaml
build:
  context: ./content
  dockerfile: Dockerfile
```

**2. Container Naming**:

```yaml
container_name: its-content-service
```

- Prefix `its-` để dễ identify
- Consistent naming cho troubleshooting

**3. Restart Policy**:

```yaml
restart: unless-stopped
```

- Services tự động restart nếu crash
- Không restart nếu manually stopped

**4. Environment Variables**:

```yaml
environment:
  POSTGRES_HOST: postgres # Service name resolution
  POSTGRES_PORT: 5432
  POSTGRES_DB: content_db
```

**5. Port Mapping**:

```yaml
ports:
  - "8081:8081" # host:container
```

**6. Health Checks** (cho services hỗ trợ):

```yaml
healthcheck:
  test: ["CMD", "curl", "-f", "http://localhost:8081/health"]
  interval: 30s
  timeout: 5s
  retries: 3
  start_period: 60s
```

### Network Configuration

#### Docker Network: its-network

```yaml
networks:
  its-network:
    driver: bridge
    name: its-network
```

**Đặc điểm**:

- **Bridge driver**: Default network driver cho single-host deployment
- **Service discovery**: Services có thể gọi nhau bằng container name
- **DNS resolution**: Docker tự động resolve service names thành IP addresses
- **Network isolation**: Services chỉ có thể communicate trong cùng network

**Service Name Resolution**:

```bash
# Từ bên trong Content Service container
curl http://postgres:5432              # Kết nối PostgreSQL
curl http://rabbitmq:5672              # Kết nối RabbitMQ

# Từ bên trong Adaptive Engine container
curl http://content-service:8081/api   # Gọi Content Service
curl http://learner-model-api:8083/api # Gọi Learner Model
```

**External Network**:

Infrastructure layer tạo network:

```yaml
# docker-compose.infra.yml
networks:
  its-network:
    driver: bridge
    name: its-network
```

Application layer sử dụng existing network:

```yaml
# docker-compose.yml
networks:
  its-network:
    external: true
```

**Port Exposure**:

| Service               | Internal Port | External Port | Purpose            |
| --------------------- | ------------- | ------------- | ------------------ |
| PostgreSQL            | 5432          | 5432          | Database access    |
| RabbitMQ (AMQP)       | 5672          | 5672          | Message queue      |
| RabbitMQ (Management) | 15672         | 15672         | Web UI             |
| Content Service       | 8081          | 8081          | REST API           |
| Scoring Service       | 8082          | 8082          | REST API           |
| Learner Model API     | 8083          | 8083          | REST API           |
| Adaptive Engine       | 8084          | 8084          | REST API (Gateway) |
| Client                | 3001          | 3001          | Web UI             |

**Network Security**:

- Services chỉ expose ports cần thiết
- Internal communication không cần expose ports
- Production nên sử dụng reverse proxy (nginx) và chỉ expose port 80/443

### Volume Management

#### Persistent Volumes

Hệ thống sử dụng **named volumes** để persist data:

```yaml
volumes:
  postgres-data: # PostgreSQL data
  rabbitmq-data: # RabbitMQ queues và messages
```

**Lợi ích của Named Volumes**:

- **Data persistence**: Data không bị mất khi container restart hoặc rebuild
- **Performance**: Tốt hơn bind mounts trên Docker Desktop
- **Portability**: Dễ backup và migrate
- **Management**: Docker quản lý lifecycle

#### Volume Locations

```bash
# Xem tất cả volumes
docker volume ls

# Inspect volume để xem location
docker volume inspect postgres-data

# Output (macOS/Linux):
{
    "Mountpoint": "/var/lib/docker/volumes/postgres-data/_data",
    ...
}
```

#### Volume Operations

**Backup Volume**:

```bash
# Backup PostgreSQL data
docker run --rm \
  -v postgres-data:/data \
  -v $(pwd):/backup \
  alpine tar czf /backup/postgres-backup.tar.gz -C /data .

# Backup RabbitMQ data
docker run --rm \
  -v rabbitmq-data:/data \
  -v $(pwd):/backup \
  alpine tar czf /backup/rabbitmq-backup.tar.gz -C /data .
```

**Restore Volume**:

```bash
# Restore PostgreSQL data
docker run --rm \
  -v postgres-data:/data \
  -v $(pwd):/backup \
  alpine sh -c "cd /data && tar xzf /backup/postgres-backup.tar.gz"
```

**Delete Volume** (cẩn thận - mất data!):

```bash
# Dừng services trước
docker-compose -f docker-compose.infra.yml down

# Xóa volume
docker volume rm postgres-data
docker volume rm rabbitmq-data

# Hoặc xóa tất cả unused volumes
docker volume prune
```

**Volume Size Monitoring**:

```bash
# Kiểm tra disk usage
docker system df -v

# Output:
# VOLUME NAME          SIZE
# postgres-data        150MB
# rabbitmq-data        50MB
```

#### Bind Mounts (Development)

Cho local development, có thể sử dụng bind mounts để hot reload:

```yaml
# Example: Mount source code
services:
  content-service:
    volumes:
      - ./content/src:/app/src:ro # Read-only mount
```

**Lưu ý**: Bind mounts có performance issues trên Docker Desktop (macOS/Windows). Chỉ dùng cho development.

### Environment Variables Strategy

#### Configuration Hierarchy

1. **Default values** trong Dockerfile
2. **docker-compose.yml** environment section
3. **.env files** (nếu có)
4. **Runtime overrides** với `-e` flag

#### Service-Specific Environment Variables

**Content Service (Java/Spring Boot)**:

```yaml
environment:
  SPRING_DATASOURCE_URL: jdbc:postgresql://postgres:5432/content_db
  SPRING_DATASOURCE_USERNAME: postgres
  SPRING_DATASOURCE_PASSWORD: postgres
  SPRING_JPA_HIBERNATE_DDL_AUTO: update
  SERVER_PORT: 8081
  LOGGING_LEVEL_ROOT: INFO
```

**Go Services (Scoring, Learner Model, Adaptive Engine)**:

```yaml
environment:
  APP_HOST: 0.0.0.0
  APP_PORT: 8082
  POSTGRES_HOST: postgres
  POSTGRES_PORT: 5432
  POSTGRES_USER: postgres
  POSTGRES_PASSWORD: postgres
  POSTGRES_DB: scoring_db
  POSTGRES_SSLMODE: disable
  LOG_LEVEL: debug
```

**Client (Next.js)**:

```yaml
environment:
  NODE_ENV: production
  PORT: 3001
  NEXT_PUBLIC_API_BASE_URL: http://localhost:8084
  NEXT_PUBLIC_CONTENT_API: http://localhost:8081
```

**Lưu ý**: Next.js environment variables với prefix `NEXT_PUBLIC_` được expose cho browser.

#### Secrets Management

**Development**: Credentials trong docker-compose.yml (OK cho local)

**Production**: Nên sử dụng:

- Docker secrets
- Environment variable injection từ CI/CD
- Secret management tools (Vault, AWS Secrets Manager)

```yaml
# Production example với Docker secrets
services:
  content-service:
    secrets:
      - db_password
    environment:
      SPRING_DATASOURCE_PASSWORD_FILE: /run/secrets/db_password

secrets:
  db_password:
    external: true
```

### Multi-Stage Builds

Tất cả services sử dụng **multi-stage builds** để optimize image size:

#### Java Service Example (Content Service)

```dockerfile
# Stage 1: Build
FROM maven:3.9-eclipse-temurin-17 AS builder
WORKDIR /app
COPY pom.xml .
RUN mvn dependency:go-offline
COPY src ./src
RUN mvn clean package -DskipTests

# Stage 2: Runtime
FROM eclipse-temurin:17-jre-alpine
WORKDIR /app
COPY --from=builder /app/target/*.jar app.jar
EXPOSE 8081
ENTRYPOINT ["java", "-jar", "app.jar"]
```

**Lợi ích**:

- Build dependencies không có trong final image
- Image size nhỏ hơn (JRE thay vì JDK)
- Faster deployment và pull times

#### Go Service Example (Scoring Service)

```dockerfile
# Stage 1: Build
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/api/main.go

# Stage 2: Runtime
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8082
CMD ["./main"]
```

**Lợi ích**:

- Static binary, không cần Go runtime
- Alpine base image (~5MB)
- Final image size < 20MB

#### Next.js Example (Client)

```dockerfile
# Stage 1: Dependencies
FROM node:20-alpine AS deps
WORKDIR /app
COPY package*.json ./
RUN npm ci

# Stage 2: Build
FROM node:20-alpine AS builder
WORKDIR /app
COPY --from=deps /app/node_modules ./node_modules
COPY . .
RUN npm run build

# Stage 3: Runtime
FROM node:20-alpine AS runner
WORKDIR /app
ENV NODE_ENV production
COPY --from=builder /app/public ./public
COPY --from=builder /app/.next/standalone ./
COPY --from=builder /app/.next/static ./.next/static
EXPOSE 3001
CMD ["node", "server.js"]
```

### Docker Compose Commands Reference

#### Infrastructure Management

```bash
# Khởi động infrastructure
docker-compose -f docker-compose.infra.yml up -d

# Xem logs
docker-compose -f docker-compose.infra.yml logs -f

# Dừng infrastructure
docker-compose -f docker-compose.infra.yml down

# Dừng và xóa volumes (mất data!)
docker-compose -f docker-compose.infra.yml down -v
```

#### Application Management

```bash
# Build và khởi động tất cả services
docker-compose up -d --build

# Khởi động specific service
docker-compose up -d content-service

# Rebuild specific service
docker-compose up -d --build content-service

# Xem logs
docker-compose logs -f content-service

# Dừng tất cả services
docker-compose down

# Restart service
docker-compose restart content-service
```

#### Scaling Services

```bash
# Scale stateless services (Adaptive Engine)
docker-compose up -d --scale adaptive-engine=3

# Verify
docker-compose ps
```

**Lưu ý**: Chỉ scale được stateless services. Services có database connection pool cần configure cẩn thận.

### Deployment Best Practices

#### 1. Health Checks

Luôn implement health check endpoints:

```go
// Go example
func HealthHandler(c *gin.Context) {
    c.JSON(200, gin.H{
        "status": "UP",
        "timestamp": time.Now(),
        "service": "content-service",
    })
}
```

#### 2. Graceful Shutdown

Handle SIGTERM signals:

```go
// Go example
func main() {
    srv := &http.Server{Addr: ":8082"}

    go func() {
        if err := srv.ListenAndServe(); err != nil {
            log.Fatal(err)
        }
    }()

    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    srv.Shutdown(ctx)
}
```

#### 3. Resource Limits

Set memory và CPU limits:

```yaml
services:
  content-service:
    deploy:
      resources:
        limits:
          cpus: "1.0"
          memory: 1G
        reservations:
          cpus: "0.5"
          memory: 512M
```

#### 4. Logging

Centralized logging với Docker:

```yaml
services:
  content-service:
    logging:
      driver: "json-file"
      options:
        max-size: "10m"
        max-file: "3"
```

#### 5. Security

- Không hardcode credentials
- Sử dụng non-root users trong containers
- Scan images cho vulnerabilities
- Keep base images updated

```dockerfile
# Run as non-root user
FROM alpine:latest
RUN addgroup -g 1000 appuser && \
    adduser -D -u 1000 -G appuser appuser
USER appuser
```

## Health Checks

### Kiểm tra Health của Tất cả Services

```bash
# Sử dụng Makefile
make health

# Hoặc kiểm tra thủ công
curl http://localhost:8081/health  # Content Service
curl http://localhost:8082/health  # Scoring Service
curl http://localhost:8083/health  # Learner Model API
curl http://localhost:8084/health  # Adaptive Engine
```

### Kết quả Mong đợi

Mỗi service sẽ trả về JSON response:

```json
{
  "status": "UP",
  "timestamp": "2024-12-07T10:30:00Z",
  "service": "content-service",
  "version": "1.0.0"
}
```

### Kiểm tra Infrastructure

```bash
# PostgreSQL
docker exec -it its-postgres psql -U postgres -c "SELECT version();"

# RabbitMQ
curl -u admintest:adminTest2025 http://localhost:15672/api/overview

# MinIO
curl http://localhost:9000/minio/health/live
```

### Kiểm tra Database Connections

```bash
# Content Database
docker exec -it its-postgres psql -U postgres -d content_db -c "\dt"

# Scoring Database
docker exec -it its-postgres psql -U postgres -d scoring_db -c "\dt"

# Learner Model Database
docker exec -it its-postgres psql -U postgres -d learner_db -c "\dt"
```

### Kiểm tra RabbitMQ Queues

```bash
# Xem tất cả queues
curl -u admintest:adminTest2025 http://localhost:15672/api/queues

# Xem queue cụ thể
curl -u admintest:adminTest2025 http://localhost:15672/api/queues/%2F/learner_queue
```

## Testing

### Unit Tests

Mỗi service có unit tests riêng:

```bash
# Content Service (Java/JUnit)
cd content
mvn test

# Scoring Service (Go)
cd scoring
go test ./... -v

# Learner Model Service (Go)
cd learner-model
go test ./... -v

# Adaptive Engine (Go)
cd adaptive-engine
go test ./... -v

# Client (Jest)
cd client
npm test
```

### Integration Tests

Integration tests kiểm tra tương tác giữa các services:

```bash
# Chạy integration tests
cd tests/integration
go test -v

# Hoặc sử dụng script
./tests/integration/run_tests.sh
```

### System Tests

System tests kiểm tra toàn bộ hệ thống end-to-end:

```bash
# Chạy system tests
cd tests/system
go test -v

# Hoặc sử dụng script
./tests/system/run_tests.sh
```

### E2E Tests (Playwright)

E2E tests cho frontend:

```bash
cd client

# Chạy E2E tests
npm run test:e2e

# Chạy với UI mode
npm run test:e2e:ui

# Chạy specific test file
npx playwright test e2e/learning-flow.spec.ts
```

### Test Coverage

```bash
# Content Service
cd content
mvn test jacoco:report
# Report: target/site/jacoco/index.html

# Go Services
cd scoring
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out

# Client
cd client
npm run test:coverage
```

### Manual Testing với curl

#### Test Content Service

```bash
# Lấy danh sách courses
curl http://localhost:8081/api/courses

# Lấy course theo ID
curl http://localhost:8081/api/courses/1

# Lấy questions theo skill
curl http://localhost:8081/api/questions?skill=math_algebra
```

#### Test Scoring Service

```bash
# Submit answer
curl -X POST http://localhost:8082/api/scoring/submit \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "test_user_123",
    "question_id": 1,
    "answer": "A"
  }'
```

#### Test Learner Model Service

```bash
# Lấy skill mastery
curl http://localhost:8083/api/learner/test_user_123/mastery?skill=math_algebra

# Lấy learner profile
curl http://localhost:8083/api/learner/test_user_123/profile
```

#### Test Adaptive Engine

```bash
# Lấy recommended content
curl http://localhost:8084/api/adaptive/recommend?user_id=test_user_123&skill=math_algebra
```

## Troubleshooting

### Common Issues và Solutions

#### 1. Port Already in Use (Port đã được sử dụng)

**Triệu chứng**: Error khi khởi động service: `bind: address already in use`

**Nguyên nhân**: Port đã được process khác sử dụng

**Giải pháp**:

```bash
# Kiểm tra port nào đang được sử dụng
lsof -i :8081  # Thay 8081 bằng port bị conflict
# Hoặc trên Linux
netstat -tulpn | grep 8081

# Dừng process đang chiếm port
kill -9 <PID>

# Hoặc thay đổi port trong docker-compose.yml
# Ví dụ: "8081:8081" -> "8091:8081"
```

#### 2. Database Connection Failed

**Triệu chứng**: Service không thể kết nối database: `connection refused`

**Nguyên nhân**: PostgreSQL chưa khởi động hoàn toàn hoặc credentials sai

**Giải pháp**:

```bash
# Kiểm tra PostgreSQL đang chạy
docker ps | grep postgres

# Kiểm tra logs
docker logs its-postgres

# Restart PostgreSQL
docker-compose -f docker-compose.infra.yml restart postgres

# Kiểm tra connection
docker exec -it its-postgres psql -U postgres -c "SELECT 1;"

# Verify database tồn tại
docker exec -it its-postgres psql -U postgres -c "\l"
```

#### 3. RabbitMQ Connection Failed

**Triệu chứng**: Learner Model Consumer không thể kết nối RabbitMQ

**Nguyên nhân**: RabbitMQ chưa sẵn sàng hoặc credentials sai

**Giải pháp**:

```bash
# Kiểm tra RabbitMQ đang chạy
docker ps | grep rabbitmq

# Kiểm tra logs
docker logs its-rabbitmq

# Restart RabbitMQ
docker-compose -f docker-compose.infra.yml restart rabbitmq

# Kiểm tra queues
curl -u admintest:adminTest2025 http://localhost:15672/api/queues

# Reset RabbitMQ (xóa tất cả queues)
./scripts/reset_rabbitmq.sh
```

#### 4. Docker Out of Memory

**Triệu chứng**: Services bị crash hoặc không khởi động: `OOMKilled`

**Nguyên nhân**: Docker không đủ memory

**Giải pháp**:

```bash
# Kiểm tra Docker memory usage
docker stats

# Tăng memory limit cho Docker Desktop
# Settings → Resources → Memory → Tăng lên 8GB hoặc 16GB

# Giảm số lượng services chạy đồng thời
# Chỉ chạy services cần thiết cho development
```

#### 5. Build Failed - Maven/Go/npm

**Triệu chứng**: Docker build failed với compilation errors

**Nguyên nhân**: Dependencies không tải được hoặc code lỗi

**Giải pháp**:

```bash
# Content Service (Maven)
cd content
mvn clean install -DskipTests
docker-compose build content-service --no-cache

# Go Services
cd scoring
go mod download
go mod tidy
docker-compose build scoring-service --no-cache

# Client (npm)
cd client
rm -rf node_modules package-lock.json
npm install
docker-compose build client --no-cache
```

#### 6. Service Unhealthy

**Triệu chứng**: `docker ps` hiển thị service status là `unhealthy`

**Nguyên nhân**: Health check endpoint không response hoặc service crash

**Giải pháp**:

```bash
# Xem logs chi tiết
docker logs its-content-service --tail 100

# Kiểm tra health endpoint thủ công
docker exec -it its-content-service curl http://localhost:8081/health

# Restart service
docker-compose restart content-service

# Rebuild nếu cần
docker-compose up -d --build content-service
```

#### 7. Database Migration Failed

**Triệu chứng**: Service khởi động nhưng không có tables trong database

**Nguyên nhân**: Migration scripts không chạy hoặc lỗi

**Giải pháp**:

```bash
# Kiểm tra tables
docker exec -it its-postgres psql -U postgres -d content_db -c "\dt"

# Chạy migration thủ công
docker exec -it its-postgres psql -U postgres -d content_db -f /docker-entrypoint-initdb.d/01-init-content-db.sql

# Hoặc sử dụng Makefile
make db-init
```

#### 8. Events Not Processing (RabbitMQ)

**Triệu chứng**: Scoring Service publish events nhưng Learner Model không update

**Nguyên nhân**: Consumer không chạy hoặc queue configuration sai

**Giải pháp**:

```bash
# Kiểm tra Consumer đang chạy
docker ps | grep learner-model-consumer

# Kiểm tra logs
docker logs its-learner-model-consumer --tail 50

# Kiểm tra queue có messages
curl -u admintest:adminTest2025 http://localhost:15672/api/queues/%2F/learner_queue

# Restart consumer
docker-compose restart learner-model-consumer

# Test event flow
./scripts/integration_test.sh
```

#### 9. Client Cannot Connect to Backend

**Triệu chứng**: Frontend hiển thị network errors

**Nguyên nhân**: Environment variables sai hoặc CORS issues

**Giải pháp**:

```bash
# Kiểm tra environment variables
docker exec -it its-client env | grep NEXT_PUBLIC

# Verify backend services đang chạy
make health

# Kiểm tra CORS headers
curl -I http://localhost:8081/api/courses

# Rebuild client với correct env vars
cd client
docker-compose up -d --build client
```

#### 10. Disk Space Full

**Triệu chứng**: Docker operations failed: `no space left on device`

**Nguyên nhân**: Docker images/volumes chiếm quá nhiều disk

**Giải pháp**:

```bash
# Kiểm tra disk usage
df -h
docker system df

# Xóa unused images và containers
docker system prune -a

# Xóa unused volumes (cẩn thận - mất data!)
docker volume prune

# Xóa specific volumes
docker volume rm its-postgres-data
```

### Debugging Tips

#### Xem Logs Real-time

```bash
# Tất cả services
docker-compose logs -f

# Specific service
docker-compose logs -f content-service

# Với timestamp
docker-compose logs -f --timestamps content-service

# Last N lines
docker-compose logs --tail 100 content-service
```

#### Truy cập Container Shell

```bash
# Content Service (Java)
docker exec -it its-content-service bash

# Go Services
docker exec -it its-scoring-service sh

# PostgreSQL
docker exec -it its-postgres psql -U postgres

# RabbitMQ
docker exec -it its-rabbitmq rabbitmqctl status
```

#### Kiểm tra Network Connectivity

```bash
# Ping giữa containers
docker exec -it its-client ping content-service

# Kiểm tra DNS resolution
docker exec -it its-client nslookup content-service

# Kiểm tra port connectivity
docker exec -it its-client nc -zv content-service 8081
```

#### Reset Toàn bộ Hệ thống

```bash
# Dừng tất cả services
make stop

# Xóa containers và volumes
make clean

# Rebuild từ đầu
make setup
```

## Development Workflow

### Local Development Setup

#### Phương pháp 1: Hybrid (Infrastructure trong Docker, Services chạy local)

Phương pháp này cho phép bạn develop và debug services dễ dàng hơn:

```bash
# 1. Khởi động infrastructure
make infra

# 2. Khởi tạo databases
make db-init

# 3. Chạy services locally
# Content Service
cd content
mvn spring-boot:run

# Scoring Service (terminal mới)
cd scoring
go run cmd/api/main.go

# Learner Model API (terminal mới)
cd learner-model
go run cmd/api/main.go

# Learner Model Consumer (terminal mới)
cd learner-model
go run cmd/consumer/main.go

# Adaptive Engine (terminal mới)
cd adaptive-engine
go run cmd/api/main.go

# Client (terminal mới)
cd client
npm run dev
```

#### Phương pháp 2: Full Docker với Hot Reload

Sử dụng Docker volumes để mount source code:

```bash
# Chỉnh sửa docker-compose.yml để thêm volumes
# Ví dụ cho Content Service:
volumes:
  - ./content/src:/app/src

# Khởi động với hot reload
docker-compose up -d
```

### Code Style và Conventions

#### Java (Content Service)

```bash
# Format code
cd content
mvn spotless:apply

# Check style
mvn checkstyle:check

# Run static analysis
mvn pmd:check
```

#### Go (Scoring, Learner Model, Adaptive Engine)

```bash
# Format code
go fmt ./...

# Run linter
golangci-lint run

# Check imports
goimports -w .

# Vet code
go vet ./...
```

#### JavaScript/TypeScript (Client)

```bash
cd client

# Format code
npm run format

# Lint code
npm run lint

# Type check
npm run type-check
```

### Database Migrations

#### Content Service (Flyway)

```bash
cd content

# Run migrations
mvn flyway:migrate

# Check migration status
mvn flyway:info

# Rollback (nếu cần)
mvn flyway:undo
```

#### Go Services (golang-migrate)

```bash
# Scoring Service
cd scoring
migrate -path migration -database "postgresql://postgres:postgres@localhost:5432/scoring_db?sslmode=disable" up

# Learner Model Service
cd learner-model
migrate -path migration -database "postgresql://postgres:postgres@localhost:5432/learner_db?sslmode=disable" up
```

### Adding New Features

#### 1. Thêm API Endpoint Mới

**Content Service (Java)**:

```bash
# 1. Tạo entity
src/main/java/com/its/content/entity/NewEntity.java

# 2. Tạo repository
src/main/java/com/its/content/repository/NewEntityRepository.java

# 3. Tạo service
src/main/java/com/its/content/service/NewEntityService.java

# 4. Tạo controller
src/main/java/com/its/content/controller/NewEntityController.java

# 5. Viết tests
src/test/java/com/its/content/service/NewEntityServiceTest.java
```

**Go Services**:

```bash
# 1. Định nghĩa types
internal/domain/type.go

# 2. Tạo repository interface
internal/domain/interface.go

# 3. Implement repository
internal/repository/postgres/repository.go

# 4. Tạo use case
internal/usecase/usecase.go

# 5. Tạo HTTP handler
internal/delivery/http/handler.go

# 6. Viết tests
internal/usecase/usecase_test.go
```

#### 2. Thêm Database Table Mới

```bash
# 1. Tạo migration file
# Content Service
cd content/src/main/resources/db/migration
touch V2__add_new_table.sql

# Go Services
cd scoring/migration
touch 02_add_new_table.sql

# 2. Viết SQL
CREATE TABLE new_table (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

# 3. Run migration
make db-migrate
```

#### 3. Thêm RabbitMQ Event Mới

```bash
# 1. Định nghĩa event type
# Scoring Service (Publisher)
internal/publisher/event.go

type NewEvent struct {
    EventType string    `json:"event_type"`
    UserID    string    `json:"user_id"`
    Data      EventData `json:"data"`
}

# 2. Publish event
publisher.PublishNewEvent(ctx, event)

# 3. Consume event
# Learner Model Consumer
internal/consumer/handler.go

func (h *Handler) HandleNewEvent(event NewEvent) error {
    // Process event
}
```

### Debugging

#### Java (Content Service)

```bash
# Run với debug mode
mvn spring-boot:run -Dspring-boot.run.jvmArguments="-Xdebug -Xrunjdwp:transport=dt_socket,server=y,suspend=n,address=5005"

# Attach debugger từ IDE (IntelliJ/Eclipse)
# Remote Debug Configuration: localhost:5005
```

#### Go Services

```bash
# Install Delve debugger
go install github.com/go-delve/delve/cmd/dlv@latest

# Run với debugger
cd scoring
dlv debug cmd/api/main.go

# Hoặc attach vào running process
dlv attach <PID>
```

#### Client (Next.js)

```bash
cd client

# Run với debug mode
NODE_OPTIONS='--inspect' npm run dev

# Attach Chrome DevTools
# chrome://inspect
```

### Performance Profiling

#### Go Services

```bash
# CPU profiling
go test -cpuprofile=cpu.prof -bench=.
go tool pprof cpu.prof

# Memory profiling
go test -memprofile=mem.prof -bench=.
go tool pprof mem.prof

# HTTP profiling (add to main.go)
import _ "net/http/pprof"
go tool pprof http://localhost:8082/debug/pprof/profile
```

#### Database Query Analysis

```bash
# PostgreSQL slow query log
docker exec -it its-postgres psql -U postgres -c "ALTER SYSTEM SET log_min_duration_statement = 100;"
docker exec -it its-postgres psql -U postgres -c "SELECT pg_reload_conf();"

# Analyze query
docker exec -it its-postgres psql -U postgres -d content_db -c "EXPLAIN ANALYZE SELECT * FROM courses;"
```

### Continuous Integration

#### Pre-commit Checks

```bash
# Tạo pre-commit hook
cat > .git/hooks/pre-commit << 'EOF'
#!/bin/bash
# Run tests
make test

# Check formatting
make format-check

# Run linters
make lint
EOF

chmod +x .git/hooks/pre-commit
```

#### CI/CD Pipeline (GitHub Actions)

```yaml
# .github/workflows/ci.yml
name: CI

on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - name: Run tests
        run: |
          cd sources
          make test
```

## API Gateway

### Adaptive Engine as API Gateway

Adaptive Engine đóng vai trò là API Gateway cho Client, cung cấp một điểm truy cập duy nhất (single entry point) cho tất cả các operations.

#### Lợi ích của Pattern này

1. **Simplified Client**: Client chỉ cần biết về Adaptive Engine
2. **Service Abstraction**: Ẩn đi complexity của microservices architecture
3. **Centralized Logic**: Recommendation và orchestration logic ở một nơi
4. **Reduced Coupling**: Client không phụ thuộc trực tiếp vào nhiều services

### API Endpoints

#### 1. Recommendation Endpoints

```bash
# Lấy recommended content cho user
GET /api/adaptive/recommend?user_id={user_id}&skill={skill}

# Response
{
  "user_id": "user_123",
  "recommended_content": [
    {
      "content_id": "content_456",
      "title": "Algebra Basics",
      "difficulty": "beginner",
      "relevance_score": 0.95
    }
  ]
}
```

#### 2. Learning Path Endpoints

```bash
# Lấy next lesson trong learning path
GET /api/adaptive/next-lesson?user_id={user_id}&course_id={course_id}

# Response
{
  "next_lesson": {
    "lesson_id": "lesson_789",
    "title": "Linear Equations",
    "estimated_time": 30
  }
}
```

#### 3. Aggregated Data Endpoints

```bash
# Lấy dashboard data (aggregates từ nhiều services)
GET /api/adaptive/dashboard?user_id={user_id}

# Response
{
  "user_id": "user_123",
  "progress": {
    "total_courses": 5,
    "completed_lessons": 23,
    "completion_percentage": 45.5
  },
  "mastery": {
    "math_algebra": 0.75,
    "math_geometry": 0.60
  },
  "gamification": {
    "total_xp": 1250,
    "level": 5,
    "badges": ["first_lesson", "week_streak"]
  }
}
```

### Service Communication Patterns

#### Synchronous Communication (REST API)

Adaptive Engine gọi các services khác qua REST API:

```go
// Adaptive Engine code
func (s *Service) GetRecommendation(userID string) (*Recommendation, error) {
    // 1. Query Learner Model Service
    mastery, err := s.learnerClient.GetMastery(userID)
    if err != nil {
        return nil, err
    }

    // 2. Query Content Service
    content, err := s.contentClient.GetAvailableContent(mastery.WeakSkills)
    if err != nil {
        return nil, err
    }

    // 3. Run recommendation algorithm
    recommendation := s.algorithm.Recommend(mastery, content)

    return recommendation, nil
}
```

#### Asynchronous Communication (RabbitMQ)

Scoring Service → Learner Model Service:

```
1. User submits answer
2. Scoring Service evaluates answer
3. Scoring Service publishes "submission.scored" event to RabbitMQ
4. Learner Model Consumer receives event
5. Learner Model updates skill mastery
6. Next time Adaptive Engine queries, it gets updated mastery
```

### Request Flow Examples

#### Example 1: User Làm Bài Tập

```
Client → Scoring Service: POST /api/scoring/submit
  ↓
Scoring Service → PostgreSQL: Save submission
  ↓
Scoring Service → RabbitMQ: Publish event
  ↓
Learner Model Consumer ← RabbitMQ: Consume event
  ↓
Learner Model → PostgreSQL: Update mastery
```

#### Example 2: User Xem Dashboard

```
Client → Adaptive Engine: GET /api/adaptive/dashboard
  ↓
Adaptive Engine → Learner Model: GET /api/learner/{id}/mastery
  ↓
Adaptive Engine → Learner Model: GET /api/learner/{id}/progress
  ↓
Adaptive Engine → Content Service: GET /api/courses/{id}
  ↓
Adaptive Engine: Aggregate data
  ↓
Adaptive Engine → Client: Return dashboard data
```

#### Example 3: User Nhận Recommendation

```
Client → Adaptive Engine: GET /api/adaptive/recommend
  ↓
Adaptive Engine → Learner Model: GET /api/learner/{id}/mastery
  ↓
Adaptive Engine → Content Service: GET /api/content/available
  ↓
Adaptive Engine: Run recommendation algorithm
  ↓
Adaptive Engine → Client: Return recommended content
```

### Error Handling

Adaptive Engine xử lý errors từ downstream services:

```go
func (s *Service) GetRecommendation(userID string) (*Recommendation, error) {
    // Circuit breaker pattern
    mastery, err := s.learnerClient.GetMastery(userID)
    if err != nil {
        // Fallback: Use cached mastery or default values
        mastery = s.cache.GetMastery(userID)
        if mastery == nil {
            return nil, errors.New("learner service unavailable")
        }
    }

    // Retry logic
    content, err := s.contentClient.GetAvailableContent(mastery.WeakSkills)
    if err != nil {
        // Retry 3 times with exponential backoff
        content, err = s.retryGetContent(mastery.WeakSkills, 3)
        if err != nil {
            return nil, errors.New("content service unavailable")
        }
    }

    return s.algorithm.Recommend(mastery, content), nil
}
```

### Rate Limiting và Caching

```go
// Rate limiting
func (s *Service) RateLimitMiddleware() gin.HandlerFunc {
    limiter := rate.NewLimiter(100, 200) // 100 req/s, burst 200
    return func(c *gin.Context) {
        if !limiter.Allow() {
            c.JSON(429, gin.H{"error": "rate limit exceeded"})
            c.Abort()
            return
        }
        c.Next()
    }
}

// Caching
func (s *Service) GetMasteryWithCache(userID string) (*Mastery, error) {
    // Check cache first
    if cached := s.cache.Get(userID); cached != nil {
        return cached, nil
    }

    // Query service
    mastery, err := s.learnerClient.GetMastery(userID)
    if err != nil {
        return nil, err
    }

    // Cache for 5 minutes
    s.cache.Set(userID, mastery, 5*time.Minute)
    return mastery, nil
}
```

## Useful Commands

### Makefile Commands

```bash
# Xem tất cả commands
make help

# Infrastructure
make infra              # Khởi động PostgreSQL, RabbitMQ, MinIO
make infra-stop         # Dừng infrastructure
make infra-restart      # Restart infrastructure

# Services
make services           # Khởi động tất cả application services
make services-stop      # Dừng tất cả services
make services-restart   # Restart tất cả services

# Database
make db-init            # Khởi tạo databases
make db-migrate         # Run migrations
make db-content         # Kết nối vào content_db
make db-scoring         # Kết nối vào scoring_db
make db-learner         # Kết nối vào learner_db
make db-backup          # Backup tất cả databases

# Testing
make test               # Chạy tất cả tests
make test-unit          # Chạy unit tests
make test-integration   # Chạy integration tests
make test-e2e           # Chạy E2E tests

# Monitoring
make health             # Kiểm tra health của tất cả services
make logs               # Xem logs của tất cả services
make status             # Kiểm tra status của containers

# Cleanup
make stop               # Dừng tất cả containers
make clean              # Dừng và xóa containers, volumes
make rebuild            # Rebuild và restart tất cả services
```

### Docker Commands

```bash
# Container management
docker ps                                    # Xem running containers
docker ps -a                                 # Xem tất cả containers
docker logs -f its-content-service           # Xem logs
docker exec -it its-content-service bash     # Truy cập container shell
docker restart its-content-service           # Restart container
docker stop its-content-service              # Dừng container

# Image management
docker images                                # Xem tất cả images
docker rmi its-content-service:latest        # Xóa image
docker build -t its-content-service .        # Build image

# Network management
docker network ls                            # Xem networks
docker network inspect its-network           # Inspect network

# Volume management
docker volume ls                             # Xem volumes
docker volume inspect postgres-data          # Inspect volume
docker volume rm postgres-data               # Xóa volume

# System cleanup
docker system prune                          # Xóa unused resources
docker system prune -a                       # Xóa tất cả unused resources
docker volume prune                          # Xóa unused volumes
```

## Project Structure

```
sources/
├── docker-compose.yml              # Application services configuration
├── docker-compose.infra.yml        # Infrastructure configuration
├── Makefile                        # Build và management commands
│
├── scripts/                        # Utility scripts
│   ├── start_services.sh           # Khởi động tất cả services
│   ├── stop_services.sh            # Dừng tất cả services
│   ├── integration_test.sh         # Integration test script
│   ├── init-multiple-postgresql-databases.sh
│   └── insert_test_data.sql        # Test data
│
├── content/                        # Content Service (Java/Spring Boot)
│   ├── README.md                   # Service documentation
│   ├── Dockerfile
│   ├── pom.xml                     # Maven configuration
│   └── src/
│       ├── main/java/              # Source code
│       └── test/java/              # Tests
│
├── scoring/                        # Scoring Service (Go)
│   ├── README.md
│   ├── Dockerfile
│   ├── go.mod
│   ├── cmd/api/                    # Main application
│   ├── internal/                   # Internal packages
│   └── pkg/                        # Shared packages
│
├── learner-model/                  # Learner Model Service (Go)
│   ├── README.md
│   ├── cmd/
│   │   ├── api/                    # API server
│   │   └── consumer/               # RabbitMQ consumer
│   ├── internal/
│   └── pkg/
│
├── adaptive-engine/                # Adaptive Engine (Go)
│   ├── README.md
│   ├── cmd/api/
│   ├── internal/
│   └── pkg/
│
├── client/                         # Frontend Client (Next.js)
│   ├── README.md
│   ├── Dockerfile
│   ├── package.json
│   ├── app/                        # Next.js App Router
│   ├── components/                 # React components
│   ├── e2e/                        # Playwright E2E tests
│   └── public/
│
└── tests/                          # Integration & System tests
    ├── integration/                # Integration tests
    └── system/                     # System tests
```

## Liên kết Liên quan (Related Links)

### Service READMEs

| Service                   | Đường dẫn                                                | Công nghệ        | Port | Mô tả                                       |
| ------------------------- | -------------------------------------------------------- | ---------------- | ---- | ------------------------------------------- |
| **Content Service**       | [content/README.md](./content/README.md)                 | Java/Spring Boot | 8081 | Quản lý nội dung học tập, API documentation |
| **Scoring Service**       | [scoring/README.md](./scoring/README.md)                 | Go/Gin           | 8082 | Chấm điểm, RabbitMQ event publishing        |
| **Learner Model Service** | [learner-model/README.md](./learner-model/README.md)     | Go/Gin           | 8083 | BKT algorithm, skill mastery tracking       |
| **Adaptive Engine**       | [adaptive-engine/README.md](./adaptive-engine/README.md) | Go/Gin           | 8084 | Recommendation algorithm, orchestration     |
| **Client**                | [client/README.md](./client/README.md)                   | Next.js/React    | 3001 | Frontend UI/UX, E2E tests                   |

### Tài liệu Dự án

| Tài liệu                | Đường dẫn                                              | Mô tả                                |
| ----------------------- | ------------------------------------------------------ | ------------------------------------ |
| **Root README**         | [../README.md](../README.md)                           | Tổng quan dự án, cấu trúc repository |
| **Report README**       | [../report/README.md](../report/README.md)             | Hướng dẫn build báo cáo LaTeX        |
| **Presentation README** | [../presentation/README.md](../presentation/README.md) | Hướng dẫn xem/build slides           |

### Tài liệu Phân tích Kiến trúc

| Tài liệu                         | Đường dẫn                                                                                                    | Nội dung                               |
| -------------------------------- | ------------------------------------------------------------------------------------------------------------ | -------------------------------------- |
| **Microservices Analysis**       | [../markdown/microservices.md](../markdown/microservices.md)                                                 | Chi tiết domain model, database schema |
| **Architecture Decisions**       | [../markdown/report/5-architecture-decisions.md](../markdown/report/5-architecture-decisions.md)             | Architecture Decision Records (ADRs)   |
| **SOLID Principles**             | [../markdown/report/6-SOLID-principles.md](../markdown/report/6-SOLID-principles.md)                         | Ví dụ áp dụng SOLID trong code         |
| **Architecture Characteristics** | [../markdown/report/2-architecture-characteristics.md](../markdown/report/2-architecture-characteristics.md) | Đặc tính kiến trúc ưu tiên             |
| **Architecture Styles**          | [../markdown/report/3-architecture-styles.md](../markdown/report/3-architecture-styles.md)                   | So sánh các kiểu kiến trúc             |

### Diagrams

| Diagram                       | Đường dẫn                                                                                                                                | Loại               |
| ----------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ------------------ |
| **Domain Model**              | [../markdown/diagrams/domain_model_class_diagram.md](../markdown/diagrams/domain_model_class_diagram.md)                                 | Class Diagram      |
| **Deployment Architecture**   | [../markdown/diagrams/deployment_architecture_onprem.md](../markdown/diagrams/deployment_architecture_onprem.md)                         | Deployment Diagram |
| **Adaptive Content Delivery** | [../markdown/diagrams/adaptive_content_delivery_sequence.md](../markdown/diagrams/adaptive_content_delivery_sequence.md)                 | Sequence Diagram   |
| **Assessment Submission**     | [../markdown/diagrams/assessment_submission_and_scoring_sequence.md](../markdown/diagrams/assessment_submission_and_scoring_sequence.md) | Sequence Diagram   |

### Testing

| Test Suite            | Đường dẫn                                  | Mô tả                             |
| --------------------- | ------------------------------------------ | --------------------------------- |
| **Integration Tests** | [tests/integration/](./tests/integration/) | Integration tests cho services    |
| **System Tests**      | [tests/system/](./tests/system/)           | System-level tests                |
| **E2E Tests**         | [client/e2e/](./client/e2e/)               | Playwright E2E tests cho frontend |

### External Documentation

| Tài liệu        | Link                                                        | Mô tả                     |
| --------------- | ----------------------------------------------------------- | ------------------------- |
| **Docker**      | [docs.docker.com](https://docs.docker.com/)                 | Docker documentation      |
| **Spring Boot** | [spring.io](https://spring.io/projects/spring-boot)         | Spring Boot documentation |
| **Go**          | [golang.org](https://golang.org/doc/)                       | Go documentation          |
| **Next.js**     | [nextjs.org](https://nextjs.org/docs)                       | Next.js documentation     |
| **RabbitMQ**    | [rabbitmq.com](https://www.rabbitmq.com/documentation.html) | RabbitMQ documentation    |
| **PostgreSQL**  | [postgresql.org](https://www.postgresql.org/docs/)          | PostgreSQL documentation  |

## Support

Nếu gặp vấn đề hoặc có câu hỏi:

1. Kiểm tra [Troubleshooting](#troubleshooting) section
2. Xem logs: `make logs` hoặc `docker-compose logs -f`
3. Kiểm tra health: `make health`
4. Tham khảo service-specific README trong từng thư mục service

## License

Dự án này được phát triển cho mục đích học tập tại HCMUT - Môn CO3017 Kiến Trúc Phần Mềm.

---

**Phát triển bởi**: Nhóm sinh viên HCMUT  
**Môn học**: CO3017 - Kiến Trúc Phần Mềm  
**Học kỳ**: HK251 (2024-2025)
