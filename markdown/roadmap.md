# Roadmap Hoàn Thành Đề Tài ITS (Intelligent Tutoring System)

## Tổng Quan Đề Tài

**Mục tiêu chính:** Thiết kế kiến trúc phần mềm và triển khai Hệ thống Gia sư Thông minh (ITS) sử dụng các nguyên tắc SOLID.

**Yêu cầu cốt lõi:**
- Kiến trúc linh hoạt và có khả năng mở rộng (scalable)
- Xử lý nhiều ngữ cảnh học tập khác nhau
- Đánh giá kiến thức, cung cấp phản hồi, đề xuất lộ trình học tập tùy chỉnh
- Áp dụng nguyên tắc SOLID

---

## Phase 1: Phân Tích và Thiết Kế Kiến Trúc (Software Architecture Design)

### 1.1 Phân Tích Yêu Cầu (Requirements Analysis)

**Mục tiêu:** Xác định rõ ràng các yêu cầu chức năng và phi chức năng

#### Kiến thức cần có trước:
- Hiểu về phân tích yêu cầu phần mềm
- Biết phân biệt functional vs non-functional requirements
- Hiểu về domain modeling

#### Công việc cần làm:
1. **Xác định Functional Requirements:**
   - Học tập cá nhân hóa (Personalized Learning)
   - Hệ thống phản hồi (Feedback System)
   - Đánh giá và thẩm định (Assessment & Evaluation)
   - Bảng điều khiển giảng viên (Instructor Dashboard)
   - Quản lý nội dung học tập (Content Management)

2. **Xác định Non-Functional Requirements:**
   - Performance (hiệu suất)
   - Scalability (khả năng mở rộng)
   - Usability (khả năng sử dụng)
   - Reliability (độ tin cậy)
   - Security (bảo mật)

#### Kiến thức sẽ có sau:
- Hiểu rõ về domain ITS
- Biết cách phân tích và mô hình hóa yêu cầu
- Có bản spec chi tiết cho hệ thống

### 1.2 Thiết Kế Architecture Characteristics

**Mục tiêu:** Xác định các tiêu chí thành công của ITS

#### Kiến thức cần có trước:
- Hiểu về architecture characteristics
- Biết các loại quality attributes
- Hiểu về trade-offs trong thiết kế

#### Công việc cần làm:
1. **Xác định Primary Characteristics:**
   - Performance (response time, throughput)
   - Scalability (concurrent users, data volume)
   - Availability (uptime, fault tolerance)
   - Usability (user experience, accessibility)

2. **Xác định Secondary Characteristics:**
   - Security (authentication, authorization)
   - Maintainability (code quality, documentation)
   - Extensibility (new features, integrations)
   - Testability (unit tests, integration tests)

#### Kiến thức sẽ có sau:
- Hiểu cách đánh giá và ưu tiên các đặc tính kiến trúc
- Biết cách balance giữa các yêu cầu mâu thuẫn
- Có framework để đánh giá thiết kế

### 1.3 So Sánh và Lựa Chọn Architecture Styles

**Mục tiêu:** Chọn kiến trúc phù hợp cho ITS

#### Kiến thức cần có trước:
- Hiểu về các architecture patterns phổ biến
- Biết ưu nhược điểm của từng pattern
- Hiểu về distributed systems

#### Công việc cần làm:
1. **Phân tích các Architecture Styles:**
   - Layered Architecture
   - Microservices Architecture
   - Event-Driven Architecture
   - Hexagonal Architecture
   - Clean Architecture

2. **So sánh và Đánh giá:**
   - Tạo bảng so sánh các đặc điểm
   - Đánh giá theo architecture characteristics đã định
   - Lựa chọn và justify decision

#### Kiến thức sẽ có sau:
- Hiểu sâu về các architecture patterns
- Biết cách đánh giá và lựa chọn architecture
- Có khả năng design system architecture

### 1.4 Thiết Kế Architecture Views

**Mục tiêu:** Trình bày kiến trúc theo các góc nhìn khác nhau

#### Kiến thức cần có trước:
- Hiểu về 4+1 View Model
- Biết về UML diagrams
- Hiểu về component modeling

#### Công việc cần làm:
1. **Module Views:**
   - Package diagram
   - Module dependency diagram
   - Layered architecture diagram

2. **Component-and-Connector Views:**
   - Component diagram
   - Deployment diagram
   - Sequence diagram cho key scenarios

3. **Allocation Views:**
   - Deployment diagram
   - Process diagram
   - Development allocation diagram

#### Kiến thức sẽ có sau:
- Biết cách document architecture đầy đủ
- Hiểu về multiple views của system
- Có skill vẽ và đọc architecture diagrams

### 1.5 Architecture Decisions và Design Principles

**Mục tiêu:** Xác định quy tắc và nguyên tắc thiết kế

#### Kiến thức cần có trước:
- Hiểu về Architecture Decision Records (ADRs)
- Biết về design principles
- Hiểu về architectural patterns

#### Công việc cần làm:
1. **Architecture Decisions:**
   - Document key architectural decisions
   - Rationale và alternatives considered
   - Consequences và trade-offs

2. **Design Principles:**
   - SOLID principles application
   - Domain-driven design principles
   - Clean code principles

#### Kiến thức sẽ có sau:
- Biết cách document architectural decisions
- Hiểu cách apply design principles
- Có framework cho architectural thinking

### 1.6 Áp Dụng SOLID Principles

**Mục tiêu:** Giải thích cách áp dụng SOLID trong thiết kế

#### Kiến thức cần có trước:
- Hiểu rõ 5 nguyên tắc SOLID
- Biết cách identify violations
- Hiểu về design patterns

#### Công việc cần làm:
1. **Single Responsibility Principle (SRP):**
   - Identify responsibilities trong ITS
   - Design classes với single responsibility
   - Examples và diagrams

2. **Open/Closed Principle (OCP):**
   - Design cho extension không modification
   - Strategy pattern applications
   - Plugin architecture considerations

3. **Liskov Substitution Principle (LSP):**
   - Design inheritance hierarchies
   - Interface segregation
   - Polymorphism examples

4. **Interface Segregation Principle (ISP):**
   - Design focused interfaces
   - Avoid fat interfaces
   - Client-specific interfaces

5. **Dependency Inversion Principle (DIP):**
   - Design với dependency injection
   - Abstract dependencies
   - Inversion of control containers

#### Kiến thức sẽ có sau:
- Master SOLID principles
- Biết cách apply SOLID trong real projects
- Có skill refactor code theo SOLID

### 1.7 Reflection Report

**Mục tiêu:** Viết báo cáo phản ánh về việc áp dụng SOLID

#### Kiến thức cần có trước:
- Đã hoàn thành các phần trên
- Có kinh nghiệm thực tế với SOLID
- Biết cách viết technical reports

#### Công việc cần làm:
1. **Reflection Content:**
   - How SOLID improved the design
   - Challenges encountered
   - Trade-offs made
   - Lessons learned

2. **Report Structure:**
   - Executive summary
   - Detailed analysis
   - Examples và code snippets
   - Conclusions và recommendations

#### Kiến thức sẽ có sau:
- Kỹ năng technical writing
- Khả năng self-reflection về design
- Experience với architectural thinking

---

## Phase 2: Implementation (Code Implementation)

### 2.1 Lựa Chọn Technology Stack

**Mục tiêu:** Chọn công nghệ phù hợp cho implementation

#### Kiến thức cần có trước:
- Hiểu về modern programming languages
- Biết về frameworks và libraries
- Hiểu về database technologies

#### Công việc cần làm:
1. **Backend Technology:**
   - Programming language (Java, C#, Python, etc.)
   - Framework (Spring, .NET, Django, etc.)
   - Database (PostgreSQL, MongoDB, etc.)
   - API framework (REST, GraphQL)

2. **Frontend Technology:**
   - JavaScript framework (React, Vue, Angular)
   - UI/UX libraries
   - State management

3. **Infrastructure:**
   - Containerization (Docker)
   - CI/CD pipeline
   - Cloud platform (AWS, Azure, GCP)

#### Kiến thức sẽ có sau:
- Hiểu về technology selection process
- Biết về modern development stack
- Có experience với chosen technologies

### 2.2 Setup Development Environment

**Mục tiêu:** Chuẩn bị môi trường phát triển

#### Kiến thức cần có trước:
- Biết về version control (Git)
- Hiểu về IDE và development tools
- Biết về project structure

#### Công việc cần làm:
1. **Repository Setup:**
   - Initialize Git repository
   - Setup .gitignore
   - Create branch strategy

2. **Project Structure:**
   - Create project folders
   - Setup build configuration
   - Configure IDE settings

3. **Development Tools:**
   - Code formatter và linter
   - Testing framework setup
   - Documentation tools

#### Kiến thức sẽ có sau:
- Professional development setup
- Best practices cho project organization
- Experience với development tools

### 2.3 Implement Core Module (Chọn ít nhất 1 module)

**Mục tiêu:** Triển khai ít nhất một module chính của ITS

#### Kiến thức cần có trước:
- Đã hoàn thành architecture design
- Có experience với chosen technology
- Hiểu về SOLID principles

#### Lựa chọn Modules:

**Option 1: Personalized Learning Module**
- User profiling system
- Learning path recommendation
- Content adaptation engine
- Progress tracking

**Option 2: Assessment & Evaluation Module**
- Question bank management
- Quiz/exam creation
- Automated grading
- Performance analytics

**Option 3: Feedback System Module**
- Real-time feedback generation
- Hint system
- Progress indicators
- Motivational messages

#### Công việc cần làm cho Module đã chọn:
1. **Domain Modeling:**
   - Identify entities và value objects
   - Define domain services
   - Create domain events

2. **Application Layer:**
   - Use cases implementation
   - Application services
   - DTOs và mappers

3. **Infrastructure Layer:**
   - Repository implementations
   - External service integrations
   - Database access

4. **Presentation Layer:**
   - REST API endpoints
   - Request/response handling
   - Error handling

5. **Testing:**
   - Unit tests cho domain logic
   - Integration tests cho APIs
   - End-to-end tests cho user scenarios

#### Kiến thức sẽ có sau:
- Practical experience với clean architecture
- SOLID principles application in real code
- Full-stack development skills
- Testing best practices

### 2.4 Code Quality và Documentation

**Mục tiêu:** Đảm bảo code quality và documentation

#### Kiến thức cần có trước:
- Hiểu về code quality metrics
- Biết về documentation standards
- Experience với code review

#### Công việc cần làm:
1. **Code Quality:**
   - Code review process
   - Static analysis tools
   - Code coverage metrics
   - Performance profiling

2. **Documentation:**
   - API documentation (Swagger/OpenAPI)
   - Code comments và JSDoc
   - README files
   - Architecture documentation

3. **SOLID Compliance:**
   - Verify SOLID principles application
   - Refactor nếu cần
   - Document design decisions

#### Kiến thức sẽ có sau:
- Professional code quality practices
- Technical writing skills
- Code review experience

---

## Phase 3: Documentation và Submission

### 3.1 Final Report Compilation

**Mục tiêu:** Tổng hợp báo cáo cuối cùng

#### Kiến thức cần có trước:
- Đã hoàn thành tất cả phases trước
- Có technical writing skills
- Hiểu về report structure

#### Công việc cần làm:
1. **Report Structure:**
   - Executive summary
   - Problem statement và objectives
   - Architecture design details
   - Implementation details
   - SOLID principles application
   - Reflection và lessons learned
   - Conclusions và future work

2. **Supporting Materials:**
   - Architecture diagrams
   - Code snippets
   - Screenshots của system
   - Test results

#### Kiến thức sẽ có sau:
- Professional technical writing
- Project presentation skills
- Comprehensive understanding của ITS

### 3.2 GitHub Repository Setup

**Mục tiêu:** Chuẩn bị repository cho submission

#### Kiến thức cần có trước:
- Git proficiency
- Repository management
- Open source best practices

#### Công việc cần làm:
1. **Repository Organization:**
   - Clear folder structure
   - Comprehensive README
   - Contributing guidelines
   - License information

2. **Documentation:**
   - Installation instructions
   - Usage examples
   - API documentation
   - Architecture overview

3. **Code Quality:**
   - Clean commit history
   - Proper branch management
   - Issue tracking setup

#### Kiến thức sẽ có sau:
- Professional repository management
- Open source contribution experience
- Portfolio development

---

## Timeline và Milestones

### Week 1-2: Requirements Analysis & Architecture Design
- [ ] Complete requirements analysis
- [ ] Define architecture characteristics
- [ ] Select architecture style

### Week 3-4: Architecture Views & SOLID Application
- [ ] Create architecture views
- [ ] Document architecture decisions
- [ ] Apply SOLID principles in design

### Week 5-6: Technology Selection & Environment Setup
- [ ] Choose technology stack
- [ ] Setup development environment
- [ ] Create project structure

### Week 7-10: Core Module Implementation
- [ ] Implement chosen module
- [ ] Apply SOLID principles in code
- [ ] Write comprehensive tests

### Week 11-12: Documentation & Submission
- [ ] Complete reflection report
- [ ] Finalize documentation
- [ ] Prepare GitHub repository
- [ ] Submit final report

---

## Kiến Thức Tổng Quan Cần Có

### Trước khi bắt đầu:
- [ ] Software Architecture fundamentals
- [ ] SOLID principles understanding
- [ ] Object-oriented programming
- [ ] Database design basics
- [ ] REST API concepts
- [ ] Version control (Git)

### Sau khi hoàn thành:
- [ ] Advanced software architecture design
- [ ] SOLID principles mastery
- [ ] Clean architecture implementation
- [ ] Full-stack development experience
- [ ] Technical writing skills
- [ ] Project management skills
- [ ] Code quality best practices

---

## Resources và References

### Books:
- "Clean Architecture" by Robert C. Martin
- "Software Architecture in Practice" by Len Bass
- "Domain-Driven Design" by Eric Evans
- "Patterns of Enterprise Application Architecture" by Martin Fowler

### Online Resources:
- Microsoft Architecture Center
- AWS Well-Architected Framework
- Google Cloud Architecture Framework
- SOLID Principles tutorials

### Tools:
- Draw.io (architecture diagrams)
- Visual Studio Code
- Postman (API testing)
- Docker (containerization)
- Git/GitHub (version control)

## Kế Hoạch Chi Tiết 14 Ngày & Giao Phẩm (Phase 1 Finalization)

---

### Giai đoạn I: Đảm bảo Nền tảng và Quyết định Cấu trúc (Ngày 1–5)

Giai đoạn này tập trung vào sự đồng thuận về yêu cầu, mô hình hóa domain, và quyết định kiến trúc cấp cao.

| **Ngày** | **Nhiệm vụ chi tiết** | **Giao phẩm Cần đạt (Output)** | **Phân công** |
|----------|----------------------|-------------------------------|---------------|
| **1** | **Chốt Yêu cầu (1.1 Finalization):**<br>Kiểm tra chéo toàn bộ FRs, NFRs/ACs và User Stories. Đảm bảo mỗi FR đều có ACs hỗ trợ. | **Tài liệu Specification Final:**<br>Bảng FRs, ACs/Trade-offs đã được ký xác nhận. | Dev 1 |
| **2-3** | **Kiểm tra Domain Model:**<br>Review lại các Aggregates (LearnerModel, Content), Domain Services (ScoringEngine). Xác nhận ranh giới này tuân thủ SRP. | **Domain Model Verification Report:**<br>Xác nhận các Boundary của 5 Services là hợp lý và tuân thủ SRP. | Dev 2 |
| **4-5** | **Chốt Quyết định Cấu trúc (1.3 & 1.5):**<br>Xác nhận ADR-1 (Microservices) và ADR-5 (Polyglot Programming). Phân bổ Service và Stack cuối cùng. | **ADR-1 & ADR-5 Final:**<br>Chốt 5 Microservice và Stack (Golang/Java/Postgres) cho từng Service. | Dev 1 & Lead |

**Mục tiêu Giai đoạn I:**
- ✅ Specification hoàn chỉnh và được phê duyệt
- ✅ Domain boundaries rõ ràng (5 services)
- ✅ Architecture decisions đã được documented (ADR)

---

### Giai đoạn II: Trực quan hóa và Triển khai (Ngày 6–10)

Giai đoạn này tập trung vào việc tạo ra các sơ đồ kiến trúc (Views) và quyết định quy tắc kỹ thuật.

| **Ngày** | **Nhiệm vụ chi tiết** | **Giao phẩm Cần đạt (Output)** | **Phân công** |
|----------|----------------------|-------------------------------|---------------|
| **6-7** | **Module Views (1.4.1):**<br>Vẽ Internal Clean Architecture Diagram cho Service Golang (Scoring Engine). Thể hiện rõ các Interfaces (Abstraction) và Implementations (Concretion). | **Clean Architecture Diagram:**<br>Sơ đồ 4 lớp (Domain → Infrastructure) với các Interfaces cốt lõi. | Dev 2 |
| **8-9** | **Component-and-Connector & Allocation Views (1.4.2 & 1.4.3):**<br>Vẽ Container Diagram (dựa trên ADR-1, ADR-2) và Deployment Diagram (Kubernetes, Polyglot DBs). | **Diagrams Set:**<br>1. Container Diagram<br>2. Deployment Diagram (K8s/DBs)<br>3. Sequence Diagram (UC-L-02) | Dev 3 |
| **10** | **Chốt Quyết định Kỹ thuật (1.5):**<br>Lead review ADR-2 (Kafka), ADR-3 (Clean Arch), và ADR-4 (Polyglot Persistence). Đảm bảo chúng phù hợp với các Diagrams (I-II). | **ADR Finalization Report:**<br>Ký xác nhận các quyết định về Công nghệ Giao tiếp (Kafka) và Chiến lược DB (Postgres/NoSQL/Redis). | Lead |

**Mục tiêu Giai đoạn II:**
- ✅ Architecture views hoàn chỉnh (Module, C&C, Allocation)
- ✅ Technical ADRs được finalized
- ✅ Diagrams ready cho presentation

---

### Giai đoạn III: Chứng minh SOLID và Tổng kết (Ngày 11–14)

Giai đoạn này tập trung vào việc chứng minh sự tuân thủ SOLID (1.6) và báo cáo cuối cùng (1.7).

| **Ngày** | **Nhiệm vụ chi tiết** | **Giao phẩm Cần đạt (Output)** | **Phân công** |
|----------|----------------------|-------------------------------|---------------|
| **11-12** | **Chứng minh SOLID (1.6) - Part 1:**<br>Lấy ví dụ SRP & DIP (Golang/Java) cho Adaptive Engine và OCP (Strategy Pattern) cho Scoring Engine. | **SOLID Code Examples:**<br>Tài liệu hóa các ví dụ thiết kế Class/Interface để chứng minh tuân thủ SRP, OCP, DIP. | Dev 1 |
| **13** | **Chứng minh SOLID (1.6) - Part 2:**<br>Lấy ví dụ ISP (Interface Segregation) cho LearnerModel Repository và LSP (Inheritance Design) cho các loại Assessment. | **SOLID Design Documentation:**<br>Tài liệu hóa các ví dụ thiết kế Class/Interface để chứng minh tuân thủ ISP, LSP. | Dev 2 |
| **14** | **Kiểm tra cuối cùng & Báo cáo (1.7):**<br>Lead thực hiện kiểm tra chéo cuối cùng của toàn bộ tài liệu Phase 1. Dev 3 tổng hợp Reflection Report. | **Tài liệu Phase 1 Hoàn chỉnh (Final Delivery):**<br>Bao gồm Reflection Report. Sẵn sàng chuyển sang Phase 2: Code. | Lead & Dev 3 |

**Mục tiêu Giai đoạn III:**
- ✅ SOLID principles được chứng minh với code examples
- ✅ Reflection report hoàn chỉnh
- ✅ Phase 1 deliverables sẵn sàng submit

---

### Checklist Giao Phẩm Phase 1 (14-Day Sprint)

#### 📋 **Documentation Deliverables**
- [ ] 1.1 Specification Final (FRs, NFRs, ACs, Trade-offs)
- [ ] 1.2 Domain Model Verification Report
- [ ] 1.3 Architecture Style Decision (ADR-1)
- [ ] 1.4 Architecture Views:
  - [ ] Module View (Clean Architecture)
  - [ ] Component & Connector View (Container Diagram)
  - [ ] Allocation View (Deployment Diagram)
  - [ ] Sequence Diagram (UC-L-02)
- [ ] 1.5 ADR Finalization:
  - [ ] ADR-2: Event-Driven (Kafka)
  - [ ] ADR-3: Clean Architecture
  - [ ] ADR-4: Polyglot Persistence
  - [ ] ADR-5: Polyglot Programming
- [ ] 1.6 SOLID Principles Documentation:
  - [ ] SRP examples (Golang/Java)
  - [ ] OCP examples (Strategy Pattern)
  - [ ] LSP examples (Assessment types)
  - [ ] ISP examples (Repository interfaces)
  - [ ] DIP examples (Adaptive Engine)
- [ ] 1.7 Reflection Report

#### 🎯 **Quality Gates**
- [ ] All FRs mapped to ACs
- [ ] All 5 service boundaries verified (SRP)
- [ ] All ADRs reviewed and signed off
- [ ] All diagrams peer-reviewed
- [ ] SOLID examples validated with code

---

### Timeline Visualization

```
Week 1: Foundation & Structure
├── Day 1    │ Specification Finalization
├── Day 2-3  │ Domain Model Verification  
├── Day 4-5  │ ADR-1 & ADR-5 Finalization
└─────────────────────────────────────────

Week 2: Visualization & SOLID
├── Day 6-7  │ Module Views (Clean Arch)
├── Day 8-9  │ C&C + Allocation Views
├── Day 10   │ ADR-2,3,4 Finalization
├── Day 11-12│ SOLID Examples (SRP,OCP,DIP)
├── Day 13   │ SOLID Examples (ISP,LSP)
└── Day 14   │ Final Review & Reflection Report
```

**Critical Path:** Ngày 1 → Ngày 2-3 → Ngày 4-5 → Ngày 10 → Ngày 14

---
