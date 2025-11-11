# Áp Dụng SOLID Principles Chi Tiết Cho ITS

## Mục Tiêu

Đảm bảo **Modularity (AC1)** và **Testability (AC4)** bằng cách tuân thủ các nguyên tắc SOLID, đặc biệt là **DIP (Dependency Inversion Principle)** - phụ thuộc vào abstraction, không phải concretion.

**Tech Stack:** Java (Spring Boot), Golang

**SOLID Principles:**
- **S**ingle Responsibility Principle (SRP)
- **O**pen/Closed Principle (OCP)
- **L**iskov Substitution Principle (LSP)
- **I**nterface Segregation Principle (ISP)
- **D**ependency Inversion Principle (DIP)

---

## 1. Single Responsibility Principle (SRP)

### 1.1. Định Nghĩa

> "A class should have one, and only one, reason to change."

Một module (class/service) chỉ nên có **một lý do để thay đổi**. Trong kiến trúc Microservices, điều này mở rộng thành một Service chỉ chịu trách nhiệm cho **một Domain Aggregate**.

### 1.2. Áp Dụng Trong ITS

#### **Ví Dụ 1: AI Computation Service**

| **Aspect** | **Problem (Violation)** | **Solution (SRP Applied)** |
|------------|-------------------------|----------------------------|
| **Scenario** | `LearnerModelService` chứa cả:<br>- Logic quản lý Profile (CRUD)<br>- Logic cập nhật Skill Mastery Score<br>- Logic AI computation | **Tách thành 2 services:**<br>1. **User Management Service** (Java)<br>   - Quản lý `LearnerProfile`<br>   - CRUD operations<br>2. **Learner Model Service** (Golang)<br>   - Chỉ quản lý `SkillMasteryScore`<br>   - AI/ML computations |
| **Reason to Change** | Violation: Changes in profile schema OR AI algorithm<br>→ 2 reasons to change | SRP: Each service has 1 reason:<br>- Profile schema changes<br>- AI algorithm changes |
| **ACs Supported** | - | **AC1:** Modularity<br>**AC2:** Independent scaling |

**Code Example - Violation:**

```java
// ❌ BAD: Multiple responsibilities
@Service
public class LearnerModelService {
    @Autowired
    private LearnerRepository learnerRepo;
    
    // Responsibility 1: Profile management
    public Learner createProfile(String name, String email) {
        Learner learner = new Learner(name, email);
        return learnerRepo.save(learner);
    }
    
    public Learner updateProfile(Long id, ProfileDTO dto) {
        // Profile update logic
    }
    
    // Responsibility 2: Skill mastery calculation
    public void updateSkillMastery(Long learnerId, String skill, double score) {
        // AI logic to calculate mastery
        double mastery = calculateMastery(score); // Complex AI algorithm
        learnerRepo.updateMastery(learnerId, skill, mastery);
    }
    
    // Responsibility 3: Analytics
    public Report generateLearningReport(Long learnerId) {
        // Reporting logic
    }
}
```

**Code Example - SRP Applied:**

```java
// ✅ GOOD: Separated responsibilities

// Service 1: Profile Management (User Management Service)
@Service
public class LearnerProfileService {
    @Autowired
    private LearnerProfileRepository repository;
    
    public LearnerProfile createProfile(String name, String email) {
        LearnerProfile profile = new LearnerProfile(name, email);
        return repository.save(profile);
    }
    
    public LearnerProfile updateProfile(Long id, ProfileUpdateDTO dto) {
        // Only profile management logic
        LearnerProfile profile = repository.findById(id)
            .orElseThrow(() -> new NotFoundException());
        profile.setName(dto.getName());
        profile.setEmail(dto.getEmail());
        return repository.save(profile);
    }
}

// Service 2: Skill Mastery (Learner Model Service - Golang)
package application

type UpdateSkillMasteryUseCase struct {
    repository LearnerModelRepository
    calculator SkillMasteryCalculator // AI logic separated
}

func (uc *UpdateSkillMasteryUseCase) Execute(ctx context.Context, req UpdateRequest) error {
    // Only skill mastery update logic
    mastery := uc.calculator.Calculate(req.Score)
    return uc.repository.UpdateMastery(ctx, req.LearnerID, req.Skill, mastery)
}

// Service 3: Analytics (Separate service)
@Service
public class LearningAnalyticsService {
    public Report generateReport(Long learnerId) {
        // Only reporting logic
    }
}
```

---

#### **Ví Dụ 2: Scoring & Storage Separation**

| **Aspect** | **Problem** | **Solution** |
|------------|-------------|--------------|
| **Scenario** | `QuizController` vừa:<br>- Xử lý HTTP Request<br>- Gọi logic chấm điểm<br>- Gọi trực tiếp PostgreSQL | **Tách theo Clean Architecture:**<br>- `ScoringController` (HTTP)<br>- `ScoringUseCase` (Business Logic)<br>- `ResultRepositoryImpl` (DB) |
| **Stack** | Golang với interfaces | Golang with Clean Architecture |

**Code Example:**

```go
// ❌ BAD: Controller does everything
type QuizController struct {
    db *sql.DB
}

func (c *QuizController) SubmitAnswer(ctx *gin.Context) {
    var req SubmitRequest
    ctx.BindJSON(&req)
    
    // HTTP handling + Business logic + DB access all mixed
    score := c.calculateScore(req.Answer) // Business logic in controller
    
    // Direct DB access in controller
    _, err := c.db.Exec("INSERT INTO scores (user_id, score) VALUES (?, ?)", 
        req.UserID, score)
    
    ctx.JSON(200, gin.H{"score": score})
}

// ✅ GOOD: Separated by Clean Architecture layers

// Controller (Adapter Layer) - Only HTTP handling
type ScoringController struct {
    useCase *application.ScoringUseCase
}

func (c *ScoringController) SubmitAnswer(ctx *gin.Context) {
    var req SubmitAnswerDTO
    if err := ctx.BindJSON(&req); err != nil {
        ctx.JSON(400, gin.H{"error": err.Error()})
        return
    }
    
    // Delegate to use case
    result, err := c.useCase.Execute(ctx.Request.Context(), toRequest(req))
    if err != nil {
        ctx.JSON(500, gin.H{"error": err.Error()})
        return
    }
    
    ctx.JSON(200, toDTO(result))
}

// Use Case (Application Layer) - Business logic only
package application

type ScoringUseCase struct {
    repository ScoreRepository // Interface (DIP)
}

func (uc *ScoringUseCase) Execute(ctx context.Context, req SubmitRequest) (*ScoreResult, error) {
    // Pure business logic
    score := uc.calculateScore(req.Answer)
    
    // Delegate storage to repository
    if err := uc.repository.Save(ctx, score); err != nil {
        return nil, err
    }
    
    return &ScoreResult{Score: score}, nil
}

// Repository (Infrastructure Layer) - DB access only
package infrastructure

type PostgresScoreRepository struct {
    db *sql.DB
}

func (r *PostgresScoreRepository) Save(ctx context.Context, score Score) error {
    _, err := r.db.ExecContext(ctx, 
        "INSERT INTO scores (user_id, score) VALUES ($1, $2)",
        score.UserID, score.Value)
    return err
}
```

---

#### **Ví Dụ 3: Reporting Separation**

```java
// ❌ BAD
@Service
public class InstructorService {
    public CohortReport generateCohortReport(Long cohortId) { }
    public IndividualReport generateIndividualReport(Long studentId) { }
    public ComparisonReport generateComparisonReport(...) { }
    // Too many responsibilities
}

// ✅ GOOD
@Service
public class CohortReportGenerator {
    public CohortReport generate(Long cohortId) {
        // Only cohort-level reporting
    }
}

@Service
public class IndividualDiagnosticReportGenerator {
    public IndividualReport generate(Long studentId) {
        // Only individual-level reporting
    }
}
```

---

## 2. Open/Closed Principle (OCP)

### 2.1. Định Nghĩa

> "Software entities should be open for extension but closed for modification."

Một thực thể nên **mở để mở rộng** nhưng **đóng để sửa đổi**. Điều này rất quan trọng cho các **Policy Modules** như thuật toán AI.

### 2.2. Áp Dụng Trong ITS

#### **Ví Dụ 1: Hint Generation Strategy**

| **Aspect** | **Problem** | **Solution** |
|------------|-------------|--------------|
| **Scenario** | `HintGenerator` sử dụng if/else:<br>```if (type == 'syntax') { ... }```<br>```else if (type == 'concept') { ... }``` | **Strategy Pattern:**<br>- `HintStrategy` interface<br>- Multiple implementations |
| **Adding New Type** | Violation: Phải sửa `HintGenerator`<br>→ Risk of breaking existing logic | OCP: Chỉ cần add new strategy class<br>→ Existing code untouched |

**Code Example - Violation:**

```go
// ❌ BAD: if/else chain (closed for extension)
type HintGenerator struct{}

func (h *HintGenerator) GenerateHint(hintType string, context Context) string {
    if hintType == "syntax" {
        return "Check your syntax at line " + context.Line
    } else if hintType == "concept" {
        return "Review the concept of " + context.Topic
    } else if hintType == "example" {
        return "Here's an example: " + context.Example
    }
    // Adding new type requires modifying this function
    return "No hint available"
}
```

**Code Example - OCP Applied:**

```go
// ✅ GOOD: Strategy Pattern (open for extension)

// Interface (Abstraction)
type HintStrategy interface {
    Generate(context Context) string
}

// Implementation 1
type SyntaxHintStrategy struct{}

func (s *SyntaxHintStrategy) Generate(ctx Context) string {
    return "Check your syntax at line " + ctx.Line
}

// Implementation 2
type ConceptHintStrategy struct{}

func (c *ConceptHintStrategy) Generate(ctx Context) string {
    return "Review the concept of " + ctx.Topic
}

// Implementation 3 (NEW - added without modifying existing code)
type ExampleHintStrategy struct{}

func (e *ExampleHintStrategy) Generate(ctx Context) string {
    return "Here's an example: " + ctx.Example
}

// Use Case (closed for modification, open for extension)
type GenerateHintUseCase struct {
    strategy HintStrategy // Interface
}

func (uc *GenerateHintUseCase) Execute(ctx Context) string {
    return uc.strategy.Generate(ctx) // Delegates to strategy
}

// Configuration (dependency injection)
func NewGenerateHintUseCase(strategyType string) *GenerateHintUseCase {
    var strategy HintStrategy
    
    switch strategyType {
    case "syntax":
        strategy = &SyntaxHintStrategy{}
    case "concept":
        strategy = &ConceptHintStrategy{}
    case "example":
        strategy = &ExampleHintStrategy{} // NEW strategy easily added
    default:
        strategy = &ConceptHintStrategy{}
    }
    
    return &GenerateHintUseCase{strategy: strategy}
}
```

---

#### **Ví Dụ 2: Adaptive Path Generation (FR9 - Live Model Swapping)**

| **Aspect** | **Problem** | **Solution** |
|------------|-------------|--------------|
| **Scenario** | Cần thêm thuật toán AI mới<br>(AdaptivePathGenerator V2) | **Factory Pattern + OCP:**<br>Use `PathGenerationAlgorithm` interface |
| **Blue/Green Deploy** | Violation: Hard to swap algorithms<br>→ Downtime required | OCP: Easy to swap V1 ↔ V2<br>→ Zero downtime (FR9) |

**Code Example:**

```java
// ✅ GOOD: OCP for AI algorithm swapping

// Interface (open for extension)
public interface PathGenerationAlgorithm {
    AdaptivePath generate(LearnerModel learner, List<Content> availableContent);
}

// V1 Implementation (closed for modification)
@Component("pathGenV1")
public class BasicPathGenerator implements PathGenerationAlgorithm {
    @Override
    public AdaptivePath generate(LearnerModel learner, List<Content> content) {
        // V1 algorithm: Rule-based
        List<String> weakSkills = learner.getWeakSkills();
        return new AdaptivePath(filterContentBySkills(content, weakSkills));
    }
}

// V2 Implementation (NEW - extension without modification)
@Component("pathGenV2")
public class AIPathGenerator implements PathGenerationAlgorithm {
    @Autowired
    private MLModel mlModel;
    
    @Override
    public AdaptivePath generate(LearnerModel learner, List<Content> content) {
        // V2 algorithm: ML-based
        Prediction prediction = mlModel.predict(learner);
        return new AdaptivePath(
            filterContentByPrediction(content, prediction),
            prediction.getConfidence()
        );
    }
}

// Use Case (closed for modification)
@Service
public class GenerateAdaptivePathUseCase {
    private final PathGenerationAlgorithm algorithm;
    
    // Dependency Injection (can inject V1 or V2)
    @Autowired
    public GenerateAdaptivePathUseCase(
        @Qualifier("pathGenV2") PathGenerationAlgorithm algorithm // V2 active
    ) {
        this.algorithm = algorithm;
    }
    
    public AdaptivePath execute(Long learnerId) {
        LearnerModel learner = // fetch learner
        List<Content> content = // fetch content
        return algorithm.generate(learner, content); // OCP in action
    }
}

// Configuration (Blue/Green switching)
@Configuration
public class AlgorithmConfig {
    @Value("${adaptive.algorithm.version:v2}") // Configurable
    private String version;
    
    @Bean
    public PathGenerationAlgorithm pathGenerationAlgorithm(
        @Qualifier("pathGenV1") PathGenerationAlgorithm v1,
        @Qualifier("pathGenV2") PathGenerationAlgorithm v2
    ) {
        return "v2".equals(version) ? v2 : v1; // Easy switching
    }
}
```

---

## 3. Liskov Substitution Principle (LSP)

### 3.1. Định Nghĩa

> "Objects of a superclass should be replaceable with objects of a subclass without breaking the application."

Lớp con phải **thay thế được** lớp cha mà không phá vỡ logic chương trình. Điều này đảm bảo **polymorphism** đáng tin cậy.

### 3.2. Áp Dụng Trong ITS

#### **Ví Dụ 1: Assessment Types**

| **Aspect** | **Problem (LSP Violation)** | **Solution (LSP Applied)** |
|------------|----------------------------|----------------------------|
| **Scenario** | `ProjectAssessment extends QuizAssessment`<br>Nhưng `autoScore()` ném exception<br>(vì project cần chấm thủ công) | **Tách interfaces:**<br>- `ScorableAssessment` (auto)<br>- `ManualAssessment` (manual) |
| **Issue** | Code expecting `QuizAssessment`<br>breaks with `ProjectAssessment` | Each type implements<br>appropriate interface |

**Code Example - Violation:**

```java
// ❌ BAD: LSP violation

// Base class
public abstract class Assessment {
    public abstract double autoScore(String answer);
}

// Subclass 1: Works fine
public class QuizAssessment extends Assessment {
    @Override
    public double autoScore(String answer) {
        return answer.equals(correctAnswer) ? 10.0 : 0.0;
    }
}

// Subclass 2: VIOLATION - breaks LSP
public class ProjectAssessment extends Assessment {
    @Override
    public double autoScore(String answer) {
        throw new UnsupportedOperationException(
            "Projects require manual grading"); // LSP violation!
    }
}

// Client code breaks when using ProjectAssessment
public class GradingService {
    public void gradeAll(List<Assessment> assessments) {
        for (Assessment assessment : assessments) {
            // Will throw exception for ProjectAssessment
            double score = assessment.autoScore(userAnswer);
        }
    }
}
```

**Code Example - LSP Applied:**

```java
// ✅ GOOD: Segregated interfaces

// Base interface (common behavior)
public interface Assessment {
    String getTitle();
    String getDescription();
}

// Interface for auto-gradable assessments
public interface ScorableAssessment extends Assessment {
    double autoScore(String answer);
}

// Interface for manual assessments
public interface ManualAssessment extends Assessment {
    void sendForReview(String submission, Instructor reviewer);
}

// Quiz implements ScorableAssessment
public class QuizAssessment implements ScorableAssessment {
    @Override
    public double autoScore(String answer) {
        return answer.equals(correctAnswer) ? 10.0 : 0.0;
    }
    
    @Override
    public String getTitle() { return title; }
    
    @Override
    public String getDescription() { return description; }
}

// Project implements ManualAssessment
public class ProjectAssessment implements ManualAssessment {
    @Override
    public void sendForReview(String submission, Instructor reviewer) {
        // Send for manual review
    }
    
    @Override
    public String getTitle() { return title; }
    
    @Override
    public String getDescription() { return description; }
}

// Client code is type-safe
public class GradingService {
    public void gradeAutomatic(List<ScorableAssessment> assessments) {
        for (ScorableAssessment assessment : assessments) {
            double score = assessment.autoScore(userAnswer); // Safe
        }
    }
    
    public void gradManual(List<ManualAssessment> assessments) {
        for (ManualAssessment assessment : assessments) {
            assessment.sendForReview(submission, instructor); // Safe
        }
    }
}
```

---

#### **Ví Dụ 2: Repository Interface**

```go
// ❌ BAD: LSP violation

type Repository interface {
    FindByID(id string) (*Entity, error)
    FindBatch(ids []string) ([]*Entity, error) // Not all DBs support batch
}

// PostgreSQL implementation - OK
type PostgresRepository struct{}

func (r *PostgresRepository) FindBatch(ids []string) ([]*Entity, error) {
    // Uses SQL "WHERE id IN (...)" - works fine
}

// Redis implementation - VIOLATION
type RedisRepository struct{}

func (r *RedisRepository) FindBatch(ids []string) ([]*Entity, error) {
    return nil, errors.New("batch not supported") // LSP violation!
}

// ✅ GOOD: Only common operations in interface

type Repository interface {
    FindByID(id string) (*Entity, error)
    // Only methods ALL implementations can support
}

// Batch capability is optional extension
type BatchRepository interface {
    Repository
    FindBatch(ids []string) ([]*Entity, error)
}

// PostgreSQL implements both
type PostgresRepository struct{}

func (r *PostgresRepository) FindByID(id string) (*Entity, error) { ... }
func (r *PostgresRepository) FindBatch(ids []string) ([]*Entity, error) { ... }

// Redis implements base only
type RedisRepository struct{}

func (r *RedisRepository) FindByID(id string) (*Entity, error) { ... }

// Client code checks capability
func processBatch(repo Repository, ids []string) {
    if batchRepo, ok := repo.(BatchRepository); ok {
        // Use batch if available
        return batchRepo.FindBatch(ids)
    } else {
        // Fallback to individual queries
        for _, id := range ids {
            repo.FindByID(id)
        }
    }
}
```

---

## 4. Interface Segregation Principle (ISP)

### 4.1. Định Nghĩa

> "Clients should not be forced to depend on interfaces they don't use."

Client không nên bị buộc phụ thuộc vào các phương thức không sử dụng. Tránh **"Fat Interface"** (Giao diện Quá Béo).

### 4.2. Áp Dụng Trong ITS

#### **Ví Dụ 1: Learner Data Access**

| **Aspect** | **Problem** | **Solution** |
|------------|-------------|--------------|
| **Scenario** | `LearnerRepository` chứa:<br>- `saveProfile()`<br>- `getProfile()`<br>- `updateMasteryScore()`<br>- `getHistory()` | **Tách interfaces:**<br>- `LearnerReadRepository`<br>- `LearnerWriteRepository` |
| **Issue** | Scoring Service chỉ cần đọc<br>nhưng phải implement write methods | Each service depends only<br>on what it needs |

**Code Example:**

```go
// ❌ BAD: Fat Interface

type LearnerRepository interface {
    // Read operations
    GetProfile(id string) (*Profile, error)
    GetSkillMastery(id string) (map[string]float64, error)
    GetHistory(id string) ([]*Event, error)
    
    // Write operations
    SaveProfile(profile *Profile) error
    UpdateProfile(id string, data *ProfileData) error
    UpdateMasteryScore(id string, skill string, score float64) error
    
    // Analytics
    GenerateReport(id string) (*Report, error)
    
    // Notification
    SendEmail(id string, subject, body string) error // ???
}

// Scoring Service forced to implement ALL methods
type ScoringService struct {
    learnerRepo LearnerRepository // Only needs GetSkillMastery!
}

// ✅ GOOD: Segregated Interfaces

// Read-only interface
type ProfileReader interface {
    GetProfile(id string) (*Profile, error)
}

// Skill mastery read
type SkillMasteryReader interface {
    GetSkillMastery(id string) (map[string]float64, error)
}

// Skill mastery write
type SkillMasteryWriter interface {
    UpdateMasteryScore(id string, skill string, score float64) error
}

// Profile write
type ProfileWriter interface {
    SaveProfile(profile *Profile) error
    UpdateProfile(id string, data *ProfileData) error
}

// Services depend ONLY on what they need

type ScoringService struct {
    skillReader SkillMasteryReader // Small, focused interface
}

type AdaptiveEngine struct {
    profileReader ProfileReader
    skillReader   SkillMasteryReader
}

type ProfileManagementService struct {
    profileWriter ProfileWriter
}

// Implementation can implement multiple interfaces
type LearnerRepository struct {
    db *sql.DB
}

// Implements ProfileReader
func (r *LearnerRepository) GetProfile(id string) (*Profile, error) { ... }

// Implements SkillMasteryReader
func (r *LearnerRepository) GetSkillMastery(id string) (map[string]float64, error) { ... }

// Implements SkillMasteryWriter
func (r *LearnerRepository) UpdateMasteryScore(id string, skill string, score float64) error { ... }

// Implements ProfileWriter
func (r *LearnerRepository) SaveProfile(profile *Profile) error { ... }
func (r *LearnerRepository) UpdateProfile(id string, data *ProfileData) error { ... }

// Compile-time verification
var _ ProfileReader = (*LearnerRepository)(nil)
var _ SkillMasteryReader = (*LearnerRepository)(nil)
var _ SkillMasteryWriter = (*LearnerRepository)(nil)
var _ ProfileWriter = (*LearnerRepository)(nil)
```

---

#### **Ví Dụ 2: Content Management**

```java
// ❌ BAD: Fat Interface
public interface ContentManager {
    void createQuiz(QuizDTO dto);
    void createVideo(VideoDTO dto);
    void publishContent(Long id);
    void archiveContent(Long id);
    void tagMetadata(Long id, List<String> tags);
    void assignToLearner(Long contentId, Long learnerId);
    void trackProgress(Long contentId, Long learnerId, double progress);
}

// ✅ GOOD: Segregated Interfaces
public interface ContentCreator {
    void createQuiz(QuizDTO dto);
    void createVideo(VideoDTO dto);
}

public interface ContentPublisher {
    void publish(Long id);
    void archive(Long id);
}

public interface ContentTagger {
    void tagMetadata(Long id, List<String> tags);
}

public interface LearnerContentAssigner {
    void assign(Long contentId, Long learnerId);
}

public interface ProgressTracker {
    void track(Long contentId, Long learnerId, double progress);
}

// Services use only what they need
@Service
public class QuizCreationService {
    private final ContentCreator creator; // Small interface
}

@Service
public class ContentPublishingService {
    private final ContentPublisher publisher; // Small interface
}
```

---

## 5. Dependency Inversion Principle (DIP)

### 5.1. Định Nghĩa

> "High-level modules should not depend on low-level modules. Both should depend on abstractions."

Module cấp cao (Policy) không phụ thuộc Module cấp thấp (Detail); cả hai phụ thuộc vào **Abstraction (Interface)**. Đây là **nguyên tắc bảo vệ lõi** của Clean Architecture.

### 5.2. Áp Dụng Trong ITS

#### **Ví Dụ 1: Adaptive Logic**

| **Aspect** | **Problem** | **Solution** |
|------------|-------------|--------------|
| **Scenario** | `AdaptivePathGenerator` (Policy)<br>gọi trực tiếp `PostgresConnection` | **Đảo ngược dependency:**<br>Policy → Interface ← Implementation |
| **Issue** | Policy coupled to DB<br>→ Hard to test, hard to change DB | Policy depends on interface<br>→ Easy to test, DB-agnostic |

**Code Example:**

```java
// ❌ BAD: High-level depends on low-level

// High-level module (Policy)
@Service
public class AdaptivePathGenerator {
    @Autowired
    private PostgresConnection postgres; // Direct dependency on DB!
    
    public AdaptivePath generate(Long learnerId) {
        // Directly calls PostgreSQL
        ResultSet rs = postgres.query(
            "SELECT * FROM learner_models WHERE id = " + learnerId);
        
        LearnerModel model = parseResultSet(rs);
        return generatePath(model);
    }
}

// Low-level module (Detail)
@Component
public class PostgresConnection {
    public ResultSet query(String sql) {
        // PostgreSQL-specific code
    }
}

// ✅ GOOD: Both depend on abstraction (DIP)

// Abstraction (Interface) - in application layer
package application.ports;

public interface LearnerModelRepository {
    LearnerModel findById(Long id);
}

// High-level module (Policy) - depends on INTERFACE
package application;

@Service
public class AdaptivePathGenerator {
    private final LearnerModelRepository repository; // Interface!
    
    // Constructor Injection (DIP)
    @Autowired
    public AdaptivePathGenerator(LearnerModelRepository repository) {
        this.repository = repository;
    }
    
    public AdaptivePath generate(Long learnerId) {
        LearnerModel model = repository.findById(learnerId); // Interface call
        return generatePath(model);
    }
    
    private AdaptivePath generatePath(LearnerModel model) {
        // Pure business logic, no DB dependencies
        List<String> weakSkills = model.getWeakSkills();
        return new AdaptivePath(weakSkills);
    }
}

// Low-level module (Detail) - implements INTERFACE
package infrastructure.persistence;

@Repository
public class PostgresLearnerModelRepository implements LearnerModelRepository {
    @Autowired
    private JdbcTemplate jdbcTemplate;
    
    @Override
    public LearnerModel findById(Long id) {
        // PostgreSQL-specific implementation
        return jdbcTemplate.queryForObject(
            "SELECT * FROM learner_models WHERE id = ?",
            new Object[]{id},
            new LearnerModelRowMapper()
        );
    }
}
```

**Benefits:**

```java
// Easy to test - mock the interface
@Test
public void testAdaptivePathGenerator() {
    // Mock repository
    LearnerModelRepository mockRepo = mock(LearnerModelRepository.class);
    when(mockRepo.findById(1L)).thenReturn(
        new LearnerModel(1L, Map.of("algebra", 0.3, "calculus", 0.9))
    );
    
    // Test business logic WITHOUT database
    AdaptivePathGenerator generator = new AdaptivePathGenerator(mockRepo);
    AdaptivePath path = generator.generate(1L);
    
    // Assert focuses on weak skills
    assertTrue(path.getRecommendedTopics().contains("algebra"));
}
```

---

#### **Ví Dụ 2: Factory Pattern (Object Creation)**

```java
// ❌ BAD: Direct instantiation (depends on concretion)

@Service
public class AIServiceFactory {
    public PathGenerator createGenerator(String type) {
        if ("basic".equals(type)) {
            return new BasicPathGenerator(); // Direct dependency!
        } else {
            return new AIPathGenerator(new MLModel()); // Direct dependency!
        }
    }
}

// ✅ GOOD: IoC Container (Spring DI) - depends on abstraction

// Configuration class
@Configuration
public class AIConfig {
    @Value("${ai.generator.type:ai}")
    private String generatorType;
    
    @Bean
    public PathGenerationAlgorithm pathGenerator(
        @Qualifier("basic") PathGenerationAlgorithm basic,
        @Qualifier("ai") PathGenerationAlgorithm ai
    ) {
        return "ai".equals(generatorType) ? ai : basic; // DIP
    }
    
    @Bean("basic")
    public PathGenerationAlgorithm basicPathGenerator() {
        return new BasicPathGenerator();
    }
    
    @Bean("ai")
    public PathGenerationAlgorithm aiPathGenerator(MLModel model) {
        return new AIPathGenerator(model);
    }
}

// Use Case depends on INTERFACE (not implementation)
@Service
public class GeneratePathUseCase {
    private final PathGenerationAlgorithm algorithm;
    
    @Autowired // Spring automatically injects the right implementation
    public GeneratePathUseCase(PathGenerationAlgorithm algorithm) {
        this.algorithm = algorithm; // DIP: depends on abstraction
    }
}
```

**Golang Implementation:**

```go
// Manual Dependency Injection (Golang)

// Interface (Abstraction)
type LearnerModelRepository interface {
    FindByID(ctx context.Context, id string) (*LearnerModel, error)
}

// High-level module
type AdaptivePathGenerator struct {
    repo LearnerModelRepository // Interface
}

// Constructor injection (manual DI)
func NewAdaptivePathGenerator(repo LearnerModelRepository) *AdaptivePathGenerator {
    return &AdaptivePathGenerator{repo: repo}
}

func (g *AdaptivePathGenerator) Generate(ctx context.Context, learnerID string) (*Path, error) {
    model, err := g.repo.FindByID(ctx, learnerID)
    if err != nil {
        return nil, err
    }
    return g.generatePath(model), nil
}

// Low-level module (Implementation)
type MongoLearnerRepository struct {
    client *mongo.Client
}

func (r *MongoLearnerRepository) FindByID(ctx context.Context, id string) (*LearnerModel, error) {
    // MongoDB implementation
}

// Wiring (main.go)
func main() {
    // Infrastructure setup
    mongoClient := setupMongo()
    
    // Create implementation
    repo := &MongoLearnerRepository{client: mongoClient}
    
    // Inject into high-level module (DIP)
    generator := NewAdaptivePathGenerator(repo)
    
    // Use
    path, _ := generator.Generate(context.Background(), "learner123")
}
```

---

## 6. Summary: SOLID Benefits for ITS

### 6.1. Architecture Characteristics Achieved

| **Principle** | **AC Supported** | **Benefits** |
|---------------|------------------|--------------|
| **SRP** | **AC1:** Modularity | - Clear service boundaries<br>- Independent scaling<br>- Easier to understand |
| **OCP** | **AC5:** Deployability<br>**AC1:** Modularity | - Add features without breaking existing code<br>- Support Blue/Green deployment (FR9)<br>- Easy A/B testing |
| **LSP** | **AC4:** Testability | - Reliable polymorphism<br>- Type-safe substitution<br>- Fewer runtime errors |
| **ISP** | **AC4:** Testability<br>**AC1:** Modularity | - Small, focused interfaces<br>- Easy to mock<br>- Reduced coupling |
| **DIP** | **AC4:** Testability<br>**AC1:** Modularity<br>**AC7:** Maintainability | - **Testable business logic**<br>- Framework independence<br>- Database independence<br>- **Stable architecture (I≈0)** |

### 6.2. Implementation Strategy

**Java (Spring Boot):**
- ✅ Use `@Autowired` constructor injection for DIP
- ✅ Use `@Qualifier` for multiple implementations (OCP)
- ✅ Use `@Service`, `@Repository` stereotypes (SRP)
- ✅ Use interfaces liberally (ISP, LSP, DIP)

**Golang:**
- ✅ Use interfaces extensively (all SOLID principles)
- ✅ Use constructor functions for DI
- ✅ Use `interface{}` sparingly (prefer typed interfaces)
- ✅ Use struct composition over inheritance

### 6.3. Key Takeaways

1. **DIP is the Foundation:**
   - Most important principle for Clean Architecture
   - Enables Testability (AC4)
   - Protects core business logic

2. **SRP at All Levels:**
   - Service level (microservices)
   - Class level (single responsibility)
   - Method level (do one thing well)

3. **OCP for Extensibility:**
   - Critical for AI/ML features (FR9)
   - Supports Blue/Green deployment
   - Reduces risk of changes

4. **ISP for Flexibility:**
   - Small interfaces easier to implement
   - Easier to mock in tests
   - Reduces coupling

5. **LSP for Reliability:**
   - Ensures polymorphism works
   - Prevents runtime surprises
   - Type-safe substitution

---

**Tài liệu tham khảo:**
- Clean Architecture (Robert C. Martin)
- Agile Software Development: Principles, Patterns, and Practices (Robert C. Martin)
- Design Patterns: Elements of Reusable Object-Oriented Software (Gang of Four)
- Effective Java (Joshua Bloch)
- Go Design Patterns (Mario Castro Contreras)
