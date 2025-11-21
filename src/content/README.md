# CO3017 Content Service - Clean Architecture Implementation

## 📋 Tổng quan

Microservice quản lý nội dung, tài liệu và dữ liệu cho môn CO3017 - HCMUT.

Dự án được xây dựng theo **Clean Architecture** pattern với **Spring Boot 3.5.6**, **Java 17**, và **Maven**.

## 🏗️ Kiến trúc Clean Architecture

### 🎯 Nguyên tắc Clean Architecture

1. **Dependency Rule**: Dependencies chỉ trỏ vào trong (inward)
2. **Domain Independence**: Core business logic không phụ thuộc framework
3. **Testability**: Dễ dàng test từng layer độc lập
4. **Flexibility**: Dễ thay đổi implementation mà không ảnh hưởng core logic

### 📐 Layer Structure

```
┌─────────────────────────────────────┐
│         Domain Layer                │  ← Core Business Logic
│   (Content, Document entities)      │
└─────────────────────────────────────┘
           ↑
┌─────────────────────────────────────┐
│      Application Layer              │  ← Use Cases & Business Rules
│  (ContentUseCase, DocumentUseCase) │
└─────────────────────────────────────┘
           ↑
┌──────────────────┬──────────────────┐
│  Adapter Layer   │ Infrastructure   │  ← External Communication
│  (REST API)      │  (Database)      │
└──────────────────┴──────────────────┘
```

## 📁 Cấu trúc Project

```
src/main/java/co3017/microservices/content_service/
├── models/                    ← Domain Entities (Pure Business Logic)
│   ├── Content.java
│   └── Document.java
├── usecase/                   ← Application Layer
│   ├── ContentUseCase.java    ← Interface cho Content domain
│   ├── DocumentUseCase.java   ← Interface cho Document domain
│   ├── service/               ← Use Case Implementations
│   │   ├── ContentService.java
│   │   └── DocumentService.java
│   └── types/                 ← Command Objects
│       ├── CreateContentCommand.java
│       └── CreateDocumentCommand.java
├── repository/                ← Repository Layer
│   ├── ContentRepository.java ← Repository Interface
│   ├── DocumentRepository.java ← Repository Interface
│   └── postgresql/            ← Database Implementation
│       ├── entity/            ← JPA Entities
│       │   ├── ContentEntity.java
│       │   └── DocumentEntity.java
│       ├── mapper/            ← Domain ↔ Entity Mapping
│       │   ├── ContentMapper.java
│       │   └── DocumentMapper.java
│       ├── JpaContentRepository.java
│       ├── JpaDocumentRepository.java
│       ├── SpringDataContentRepository.java
│       └── SpringDataDocumentRepository.java
├── adapter/http/              ← Presentation Layer
│   ├── ContentController.java ← REST Controller
│   ├── DocumentController.java ← REST Controller
│   ├── dto/                   ← Data Transfer Objects
│   │   ├── ApiResponse.java   ← Standardized Response Format
│   │   ├── CreateContentRequest.java
│   │   ├── ContentResponse.java
│   │   ├── CreateDocumentRequest.java
│   │   └── DocumentResponse.java
│   └── response/              ← Response Builders
│       ├── CommandBuilder.java
│       ├── ContentResponseBuilder.java
│       └── DocumentResponseBuilder.java
└── config/                    ← Configuration
    ├── ContentServiceApplication.java
    └── CorsConfig.java        ← CORS Configuration
```

## 🔄 Luồng Request Flow

### 1️⃣ HTTP Request Flow

```
Client Request → Controller → UseCase → Repository → Database
     ↓              ↓           ↓          ↓
Response ← ResponseBuilder ← Domain ← Entity ← Database
```

### 2️⃣ Chi tiết từng bước

1. **HTTP Request** đến `ContentController`
2. **Controller** nhận DTO, validate input
3. **CommandBuilder** chuyển DTO → Command
4. **UseCase Interface** định nghĩa contract
5. **UseCase Service** thực thi business logic
6. **Repository Interface** định nghĩa data contract
7. **JPA Repository** implement database operations
8. **Entity** ↔ **Domain** mapping qua Mapper
9. **ResponseBuilder** chuyển Domain → Response DTO
10. **ApiResponse** wrap kết quả với format chuẩn

### 3️⃣ Ví dụ: Tạo Content

```java
// 1. Controller nhận request
@PostMapping
public ResponseEntity<ApiResponse<ContentResponse>> createContent(@RequestBody CreateContentRequest request) {
    // 2. Chuyển DTO → Command
    Content content = contentUseCase.createContent(CommandBuilder.toCreateContentCommand(request));
    // 3. Chuyển Domain → Response
    ContentResponse response = ContentResponseBuilder.toResponse(content);
    return ResponseEntity.ok(ApiResponse.success(response));
}

// 4. UseCase thực thi business logic
public Content createContent(CreateContentCommand command) {
    // Business validation
    if (command.getTitle().isEmpty()) {
        throw new IllegalArgumentException("Title không được để trống");
    }
    // Tạo domain entity
    Content content = new Content(command.getTitle(), command.getBody());
    // Lưu qua repository
    return contentRepository.save(content);
}
```

## 🎯 Domains

### 1️⃣ Content Domain
**Chức năng**: Quản lý nội dung và tài liệu
- **Endpoints**: `/api/contents`
- **Features**: Create, Get by ID/Title, List all
- **Business Rules**: Title unique, content validation

### 2️⃣ Document Domain
**Chức năng**: Quản lý tài liệu và file
- **Endpoints**: `/api/documents`
- **Features**: Create, Get by ID, Search by title, List all
- **Business Rules**: Title unique, file size > 0, format validation

## 🚀 Hướng dẫn Chạy Source

### Prerequisites
- **Java 17+**
- **Maven 3.6+**
- **PostgreSQL**
- **IDE**: IntelliJ IDEA, VS Code, hoặc Eclipse

### 1️⃣ Setup Database

```bash
# Tạo database
createdb co3017

# Tạo tables (vì dùng ddl-auto: validate)
psql -U postgres -d co3017 -c "
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE tests (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    duration INTEGER NOT NULL,
    max_score INTEGER NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
"
```

### 2️⃣ Chạy Application

#### Cách 1: Maven Command Line
```bash
# Clone project
git clone <repository-url>
cd co3017-content-service

# Run application
mvn spring-boot:run
```

#### Cách 2: IDE (Khuyến nghị)
1. **IntelliJ IDEA**:
   - Open project folder
   - Tìm `ContentServiceApplication.java`
   - Click chuột phải → **Run 'ContentServiceApplication'**

2. **VS Code**:
   - Cài extension "Spring Boot Extension Pack"
   - Open project
   - Nhấn F5 hoặc click "Run" ở file `ContentServiceApplication.java`

3. **Eclipse**:
   - Import as Maven project
   - Right-click project → **Run As** → **Spring Boot App**

### 3️⃣ Kiểm tra Application

```bash
# Application sẽ chạy tại: http://localhost:9000

# Test API
curl http://localhost:9000/api/contents
curl http://localhost:9000/api/documents

# Test CORS
curl -H "Origin: http://localhost:3000" -X GET http://localhost:9000/api/contents
```

## 🔌 API Examples

### Content APIs

```bash
# Get all contents
GET http://localhost:9000/api/contents

# Get content by ID
GET http://localhost:9000/api/contents/1

# Get content by title
GET http://localhost:9000/api/contents/title/math-content

# Create content
POST http://localhost:9000/api/contents
Content-Type: application/json

{
  "title": "Math Content",
  "body": "This is math content body"
}
```

### Document APIs

```bash
# Get all documents
GET http://localhost:9000/api/documents

# Get document by ID
GET http://localhost:9000/api/documents/1

# Search documents by title
GET http://localhost:9000/api/documents/search?title=math

# Create document
POST http://localhost:9000/api/documents
Content-Type: application/json

{
  "title": "Math Document",
  "description": "Basic mathematics document",
  "fileSize": 1024,
  "format": "pdf"
}
```

### Response Format

Tất cả API trả về format chuẩn:

```json
{
  "errorCode": 0,
  "message": "Success",
  "data": {
    "id": 1,
    "title": "Math Content",
    "body": "This is math content body"
  }
}
```

## 🛠️ Development Guide

### Thêm Domain Mới

1. **Tạo Domain Entity**:
```java
// models/Article.java
public class Article {
    private Long id;
    private String title;
    private String content;
    
    // Business methods
    public boolean isLongContent() { return content.length() > 1000; }
}
```

2. **Tạo UseCase Interface**:
```java
// usecase/ArticleUseCase.java
public interface ArticleUseCase {
    Article createArticle(CreateArticleCommand command);
    Optional<Article> getArticleById(Long id);
    List<Article> getAllArticles();
}
```

3. **Tạo UseCase Service**:
```java
// usecase/service/ArticleService.java
@Service
public class ArticleService implements ArticleUseCase {
    private final ArticleRepository articleRepository;
    // Implementation...
}
```

4. **Tạo Repository Interface**:
```java
// repository/ArticleRepository.java
public interface ArticleRepository {
    Article save(Article article);
    Optional<Article> findById(Long id);
    List<Article> findAll();
}
```

5. **Tạo JPA Implementation**:
```java
// repository/postgresql/JpaArticleRepository.java
@Repository
public class JpaArticleRepository implements ArticleRepository {
    // JPA implementation...
}
```

6. **Tạo Controller**:
```java
// adapter/http/ArticleController.java
@RestController
@RequestMapping("/api/articles")
public class ArticleController {
    // REST endpoints...
}
```

### Testing Strategy

```java
// Test UseCase (Unit Test)
@Test
public void shouldCreateContentWithValidTitle() {
    // Given
    CreateContentCommand command = new CreateContentCommand("Math Content", "This is content body");
    
    // When
    Content content = contentService.createContent(command);
    
    // Then
    assertThat(content.getTitle()).isEqualTo("Math Content");
    assertThat(content.getBody()).isEqualTo("This is content body");
}

// Test Controller (Integration Test)
@SpringBootTest
@TestPropertySource(locations = "classpath:application-test.properties")
class ContentControllerTest {
    @Test
    public void shouldCreateContent() {
        // Test HTTP request/response
    }
}
```

## 🏭 Production Considerations

### Security
- Thay đổi CORS config từ `*` thành specific origins
- Thêm authentication/authorization
- Validate input data
- Rate limiting

### Performance
- Database connection pooling
- Caching (Redis)
- Pagination cho list APIs
- Database indexing

### Monitoring
- Health check endpoints
- Metrics collection
- Logging configuration
- Error tracking

## 📚 Tech Stack

- **Framework**: Spring Boot 3.5.6
- **Language**: Java 17
- **Build Tool**: Maven
- **Database**: PostgreSQL
- **ORM**: Spring Data JPA / Hibernate
- **Architecture**: Clean Architecture
- **CORS**: Enabled for all origins (development)
- **Utilities**: Lombok
- **gRPC**: Spring gRPC 0.11.0 (prepared)

## 📚 Documentation

- [Spring Boot Reference](https://docs.spring.io/spring-boot/docs/current/reference/htmlsingle/)
- [Clean Architecture by Uncle Bob](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Spring Data JPA](https://docs.spring.io/spring-data/jpa/docs/current/reference/html/)

## 🧪 Testing

```bash
# Run all tests
mvn test

# Run specific test class
mvn test -Dtest=ContentServiceTest

# Run with coverage
mvn test jacoco:report
```

## 📦 Build

```bash
# Clean build
mvn clean package

# Skip tests
mvn clean package -DskipTests

# Run JAR
java -jar target/content-service-0.0.1-SNAPSHOT.jar
```

Build output: `target/content-service-0.0.1-SNAPSHOT.jar`

**Happy Coding! 🚀**