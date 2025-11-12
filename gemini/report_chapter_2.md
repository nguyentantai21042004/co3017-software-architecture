Chắc chắn rồi. Dưới đây là nội dung chi tiết cho Mục 2.3, được xây dựng chính xác dựa trên nội dung các tệp bạn đã cung cấp và tuân thủ mẫu ADR (Architecture Decision Record) mà bạn yêu cầu.

-----

### 2.3 Architecture Decision Records

Phần này ghi lại các quyết định kiến trúc quan trọng (ADR) đã được đưa ra để định hình hệ thống ITS. Mỗi ADR tuân theo một cấu trúc thống nhất để làm rõ bối cảnh, quyết định, lý do và các hậu quả đi kèm.

## ADR-1: Polyglot Programming Strategy

### Status

✅ **Accepted**
*(Date: 2025-10-13, Deciders: Architecture Team)*

### Context

Hệ thống ITS bao gồm nhiều loại dịch vụ (services) với các yêu cầu phi chức năng khác nhau:

  * **Management Services** (Quản lý User, Content): Đòi hỏi logic nghiệp vụ phức tạp, cần hệ sinh thái (ecosystem) rộng lớn và khả năng bảo trì (Maintainability) cao.
  * **Computation Services** (Chấm điểm, Phản hồi): Đòi hỏi hiệu suất (Performance) rất cao, độ trễ thấp (≤500ms), và khả năng xử lý đồng thời (Concurrency) xuất sắc để phản hồi thời gian thực.
  * **AI/ML Services** (Lõi thích ứng): Yêu cầu xử lý tính toán CPU chuyên sâu (CPU-intensive) và lặp lại thuật toán nhanh.
  * **Ràng buộc:** Đội ngũ phát triển có kinh nghiệm với cả Java và Golang.

### Decision

Sử dụng chiến lược **Polyglot Programming (Đa ngôn ngữ)**. Phân chia công nghệ dựa trên vai trò của service:

  * **Java 17+ (Spring Boot 3.x):** Dùng cho các service thiên về nghiệp vụ và bảo trì:
      * `User Management Service`
      * `Content Service`
  * **Golang 1.21+ (Gin/Echo):** Dùng cho các service thiên về hiệu năng, đồng thời và độ trễ thấp:
      * `Scoring/Feedback Service`
      * `Adaptive Engine`
      * `Learner Model Service`
      * `API Gateway`

### Rationale

  * **Supports AC-3 (Performance) & AC-2 (Scalability):** Golang cung cấp hiệu suất gần bằng C, thời gian khởi động nhanh, và mô hình đồng thời (goroutines) xuất sắc, lý tưởng cho các dịch vụ thời gian thực (Scoring) và xử lý AI (Adaptive Engine).
  * **Supports AC-7 (Maintainability):** Java/Spring Boot có một hệ sinh thái trưởng thành (Spring Security, Spring Data JPA/Hibernate) để xử lý logic nghiệp vụ phức tạp, RBAC, và các truy vấn ORM phức tạp, giúp code dễ đọc và bảo trì.
  * **Addresses requirement:** Tối ưu hóa việc sử dụng công nghệ phù hợp nhất cho từng bài toán cụ thể, thay vì dùng "một búa cho mọi loại đinh".
  * **Mitigates risk:** Giảm thiểu rủi ro dùng một ngôn ngữ duy nhất và thất bại:
      * Nếu chỉ dùng Java: Rủi ro thất bại về **AC-3 (Performance)** ở các dịch vụ AI/scoring.
      * Nếu chỉ dùng Golang: Rủi ro thất bại về **AC-7 (Maintainability)** khi xử lý nghiệp vụ phức tạp (ví dụ: RBAC, quản lý nội dung).

### Consequences

**Positive:**

  * ✅ Tối ưu hóa hiệu năng cho các dịch vụ thời gian thực (real-time services).
  * ✅ Tăng khả năng bảo trì cho các dịch vụ logic nghiệp vụ (business logic services).
  * ✅ Cho phép đội ngũ tận dụng điểm mạnh của từng ngôn ngữ.

**Negative:**

  * ❌ Đội ngũ cần có chuyên môn ở cả hai ngôn ngữ.
  * ❌ Sử dụng các bộ công cụ (tooling) khác nhau (Maven/Gradle vs Go modules).
  * ❌ Các framework kiểm thử khác nhau (JUnit vs Go testing).
  * ❌ Tăng độ phức tạp trong pipeline CI/CD.

**Risks:**

  * **Risk 1:** Sự không nhất quán và tăng độ phức tạp trong phát triển và vận hành do sử dụng hai hệ sinh thái song song.
  * **Mitigation:** Tổ chức các buổi đào tạo chéo; chuẩn hóa cấu trúc dự án cho cả hai ngôn ngữ; sử dụng các mẫu (templates) CI/CD dùng chung; và áp dụng định dạng monitoring/logging thống nhất.

### Alternatives Considered

1.  **Option A: Tất cả bằng Java**
      * **Pros:** Chỉ cần chuyên môn một ngôn ngữ.
      * **Cons:** Hiệu năng thấp hơn cho các dịch vụ thời gian thực; footprint bộ nhớ cao hơn (ảnh hưởng chi phí scaling).
      * **Reason rejected:** Không đáp ứng mục tiêu **AC-3 (Performance)**.
2.  **Option B: Tất cả bằng Golang**
      * **Pros:** Bộ công cụ nhất quán.
      * **Cons:** Hệ sinh thái chưa trưởng thành bằng Java cho các nghiệp vụ phức tạp; không có DI (Dependency Injection) tích hợp sẵn.
      * **Reason rejected:** Không đáp ứng mục tiêu **AC-7 (Maintainability)** cho các nghiệp vụ phức tạp.

### Related Decisions

  * **Influences:** ADR-3 (Clean Architecture phải được triển khai ở cả hai ngôn ngữ), ADR-5 (Chiến lược kiểm thử phải bao gồm cả hai hệ sinh thái Java và Golang).

-----

## ADR-2: PostgreSQL as Primary Relational Database

### Status

✅ **Accepted**
*(Date: 2025-10-13, Deciders: Architecture Team)*

### Context

Các dịch vụ `User Management` và `Content Service` cần một cơ sở dữ liệu quan hệ (relational database) để đảm bảo:

  * Tính toàn vẹn giao dịch (ACID compliance).
  * Hỗ trợ các truy vấn phức tạp (JOINs, aggregations).
  * Quản lý xác thực và phân quyền (RBAC).
  * Hỗ trợ lưu trữ dữ liệu linh hoạt (JSON) cho metadata.
  * Yêu cầu là mã nguồn mở và có cơ chế replication/backup trưởng thành.

### Decision

Sử dụng **PostgreSQL 15+** làm cơ sở dữ liệu quan hệ chính cho `User Management Service` và `Content Service`.
Cấu hình sẽ bao gồm:

  * Primary-Standby replication.
  * Connection pooling (ví dụ: PgBouncer).
  * WAL archiving (lưu trữ WAL) để phục hồi theo thời điểm (point-in-time recovery).

### Rationale

  * **Supports AC-6 (Security):** PostgreSQL cung cấp các tính năng bảo mật nâng cao như Row-level security (Bảo mật cấp độ hàng).
  * **Supports AC-7 (Maintainability):** Đảm bảo ACID và tính toàn vẹn dữ liệu mạnh mẽ.
  * **Addresses requirement:** PostgreSQL có khả năng hỗ trợ JSON/JSONB vượt trội (cho metadata linh hoạt) và xử lý các truy vấn phức tạp tốt hơn MySQL.
  * **Mitigates risk:** Giảm thiểu rủi ro mất mát hoặc không nhất quán dữ liệu cho các thông tin quan trọng (user, content) mà NoSQL có thể gặp phải.

### Consequences

**Positive:**

  * ✅ Đảm bảo tính toàn vẹn dữ liệu mạnh mẽ (ACID).
  * ✅ Khả năng truy vấn phong phú, bao gồm hỗ trợ JSONB hiệu quả.
  * ✅ Không bị khóa bởi nhà cung cấp (vendor lock-in) vì là mã nguồn mở.

**Negative:**

  * ❌ Gặp giới hạn về mở rộng theo chiều dọc (Vertical scaling).
  * ❌ Yêu cầu tối ưu hóa index cẩn thận cho các truy vấn phức tạp.
  * ❌ Việc di trú (migrate) schema có thể phức tạp.

**Risks:**

  * **Risk 1:** Hiệu suất hệ thống suy giảm do tải nặng hoặc các truy vấn không được tối ưu.
  * **Mitigation:** Sử dụng read replicas (bản sao chỉ đọc) cho các workload đọc nhiều; triển khai chiến lược đánh index phù hợp; sử dụng connection pooling; và giám sát các truy vấn chậm (slow queries).

### Alternatives Considered

1.  **Option A: MySQL**
      * **Pros:** Phổ biến.
      * **Cons:** Hỗ trợ JSON kém hơn PostgreSQL; xử lý truy vấn phức tạp kém hiệu quả hơn.
      * **Reason rejected:** PostgreSQL phù hợp hơn với yêu cầu hỗ trợ metadata linh hoạt (JSON) và các truy vấn phức tạp của ITS.
2.  **Option B: NoSQL (ví dụ: MongoDB)**
      * **Pros:** Mở rộng theo chiều ngang tốt.
      * **Cons:** Cần tính toàn vẹn quan hệ (users/roles); cần các truy vấn JOIN phức tạp; cần ACID cho các dữ liệu quan trọng.
      * **Reason rejected:** Không đáp ứng yêu cầu về tính toàn vẹn dữ liệu và nghiệp vụ quan hệ cho các service này.

### Related Decisions

  * **Influences:** ADR-4 (Repository Pattern sẽ có các triển khai cụ thể cho PostgreSQL); ADR-7 (Chiến lược Data Privacy sẽ tận dụng các extension của PostgreSQL như `pgcrypto`).

-----

## ADR-3: Clean/Hexagonal Architecture for All Services

### Status

✅ **Accepted**
*(Date: 2025-10-13, Deciders: Architecture Team)*

### Context

Hệ thống ITS cần phải đảm bảo các ACs quan trọng:

  * **AC-4 (Testability):** Logic nghiệp vụ và thuật toán AI phải dễ dàng kiểm thử độc lập mà không cần khởi động framework hay cơ sở dữ liệu.
  * **AC-1 (Modularity):** Tách biệt rõ ràng các mối quan tâm (separation of concerns).
  * **AC-7 (Maintainability):** Code phải dễ hiểu, dễ sửa đổi, và độc lập với các chi tiết hạ tầng (DB, framework).
  * **Vấn đề:** Các kiến trúc tầng (layered) truyền thống thường tạo ra sự kết dính chặt (tight coupling) giữa logic nghiệp vụ và cơ sở dữ liệu/framework, làm cho việc kiểm thử và thay đổi trở nên khó khăn.

### Decision

Áp dụng **Clean Architecture** (hoặc Hexagonal/Onion Architecture) cho **TẤT CẢ** các microservices (cả Java và Golang).
Cấu trúc này tuân thủ **Dependency Rule** (Quy tắc Phụ thuộc): Mọi phụ thuộc phải hướng vào trong, từ các lớp ngoài (hạ tầng, adapters) vào các lớp trong (ứng dụng, domain).

```
Infrastructure → Adapters → Application → Domain
(Framework/DB) → (Controllers) → (Use Cases) → (Entities)
```

### Rationale

  * **Supports AC-4 (Testability):** Đây là lợi ích lớn nhất. Logic nghiệp vụ (Domain) và các Use Cases (Application) có thể được unit test một cách độc lập, không cần DB hay framework.
  * **Supports AC-1 (Modularity) & AC-7 (Maintainability):** Tách biệt rõ ràng các mối quan tâm. Logic nghiệp vụ cốt lõi (domain) không biết gì về cơ sở dữ liệu đang được sử dụng (Postgres hay Mongo) hoặc cách nó được kích hoạt (REST hay gRPC).
  * **Addresses requirement:** Đảm bảo tuân thủ **Dependency Inversion Principle (DIP)**. Cho phép thay đổi hạ tầng (ví dụ: đổi từ Spring sang Micronaut, hoặc từ Postgres sang MySQL) mà không ảnh hưởng đến logic nghiệp vụ.
  * **Mitigates risk:** Ngăn ngừa rủi ro hệ thống bị khóa cứng vào một công nghệ cụ thể (vendor lock-in) và rủi ro code trở nên "rối như spaghetti" (big ball of mud) khi dự án phát triển.

### Consequences

**Positive:**

  * ✅ Logic nghiệp vụ ở tầng `application` là "thuần túy" (không phụ thuộc framework).
  * ✅ Có thể kiểm thử Use Cases với các repository giả (mock repositories).
  * ✅ Có thể hoán đổi DB mà không cần chạm đến logic nghiệp vụ.
  * ✅ Ranh giới rõ ràng giữa các tầng.

**Negative:**

  * ❌ Nhiều code boilerplate hơn (phải định nghĩa interfaces, DTOs).
  * ❌ Đường cong học tập (learning curve) dốc hơn cho các lập trình viên mới.
  * ❌ Nhiều tệp/packages hơn để điều hướng.

**Risks:**

  * **Risk 1:** Đội ngũ phát triển (đặc biệt là junior) có thể thấy khó khăn và áp dụng không nhất quán, dẫn đến vi phạm quy tắc phụ thuộc.
  * **Mitigation:** Cung cấp các mẫu code (templates) chi tiết cho cả Java và Golang (như trong tệp ADR); tài liệu hóa rõ ràng; và thực hiện code reviews nghiêm ngặt để đảm bảo sự tuân thủ.

### Alternatives Considered

1.  **Option A: Traditional Layered Architecture** (Kiến trúc tầng truyền thống)
      * **Pros:** Đơn giản, quen thuộc.
      * **Cons:** Kết dính chặt (tight coupling) logic nghiệp vụ với framework/DB; khó kiểm thử độc lập.
      * **Reason rejected:** Thất bại hoàn toàn trong việc đáp ứng **AC-4 (Testability)**, vốn là ưu tiên hàng đầu.

### Related Decisions

  * **Depends on:** Quyết định chọn kiến trúc tổng thể (Modular Monolith / Microservices).
  * **Influences:** ADR-1 (Phải triển khai Clean Architecture ở cả hai stack Java và Go); ADR-4 (Repository Pattern là một hệ quả kỹ thuật trực tiếp của ADR này); ADR-5 (Chiến lược Unit Test phụ thuộc hoàn toàn vào ADR này).

-----

## ADR-4: Repository Pattern with Interface Abstraction

### Status

✅ **Accepted**
*(Date: 2025-10-13, Deciders: Architecture Team)*

### Context

  * **Vấn đề:** Theo sau ADR-3 (Clean Architecture), logic nghiệp vụ (Application Layer) cần truy cập dữ liệu, nhưng không được phép phụ thuộc trực tiếp vào các chi tiết hạ tầng (như ORM hoặc database driver).
  * **Mục tiêu:** Cần một mẫu thiết kế (design pattern) để đảo ngược sự phụ thuộc (DIP), giúp:
      * **AC-1 (Modularity):** Tách rời logic nghiệp vụ khỏi logic truy cập dữ liệu.
      * **AC-4 (Testability):** Dễ dàng giả lập (mock) tầng dữ liệu trong khi kiểm thử.

### Decision

Triển khai **Repository Pattern** với **Interface Abstraction**:

1.  **Định nghĩa Repository Interfaces** (Ports) trong tầng `application`. Các interface này là "thuần túy" (không có chi tiết về framework).
2.  **Triển khai (Implement) các Interfaces** (Adapters) trong tầng `infrastructure`. Các lớp triển khai này sẽ chứa logic ORM (Hibernate, GORM) hoặc SQL thô.
3.  **Sử dụng Dependency Injection** để "tiêm" các triển khai cụ thể (ví dụ: `PostgresUserRepository`) vào các Use Cases (vốn chỉ biết đến interface `UserRepository`).

### Rationale

  * **Supports AC-4 (Testability):** Cho phép các Use Cases được kiểm thử bằng cách "tiêm" vào một mock repository (ví dụ: `MockUserRepository`) thay vì một repository thật.
  * **Supports AC-1 (Modularity):** Cho phép thay đổi công nghệ cơ sở dữ liệu (ví dụ: đổi từ `PostgresContentRepository` sang `MongoContentRepository`) mà không cần thay đổi một dòng code nào trong tầng `application`.
  * **Addresses requirement:** Đây là cách triển khai kỹ thuật cụ thể của **Dependency Inversion Principle (DIP)** mà ADR-3 yêu cầu.
  * **Mitigates risk:** Giảm thiểu rủi ro logic nghiệp vụ bị phụ thuộc vào các chi tiết của ORM (ví dụ: các `UserEntity` của JPA).

### Consequences

**Positive:**

  * ✅ Logic nghiệp vụ có thể kiểm thử mà không cần cơ sở dữ liệu.
  * ✅ Có thể hoán đổi các triển khai cơ sở dữ liệu.
  * ✅ Ranh giới truy cập dữ liệu rõ ràng.

**Negative:**

  * ❌ Nhiều interface hơn để bảo trì.
  * ❌ Cần thêm một lớp ánh xạ (mapping) giữa các domain entities (ví dụ: `User`) và các DB entities (ví dụ: `UserEntity`).

**Risks:**

  * **Risk 1:** Việc ánh xạ (mapping) giữa các đối tượng domain và đối tượng DB có thể trở nên tẻ nhạt và tốn thời gian.
  * **Mitigation:** Sử dụng các thư viện ánh xạ (ví dụ: MapStruct cho Java) hoặc tự động sinh code (code generation) để giảm bớt code boilerplate.

### Alternatives Considered

1.  **Option A: Sử dụng ORM trực tiếp trong Use Cases**
      * **Pros:** Ít code boilerplate hơn.
      * **Cons:** Logic nghiệp vụ bị kết dính chặt với ORM; không thể unit test độc lập (phải dùng integration test).
      * **Reason rejected:** Vi phạm nghiêm trọng ADR-3 và thất bại trong việc đáp ứng **AC-4 (Testability)**.

### Related Decisions

  * **Depends on:** ADR-3 (Repository Pattern là một chiến thuật để thực thi Clean Architecture).
  * **Influences:** ADR-5 (Chiến lược Unit Test dựa hoàn toàn vào việc mocking các interface repository này); ADR-2 (Cần phải có các triển khai repository cụ thể cho PostgreSQL).

-----

## ADR-5: Testing Strategy (Testing Pyramid)

### Status

✅ **Accepted**
*(Date: 2025-10-14, Deciders: Architecture Team, QA Lead)*

### Context

Cần một chiến lược kiểm thử rõ ràng để đảm bảo **AC-4 (Testability)** và chất lượng code trong môi trường Microservices + Polyglot (ADR-1). Logic AI (Adaptive Engine, Scoring) đòi hỏi độ chính xác và độ tin cậy rất cao.

### Decision

Áp dụng mô hình **"Testing Pyramid" (Kim tự tháp Kiểm thử)**:

1.  **Unit Tests (Nền tảng - 80%+):**
      * Mục tiêu: Kiểm thử logic (Domain) và (Application) độc lập.
      * Công nghệ: JUnit 5/Mockito (Java), `go test`/`testify/mock` (Go).
      * Quy tắc: Bắt buộc mock tất cả I/O (database, network).
      * SLO: Code coverage **\> 80%** cho `domain` và `application` layers.
2.  **Integration Tests (Tầng giữa):**
      * Mục tiêu: Kiểm thử tích hợp của service với hạ tầng (Database, Message Broker).
      * Công nghệ: `@SpringBootTest` (Java), `go test` + **Testcontainers** (cho cả hai).
      * Quy tắc: Kiểm tra các repository (ADR-4) có thể ghi/đọc đúng từ DB thật (chạy trong Docker).
3.  **End-to-End (E2E) Tests (Đỉnh tháp):**
      * Mục tiêu: Xác thực các luồng nghiệp vụ quan trọng qua toàn bộ hệ thống (giả lập từ API Gateway).
      * Công nghệ: Cypress, Playwright, hoặc Postman/K6.
      * Quy tắc: Chỉ test các luồng chính (ví dụ: UC-10: Nộp bài & nhận phản hồi). Giữ số lượng test ít để tránh không ổn định (flakiness).

### Rationale

  * **Supports AC-4 (Testability):** Cung cấp một chiến lược toàn diện để đạt được AC-4. ADR-3/ADR-4 là nền tảng kỹ thuật cho phép Unit Test hiệu quả, và ADR-5 định nghĩa cách thực thi nó.
  * **Addresses requirement:** Đảm bảo độ tin cậy cao cho các thuật toán AI/Scoring.
  * **Mitigates risk:** Phát hiện lỗi sớm ở tầng Unit Test (rẻ nhất, nhanh nhất) thay vì ở E2E Test (đắt nhất, chậm nhất). Integration test với Testcontainers giúp giảm rủi ro lỗi tích hợp (ví dụ: sai câu query SQL) mà Unit Test bỏ lỡ.

### Consequences

**Positive:**

  * ✅ Độ tin cậy cao vào chất lượng code.
  * ✅ Phát hiện lỗi sớm (Unit/Integration tests chạy trong CI pipeline).
  * ✅ Logic AI được kiểm thử kỹ lưỡng bằng Unit Test.

**Negative:**

  * ❌ `Testcontainers` làm tăng thời gian chạy CI/CD pipeline (vì phải khởi động Docker).
  * ❌ E2E tests có thể không ổn định (flaky) và khó debug.
  * ❌ Yêu cầu đội ngũ phải học cách sử dụng Testcontainers.

**Risks:**

  * **Risk 1:** CI pipeline chạy quá chậm do Integration tests.
  * **Mitigation:** Phân tách các bước (stage) trong CI: chạy Unit Test trên mỗi commit, chỉ chạy Integration/E2E Test khi tạo Pull Request hoặc merge vào main.
  * **Risk 2:** E2E tests không ổn định (flaky).
  * **Mitigation:** Giới hạn số lượng E2E tests chỉ cho các luồng "happy path" quan trọng nhất (ví dụ: đăng nhập, nộp bài).

### Alternatives Considered

1.  **Option A: Chỉ Unit Tests**
      * **Pros:** Rất nhanh.
      * **Cons:** Bỏ lỡ các lỗi tích hợp (ví dụ: sai cấu hình DB, sai câu query SQL).
      * **Reason rejected:** Không đủ độ tin cậy.
2.  **Option B: Chỉ E2E Tests (Ice Cream Cone Pattern)**
      * **Pros:** Không có.
      * **Cons:** Rất chậm, đắt đỏ, khó bảo trì, và khó xác định nguyên nhân gốc rễ của lỗi.
      * **Reason rejected:** Cực kỳ không hiệu quả và giòn (brittle).
3.  **Option C: Tách Contract Tests (ví dụ: Pact.io)**
      * **Pros:** Rất tốt cho việc xác thực giao tiếp giữa các microservices.
      * **Cons:** Quá phức tạp cho giai đoạn MVP.
      * **Reason rejected:** Sẽ được xem xét trong tương lai, nhưng hiện tại là quá phức tạp.

### Related Decisions

  * **Depends on:** ADR-1 (Phải hỗ trợ cả Java và Go), ADR-3 và ADR-4 (Unit Test chỉ hiệu quả khi có Clean Architecture và Repository Pattern).

-----

## ADR-6: Security Architecture (AuthN & AuthZ)

### Status

✅ **Accepted**
*(Date: 2025-10-14, Deciders: Architecture Team, Security Lead)*

### Context

Cần một cơ chế bảo mật (Authentication - AuthN và Authorization - AuthZ) mạnh mẽ cho hệ thống Microservices phân tán, đáp ứng:

  * **AC-6 (Security):** Bảo vệ tài nguyên hệ thống.
  * **FR11 (RBAC):** Hỗ trợ phân quyền dựa trên vai trò (Learner, Instructor, Admin).
  * **AC-2 (Scalability):** Cơ chế bảo mật phải là stateless (không lưu trạng thái) để hỗ trợ mở rộng theo chiều ngang.

### Decision

Áp dụng mô hình **Bảo mật Tập trung (Centralized Auth)**:

1.  **Authentication (AuthN):** Một `Auth Service` (Java/Spring Security) đóng vai trò là Identity Provider (IdP) trung tâm, tuân thủ **OAuth 2.0 / OIDC**. Dịch vụ này phát hành **JSON Web Tokens (JWTs)** (Access Token + Refresh Token).
2.  **Authorization (AuthZ) - Edge Level:** **API Gateway** (Golang) là cổng bảo mật duy nhất. Gateway sẽ **xác thực (validate) JWT** trên MỌI request đến từ bên ngoài. Nếu JWT không hợp lệ, request bị từ chối ngay lập tức.
3.  **Authorization (AuthZ) - Service Level (RBAC):** Sau khi xác thực, API Gateway chuyển tiếp thông tin người dùng (ví dụ: `X-User-ID`, `X-User-Roles`) vào header của request nội bộ. Các service bên trong (ví dụ: `ContentService`) **tin tưởng** thông tin từ Gateway và sử dụng `X-User-Roles` để kiểm tra RBAC (ví dụ: "chỉ `Instructor` mới được tạo nội dung").

### Rationale

  * **Supports AC-6 (Security):** Cung cấp một "cửa ngõ" bảo mật mạnh mẽ, tuân thủ các tiêu chuẩn ngành (OAuth 2.0/OIDC).
  * **Supports AC-1 (Modularity - SRP):** Logic AuthN/AuthZ phức tạp được tập trung tại Auth Service và API Gateway. Các service nghiệp vụ (Scoring, Adaptive) được giữ đơn giản, không cần quan tâm đến việc xác thực chữ ký JWT.
  * **Supports AC-2 (Scalability):** Sử dụng JWTs là stateless, hoàn toàn phù hợp với việc mở rộng ngang các service.
  * **Addresses requirement:** Đáp ứng FR11 (RBAC) bằng cách truyền vai trò (roles) trong header.
  * **Mitigates risk:** Giảm thiểu rủi ro logic xác thực bị triển khai lặp lại và không nhất quán ở mỗi service.

### Consequences

**Positive:**

  * ✅ Bảo mật mạnh mẽ tại "cửa ngõ" (Gateway).
  * ✅ Các service nghiệp vụ được đơn giản hóa (chỉ cần đọc header).
  * ✅ Dễ dàng scale các service (stateless).

**Negative:**

  * ❌ `Auth Service` và `API Gateway` trở thành các điểm lỗi đơn (Single Points of Failure - SPoFs). Chúng phải có độ sẵn sàng (Availability) cực cao.
  * ❌ Mô hình "tin tưởng Gateway" (passing headers) kém an toàn hơn mô hình Zero Trust (nơi mỗi service tự xác thực lẫn nhau).
  * ❌ JWTs phải có thời gian sống ngắn (ví dụ: 15 phút) và cần cơ chế refresh token phức tạp.

**Risks:**

  * **Risk 1:** `Auth Service` hoặc `API Gateway` bị sập, toàn bộ hệ thống ngừng hoạt động.
  * **Mitigation:** Đảm bảo triển khai các thành phần này với cấu hình High Availability (HA), chạy nhiều bản sao (replicas) trong Kubernetes.
  * **Risk 2:** Một kẻ tấn công nội bộ (internal) có thể bypass Gateway và gọi thẳng vào service nghiệp vụ với header giả mạo.
  * **Mitigation:** Sử dụng Mạng riêng ảo (VPC) và các Chính sách Mạng (Network Policies) của Kubernetes để đảm bảo chỉ có API Gateway mới được phép gọi các service nội bộ.

### Alternatives Considered

1.  **Option A: mTLS (Zero Trust)**
      * **Pros:** Bảo mật cao nhất, mỗi service xác thực lẫn nhau.
      * **Cons:** Cực kỳ phức tạp để triển khai và quản lý chứng chỉ (certificates) cho MVP.
      * **Reason rejected:** Quá phức tạp (overkill) cho giai đoạn hiện tại.
2.  **Option B: Session Cookies (Stateful)**
      * **Pros:** Truyền thống, đơn giản.
      * **Cons:** Không phù hợp với Microservices và **AC-2 (Scalability)** (yêu cầu sticky session hoặc shared cache).
      * **Reason rejected:** Không thể mở rộng (non-scalable).
3.  **Option C: Mỗi service tự validate JWT**
      * **Pros:** Bảo mật hơn (không tin tưởng header).
      * **Cons:** Tăng latency (mỗi service phải gọi Auth Service để lấy public key); lặp lại logic.
      * **Reason rejected:** Không hiệu quả và vi phạm SRP.

### Related Decisions

  * **Depends on:** ADR-1 (API Gateway dùng Go, Auth Service dùng Java).
  * **Influences:** ADR-7 (Thông tin `X-User-ID` từ JWT sẽ là "chìa khóa" ẩn danh để liên kết đến dữ liệu PII).

-----

## ADR-7: Data Privacy & Compliance (GDPR/FERPA)

### Status

✅ **Accepted**
*(Date: 2025-10-14, Deciders: Architecture Team, Compliance Officer)*

### Context

  * Hệ thống ITS xử lý Dữ liệu Cá nhân Nhạy cảm (PII - Personally Identifiable Information) của học sinh, bao gồm tên, email, và kết quả học tập.
  * Hệ thống phải tuân thủ các quy định về bảo mật dữ liệu như **GDPR** (Châu Âu) và **FERPA** (Mỹ).
  * Cần một chiến lược để giảm thiểu rủi ro rò rỉ PII, tuân thủ **AC-6 (Security)** ở mức cao nhất.

### Decision

Áp dụng chiến lược **"Phân tách PII (PII Isolation)"** và **"Ẩn danh hóa (Anonymization)"**:

1.  **Phân tách PII:** Dữ liệu PII (tên, email, SĐT) **CHỈ** được lưu trữ trong `User Management Service` (Database: Postgres - ADR-2).
2.  **Ẩn danh hóa:** Tất cả các service khác (ví dụ: `LearnerModelService`, `ScoringService`, `AdaptiveEngine`) **KHÔNG BAO GIỜ** được lưu trữ PII. Chúng phải tham chiếu đến người dùng thông qua một **`LearnerID` (UUID)** đã được ẩn danh.
3.  **Mã hóa PII (Encryption at Rest):** Các cột chứa PII (ví dụ: `email`, `full_name`) trong database Postgres của `User Management Service` phải được **mã hóa ở cấp độ cột** (ví dụ: sử dụng extension `pgcrypto`).
4.  **Mã hóa khi Truyền tải (Encryption in Transit):** Tất cả giao tiếp (nội bộ và bên ngoài) phải sử dụng **TLS (HTTPS)**.
5.  **Quyền được Lãng quên (Right to be Forgotten):** Triển khai một API (chỉ Admin) cho phép xóa dữ liệu người dùng dựa trên `LearnerID`, API này sẽ kích hoạt các sự kiện (event) để các service khác xóa dữ liệu liên quan.

### Rationale

  * **Supports AC-6 (Security):** Đây là biện pháp bảo mật cốt lõi để bảo vệ dữ liệu người dùng. Nó tuân thủ Nguyên tắc Đặc quyền Tối thiểu (Principle of Least Privilege) - `ScoringService` không cần biết tên học sinh để chấm bài.
  * **Addresses requirement:** Đáp ứng trực tiếp các yêu cầu tuân thủ (Compliance) của GDPR và FERPA.
  * **Mitigates risk:** Giảm thiểu rủi ro một cách đáng kể. Ngay cả khi `LearnerModelService` bị xâm nhập, kẻ tấn công cũng không thể lấy được danh tính thật của học sinh (chỉ có các UUID ẩn danh).

### Consequences

**Positive:**

  * ✅ Bảo mật PII và tuân thủ pháp lý ở mức độ cao.
  * ✅ Giảm đáng kể bề mặt tấn công (attack surface).
  * ✅ Thực thi tốt Nguyên tắc Đặc quyền Tối thiểu (Least Privilege).

**Negative:**

  * ❌ Tăng độ phức tạp. Việc "join" dữ liệu (ví dụ: hiển thị tên học sinh bên cạnh điểm số) trở nên khó khăn hơn. Nó đòi hỏi phải gọi 2 service (User Service + Scoring Service) và join ở tầng ứng dụng (API Gateway hoặc Frontend).
  * ❌ Mã hóa cấp độ cột (pgcrypto) làm giảm hiệu năng query trên các cột đó (không thể index hiệu quả).
  * ❌ Triển khai API "Right to be Forgotten" phức tạp trong hệ thống phân tán (cần dùng Saga pattern).

**Risks:**

  * **Risk 1:** Hiệu năng hệ thống bị ảnh hưởng do phải gọi nhiều service để tổng hợp dữ liệu (ví dụ: lấy tên user và điểm số).
  * **Mitigation:** Sử dụng API Gateway để tổng hợp (aggregate) các lời gọi, hoặc sử dụng cơ chế cache (Redis) cho các dữ liệu PII ít thay đổi (như tên).
  * **Risk 2:** Việc triển khai "Right to be Forgotten" không nhất quán, PII bị xóa ở User Service nhưng dữ liệu ẩn danh vẫn còn ở các service khác.
  * **Mitigation:** Sử dụng một Saga Pattern (dựa trên Event-Driven) để đảm bảo yêu cầu xóa được gửi đến tất cả các service liên quan một cách đáng tin cậy.

### Alternatives Considered

1.  **Option A: Lưu PII ở mọi nơi**
      * **Pros:** Đơn giản, dễ query.
      * **Cons:** Vi phạm pháp lý (GDPR/FERPA); rủi ro bảo mật cực kỳ cao. Một service bị hack là mất tất cả PII.
      * **Reason rejected:** Không thể chấp nhận được về mặt bảo mật và pháp lý.
2.  **Option B: Chỉ mã hóa Toàn bộ Database (Full Disk Encryption)**
      * **Pros:** Đơn giản hơn mã hóa cột.
      * **Cons:** Không đủ. Nếu service (ứng dụng) bị xâm nhập, kẻ tấn công vẫn đọc được PII (vì hệ điều hành đã giải mã đĩa). Mã hóa cột bảo vệ chống lại cả việc ứng dụng bị xâm nhập.
      * **Reason rejected:** Không đủ an toàn (Fails AC-6).

### Related Decisions

  * **Depends on:** ADR-2 (Cần PostgreSQL để sử dụng `pgcrypto`); ADR-6 (Cần JWT và `LearnerID` để làm khóa ẩn danh).