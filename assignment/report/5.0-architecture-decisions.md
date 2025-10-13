# Architecture Decision Records (ADRs)

## Mục Tiêu

Ghi lại các **quyết định kiến trúc quan trọng** (Architecture Decisions) với:
- **Context:** Bối cảnh và lý do quyết định
- **Decision:** Quyết định cụ thể
- **Rationale:** Lý luận và ACs được tối ưu
- **Consequences:** Hậu quả và trade-offs
- **Status:** Trạng thái (Proposed, Accepted, Deprecated)

**Tech Stack:** Golang, Java (Spring Boot), PostgreSQL, MongoDB, Redis, Kafka

---

## 1. ADR Overview

### 1.1. Nguyên Tắc Ra Quyết Định

Các quyết định dưới đây được đưa ra để tối ưu hóa:
- **AC1: Modularity** - Tính module hóa cao
- **AC2: Scalability** - Khả năng mở rộng
- **AC3: Performance** - Hiệu suất cao
- **AC4: Testability** - Khả năng kiểm thử
- **AC7: Maintainability** - Khả năng bảo trì

### 1.2. ADR Template

Mỗi ADR tuân theo cấu trúc:

```
## ADR-{ID}: {Title}

**Status:** [Proposed | Accepted | Deprecated | Superseded]
**Date:** YYYY-MM-DD
**Deciders:** Architecture Team

### Context
[Mô tả vấn đề, yêu cầu, constraints]

### Decision
[Quyết định cụ thể]

### Rationale
[Lý do, ACs được tối ưu]

### Consequences
[Hậu quả, trade-offs]

### Alternatives Considered
[Các lựa chọn khác đã xem xét]
```

---

## 2. Architecture Decision Records

### ADR-1: Polyglot Programming Strategy

**Status:** ✅ Accepted  
**Date:** 2025-10-13  
**Deciders:** Architecture Team

#### **Context**

ITS có các services với yêu cầu khác nhau:
- **Management Services:** Cần maintainability, ecosystem rộng
- **Computation Services:** Cần performance, concurrency
- **AI/ML Services:** Cần flexibility, fast iteration

**Constraints:**
- Team có kinh nghiệm với cả Java và Golang
- Cần tối ưu performance cho real-time scoring
- Cần maintainability cao cho business logic

#### **Decision**

**Sử dụng Polyglot Programming Strategy:**

| **Service Type** | **Language** | **Framework** | **Rationale** |
|------------------|--------------|---------------|---------------|
| **User Management** | Java 17+ | Spring Boot 3.x | - Mature ecosystem<br>- Spring Security for auth<br>- JPA/Hibernate for ORM<br>- High maintainability |
| **Content Service** | Java 17+ | Spring Boot 3.x | - Complex business rules<br>- Transactional integrity<br>- Rich query support (JPA) |
| **Scoring/Feedback** | Golang 1.21+ | Gin/Echo | - **High performance** (AC3)<br>- Fast startup time<br>- Excellent concurrency<br>- Low latency (≤500ms) |
| **Adaptive Engine** | Golang 1.21+ | Custom (Clean Arch) | - CPU-intensive AI algorithms<br>- Goroutines for parallelism<br>- Fast model inference |
| **Learner Model** | Golang 1.21+ | Gin + MongoDB Driver | - Event-driven processing<br>- NoSQL-friendly<br>- High throughput |
| **API Gateway** | Golang 1.21+ | Custom/Kong | - Low latency routing<br>- Efficient request handling |

#### **Rationale**

**Java/Spring Boot Advantages:**
- ✅ **AC7: Maintainability** - Clean, readable code với annotations
- ✅ Rich ecosystem (Spring Security, Spring Data, Spring Cloud)
- ✅ Mature ORM (Hibernate/JPA) cho complex queries
- ✅ Strong typing với compile-time checks
- ✅ Excellent tooling (IntelliJ IDEA, debugging)

**Golang Advantages:**
- ✅ **AC3: Performance** - Near C-level performance
- ✅ **AC2: Scalability** - Built-in concurrency (goroutines)
- ✅ Fast compilation và deployment
- ✅ Low memory footprint (important for containers)
- ✅ Simple dependency management (Go modules)

**Mapping to Services:**

```
┌─────────────────────────────────────────────────┐
│  Java Services (Maintainability Focus)          │
│  ✓ User Management (RBAC, complex auth)         │
│  ✓ Content Service (complex queries, txns)      │
│  ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━  │
│  ACs: Maintainability, Security, Testability    │
└─────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────┐
│  Golang Services (Performance Focus)            │
│  ✓ Scoring/Feedback (real-time, ≤500ms)        │
│  ✓ Adaptive Engine (AI algorithms, CPU)         │
│  ✓ Learner Model (event processing, high load)  │
│  ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━  │
│  ACs: Performance, Scalability, Efficiency      │
└─────────────────────────────────────────────────┘
```

#### **Consequences**

**Positive:**
- ✅ Optimized performance cho real-time services
- ✅ Better maintainability cho business logic services
- ✅ Team can use strengths of each language
- ✅ Flexibility to evolve services independently

**Negative:**
- ❌ Team needs expertise in both languages
- ❌ Different tooling (Maven/Gradle vs Go modules)
- ❌ Different testing frameworks (JUnit vs Go testing)
- ❌ Increased complexity in CI/CD pipelines

**Mitigation:**
- Training sessions for cross-language knowledge
- Standardized project structure for both languages
- Shared CI/CD templates
- Common monitoring/logging formats

#### **Alternatives Considered**

1. **All Java:**
   - ❌ Lower performance for real-time services
   - ❌ Higher memory footprint (impacts scaling cost)
   - ✅ Single language expertise needed

2. **All Golang:**
   - ❌ Less mature ecosystem for complex business logic
   - ❌ No built-in dependency injection (need manual wiring)
   - ✅ Consistent tooling

3. **Chosen: Polyglot (Java + Golang):**
   - ✅ Best of both worlds
   - ✅ Optimized for each use case

---

### ADR-2: PostgreSQL as Primary Relational Database

**Status:** ✅ Accepted  
**Date:** 2025-10-13  
**Deciders:** Architecture Team

#### **Context**

Services cần relational database cho:
- User authentication & authorization (RBAC)
- Content management với complex relationships
- Transactional integrity (ACID)

**Requirements:**
- ACID compliance
- Complex queries (JOIN, aggregations)
- JSON support (flexible metadata)
- Open-source (no licensing cost)
- Mature replication & backup

#### **Decision**

**PostgreSQL 15+** làm primary relational database cho:
- User Management Service
- Content Service

**Configuration:**
- Primary-Standby replication (1 primary + 1 standby)
- Connection pooling (PgBouncer)
- WAL archiving for point-in-time recovery

#### **Rationale**

**PostgreSQL Advantages:**
- ✅ **AC6: Security** - Row-level security, advanced auth
- ✅ **AC7: Maintainability** - ACID guarantees, data integrity
- ✅ JSON/JSONB support (flexible schema)
- ✅ Rich indexing (B-tree, GIN, GiST, BRIN)
- ✅ Excellent query optimizer
- ✅ Strong community support

**vs MySQL:**
- PostgreSQL has better JSON support
- Better handling of complex queries
- More ACID-compliant

**vs NoSQL:**
- Need relational integrity for users/roles
- Need complex queries (JOIN across tables)
- Need ACID for critical data

#### **Consequences**

**Positive:**
- ✅ Strong data integrity guarantees
- ✅ Rich query capabilities
- ✅ No vendor lock-in (open-source)
- ✅ Excellent tooling (pgAdmin, DBeaver)

**Negative:**
- ❌ Vertical scaling limits (need sharding at very high scale)
- ❌ Need careful index optimization
- ❌ Schema migrations can be complex

**Mitigation:**
- Use read replicas for read-heavy workloads
- Implement proper indexing strategy
- Use connection pooling (PgBouncer)
- Monitor slow queries and optimize

#### **Implementation Details**

**Java Services (JPA/Hibernate):**
```java
// Entity
@Entity
@Table(name = "users")
public class User {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    
    @Column(nullable = false, unique = true)
    private String email;
    
    @ManyToMany(fetch = FetchType.EAGER)
    @JoinTable(name = "user_roles",
        joinColumns = @JoinColumn(name = "user_id"),
        inverseJoinColumns = @JoinColumn(name = "role_id"))
    private Set<Role> roles;
}

// Repository Interface (DIP)
public interface UserRepository extends JpaRepository<User, Long> {
    Optional<User> findByEmail(String email);
}
```

**Golang Services (pgx driver):**
```go
// Repository Interface (DIP)
type ContentRepository interface {
    FindByID(ctx context.Context, id string) (*Content, error)
    Save(ctx context.Context, content *Content) error
}

// Implementation
type PostgresContentRepository struct {
    db *pgx.Pool
}

func (r *PostgresContentRepository) FindByID(ctx context.Context, id string) (*Content, error) {
    var content Content
    err := r.db.QueryRow(ctx, 
        "SELECT id, title, body, metadata FROM contents WHERE id = $1", 
        id).Scan(&content.ID, &content.Title, &content.Body, &content.Metadata)
    return &content, err
}
```

---

### ADR-3: Clean/Hexagonal Architecture for All Services

**Status:** ✅ Accepted  
**Date:** 2025-10-13  
**Deciders:** Architecture Team

#### **Context**

Cần đảm bảo:
- **AC4: Testability** - Dễ dàng test business logic
- **AC1: Modularity** - Tách biệt concerns
- **AC7: Maintainability** - Code dễ hiểu, dễ sửa

**Problem:**
- Traditional layered architecture tạo tight coupling với framework/DB
- Khó test business logic độc lập
- Thay đổi DB/framework ảnh hưởng toàn bộ code

#### **Decision**

**Áp dụng Clean/Hexagonal Architecture cho TẤT CẢ microservices**

**Structure:**
```
service/
├── domain/           # Entities, Value Objects (innermost)
├── application/      # Use Cases, Business Logic
├── adapters/         # Interface Adapters
│   ├── http/         # REST controllers
│   ├── grpc/         # gRPC handlers
│   └── repository/   # DB implementations
└── infrastructure/   # External dependencies
```

**Dependency Rule:**
```
Infrastructure → Adapters → Application → Domain

[Outermost]                           [Innermost]
Framework/DB → Controllers → Use Cases → Entities
```

#### **Rationale**

**Clean Architecture Benefits:**
- ✅ **AC4: Testability** - Business logic testable WITHOUT DB/framework
- ✅ **AC1: Modularity** - Clear separation of concerns
- ✅ **DIP Compliance** - Dependencies point inward
- ✅ **Framework Independence** - Can change Spring → Micronaut
- ✅ **Database Independence** - Can change Postgres → MySQL

**Example Flow:**
```
HTTP Request
    ↓
Controller (Adapter)
    ↓
Use Case (Application)
    ↓
Repository Interface (Application)
    ↑ (implements)
Repository Impl (Infrastructure)
    ↓
PostgreSQL
```

#### **Consequences**

**Positive:**
- ✅ Business logic in `application` layer is pure (no framework deps)
- ✅ Can test use cases with mock repositories
- ✅ Can swap DB without touching business logic
- ✅ Clear boundaries between layers

**Negative:**
- ❌ More boilerplate (interfaces, DTOs)
- ❌ Steeper learning curve for junior devs
- ❌ More files/packages to navigate

**Mitigation:**
- Create code templates for new features
- Documentation with examples
- Code reviews to ensure consistency

#### **Implementation Examples**

**Java Service (Scoring Service):**

```java
// Domain Layer - Pure business logic
package domain;

public class Assessment {
    private final String id;
    private final String question;
    private final String correctAnswer;
    
    public Score evaluate(String userAnswer) {
        // Pure domain logic, no dependencies
        boolean isCorrect = correctAnswer.equalsIgnoreCase(userAnswer);
        return new Score(isCorrect, isCorrect ? 10 : 0);
    }
}

// Application Layer - Use Case
package application;

public interface AssessmentRepository {
    Assessment findById(String id);
}

public class SubmitAnswerUseCase {
    private final AssessmentRepository repository;
    
    // Dependency Injection (DIP)
    public SubmitAnswerUseCase(AssessmentRepository repository) {
        this.repository = repository;
    }
    
    public ScoreResult execute(SubmitAnswerRequest request) {
        Assessment assessment = repository.findById(request.getAssessmentId());
        Score score = assessment.evaluate(request.getUserAnswer());
        return new ScoreResult(score, generateHint(assessment, score));
    }
}

// Infrastructure Layer - Repository Implementation
package infrastructure;

@Repository
public class PostgresAssessmentRepository implements AssessmentRepository {
    @Autowired
    private JdbcTemplate jdbcTemplate;
    
    @Override
    public Assessment findById(String id) {
        // SQL implementation
    }
}

// Adapter Layer - Controller
package adapters;

@RestController
public class AssessmentController {
    private final SubmitAnswerUseCase useCase;
    
    @Autowired
    public AssessmentController(SubmitAnswerUseCase useCase) {
        this.useCase = useCase;
    }
    
    @PostMapping("/api/assessments/{id}/submit")
    public ResponseEntity<ScoreDTO> submitAnswer(@RequestBody SubmitAnswerDTO dto) {
        ScoreResult result = useCase.execute(toRequest(dto));
        return ResponseEntity.ok(toDTO(result));
    }
}
```

**Golang Service (Adaptive Engine):**

```go
// Domain Layer - Entity
package domain

type LearnerModel struct {
    ID            string
    SkillMastery  map[string]float64
    LearningStyle string
}

func (l *LearnerModel) GetWeakSkills() []string {
    // Pure domain logic
    var weakSkills []string
    for skill, mastery := range l.SkillMastery {
        if mastery < 0.6 {
            weakSkills = append(weakSkills, skill)
        }
    }
    return weakSkills
}

// Application Layer - Use Case
package application

// Port (Interface) - DIP
type LearnerModelRepository interface {
    FindByID(ctx context.Context, id string) (*domain.LearnerModel, error)
}

type GenerateAdaptivePathUseCase struct {
    learnerRepo LearnerModelRepository
}

func NewGenerateAdaptivePathUseCase(repo LearnerModelRepository) *GenerateAdaptivePathUseCase {
    return &GenerateAdaptivePathUseCase{learnerRepo: repo}
}

func (uc *GenerateAdaptivePathUseCase) Execute(ctx context.Context, learnerID string) (*AdaptivePath, error) {
    learner, err := uc.learnerRepo.FindByID(ctx, learnerID)
    if err != nil {
        return nil, err
    }
    
    // Business logic
    weakSkills := learner.GetWeakSkills()
    path := uc.generatePath(learner, weakSkills)
    return path, nil
}

// Infrastructure Layer - Repository Implementation
package infrastructure

type MongoLearnerRepository struct {
    client *mongo.Client
}

func (r *MongoLearnerRepository) FindByID(ctx context.Context, id string) (*domain.LearnerModel, error) {
    // MongoDB implementation
    var learner domain.LearnerModel
    err := r.client.Database("its").Collection("learners").
        FindOne(ctx, bson.M{"_id": id}).Decode(&learner)
    return &learner, err
}

// Adapter Layer - HTTP Handler
package adapters

type AdaptiveEngineHandler struct {
    useCase *application.GenerateAdaptivePathUseCase
}

func (h *AdaptiveEngineHandler) HandleGeneratePath(c *gin.Context) {
    learnerID := c.Param("learnerId")
    
    path, err := h.useCase.Execute(c.Request.Context(), learnerID)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(200, toDTO(path))
}
```

**Testing Benefits:**

```go
// Unit Test WITHOUT database
func TestGenerateAdaptivePathUseCase(t *testing.T) {
    // Mock repository (implements interface)
    mockRepo := &MockLearnerRepository{
        learner: &domain.LearnerModel{
            ID: "123",
            SkillMastery: map[string]float64{
                "algebra": 0.3, // weak
                "calculus": 0.9, // strong
            },
        },
    }
    
    useCase := application.NewGenerateAdaptivePathUseCase(mockRepo)
    
    path, err := useCase.Execute(context.Background(), "123")
    
    assert.NoError(t, err)
    assert.Contains(t, path.RecommendedTopics, "algebra") // focuses on weak skill
}
```

---

### ADR-4: Repository Pattern with Interface Abstraction

**Status:** ✅ Accepted  
**Date:** 2025-10-13  
**Deciders:** Architecture Team

#### **Context**

**Problem:**
- Business logic trực tiếp phụ thuộc vào ORM/database driver
- Khó test business logic (cần DB running)
- Khó đổi database technology

**Goals:**
- **AC1: Modularity** - Decouple business logic from data access
- **AC4: Testability** - Easy to mock data layer
- **DIP Compliance** - High-level modules don't depend on low-level modules

#### **Decision**

**Implement Repository Pattern với Interface Abstraction:**

1. **Define Repository Interfaces** trong `application` layer
2. **Implement Repositories** trong `infrastructure` layer
3. **Use Dependency Injection** để inject implementations

**Rules:**
- ✅ Use Cases chỉ phụ thuộc vào Repository Interfaces
- ✅ Repository Interfaces không có dependencies (pure)
- ✅ Implementations có thể dùng ORM (Hibernate, GORM) hoặc raw SQL
- ❌ Use Cases KHÔNG BAO GIỜ import ORM libraries

#### **Rationale**

**Benefits:**
- ✅ **Testability:** Mock repositories trong tests
- ✅ **Flexibility:** Đổi DB không ảnh hưởng business logic
- ✅ **DIP:** Dependencies point inward
- ✅ **Single Source of Truth:** Tất cả data access qua repositories

**vs Direct ORM Usage:**
- Direct ORM: Business logic coupled to ORM entities
- Repository Pattern: Business logic coupled to domain entities only

#### **Consequences**

**Positive:**
- ✅ Business logic testable without DB
- ✅ Can swap database implementations
- ✅ Clear data access boundaries

**Negative:**
- ❌ More interfaces to maintain
- ❌ Mapping between domain entities and DB entities

**Mitigation:**
- Use mapping libraries (MapStruct for Java, manual mapping for Go)
- Code generation for boilerplate

#### **Implementation Guidelines**

**Java (JPA/Hibernate):**

```java
// Application Layer - Repository Interface (Port)
package application.ports;

public interface UserRepository {
    Optional<User> findById(Long id);
    User save(User user);
    List<User> findByRole(String role);
}

// Infrastructure Layer - Implementation (Adapter)
package infrastructure.persistence;

@Repository
public class JpaUserRepository implements UserRepository {
    @Autowired
    private UserJpaRepository jpaRepository; // Spring Data JPA
    
    @Override
    public Optional<User> findById(Long id) {
        return jpaRepository.findById(id)
            .map(this::toDomain); // Convert JPA entity → Domain entity
    }
    
    @Override
    public User save(User user) {
        UserEntity entity = toEntity(user); // Domain → JPA entity
        UserEntity saved = jpaRepository.save(entity);
        return toDomain(saved);
    }
    
    // Mapping methods
    private User toDomain(UserEntity entity) {
        return new User(entity.getId(), entity.getEmail(), entity.getRoles());
    }
    
    private UserEntity toEntity(User user) {
        UserEntity entity = new UserEntity();
        entity.setId(user.getId());
        entity.setEmail(user.getEmail());
        return entity;
    }
}

// Spring Data JPA Repository (internal to infrastructure)
interface UserJpaRepository extends JpaRepository<UserEntity, Long> {
    List<UserEntity> findByRolesContaining(String role);
}
```

**Golang (pgx/MongoDB):**

```go
// Application Layer - Repository Interface (Port)
package application

type ContentRepository interface {
    FindByID(ctx context.Context, id string) (*domain.Content, error)
    Save(ctx context.Context, content *domain.Content) error
    FindByTags(ctx context.Context, tags []string) ([]*domain.Content, error)
}

// Infrastructure Layer - PostgreSQL Implementation (Adapter)
package infrastructure

type PostgresContentRepository struct {
    pool *pgxpool.Pool
}

func NewPostgresContentRepository(pool *pgxpool.Pool) *PostgresContentRepository {
    return &PostgresContentRepository{pool: pool}
}

func (r *PostgresContentRepository) FindByID(ctx context.Context, id string) (*domain.Content, error) {
    var row ContentRow
    err := r.pool.QueryRow(ctx,
        "SELECT id, title, body, created_at FROM contents WHERE id = $1",
        id,
    ).Scan(&row.ID, &row.Title, &row.Body, &row.CreatedAt)
    
    if err != nil {
        return nil, err
    }
    
    return r.toDomain(&row), nil
}

func (r *PostgresContentRepository) Save(ctx context.Context, content *domain.Content) error {
    _, err := r.pool.Exec(ctx,
        "INSERT INTO contents (id, title, body) VALUES ($1, $2, $3) ON CONFLICT (id) DO UPDATE SET title = $2, body = $3",
        content.ID, content.Title, content.Body,
    )
    return err
}

// Mapping: DB row → Domain entity
func (r *PostgresContentRepository) toDomain(row *ContentRow) *domain.Content {
    return &domain.Content{
        ID:    row.ID,
        Title: row.Title,
        Body:  row.Body,
    }
}

// Infrastructure Layer - MongoDB Implementation (Alternative)
package infrastructure

type MongoContentRepository struct {
    client *mongo.Client
}

func (r *MongoContentRepository) FindByID(ctx context.Context, id string) (*domain.Content, error) {
    var doc ContentDocument
    err := r.client.Database("its").Collection("contents").
        FindOne(ctx, bson.M{"_id": id}).Decode(&doc)
    
    if err != nil {
        return nil, err
    }
    
    return r.toDomain(&doc), nil
}
```

**Dependency Injection (Wire-up):**

```go
// main.go or dependency injection container
func main() {
    // Infrastructure setup
    pgPool := setupPostgresPool()
    
    // Repositories (Infrastructure)
    contentRepo := infrastructure.NewPostgresContentRepository(pgPool)
    
    // Use Cases (Application)
    getContentUC := application.NewGetContentUseCase(contentRepo)
    
    // Handlers (Adapters)
    handler := adapters.NewContentHandler(getContentUC)
    
    // Setup routes
    router := gin.Default()
    router.GET("/api/contents/:id", handler.GetContent)
    router.Run(":8080")
}
```

---

## 3. Design Principles Extension (SOLID)

### 3.1. Single Responsibility Principle (SRP)

**Definition:** Một class/module chỉ có **một lý do để thay đổi**.

#### **Application in ITS:**

**Service Level (DDD):**
- ✅ User Management Service: Chỉ quản lý users/roles
- ✅ Scoring Service: Chỉ chấm điểm và generate hints
- ❌ KHÔNG: User Service quản lý users + scoring logic

**Class Level:**

**Java Example:**
```java
// ❌ BAD: Violation of SRP
public class UserService {
    public User createUser(UserDTO dto) {
        // User creation logic
    }
    
    public void sendWelcomeEmail(User user) {
        // Email sending logic (DIFFERENT responsibility)
    }
    
    public void logUserActivity(User user, String action) {
        // Logging logic (DIFFERENT responsibility)
    }
}

// ✅ GOOD: Separated responsibilities
public class UserService {
    private final EmailService emailService;
    private final AuditLogger auditLogger;
    
    public User createUser(UserDTO dto) {
        User user = // create user
        emailService.sendWelcome(user); // Delegate
        auditLogger.log(user, "USER_CREATED"); // Delegate
        return user;
    }
}

public class EmailService {
    public void sendWelcome(User user) {
        // Email logic
    }
}

public class AuditLogger {
    public void log(User user, String action) {
        // Logging logic
    }
}
```

**Golang Example:**
```go
// ❌ BAD
type ScoringService struct {
    db *sql.DB
}

func (s *ScoringService) CalculateScore(answer string) int {
    // Scoring logic
}

func (s *ScoringService) SaveToDatabase(score int) error {
    // DB logic (DIFFERENT responsibility)
}

func (s *ScoringService) SendNotification(score int) error {
    // Notification logic (DIFFERENT responsibility)
}

// ✅ GOOD
type ScoringService struct {
    repository ScoreRepository
    notifier   Notifier
}

func (s *ScoringService) CalculateScore(answer string) int {
    // Pure scoring logic
    return score
}

type ScoreRepository interface {
    Save(score int) error
}

type Notifier interface {
    Notify(score int) error
}
```

---

### 3.2. Dependency Inversion Principle (DIP)

**Definition:** 
- High-level modules không phụ thuộc vào low-level modules
- Cả hai phụ thuộc vào **abstractions (interfaces)**

#### **Application in ITS:**

**Java (Spring Framework):**
```java
// High-level module (Use Case)
package application;

public class GeneratePathUseCase {
    private final LearnerRepository repository; // Interface (abstraction)
    
    // Constructor Injection (Spring DI)
    @Autowired
    public GeneratePathUseCase(LearnerRepository repository) {
        this.repository = repository;
    }
    
    public AdaptivePath execute(String learnerId) {
        Learner learner = repository.findById(learnerId);
        // Business logic
        return generatePath(learner);
    }
}

// Abstraction (Port)
package application.ports;

public interface LearnerRepository {
    Learner findById(String id);
}

// Low-level module (Adapter)
package infrastructure;

@Repository
public class MongoLearnerRepository implements LearnerRepository {
    @Autowired
    private MongoTemplate mongoTemplate;
    
    @Override
    public Learner findById(String id) {
        // MongoDB implementation
    }
}

// Spring Configuration (Wiring)
@Configuration
public class AppConfig {
    @Bean
    public LearnerRepository learnerRepository(MongoTemplate template) {
        return new MongoLearnerRepository(template);
    }
    
    @Bean
    public GeneratePathUseCase generatePathUseCase(LearnerRepository repo) {
        return new GeneratePathUseCase(repo);
    }
}
```

**Golang (Manual DI):**
```go
// High-level module (Use Case)
package application

type GeneratePathUseCase struct {
    repo LearnerRepository // Interface (abstraction)
}

// Constructor (manual DI)
func NewGeneratePathUseCase(repo LearnerRepository) *GeneratePathUseCase {
    return &GeneratePathUseCase{repo: repo}
}

func (uc *GeneratePathUseCase) Execute(ctx context.Context, learnerID string) (*AdaptivePath, error) {
    learner, err := uc.repo.FindByID(ctx, learnerID)
    // Business logic
    return generatePath(learner), nil
}

// Abstraction (Port)
type LearnerRepository interface {
    FindByID(ctx context.Context, id string) (*Learner, error)
}

// Low-level module (Adapter)
package infrastructure

type MongoLearnerRepository struct {
    client *mongo.Client
}

func (r *MongoLearnerRepository) FindByID(ctx context.Context, id string) (*Learner, error) {
    // MongoDB implementation
}

// Wiring (main.go)
func main() {
    // Infrastructure
    mongoClient := setupMongo()
    
    // Repositories (low-level)
    learnerRepo := infrastructure.NewMongoLearnerRepository(mongoClient)
    
    // Use Cases (high-level) - Injected with abstraction
    generatePathUC := application.NewGeneratePathUseCase(learnerRepo)
    
    // Handlers
    handler := NewHandler(generatePathUC)
}
```

**Benefits:**
- ✅ Use Case testable WITHOUT MongoDB (use mock)
- ✅ Can swap MongoDB → PostgreSQL without changing Use Case
- ✅ Use Case code is stable (I ≈ 0)

---

### 3.3. Open/Closed Principle (OCP)

**Definition:** Software entities should be **open for extension** but **closed for modification**.

#### **Application: Adaptive Path Generator**

**Scenario:** Need to add new path generation algorithm (V2) without modifying existing code.

**Golang Implementation:**
```go
// Abstraction (Open for extension)
package domain

type PathGenerator interface {
    Generate(ctx context.Context, learner *Learner) (*AdaptivePath, error)
}

// V1 Implementation (Closed for modification)
type BasicPathGenerator struct {
    contentRepo ContentRepository
}

func (g *BasicPathGenerator) Generate(ctx context.Context, learner *Learner) (*AdaptivePath, error) {
    // V1 algorithm: Simple skill-based
    weakSkills := learner.GetWeakSkills()
    return &AdaptivePath{
        RecommendedContent: g.findContentForSkills(weakSkills),
    }, nil
}

// V2 Implementation (Extension, no modification of V1)
type AIPathGenerator struct {
    mlModel      MLModel
    contentRepo  ContentRepository
}

func (g *AIPathGenerator) Generate(ctx context.Context, learner *Learner) (*AdaptivePath, error) {
    // V2 algorithm: ML-based prediction
    prediction := g.mlModel.Predict(learner)
    return &AdaptivePath{
        RecommendedContent: g.findContentForPrediction(prediction),
        Confidence:         prediction.Confidence,
    }, nil
}

// Use Case (works with any PathGenerator)
package application

type GeneratePathUseCase struct {
    generator PathGenerator // Interface (OCP)
}

func (uc *GeneratePathUseCase) Execute(ctx context.Context, learnerID string) (*AdaptivePath, error) {
    learner := // fetch learner
    return uc.generator.Generate(ctx, learner) // Delegate to generator
}

// Configuration (switch implementation WITHOUT changing use case)
func main() {
    // Can switch between V1 and V2 via config
    var generator domain.PathGenerator
    
    if config.UseAI {
        generator = &domain.AIPathGenerator{...}
    } else {
        generator = &domain.BasicPathGenerator{...}
    }
    
    useCase := application.NewGeneratePathUseCase(generator)
}
```

**Benefits:**
- ✅ Add new algorithms without modifying Use Case
- ✅ Easy A/B testing (V1 vs V2)
- ✅ Supports Blue/Green deployment (FR9)

---

### 3.4. Interface Segregation Principle (ISP)

**Definition:** Clients should not be forced to depend on interfaces they don't use.

#### **Problem: "Fat Interface"**

```go
// ❌ BAD: Fat Interface
type LearnerService interface {
    GetProfile(id string) (*Profile, error)
    UpdateProfile(id string, data *ProfileData) error
    GetSkillMastery(id string) (map[string]float64, error)
    UpdateSkillMastery(id string, skill string, score float64) error
    GetLearningHistory(id string) ([]*Event, error)
    GenerateReport(id string) (*Report, error)
    SendEmail(id string, subject, body string) error // ???
}

// Scoring Service only needs skill mastery, but forced to implement ALL methods
```

#### **Solution: Segregated Interfaces**

```go
// ✅ GOOD: Segregated Interfaces
type ProfileReader interface {
    GetProfile(id string) (*Profile, error)
}

type ProfileWriter interface {
    UpdateProfile(id string, data *ProfileData) error
}

type SkillMasteryReader interface {
    GetSkillMastery(id string) (map[string]float64, error)
}

type SkillMasteryWriter interface {
    UpdateSkillMastery(id string, skill string, score float64) error
}

// Scoring Service ONLY depends on what it needs
type ScoringService struct {
    skillReader SkillMasteryReader // Small, focused interface
}

// Adaptive Engine depends on different interface
type AdaptiveEngine struct {
    profileReader ProfileReader
    skillReader   SkillMasteryReader
}

// Implementation can implement MULTIPLE interfaces
type LearnerRepository struct {
    db *sql.DB
}

func (r *LearnerRepository) GetProfile(id string) (*Profile, error) { ... }
func (r *LearnerRepository) UpdateProfile(id string, data *ProfileData) error { ... }
func (r *LearnerRepository) GetSkillMastery(id string) (map[string]float64, error) { ... }
func (r *LearnerRepository) UpdateSkillMastery(id string, skill string, score float64) error { ... }

// This single implementation satisfies all 4 interfaces
var _ ProfileReader = (*LearnerRepository)(nil)
var _ ProfileWriter = (*LearnerRepository)(nil)
var _ SkillMasteryReader = (*LearnerRepository)(nil)
var _ SkillMasteryWriter = (*LearnerRepository)(nil)
```

**Java Example:**
```java
// ❌ BAD
public interface UserRepository {
    User findById(Long id);
    void save(User user);
    void delete(Long id);
    List<User> findAll();
    void sendEmail(User user, String message); // Unrelated!
}

// ✅ GOOD
public interface UserReader {
    User findById(Long id);
    List<User> findAll();
}

public interface UserWriter {
    void save(User user);
    void delete(Long id);
}

// Read-only service only depends on UserReader
public class UserQueryService {
    private final UserReader userReader;
    
    public UserQueryService(UserReader userReader) {
        this.userReader = userReader;
    }
}

// Write service depends on UserWriter
public class UserCommandService {
    private final UserWriter userWriter;
    
    public UserCommandService(UserWriter userWriter) {
        this.userWriter = userWriter;
    }
}
```

---

## 4. Summary & Next Steps

### 4.1. Key Decisions

| **ADR** | **Decision** | **Primary ACs** |
|---------|--------------|-----------------|
| **ADR-1** | Polyglot: Java (maintainability) + Golang (performance) | AC3, AC7 |
| **ADR-2** | PostgreSQL for relational data | AC6, AC7 |
| **ADR-3** | Clean/Hexagonal Architecture for all services | AC1, AC4 |
| **ADR-4** | Repository Pattern with Interface Abstraction | AC1, AC4 |

### 4.2. SOLID Principles Applied

| **Principle** | **Java Implementation** | **Golang Implementation** |
|---------------|-------------------------|---------------------------|
| **SRP** | Spring Beans với single responsibility | Separate structs/packages |
| **OCP** | Strategy pattern với interfaces | Interface-based extension |
| **LSP** | Inheritance + interfaces | Interface implementation |
| **ISP** | Segregated interfaces | Small, focused interfaces |
| **DIP** | Spring DI (Constructor Injection) | Manual DI with constructors |

### 4.3. Benefits Achieved

✅ **Testability (AC4):**
- Business logic testable without DB/framework
- Mock repositories via interfaces

✅ **Modularity (AC1):**
- Clear boundaries between layers
- Services can evolve independently

✅ **Performance (AC3):**
- Golang for CPU-intensive services
- Java for maintainable business logic

✅ **Maintainability (AC7):**
- Clean Architecture makes code readable
- SOLID principles ensure flexible design

---

**Tài liệu tham khảo:**
- Clean Architecture (Robert C. Martin)
- Domain-Driven Design (Eric Evans)
- Architecture Decision Records (Michael Nygard)
- Go Design Patterns (Mario Castro Contreras)
- Spring Boot in Action (Craig Walls)
