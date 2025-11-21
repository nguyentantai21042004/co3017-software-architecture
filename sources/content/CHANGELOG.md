# Changelog - Content Service

All notable changes to the Content Service will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [2.0.0] - 2025-11-22

### Changed - Major Refactoring
- **BREAKING**: Complete refactoring to Clean Architecture with Layer-First approach
- **BREAKING**: Standardized API response format to match microservices standard
- **BREAKING**: Migrated from `content` package to `content_service` package
- Removed all old content package files and consolidate to content_service

### Added
- Clean Architecture implementation with Layer-First structure
- Comprehensive unit tests achieving 80%+ coverage on core business logic
- Standardized API response format matching Go services
- Domain-driven design with pure business logic in models layer
- Command and Query pattern for use case inputs
- Mapper pattern for entity-domain conversion
- JPA Specification support for dynamic queries

### Architecture - Layer-First Clean Architecture
- **Domain Layer** (`models/`):
  - `Question.java` - Pure domain model with business logic
  - No framework dependencies
  - Business validation methods
  - Domain operations (updateContent, updateDifficulty, updateSkillTag)

- **Application Layer** (`usecase/`):
  - `QuestionUseCase.java` - Business operations interface
  - `usecase/service/QuestionService.java` - Implementation with @Service, @Transactional
  - `usecase/types/` - Commands and Queries:
    - `CreateQuestionCommand.java`
    - `UpdateQuestionCommand.java`
    - `QuestionQuery.java`

- **Repository Layer** (`repository/`):
  - `QuestionRepository.java` - Domain repository interface
  - `postgresql/QuestionJpaRepository.java` - Spring Data JPA repository
  - `postgresql/QuestionRepositoryImpl.java` - Implementation bridge
  - `postgresql/entity/QuestionEntity.java` - JPA entity with annotations
  - `postgresql/mapper/QuestionMapper.java` - Domain ↔ Entity conversion

- **Adapter Layer** (`adapter/http/`):
  - `QuestionController.java` - REST controller with @RestController
  - `dto/` - Request/Response DTOs:
    - `CreateQuestionRequest.java` - with @Valid, @NotBlank
    - `UpdateQuestionRequest.java` - for partial updates
    - `QuestionResponse.java` - with fromDomain factory method
  - `response/ApiResponse.java` - Standard response wrapper

### API Response Format
Updated from custom format to standardized structure:

**Before:**
```json
{
  "success": true,
  "message": "Success",
  "data": {...}
}
```

**After (matching Go services):**
```json
{
  "error_code": 0,
  "message": "Success",
  "data": {...},
  "errors": null
}
```

- Success responses: `error_code: 0`
- Error responses: `error_code: 400/404/500`
- Snake_case JSON fields with @JsonProperty annotations
- @JsonInclude(NON_NULL) for cleaner responses

### API Endpoints
All endpoints now return standardized ApiResponse:

- `POST /api/questions` - Create question (201 Created)
- `PUT /api/questions/{id}` - Update question (200 OK, 404 Not Found)
- `GET /api/questions/{id}` - Get question by ID (200 OK, 404 Not Found)
- `GET /api/questions?difficulty={}&skillTag={}` - List with filters (200 OK)
- `DELETE /api/questions/{id}` - Delete question (200 OK, 404 Not Found)

### Testing
- **QuestionServiceTest** - 16 comprehensive unit tests:
  - Create operations (success, validation failures)
  - Update operations (full/partial, not found)
  - Read operations (by ID, with filters, empty results)
  - Delete operations (success, not found)
  - Coverage: ~95% on business logic

- **QuestionControllerTest** - 13 integration tests:
  - All HTTP endpoints covered
  - Success and error response validation
  - JSON serialization verification
  - Query parameter handling
  - Coverage: ~90% on controller layer

**Test Results:**
```
Tests run: 29, Failures: 0, Errors: 0, Skipped: 0
BUILD SUCCESS
```

### Technical Stack
- Java 17
- Spring Boot 3.5.6
- Spring Data JPA / Hibernate
- PostgreSQL database
- Jakarta Validation
- Lombok for boilerplate reduction
- Jackson for JSON processing
- JUnit 5 + Mockito for testing
- MockMvc for controller testing

### Database
- Table: `questions`
- Fields: id, content, difficulty, skill_tag, created_at, updated_at
- Indexes and constraints as per JPA entity definition

### Configuration
- `server.port` - Service port (default: 8081)
- `spring.datasource.*` - PostgreSQL connection
- `spring.jpa.*` - JPA/Hibernate settings

### Removed
- Old `co3017/microservices/content` package (6 files)
- Legacy QuestionService, QuestionController, QuestionRepository implementations
- Old DTO classes (ApiResponse, QuestionResponse)
- Outdated architecture mixing concerns

### Migration Notes
For developers updating from v1.x:
1. Update import statements from `content` to `content_service` package
2. Update API response parsing to handle new format (error_code, snake_case fields)
3. Use new endpoint paths under `/api/questions`
4. Expect standardized error responses with error_code field

## [1.0.0] - 2025-11-21

### Added
- Initial implementation with basic CRUD operations
- Question management API
- PostgreSQL integration
- Sample data seeding
