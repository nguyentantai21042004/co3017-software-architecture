# Project Architecture - Content Service

## Table of Contents

1. [Architecture Overview](#architecture-overview)
2. [Architecture Layers](#architecture-layers)
3. [Directory Structure](#directory-structure)
4. [Request Flow](#request-flow)
5. [Implementation Guide](#implementation-guide)
6. [Rules and Best Practices](#rules-and-best-practices)
7. [Real Examples](#real-examples)

---

## Architecture Overview

This project follows **Clean Architecture** principles with 4 main layers:

```
┌─────────────────────────────────────────────────────────┐
│                    Domain Layer                         │
│  (models/) - Pure Business Logic, No Dependencies       │
└─────────────────────────────────────────────────────────┘
                        ↑
┌─────────────────────────────────────────────────────────┐
│                Application Layer                        │
│  (usecase/) - Business Rules & Orchestration            │
└─────────────────────────────────────────────────────────┘
                        ↑
┌─────────────────────────────────────────────────────────┐
│              Interface Adapters Layer                   │
│  (adapter/http/) - HTTP Controllers, DTOs, Mappers     │
└─────────────────────────────────────────────────────────┘
                        ↑
┌─────────────────────────────────────────────────────────┐
│              Infrastructure Layer                       │
│  (repository/postgresql/) - Database, JPA, Entities     │
└─────────────────────────────────────────────────────────┘
```

### Dependency Rule

- **Dependencies point inward only**
- Domain layer has no dependencies on any other layer
- Application layer depends only on Domain layer
- Interface Adapters depend on Application and Domain layers
- Infrastructure depends on all layers above

---

## Architecture Layers

This section provides a detailed explanation of each layer in the Clean Architecture. Understanding these layers is crucial for maintaining the architecture and adding new features correctly.

### Layer 1: Domain Layer (`models/`)

**Purpose**: Contains pure business logic with no framework dependencies

**Location**: `src/main/java/co3017/microservices/content_service/models/`

**Key Characteristics**:
- Pure Java classes with no Spring or JPA annotations
- Contains business rules and validation logic
- Provides methods to manipulate domain objects (e.g., `updateTitle()`, `isValid()`)
- No dependencies on database, HTTP, or any framework
- Can be tested independently without Spring context

**Example**:
```java
// models/Course.java
public class Course {
    private UUID courseId;
    private String title;
    private String description;
    private UUID instructorId;
    private StructureType structureType;
    private LocalDateTime createdAt;
    private LocalDateTime updatedAt;
    
    // Business logic methods
    public boolean isValid() {
        return hasValidTitle() && hasValidDescription() &&
               hasValidInstructorId() && hasValidStructureType();
    }
    
    public void updateTitle(String newTitle) {
        if (newTitle == null || newTitle.trim().isEmpty()) {
            throw new IllegalArgumentException("Title cannot be null or empty");
        }
        this.title = newTitle;
        this.updatedAt = LocalDateTime.now();
    }
    
    public boolean hasValidTitle() {
        return title != null && !title.trim().isEmpty() && title.length() >= 3;
    }
}
```

**Rules**:
- DO NOT import Spring, JPA, or any framework classes
- DO NOT use `@Entity`, `@Table`, `@Column` annotations
- DO contain business logic and validation only
- CAN be tested independently without Spring context
- MUST be framework-agnostic

**Domain Models in Project**:
- `Course.java` - Course aggregate root
- `Chapter.java` - Chapter aggregate root
- `ContentUnit.java` - ContentUnit aggregate root
- `ContentVersion.java` - ContentVersion aggregate root
- `MetadataTag.java` - Metadata tag value object
- `PathCondition.java` - Path condition value object
- `UnitTag.java` - Unit tag value object
- `Test.java` - Test aggregate root

---

### Layer 2: Application Layer (`usecase/`)

**Purpose**: Orchestrates business logic, enforces business rules, and coordinates domain operations

**Location**: `src/main/java/co3017/microservices/content_service/usecase/`

**Key Characteristics**:
- Defines use case contracts through interfaces
- Implements business orchestration logic
- Validates business rules before domain operations
- Coordinates between domain models and repositories
- Uses `@Service` and `@Transactional` annotations
- Depends only on Domain layer and Repository interfaces

#### 2.1. Use Case Interface (`usecase/*UseCase.java`)

**Purpose**: Defines the contract for all use cases of a domain

**Example**:
```java
// usecase/CourseUseCase.java
public interface CourseUseCase {
    Course create(CreateCourseCommand command);
    Optional<Course> detail(UUID courseId);
    CoursePageResult list(CourseSearchCriteria criteria);
    Optional<Course> update(UUID courseId, UpdateCourseCommand command);
    int deletes(List<UUID> courseIds);
    boolean existsById(UUID courseId);
    boolean existsByTitle(String title);
}
```

**Rules**:
- Defines all business operations for a domain
- Works only with Domain models and Command/Query objects
- No implementation, interface only
- One interface per domain

#### 2.2. Use Case Service (`usecase/service/*Service.java`)

**Purpose**: Implements use case interfaces and contains business orchestration logic

**Example**:
```java
// usecase/service/CourseService.java
@Service
@Transactional
public class CourseService implements CourseUseCase {
    private final CourseRepository courseRepository;
    
    public CourseService(CourseRepository courseRepository) {
        this.courseRepository = courseRepository;
    }
    
    @Override
    public Course create(CreateCourseCommand command) {
        // Business validation
        if (command.getTitle() == null || command.getTitle().trim().isEmpty()) {
            throw new IllegalArgumentException("Course title cannot be null or empty");
        }
        
        // Check duplicate
        if (courseRepository.existsByTitle(command.getTitle())) {
            throw new IllegalArgumentException("Course with this title already exists");
        }
        
        // Create domain entity
        Course course = new Course(
            command.getTitle(),
            command.getDescription(),
            command.getInstructorId(),
            command.getStructureType()
        );
        
        // Domain validation
        if (!course.isValid()) {
            throw new IllegalArgumentException("Invalid course data");
        }
        
        // Save via repository
        return courseRepository.save(course);
    }
    
    @Override
    @Transactional(readOnly = true)
    public Optional<Course> detail(UUID courseId) {
        if (courseId == null) {
            throw new IllegalArgumentException("Course ID cannot be null");
        }
        return courseRepository.findById(courseId);
    }
}
```

**Rules**:
- Implements UseCase interface
- Uses `@Service` and `@Transactional` annotations
- Contains business validation and orchestration
- Calls Repository interface (not implementation)
- Read-only operations use `@Transactional(readOnly = true)`
- Write operations use `@Transactional` (default)

#### 2.3. Command/Query Types (`usecase/types/`)

**Purpose**: Pure Java DTOs for commands and queries, used to transfer data between layers

**Example**:
```java
// usecase/types/CreateCourseCommand.java
public class CreateCourseCommand {
    private final String title;
    private final String description;
    private final UUID instructorId;
    private final Course.StructureType structureType;
    
    public CreateCourseCommand(String title, String description, 
                               UUID instructorId, Course.StructureType structureType) {
        this.title = title;
        this.description = description;
        this.instructorId = instructorId;
        this.structureType = structureType;
    }
    
    // Getters...
}

// usecase/types/CourseSearchCriteria.java
public class CourseSearchCriteria {
    private String title;
    private UUID instructorId;
    private Course.StructureType structureType;
    private int page;
    private int size;
    
    // Constructor, getters, setters...
}
```

**Rules**:
- Contains data only, no business logic
- Immutable when possible (final fields)
- No framework annotations
- One command/query type per operation

**Use Case Services in Project**:
- `CourseService.java` - Implements CourseUseCase
- `ChapterService.java` - Implements ChapterUseCase
- `ContentUnitService.java` - Implements ContentUnitUseCase
- `ContentVersionService.java` - Implements ContentVersionUseCase
- `MetadataTagService.java` - Implements MetadataTagUseCase
- `TestService.java` - Implements TestUseCase

---

### Layer 3: Interface Adapters Layer (`adapter/http/`)

**Purpose**: Converts between HTTP world and Application world

**Location**: `src/main/java/co3017/microservices/content_service/adapter/http/`

**Key Characteristics**:
- REST endpoints with Spring `@RestController`
- Receives Request DTOs, returns Response DTOs
- Uses builders to convert between DTOs and Commands/Domain
- Wraps responses in standardized `ApiResponse<T>` format
- Handles HTTP-specific concerns (status codes, headers)

#### 3.1. Controllers (`adapter/http/*Controller.java`)

**Purpose**: Handles HTTP requests and responses

**Example**:
```java
// adapter/http/CourseController.java
@RestController
@RequestMapping("/api/courses")
public class CourseController {
    private final CourseUseCase courseUseCase;
    
    public CourseController(CourseUseCase courseUseCase) {
        this.courseUseCase = courseUseCase;
    }
    
    @PostMapping
    public ResponseEntity<ApiResponse<CourseResponse>> create(
            @RequestBody CreateCourseRequest request) {
        // Convert Request DTO → Command
        CreateCourseCommand command = CommandBuilder.toCreateCourseCommand(request);
        
        // Call use case
        Course course = courseUseCase.create(command);
        
        // Convert Domain → Response DTO
        CourseResponse response = CourseResponseBuilder.toResponse(course);
        
        // Wrap in ApiResponse
        return ResponseEntity.status(HttpStatus.CREATED)
                .body(ApiResponse.success("Course created successfully", response));
    }
    
    @GetMapping("/{courseId}")
    public ResponseEntity<ApiResponse<CourseResponse>> detail(@PathVariable UUID courseId) {
        return courseUseCase.detail(courseId)
                .map(course -> ResponseEntity.ok(
                        ApiResponse.success(CourseResponseBuilder.toResponse(course))))
                .orElse(ResponseEntity.status(HttpStatus.NOT_FOUND)
                        .body(ApiResponse.error(404, "Course not found")));
    }
    
    @ExceptionHandler(IllegalArgumentException.class)
    public ResponseEntity<ApiResponse<Void>> handleIllegalArgumentException(IllegalArgumentException e) {
        return ResponseEntity.badRequest()
                .body(ApiResponse.error(400, e.getMessage()));
    }
}
```

**Rules**:
- Controllers work only with DTOs, never directly with Domain
- Always use Builder to convert between DTOs and Commands/Domain
- Always wrap response in `ApiResponse<T>`
- Handle exceptions in Controller with `@ExceptionHandler`
- One controller per domain

#### 3.2. DTOs (`adapter/http/dto/`)

**Purpose**: Data Transfer Objects for HTTP requests and responses

**Example**:
```java
// adapter/http/dto/CreateCourseRequest.java
public class CreateCourseRequest {
    private String title;
    private String description;
    private UUID instructorId;
    private Course.StructureType structureType;
    
    // Getters, setters...
}

// adapter/http/dto/CourseResponse.java
public class CourseResponse {
    private UUID courseId;
    private String title;
    private String description;
    private UUID instructorId;
    private Course.StructureType structureType;
    private LocalDateTime createdAt;
    private LocalDateTime updatedAt;
    
    // Getters, setters...
}

// adapter/http/dto/ApiResponse.java
public class ApiResponse<T> {
    private int errorCode;
    private String message;
    private T data;
    
    public static <T> ApiResponse<T> success(T data) {
        return new ApiResponse<>(0, "Success", data);
    }
    
    public static <T> ApiResponse<T> success(String message, T data) {
        return new ApiResponse<>(0, message, data);
    }
    
    public static <T> ApiResponse<T> error(int errorCode, String message) {
        return new ApiResponse<>(errorCode, message, null);
    }
}
```

**Rules**:
- Request DTOs: Receive data from HTTP requests
- Response DTOs: Return data in HTTP responses
- `ApiResponse<T>`: Standardized response wrapper for all endpoints
- DTOs are separate from Domain models

#### 3.3. Response Builders (`adapter/http/response/`)

**Purpose**: Convert between DTOs and Commands/Domain objects

**Example**:
```java
// adapter/http/response/CommandBuilder.java
public class CommandBuilder {
    public static CreateCourseCommand toCreateCourseCommand(CreateCourseRequest request) {
        return new CreateCourseCommand(
            request.getTitle(),
            request.getDescription(),
            request.getInstructorId(),
            request.getStructureType()
        );
    }
    
    public static UpdateCourseCommand toUpdateCourseCommand(UpdateCourseRequest request) {
        return new UpdateCourseCommand(
            request.getTitle(),
            request.getDescription(),
            request.getStructureType()
        );
    }
}

// adapter/http/response/CourseResponseBuilder.java
public class CourseResponseBuilder {
    public static CourseResponse toResponse(Course course) {
        CourseResponse response = new CourseResponse();
        response.setCourseId(course.getCourseId());
        response.setTitle(course.getTitle());
        response.setDescription(course.getDescription());
        response.setInstructorId(course.getInstructorId());
        response.setStructureType(course.getStructureType());
        response.setCreatedAt(course.getCreatedAt());
        response.setUpdatedAt(course.getUpdatedAt());
        return response;
    }
}
```

**Rules**:
- `CommandBuilder`: Converts Request DTO → Command
- `*ResponseBuilder`: Converts Domain → Response DTO
- Static utility methods
- One builder per domain

**Controllers in Project**:
- `CourseController.java` - REST endpoints for Course
- `ChapterController.java` - REST endpoints for Chapter
- `ContentUnitController.java` - REST endpoints for ContentUnit
- `ContentVersionController.java` - REST endpoints for ContentVersion
- `TestController.java` - REST endpoints for Test

---

### Layer 4: Infrastructure Layer (`repository/postgresql/`)

**Purpose**: Implements data persistence with PostgreSQL and JPA

**Location**: `src/main/java/co3017/microservices/content_service/repository/postgresql/`

**Key Characteristics**:
- Implements Repository interfaces (Port Out)
- Uses Spring Data JPA for database operations
- Converts between Domain and Entity through Mappers
- Handles pagination and search with Specifications
- Contains JPA entities and database-specific code

#### 4.1. JPA Repository Implementation (`Jpa*Repository.java`)

**Purpose**: Implements Repository interface using Spring Data JPA

**Example**:
```java
// repository/postgresql/JpaCourseRepository.java
@Repository
public class JpaCourseRepository implements CourseRepository {
    private final SpringDataCourseRepository springDataCourseRepository;
    
    public JpaCourseRepository(SpringDataCourseRepository springDataCourseRepository) {
        this.springDataCourseRepository = springDataCourseRepository;
    }
    
    @Override
    public Course save(Course course) {
        // Domain → Entity
        CourseEntity entity = CourseMapper.toEntity(course);
        
        // Save to database
        CourseEntity savedEntity = springDataCourseRepository.save(entity);
        
        // Entity → Domain
        return CourseMapper.toDomain(savedEntity);
    }
    
    @Override
    public Optional<Course> findById(UUID courseId) {
        return springDataCourseRepository.findById(courseId)
            .map(CourseMapper::toDomain);
    }
    
    @Override
    public CoursePageResult search(CourseSearchCriteria criteria) {
        Pageable pageable = PageRequest.of(criteria.getPage(), criteria.getSize());
        
        // Build specification for dynamic query
        Specification<CourseEntity> spec = CourseSpecification.createSpecification(criteria);
        
        // Query with pagination
        Page<CourseEntity> page = springDataCourseRepository.findAll(spec, pageable);
        
        // Convert entities to domain
        List<Course> courses = page.getContent().stream()
            .map(CourseMapper::toDomain)
            .collect(Collectors.toList());
        
        // Build result
        return new CoursePageResult(
            courses,
            page.getNumber(),
            page.getSize(),
            page.getTotalElements(),
            page.getTotalPages()
        );
    }
    
    @Override
    public void deleteByIds(List<UUID> courseIds) {
        springDataCourseRepository.deleteAllById(courseIds);
    }
    
    @Override
    public boolean existsById(UUID courseId) {
        return springDataCourseRepository.existsById(courseId);
    }
    
    @Override
    public boolean existsByTitle(String title) {
        return springDataCourseRepository.existsByTitle(title);
    }
}
```

**Rules**:
- Implements Repository interface (Port Out)
- Always converts Domain ↔ Entity through Mapper
- Never returns Entity directly, always converts to Domain
- Uses Spring Data JPA repository for database operations
- Handles pagination and search with Specifications

#### 4.2. Spring Data JPA Interface (`SpringData*Repository.java`)

**Purpose**: Spring Data JPA repository interface

**Example**:
```java
// repository/postgresql/SpringDataCourseRepository.java
public interface SpringDataCourseRepository extends JpaRepository<CourseEntity, UUID>, 
                                                   JpaSpecificationExecutor<CourseEntity> {
    boolean existsByTitle(String title);
}
```

**Rules**:
- Extends `JpaRepository<Entity, ID>`
- Extends `JpaSpecificationExecutor<Entity>` for dynamic queries
- Contains only method signatures, Spring implements automatically
- Can have custom query methods

#### 4.3. JPA Entities (`repository/postgresql/entity/`)

**Purpose**: Database representation with JPA annotations

**Example**:
```java
// repository/postgresql/entity/CourseEntity.java
@Entity
@Table(name = "courses")
public class CourseEntity {
    @Id
    @Column(name = "course_id")
    private UUID courseId;
    
    @Column(name = "title", nullable = false)
    private String title;
    
    @Column(name = "description")
    private String description;
    
    @Column(name = "instructor_id", nullable = false)
    private UUID instructorId;
    
    @Column(name = "structure_type", nullable = false)
    private String structureType;
    
    @Column(name = "created_at")
    private LocalDateTime createdAt;
    
    @Column(name = "updated_at")
    private LocalDateTime updatedAt;
    
    // Getters, setters...
}
```

**Rules**:
- Uses JPA annotations (`@Entity`, `@Table`, `@Column`)
- Represents database structure
- No business logic
- Separate from Domain models

#### 4.4. Mappers (`repository/postgresql/mapper/`)

**Purpose**: Converts between Domain and Entity

**Example**:
```java
// repository/postgresql/mapper/CourseMapper.java
public class CourseMapper {
    public static CourseEntity toEntity(Course course) {
        CourseEntity entity = new CourseEntity();
        entity.setCourseId(course.getCourseId());
        entity.setTitle(course.getTitle());
        entity.setDescription(course.getDescription());
        entity.setInstructorId(course.getInstructorId());
        entity.setStructureType(course.getStructureType().name());
        entity.setCreatedAt(course.getCreatedAt());
        entity.setUpdatedAt(course.getUpdatedAt());
        return entity;
    }
    
    public static Course toDomain(CourseEntity entity) {
        return new Course(
            entity.getCourseId(),
            entity.getTitle(),
            entity.getDescription(),
            entity.getInstructorId(),
            Course.StructureType.valueOf(entity.getStructureType()),
            entity.getCreatedAt(),
            entity.getUpdatedAt()
        );
    }
}
```

**Rules**:
- Static utility methods
- Converts Domain → Entity for saving
- Converts Entity → Domain for reading
- One mapper per domain

#### 4.5. Specifications (`repository/postgresql/specification/`)

**Purpose**: Builds dynamic queries using JPA Specifications

**Example**:
```java
// repository/postgresql/specification/CourseSpecification.java
public class CourseSpecification {
    public static Specification<CourseEntity> createSpecification(CourseSearchCriteria criteria) {
        return Specification.where(titleLike(criteria.getTitle()))
            .and(instructorIdEquals(criteria.getInstructorId()))
            .and(structureTypeEquals(criteria.getStructureType()));
    }
    
    private static Specification<CourseEntity> titleLike(String title) {
        return (root, query, cb) -> 
            title == null ? null : 
            cb.like(cb.lower(root.get("title")), "%" + title.toLowerCase() + "%");
    }
    
    private static Specification<CourseEntity> instructorIdEquals(UUID instructorId) {
        return (root, query, cb) -> 
            instructorId == null ? null : 
            cb.equal(root.get("instructorId"), instructorId);
    }
    
    private static Specification<CourseEntity> structureTypeEquals(Course.StructureType structureType) {
        return (root, query, cb) -> 
            structureType == null ? null : 
            cb.equal(root.get("structureType"), structureType.name());
    }
}
```

**Rules**:
- Builds dynamic query criteria from SearchCriteria
- Uses JPA Specifications for type-safe queries
- One specification per domain

**Repository Implementations in Project**:
- `JpaCourseRepository.java` - Implements CourseRepository
- `JpaChapterRepository.java` - Implements ChapterRepository
- `JpaContentUnitRepository.java` - Implements ContentUnitRepository
- `JpaContentVersionRepository.java` - Implements ContentVersionRepository
- `JpaMetadataTagRepository.java` - Implements MetadataTagRepository
- `JpaTestRepository.java` - Implements TestRepository

---

## Directory Structure

```
src/main/java/co3017/microservices/content_service/
│
├── models/                          # Domain Layer - Pure Business Logic
│   ├── Course.java                  # Aggregate Root - Course domain
│   ├── Chapter.java                 # Aggregate Root - Chapter domain
│   ├── ContentUnit.java            # Aggregate Root - ContentUnit domain
│   ├── ContentVersion.java         # Aggregate Root - ContentVersion domain
│   ├── MetadataTag.java            # Value Object - Metadata tag
│   ├── PathCondition.java          # Value Object - Path condition
│   ├── UnitTag.java                # Value Object - Unit tag
│   └── Test.java                   # Aggregate Root - Test domain
│
├── usecase/                         # Application Layer - Business Rules
│   ├── CourseUseCase.java          # Interface - Course use cases contract
│   ├── ChapterUseCase.java         # Interface - Chapter use cases contract
│   ├── ContentUnitUseCase.java     # Interface - ContentUnit use cases
│   ├── ContentVersionUseCase.java  # Interface - ContentVersion use cases
│   ├── MetadataTagUseCase.java     # Interface - MetadataTag use cases
│   ├── PathConditionUseCase.java  # Interface - PathCondition use cases
│   ├── UnitTagUseCase.java         # Interface - UnitTag use cases
│   ├── TestUseCase.java            # Interface - Test use cases
│   │
│   ├── service/                    # Use Case Implementations
│   │   ├── CourseService.java      # Implements CourseUseCase
│   │   ├── ChapterService.java     # Implements ChapterUseCase
│   │   ├── ContentUnitService.java # Implements ContentUnitUseCase
│   │   ├── ContentVersionService.java
│   │   ├── MetadataTagService.java
│   │   └── TestService.java
│   │
│   └── types/                      # Commands, Queries, Results
│       ├── CreateCourseCommand.java      # Command for creating course
│       ├── UpdateCourseCommand.java      # Command for updating course
│       ├── CourseSearchCriteria.java     # Query criteria for search
│       ├── CoursePageResult.java         # Paginated result
│       ├── CreateChapterCommand.java
│       ├── UpdateChapterCommand.java
│       ├── ChapterSearchCriteria.java
│       ├── ChapterPageResult.java
│       └── ... (similar for other domains)
│
├── repository/                      # Repository Layer - Data Access
│   ├── CourseRepository.java       # Port Out - Course repository interface
│   ├── ChapterRepository.java      # Port Out - Chapter repository interface
│   ├── ContentUnitRepository.java
│   ├── ContentVersionRepository.java
│   ├── MetadataTagRepository.java
│   ├── TestRepository.java
│   │
│   └── postgresql/                 # Infrastructure - PostgreSQL Implementation
│       ├── JpaCourseRepository.java        # Implements CourseRepository
│       ├── JpaChapterRepository.java       # Implements ChapterRepository
│       ├── JpaContentUnitRepository.java
│       ├── JpaContentVersionRepository.java
│       ├── JpaMetadataTagRepository.java
│       ├── JpaTestRepository.java
│       │
│       ├── SpringDataCourseRepository.java    # Spring Data JPA interface
│       ├── SpringDataChapterRepository.java
│       ├── SpringDataTestRepository.java
│       │
│       ├── entity/                  # JPA Entities (Database representation)
│       │   ├── CourseEntity.java
│       │   ├── ChapterEntity.java
│       │   ├── ContentUnitEntity.java
│       │   ├── ContentVersionEntity.java
│       │   ├── MetadataTagEntity.java
│       │   └── TestEntity.java
│       │
│       ├── mapper/                  # Domain ↔ Entity Mappers
│       │   ├── CourseMapper.java    # Course ↔ CourseEntity
│       │   ├── ChapterMapper.java   # Chapter ↔ ChapterEntity
│       │   └── TestMapper.java
│       │
│       └── specification/            # JPA Specifications (Dynamic Queries)
│           ├── CourseSpecification.java
│           ├── ChapterSpecification.java
│           └── ...
│
├── adapter/                         # Interface Adapters Layer
│   └── http/                       # HTTP Adapter (REST API)
│       ├── CourseController.java    # REST Controller for Course
│       ├── ChapterController.java   # REST Controller for Chapter
│       ├── ContentUnitController.java
│       ├── ContentVersionController.java
│       ├── TestController.java
│       │
│       ├── dto/                    # Data Transfer Objects
│       │   ├── ApiResponse.java    # Standardized API response wrapper
│       │   ├── ErrorResponse.java  # Error response format
│       │   ├── CreateCourseRequest.java    # Request DTO for creating course
│       │   ├── UpdateCourseRequest.java    # Request DTO for updating course
│       │   ├── CourseResponse.java         # Response DTO for course
│       │   ├── CoursePageResponse.java     # Paginated response DTO
│       │   ├── CreateChapterRequest.java
│       │   ├── UpdateChapterRequest.java
│       │   ├── ChapterResponse.java
│       │   └── ... (similar for other domains)
│       │
│       └── response/               # Response Builders
│           ├── CommandBuilder.java          # Request DTO → Command
│           ├── CourseResponseBuilder.java   # Domain → Response DTO
│           ├── ChapterResponseBuilder.java
│           ├── ContentUnitResponseBuilder.java
│           ├── ContentVersionResponseBuilder.java
│           └── TestResponseBuilder.java
│
├── mappers/                         # Cross-layer Mappers
│   ├── ContentUnitMapper.java      # Domain ↔ DTO mappers
│   ├── ContentVersionMapper.java
│   └── MetadataTagMapper.java
│
└── config/                          # Configuration Layer
    ├── ContentServiceApplication.java  # Spring Boot entry point
    └── CorsConfig.java                 # CORS configuration
```

---

## Request Flow

### Flow Diagram

```
HTTP Request
    ↓
Controller (adapter/http/CourseController)
    ↓ [Request DTO → Command]
CommandBuilder.toCreateCourseCommand(request)
    ↓
UseCase Interface (usecase/CourseUseCase)
    ↓
UseCase Service (usecase/service/CourseService)
    ↓ [Business Validation]
Domain Model (models/Course)
    ↓ [Domain Validation]
Repository Interface (repository/CourseRepository)
    ↓
JPA Repository (repository/postgresql/JpaCourseRepository)
    ↓ [Domain → Entity]
Mapper (repository/postgresql/mapper/CourseMapper)
    ↓
Spring Data JPA (repository/postgresql/SpringDataCourseRepository)
    ↓
PostgreSQL Database
    ↓ [Entity → Domain]
Mapper.toDomain(entity)
    ↓
Domain Model
    ↓
Repository returns Domain
    ↓
UseCase returns Domain
    ↓ [Domain → Response DTO]
ResponseBuilder.toResponse(domain)
    ↓
Controller wraps in ApiResponse
    ↓
HTTP Response
```

### Detailed Example: Creating a Course

```java
// 1. HTTP Request to Controller
POST /api/courses
{
  "title": "Java Programming",
  "description": "Learn Java",
  "instructorId": "123e4567-e89b-12d3-a456-426614174000",
  "structureType": "LINEAR"
}

// 2. Controller receives Request DTO
@PostMapping
public ResponseEntity<ApiResponse<CourseResponse>> create(
        @RequestBody CreateCourseRequest request) {
    
    // 3. Convert Request DTO → Command
    CreateCourseCommand command = CommandBuilder.toCreateCourseCommand(request);
    
    // 4. Call UseCase
    Course course = courseUseCase.create(command);
    
    // 5. Convert Domain → Response DTO
    CourseResponse response = CourseResponseBuilder.toResponse(course);
    
    // 6. Wrap in ApiResponse
    return ResponseEntity.status(HttpStatus.CREATED)
            .body(ApiResponse.success("Course created successfully", response));
}

// 7. UseCase Service executes business logic
@Override
public Course create(CreateCourseCommand command) {
    // Business validation
    if (command.getTitle() == null || command.getTitle().trim().isEmpty()) {
        throw new IllegalArgumentException("Title cannot be null");
    }
    
    // Check duplicate
    if (courseRepository.existsByTitle(command.getTitle())) {
        throw new IllegalArgumentException("Title already exists");
    }
    
    // Create domain entity
    Course course = new Course(
        command.getTitle(),
        command.getDescription(),
        command.getInstructorId(),
        command.getStructureType()
    );
    
    // Domain validation
    if (!course.isValid()) {
        throw new IllegalArgumentException("Invalid course");
    }
    
    // Save via repository
    return courseRepository.save(course);
}

// 8. Repository Implementation
@Override
public Course save(Course course) {
    // Domain → Entity
    CourseEntity entity = CourseMapper.toEntity(course);
    
    // Save to database
    CourseEntity savedEntity = springDataCourseRepository.save(entity);
    
    // Entity → Domain
    return CourseMapper.toDomain(savedEntity);
}

// 9. Response
{
  "errorCode": 0,
  "message": "Course created successfully",
  "data": {
    "courseId": "123e4567-e89b-12d3-a456-426614174001",
    "title": "Java Programming",
    "description": "Learn Java",
    "instructorId": "123e4567-e89b-12d3-a456-426614174000",
    "structureType": "LINEAR",
    "createdAt": "2024-01-01T10:00:00",
    "updatedAt": "2024-01-01T10:00:00"
  }
}
```

---

## Implementation Guide

### Adding a New Domain

Suppose you want to add a new domain `Article`:

#### Step 1: Create Domain Model

```java
// models/Article.java
package co3017.microservices.content_service.models;

import java.time.LocalDateTime;
import java.util.UUID;

public class Article {
    private UUID articleId;
    private String title;
    private String content;
    private UUID authorId;
    private LocalDateTime createdAt;
    private LocalDateTime updatedAt;
    
    // Constructors
    public Article(String title, String content, UUID authorId) {
        this.articleId = UUID.randomUUID();
        this.title = title;
        this.content = content;
        this.authorId = authorId;
        this.createdAt = LocalDateTime.now();
        this.updatedAt = LocalDateTime.now();
    }
    
    // Business logic
    public boolean isValid() {
        return title != null && !title.trim().isEmpty() &&
               content != null && !content.trim().isEmpty() &&
               authorId != null;
    }
    
    public void updateTitle(String newTitle) {
        if (newTitle == null || newTitle.trim().isEmpty()) {
            throw new IllegalArgumentException("Title cannot be null or empty");
        }
        this.title = newTitle;
        this.updatedAt = LocalDateTime.now();
    }
    
    // Getters...
}
```

#### Step 2: Create UseCase Interface

```java
// usecase/ArticleUseCase.java
package co3017.microservices.content_service.usecase;

import co3017.microservices.content_service.models.Article;
import co3017.microservices.content_service.usecase.types.*;
import java.util.List;
import java.util.Optional;
import java.util.UUID;

public interface ArticleUseCase {
    Article create(CreateArticleCommand command);
    Optional<Article> detail(UUID articleId);
    ArticlePageResult list(ArticleSearchCriteria criteria);
    Optional<Article> update(UUID articleId, UpdateArticleCommand command);
    int deletes(List<UUID> articleIds);
    boolean existsById(UUID articleId);
}
```

#### Step 3: Create Command/Query Types

```java
// usecase/types/CreateArticleCommand.java
package co3017.microservices.content_service.usecase.types;

import java.util.UUID;

public class CreateArticleCommand {
    private final String title;
    private final String content;
    private final UUID authorId;
    
    public CreateArticleCommand(String title, String content, UUID authorId) {
        this.title = title;
        this.content = content;
        this.authorId = authorId;
    }
    
    // Getters...
}

// usecase/types/UpdateArticleCommand.java
// usecase/types/ArticleSearchCriteria.java
// usecase/types/ArticlePageResult.java
```

#### Step 4: Create UseCase Service Implementation

```java
// usecase/service/ArticleService.java
package co3017.microservices.content_service.usecase.service;

import co3017.microservices.content_service.usecase.ArticleUseCase;
import co3017.microservices.content_service.usecase.types.*;
import co3017.microservices.content_service.repository.ArticleRepository;
import co3017.microservices.content_service.models.Article;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

@Service
@Transactional
public class ArticleService implements ArticleUseCase {
    private final ArticleRepository articleRepository;
    
    public ArticleService(ArticleRepository articleRepository) {
        this.articleRepository = articleRepository;
    }
    
    @Override
    public Article create(CreateArticleCommand command) {
        // Business validation
        if (command.getTitle() == null || command.getTitle().trim().isEmpty()) {
            throw new IllegalArgumentException("Title cannot be null or empty");
        }
        
        // Create domain entity
        Article article = new Article(
            command.getTitle(),
            command.getContent(),
            command.getAuthorId()
        );
        
        // Domain validation
        if (!article.isValid()) {
            throw new IllegalArgumentException("Invalid article");
        }
        
        // Save
        return articleRepository.save(article);
    }
    
    // Implement other methods...
}
```

#### Step 5: Create Repository Interface

```java
// repository/ArticleRepository.java
package co3017.microservices.content_service.repository;

import co3017.microservices.content_service.models.Article;
import co3017.microservices.content_service.usecase.types.ArticleSearchCriteria;
import co3017.microservices.content_service.usecase.types.ArticlePageResult;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

public interface ArticleRepository {
    Article save(Article article);
    Optional<Article> findById(UUID articleId);
    ArticlePageResult search(ArticleSearchCriteria criteria);
    void deleteByIds(List<UUID> articleIds);
    boolean existsById(UUID articleId);
}
```

#### Step 6: Create JPA Entity

```java
// repository/postgresql/entity/ArticleEntity.java
package co3017.microservices.content_service.repository.postgresql.entity;

import jakarta.persistence.*;
import java.time.LocalDateTime;
import java.util.UUID;

@Entity
@Table(name = "articles")
public class ArticleEntity {
    @Id
    @Column(name = "article_id")
    private UUID articleId;
    
    @Column(name = "title", nullable = false)
    private String title;
    
    @Column(name = "content", columnDefinition = "TEXT")
    private String content;
    
    @Column(name = "author_id", nullable = false)
    private UUID authorId;
    
    @Column(name = "created_at")
    private LocalDateTime createdAt;
    
    @Column(name = "updated_at")
    private LocalDateTime updatedAt;
    
    // Getters, setters...
}
```

#### Step 7: Create Spring Data JPA Repository

```java
// repository/postgresql/SpringDataArticleRepository.java
package co3017.microservices.content_service.repository.postgresql;

import co3017.microservices.content_service.repository.postgresql.entity.ArticleEntity;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.JpaSpecificationExecutor;

import java.util.UUID;

public interface SpringDataArticleRepository 
        extends JpaRepository<ArticleEntity, UUID>, 
                JpaSpecificationExecutor<ArticleEntity> {
}
```

#### Step 8: Create Mapper

```java
// repository/postgresql/mapper/ArticleMapper.java
package co3017.microservices.content_service.repository.postgresql.mapper;

import co3017.microservices.content_service.models.Article;
import co3017.microservices.content_service.repository.postgresql.entity.ArticleEntity;

public class ArticleMapper {
    public static ArticleEntity toEntity(Article article) {
        ArticleEntity entity = new ArticleEntity();
        entity.setArticleId(article.getArticleId());
        entity.setTitle(article.getTitle());
        entity.setContent(article.getContent());
        entity.setAuthorId(article.getAuthorId());
        entity.setCreatedAt(article.getCreatedAt());
        entity.setUpdatedAt(article.getUpdatedAt());
        return entity;
    }
    
    public static Article toDomain(ArticleEntity entity) {
        return new Article(
            entity.getArticleId(),
            entity.getTitle(),
            entity.getContent(),
            entity.getAuthorId(),
            entity.getCreatedAt(),
            entity.getUpdatedAt()
        );
    }
}
```

#### Step 9: Create JPA Repository Implementation

```java
// repository/postgresql/JpaArticleRepository.java
package co3017.microservices.content_service.repository.postgresql;

import co3017.microservices.content_service.repository.ArticleRepository;
import co3017.microservices.content_service.models.Article;
import co3017.microservices.content_service.repository.postgresql.entity.ArticleEntity;
import co3017.microservices.content_service.repository.postgresql.mapper.ArticleMapper;
import co3017.microservices.content_service.repository.postgresql.specification.ArticleSpecification;
import co3017.microservices.content_service.usecase.types.ArticleSearchCriteria;
import co3017.microservices.content_service.usecase.types.ArticlePageResult;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.domain.Specification;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;
import java.util.UUID;
import java.util.stream.Collectors;

@Repository
public class JpaArticleRepository implements ArticleRepository {
    private final SpringDataArticleRepository springDataArticleRepository;
    
    public JpaArticleRepository(SpringDataArticleRepository springDataArticleRepository) {
        this.springDataArticleRepository = springDataArticleRepository;
    }
    
    @Override
    public Article save(Article article) {
        ArticleEntity entity = ArticleMapper.toEntity(article);
        ArticleEntity savedEntity = springDataArticleRepository.save(entity);
        return ArticleMapper.toDomain(savedEntity);
    }
    
    @Override
    public Optional<Article> findById(UUID articleId) {
        return springDataArticleRepository.findById(articleId)
            .map(ArticleMapper::toDomain);
    }
    
    @Override
    public ArticlePageResult search(ArticleSearchCriteria criteria) {
        Pageable pageable = PageRequest.of(criteria.getPage(), criteria.getSize());
        Specification<ArticleEntity> spec = ArticleSpecification.createSpecification(criteria);
        Page<ArticleEntity> page = springDataArticleRepository.findAll(spec, pageable);
        
        List<Article> articles = page.getContent().stream()
            .map(ArticleMapper::toDomain)
            .collect(Collectors.toList());
        
        return new ArticlePageResult(
            articles,
            page.getNumber(),
            page.getSize(),
            page.getTotalElements(),
            page.getTotalPages()
        );
    }
    
    // Implement other methods...
}
```

#### Step 10: Create Specification (if search is needed)

```java
// repository/postgresql/specification/ArticleSpecification.java
package co3017.microservices.content_service.repository.postgresql.specification;

import co3017.microservices.content_service.repository.postgresql.entity.ArticleEntity;
import co3017.microservices.content_service.usecase.types.ArticleSearchCriteria;
import org.springframework.data.jpa.domain.Specification;

import java.util.UUID;

public class ArticleSpecification {
    public static Specification<ArticleEntity> createSpecification(ArticleSearchCriteria criteria) {
        return Specification.where(titleLike(criteria.getTitle()))
            .and(authorIdEquals(criteria.getAuthorId()));
    }
    
    private static Specification<ArticleEntity> titleLike(String title) {
        return (root, query, cb) -> 
            title == null ? null : 
            cb.like(cb.lower(root.get("title")), "%" + title.toLowerCase() + "%");
    }
    
    private static Specification<ArticleEntity> authorIdEquals(UUID authorId) {
        return (root, query, cb) -> 
            authorId == null ? null : cb.equal(root.get("authorId"), authorId);
    }
}
```

#### Step 11: Create Request/Response DTOs

```java
// adapter/http/dto/CreateArticleRequest.java
package co3017.microservices.content_service.adapter.http.dto;

import java.util.UUID;

public class CreateArticleRequest {
    private String title;
    private String content;
    private UUID authorId;
    
    // Getters, setters...
}

// adapter/http/dto/UpdateArticleRequest.java
// adapter/http/dto/ArticleResponse.java
// adapter/http/dto/ArticlePageResponse.java
```

#### Step 12: Create Response Builders

```java
// adapter/http/response/CommandBuilder.java (add method)
public static CreateArticleCommand toCreateArticleCommand(CreateArticleRequest request) {
    return new CreateArticleCommand(
        request.getTitle(),
        request.getContent(),
        request.getAuthorId()
    );
}

// adapter/http/response/ArticleResponseBuilder.java
package co3017.microservices.content_service.adapter.http.response;

import co3017.microservices.content_service.adapter.http.dto.ArticleResponse;
import co3017.microservices.content_service.models.Article;

public class ArticleResponseBuilder {
    public static ArticleResponse toResponse(Article article) {
        ArticleResponse response = new ArticleResponse();
        response.setArticleId(article.getArticleId());
        response.setTitle(article.getTitle());
        response.setContent(article.getContent());
        response.setAuthorId(article.getAuthorId());
        response.setCreatedAt(article.getCreatedAt());
        response.setUpdatedAt(article.getUpdatedAt());
        return response;
    }
}
```

#### Step 13: Create Controller

```java
// adapter/http/ArticleController.java
package co3017.microservices.content_service.adapter.http;

import co3017.microservices.content_service.adapter.http.dto.ApiResponse;
import co3017.microservices.content_service.adapter.http.dto.CreateArticleRequest;
import co3017.microservices.content_service.adapter.http.dto.UpdateArticleRequest;
import co3017.microservices.content_service.adapter.http.dto.ArticleResponse;
import co3017.microservices.content_service.adapter.http.response.CommandBuilder;
import co3017.microservices.content_service.adapter.http.response.ArticleResponseBuilder;
import co3017.microservices.content_service.usecase.ArticleUseCase;
import co3017.microservices.content_service.usecase.types.ArticleSearchCriteria;
import co3017.microservices.content_service.usecase.types.ArticlePageResult;
import co3017.microservices.content_service.models.Article;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.UUID;

@RestController
@RequestMapping("/api/articles")
public class ArticleController {
    private final ArticleUseCase articleUseCase;
    
    public ArticleController(ArticleUseCase articleUseCase) {
        this.articleUseCase = articleUseCase;
    }
    
    @PostMapping
    public ResponseEntity<ApiResponse<ArticleResponse>> create(
            @RequestBody CreateArticleRequest request) {
        Article article = articleUseCase.create(CommandBuilder.toCreateArticleCommand(request));
        ArticleResponse response = ArticleResponseBuilder.toResponse(article);
        return ResponseEntity.status(HttpStatus.CREATED)
                .body(ApiResponse.success("Article created successfully", response));
    }
    
    @GetMapping("/{articleId}")
    public ResponseEntity<ApiResponse<ArticleResponse>> detail(@PathVariable UUID articleId) {
        return articleUseCase.detail(articleId)
                .map(article -> ResponseEntity.ok(
                        ApiResponse.success(ArticleResponseBuilder.toResponse(article))))
                .orElse(ResponseEntity.status(HttpStatus.NOT_FOUND)
                        .body(ApiResponse.error(404, "Article not found")));
    }
    
    // Implement other endpoints...
}
```

---

## Rules and Best Practices

### 1. Dependency Rule

- Domain layer does not import any other layer
- Application layer imports only Domain layer
- Interface Adapters import Application and Domain layers
- Infrastructure imports all layers above

### 2. Naming Conventions

- **Domain Models**: `Course`, `Chapter`, `ContentUnit` (PascalCase, singular)
- **UseCase Interfaces**: `CourseUseCase`, `ChapterUseCase` (PascalCase + "UseCase")
- **UseCase Services**: `CourseService`, `ChapterService` (PascalCase + "Service")
- **Repository Interfaces**: `CourseRepository`, `ChapterRepository` (PascalCase + "Repository")
- **JPA Repositories**: `JpaCourseRepository`, `JpaChapterRepository` ("Jpa" + PascalCase + "Repository")
- **Spring Data Repositories**: `SpringDataCourseRepository` ("SpringData" + PascalCase + "Repository")
- **Entities**: `CourseEntity`, `ChapterEntity` (PascalCase + "Entity")
- **Controllers**: `CourseController`, `ChapterController` (PascalCase + "Controller")
- **DTOs**: `CreateCourseRequest`, `CourseResponse` (PascalCase + "Request"/"Response")
- **Commands**: `CreateCourseCommand`, `UpdateCourseCommand` (PascalCase + "Command")
- **Mappers**: `CourseMapper`, `CourseResponseBuilder` (PascalCase + "Mapper"/"Builder")

### 3. Package Structure Rules

- Each domain must have all layers: Model → UseCase → Repository → Controller
- Do not skip layers (e.g., Controller must not call Repository directly)
- Each layer only calls the layer immediately below it

### 4. Code Organization

- Each file contains only one class/interface
- Each class/interface has only one responsibility (Single Responsibility Principle)
- Use dependency injection (constructor injection)
- Use immutable objects when possible (final fields)

### 5. Validation Rules

- **Request validation**: In Controller or UseCase Service
- **Business validation**: In UseCase Service
- **Domain validation**: In Domain Model (methods like `isValid()`)
- **Database constraints**: In Entity (JPA annotations)

### 6. Error Handling

- Use `IllegalArgumentException` for validation errors
- Controllers have `@ExceptionHandler` to handle exceptions
- Always return `ApiResponse<T>` with error code and message

### 7. Transaction Management

- UseCase Services have `@Transactional`
- Read-only operations have `@Transactional(readOnly = true)`
- Write operations have `@Transactional` (default)

### 8. Testing Strategy

- **Domain Models**: Unit tests without Spring context
- **UseCase Services**: Unit tests with mocked Repository
- **Controllers**: Integration tests with `@WebMvcTest`
- **Repositories**: Integration tests with `@DataJpaTest`

---

## Real Examples

Refer to existing files in the project:

- **Domain Model**: `models/Course.java`
- **UseCase Interface**: `usecase/CourseUseCase.java`
- **UseCase Service**: `usecase/service/CourseService.java`
- **Repository Interface**: `repository/CourseRepository.java`
- **JPA Repository**: `repository/postgresql/JpaCourseRepository.java`
- **Entity**: `repository/postgresql/entity/CourseEntity.java`
- **Mapper**: `repository/postgresql/mapper/CourseMapper.java`
- **Controller**: `adapter/http/CourseController.java`
- **DTOs**: `adapter/http/dto/CreateCourseRequest.java`, `CourseResponse.java`
- **Builders**: `adapter/http/response/CommandBuilder.java`, `CourseResponseBuilder.java`

---

## Checklist for Adding New Features

- [ ] Create Domain Model in `models/`
- [ ] Create UseCase Interface in `usecase/`
- [ ] Create Command/Query types in `usecase/types/`
- [ ] Create UseCase Service implementation in `usecase/service/`
- [ ] Create Repository Interface in `repository/`
- [ ] Create JPA Entity in `repository/postgresql/entity/`
- [ ] Create Spring Data Repository in `repository/postgresql/`
- [ ] Create Mapper in `repository/postgresql/mapper/`
- [ ] Create JPA Repository implementation in `repository/postgresql/`
- [ ] Create Specification (if search needed) in `repository/postgresql/specification/`
- [ ] Create Request/Response DTOs in `adapter/http/dto/`
- [ ] Create Response Builders in `adapter/http/response/`
- [ ] Create Controller in `adapter/http/`
- [ ] Add exception handling in Controller
- [ ] Write unit tests for Domain Model
- [ ] Write unit tests for UseCase Service
- [ ] Write integration tests for Controller
- [ ] Write integration tests for Repository

---

## References

- [Clean Architecture by Robert C. Martin](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Spring Boot Documentation](https://docs.spring.io/spring-boot/docs/current/reference/htmlsingle/)
- [Spring Data JPA Documentation](https://docs.spring.io/spring-data/jpa/docs/current/reference/html/)
- [JPA Specifications](https://www.baeldung.com/spring-data-jpa-specifications)

---

**Note**: This document describes the current structure of the project. When adding new features, strictly follow the established rules and patterns to ensure consistency and maintainability of the codebase.
