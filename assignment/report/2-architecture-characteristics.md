# Architecture Characteristics (Đặc tính Kiến trúc)

> **Architecture Characteristics (ACs)** là cầu nối chiến lược giữa **Yêu cầu Nghiệp vụ** (cần làm gì) và **Quyết định Kỹ thuật** (làm như thế nào) trong quá trình thiết kế hệ thống.

---

## 1. Vai trò của Architecture Characteristics

### 1.1. Chuyển đổi Yêu cầu Phi Chức năng (NFRs) thành Tiêu chí Đo lường

- Biến những yêu cầu mơ hồ như _"hệ thống phải nhanh"_ thành tiêu chí định lượng và cụ thể.
- **Ví dụ:**
  - **Mơ hồ:** "Hệ thống phải nhanh"
  - **Định lượng:** "Độ trễ API phổ biến < 500ms"

### 1.2. Quản lý Sự Đánh Đổi (Trade-Offs)

- Giúp Kiến trúc sư xác định rõ ưu tiên giữa những đặc tính có thể mâu thuẫn:
  - **Scalability** (khả năng mở rộng) **vs** **Simplicity** (đơn giản)
  - **Performance** (hiệu năng) **vs** **Maintainability** (dễ bảo trì)
- Việc xác lập ACs là nền tảng cho quyết định **"kiến trúc ít tệ nhất"** (_the least worst architecture_).

### 1.3. Định hình Cấu trúc Hệ thống

- Việc ưu tiên các ACs sẽ quyết định chọn **Kiểu Kiến trúc** phù hợp nhất.
- **Ví dụ:**
  - Nếu yêu cầu **Modularity** và **Deployability** cao → ưu tiên **Microservices**
  - Nếu cần đơn giản hóa → có thể chọn **Monolith**

---

## 2. Xác định các Đặc tính Kiến trúc Chính (Primary Characteristics)

Dựa trên yêu cầu cốt lõi của **Hệ thống Gia sư Thông minh (ITS)**—tính AI Thích ứng và Khả năng Mở rộng cho nhiều người dùng—các đặc tính kiến trúc chính (Primary ACs) cần được ưu tiên cao nhất:

### AC1: Modularity (Tính Mô-đun)

| **Khái niệm** | **Vai trò trong ITS** | **Cơ sở Lý thuyết & Tiêu chí** |
|---------------|----------------------|--------------------------------|
| Tính Mô-đun là khả năng phân rã hệ thống thành các đơn vị độc lập (Modules/Components) với **khớp nối thấp** (Low Coupling) và **gắn kết cao** (High Cohesion). | Đây là nền tảng kỹ thuật cho **Live AI Model Swapping** (FR12). Giúp cô lập các Logic Nghiệp vụ Cốt lõi (ví dụ: `ScoringEngine`, `AdaptivePathGenerator`) thành các **Policy Modules**. | Đo lường bằng **Instability Index (I)**: Các module cốt lõi (ví dụ: Giao diện `LearnerModel`) phải đạt **I≈0** (Cₐ cao, Cₑ thấp) để đảm bảo chúng rất ổn định và được bảo vệ khỏi sự thay đổi của các chi tiết triển khai. |

### AC2: Scalability (Khả năng Mở rộng)

| **Khái niệm** | **Vai trò trong ITS** | **Cơ sở Lý thuyết & Tiêu chí** |
|---------------|----------------------|--------------------------------|
| Khả năng Mở rộng là khả năng xử lý khối lượng công việc tăng lên (ví dụ: thêm người học đồng thời) bằng cách **mở rộng ngang** (Horizontal Scaling) mà không cần can thiệp kiến trúc. | Cực kỳ quan trọng để xử lý tải người dùng tăng trưởng và tải tính toán nặng từ các thuật toán AI (FR7). Các Domain Services AI phải được thiết kế như các **Architecture Quantum** để dễ dàng nhân bản và phân phối tải thông qua Load Balancer. | **Tiêu chí:** Hỗ trợ tối thiểu **5,000 user concurrent** và có khả năng **Auto-scaling** dựa trên CPU usage. |

### AC3: Performance (Hiệu suất)

| **Khái niệm** | **Vai trò trong ITS** | **Cơ sở Lý thuyết & Tiêu chí** |
|---------------|----------------------|--------------------------------|
| Hiệu suất đo lường mức độ nhanh chóng hệ thống phản hồi hoặc hoàn thành một tác vụ. Được đo bằng **độ trễ** (latency) và **thông lượng** (throughput). | Phải đảm bảo trải nghiệm Gia sư 1-kèm-1 mượt mà, đặc biệt cho tính năng **Real-Time Remedial Feedback** (FR6). Độ trễ cao có thể làm hỏng trải nghiệm học tập. | **Mục tiêu Định lượng:** Độ trễ cho các tương tác cốt lõi của người học (chấm điểm, gợi ý) phải **< 500ms**. **Chiến lược:** Sử dụng Fast Cache (Redis) cho dữ liệu thường xuyên truy cập. |

### AC4: Testability (Khả năng Kiểm thử)

| **Khái niệm** | **Vai trò trong ITS** | **Cơ sở Lý thuyết & Tiêu chí** |
|---------------|----------------------|--------------------------------|
| Khả năng Kiểm thử là mức độ dễ dàng để xác minh hệ thống hoạt động chính xác. Được đảm bảo bằng cách **tách biệt Logic khỏi Cơ sở hạ tầng** (I/O). | Quan trọng nhất để đảm bảo tính đúng đắn và công bằng của các thuật toán chấm điểm và gợi ý AI. Nếu thuật toán bị lỗi, uy tín của hệ thống sẽ mất. | **Thực thi bằng DIP:** Logic cốt lõi phải phụ thuộc vào **Abstraction** (Interfaces) chứ không phải lớp Concrete (ví dụ: Database Repository). Điều này cho phép **Mock** (giả lập) dữ liệu dễ dàng cho Unit Test. |

# Secondary Characteristics (Đặc Tính Kiến Trúc Thứ Cấp)

## Giới thiệu

Các **Đặc tính Kiến trúc Thứ cấp** (Secondary ACs) là những yếu tố quan trọng cho sự ổn định lâu dài, bảo mật, và khả năng vận hành của hệ thống ITS. 

**Vai trò:**
- Mặc dù chúng **không định hình trực tiếp** kiểu kiến trúc như các Primary ACs (Modularity, Scalability, Performance, Testability)
- Chúng lại **quyết định cách thức triển khai** và tổ chức mã nguồn bên trong mỗi dịch vụ
- Đảm bảo hệ thống có thể **duy trì, mở rộng và vận hành** hiệu quả theo thời gian

---

## Danh sách Secondary Characteristics

### AC5: Deployability (Khả năng Triển khai)

| **Khái niệm** | **Vai trò trong ITS** |
|---------------|----------------------|
| Mức độ dễ dàng và nhanh chóng để chuyển đổi mã nguồn thành môi trường sản xuất. Đơn vị cơ bản là **Architecture Quantum**. | Hỗ trợ **Live AI Model Swapping** (FR12). Mỗi Domain Service (ví dụ: `FeedbackGenerator`) phải được đóng gói thành một **Container** và có thể triển khai độc lập mà không cần khởi động lại toàn bộ hệ thống. |

**Chiến lược thực thi:**
- Sử dụng **Containerization** (Docker)
- **CI/CD Pipeline** tự động hóa
- **Blue/Green** hoặc **Canary Deployment**
- Mỗi service là một **independent deployment unit**

---

### AC6: Security (Bảo mật)

| **Khái niệm** | **Vai trò trong ITS** |
|---------------|----------------------|
| Khả năng bảo vệ hệ thống khỏi các truy cập trái phép và sự cố dữ liệu. | Rất quan trọng để bảo vệ dữ liệu nhạy cảm của người học (`LearnerModel`) và nội dung học tập độc quyền. Phải đảm bảo phân quyền **RBAC** (FR11) nghiêm ngặt, bao gồm cả xác thực giao tiếp nội bộ (service-to-service). |

**Chiến lược thực thi:**
- **Authentication & Authorization**: JWT tokens, OAuth 2.0
- **RBAC** (Role-Based Access Control) nghiêm ngặt
- **Service-to-Service Auth**: mTLS, API Gateway
- **Data Encryption**: At rest (database) và in transit (HTTPS/TLS)
- **Audit Logs**: Ghi lại mọi thao tác nhạy cảm
- Tuân thủ **OWASP Top 10**

---

### AC7: Maintainability (Khả năng Bảo trì)

| **Khái niệm** | **Vai trò trong ITS** |
|---------------|----------------------|
| Mức độ dễ dàng để thay đổi, sửa lỗi, và mở rộng hệ thống theo thời gian. Tính năng này được củng cố bằng việc tuân thủ các nguyên tắc thiết kế. | Cần thiết để giảm **chi phí vòng đời** (life cycle cost) của hệ thống. **Cohesion cao** và **Coupling thấp** (Low Cₑ) của các Module là bắt buộc để tránh **"hiệu ứng lan truyền"** (ripple effect) khi cập nhật logic (FR4, FR6). |

**Chiến lược thực thi:**
- Áp dụng **SOLID Principles** nghiêm ngặt
- **High Cohesion**: Mỗi module có trách nhiệm rõ ràng
- **Low Coupling**: Giảm phụ thuộc giữa các module
- **Clean Code**: Naming conventions, code reviews
- **Documentation**: Technical docs, ADRs (Architecture Decision Records)
- **Refactoring** định kỳ

---

### AC8: Extensibility (Khả năng Mở rộng Chức năng)

| **Khái niệm** | **Vai trò trong ITS** |
|---------------|----------------------|
| Mức độ dễ dàng thêm các tính năng mới hoặc tích hợp với các hệ thống bên ngoài (ví dụ: Gamification, Payment Gateway) mà không sửa đổi các module cốt lõi. | Liên quan chặt chẽ đến **OCP** (Open/Closed Principle). Thiết kế phải sử dụng các **Interface/Abstraction** để cho phép mở rộng (ví dụ: thêm loại câu hỏi mới) mà không thay đổi `ScoringEngine`. |

**Chiến lược thực thi:**
- **Plugin Architecture**: Cho phép thêm module mới mà không sửa core
- **Strategy Pattern**: Cho các thuật toán có thể thay đổi
- **Dependency Injection**: Dễ dàng swap implementations
- **Event-Driven**: Thêm consumers mới mà không ảnh hưởng producers
- **API Versioning**: Hỗ trợ backward compatibility

---

### AC9: Observability (Khả năng Giám sát)

| **Khái niệm** | **Vai trò trong ITS** |
|---------------|----------------------|
| Khả năng quan sát, giám sát và chẩn đoán hệ thống thông qua **metrics**, **logs**, và **traces** để phát hiện sự cố nhanh chóng. | Cực kỳ quan trọng cho hệ thống phân tán với nhiều services. Giúp **debug**, **monitor performance**, và **phát hiện bottleneck** trong Adaptive Engine (FR7) hoặc Scoring Service (FR5). Hỗ trợ **proactive alerting** khi có vấn đề. |

**Chiến lược thực thi:**
- **Structured Logging**: ELK Stack (Elasticsearch, Logstash, Kibana) hoặc EFK (Fluentd)
- **Metrics Collection**: Prometheus + Grafana
- **Distributed Tracing**: Jaeger hoặc Zipkin để theo dõi request qua nhiều services
- **APM (Application Performance Monitoring)**: New Relic, DataDog
- **Alerting**: Cảnh báo tự động khi metrics vượt ngưỡng (CPU > 80%, latency > 500ms)
- **Health Checks**: Endpoint `/health` cho mỗi service

---

## Tổng hợp tất cả Architecture Characteristics

### Bảng phân loại đầy đủ 9 ACs cho ITS

| **AC ID** | **Tên Đặc tính**           | **Phân loại**    | **Mức độ Ưu tiên** | **Vai trò chính**                          | **NFR liên quan** |
|-----------|----------------------------|------------------|-------------------|-------------------------------------------|-------------------|
| **AC1**   | Modularity & Extensibility | **Primary**      | ⭐⭐⭐⭐⭐ Cao nhất | Định hình kiến trúc Microservices; hỗ trợ Live Model Swapping | Mở rộng chức năng |
| **AC2**   | Scalability                | **Primary**      | ⭐⭐⭐⭐⭐ Cao nhất | Xử lý tải người dùng và tính toán AI nặng  | Khả năng Mở rộng  |
| **AC3**   | Performance                | **Primary**      | ⭐⭐⭐⭐⭐ Cao nhất | Đảm bảo trải nghiệm học tập mượt mà (<500ms) | Hiệu năng         |
| **AC4**   | Testability                | **Primary**      | ⭐⭐⭐⭐ Cao       | Đảm bảo tính đúng đắn của thuật toán AI    | Kiểm thử          |
| **AC5**   | Deployability              | **Secondary**    | ⭐⭐⭐ Trung bình | Hỗ trợ triển khai độc lập từng service    | -                 |
| **AC6**   | Security                   | **Secondary**    | ⭐⭐⭐⭐ Cao       | Bảo vệ dữ liệu người học và nội dung       | Bảo mật           |
| **AC7**   | Maintainability            | **Secondary**    | ⭐⭐⭐⭐ Cao       | Giảm chi phí vòng đời, dễ sửa lỗi         | Kiểm thử          |
| **AC8**   | Observability              | **Secondary**    | ⭐⭐⭐ Trung bình | Debug, monitor, phát hiện sự cố           | Giám sát          |
| **AC9**   | Cost Efficiency            | **Secondary**    | ⭐⭐ Thấp         | Tối ưu chi phí vận hành                   | Chi phí           |

---

## Ánh xạ Architecture Characteristics → Non-Functional Requirements

Bảng này cho thấy mối liên hệ giữa các ACs và NFRs đã định nghĩa trong file `1.2-non-functional-requirements.md`:

| **AC** | **NFR tương ứng** | **Mục tiêu Định lượng** | **Chiến lược Kỹ thuật** | **User Stories** |
|--------|-------------------|------------------------|------------------------|------------------|
| **AC1: Modularity & Extensibility** | Mở rộng chức năng | Cho phép thêm tính năng mới không sửa core | Microservices, API versioning, Plugin-friendly (OCP) | US8, US0, US4 |
| **AC2: Scalability** | Khả năng Mở rộng | ≥ 5,000 user concurrent | Horizontal scaling, Kubernetes, Kafka/RabbitMQ | US2, US5, US6 |
| **AC3: Performance** | Hiệu năng | API <300ms; Grading/Report <1s | Redis caching, Indexing, Async workers | US0, US1, US2 |
| **AC4: Testability & Maintainability** | Kiểm thử | Unit test ≥ 80% | SOLID, Clean/Hexagonal Architecture, CI/CD | US8 |
| **AC6: Security & Privacy** | Bảo mật | Hash mật khẩu (bcrypt/argon2) | RBAC (FR11), Audit logs, HTTPS, OWASP Top 10 | US7 |
| **AC7: Reliability & Availability** | Độ tin cậy | SLA ≥ 99.5% uptime | Retry logic, Queues, Backup hàng ngày | US3, US7, US8 |
| **AC8: Observability** | Giám sát | Real-time monitoring & alerting | Metrics (Prometheus), Tracing (Jaeger), Logging (ELK) | US5, US6, US7 |
| **AC9: Cost Efficiency** | Chi phí | Tối ưu tài nguyên | Auto-scaling, Serverless cho tác vụ không thường xuyên | US7 |

**Lưu ý:**
- **AC5 (Deployability)** không có NFR riêng nhưng là yếu tố kỹ thuật quan trọng hỗ trợ AC1 (Modularity)
- **AC7 (Reliability)** trong NFR tương ứng với khái niệm **Availability** trong ACs

---

## So sánh Primary vs Secondary Characteristics

| **Tiêu chí** | **Primary ACs** | **Secondary ACs** |
|--------------|-----------------|-------------------|
| **Vai trò chính** | Định hình **kiểu kiến trúc** (Microservices, Monolith, etc.) | Định hình **cách triển khai** và tổ chức code |
| **Ảnh hưởng** | Cấu trúc tổng thể hệ thống | Chi tiết implementation bên trong services |
| **Ví dụ** | AC1-AC4: Modularity, Scalability, Performance, Testability | AC5-AC9: Deployability, Security, Maintainability, Observability, Cost |
| **Ưu tiên** | Phải được quyết định **đầu tiên** | Được thực thi **sau khi** chọn kiến trúc |
| **Số lượng khuyến nghị** | 3-5 ACs (tập trung) | Không giới hạn (nhưng nên ưu tiên top 3-5) |

---

## Nguyên tắc Lựa chọn Architecture Characteristics

### 1. **Nguyên tắc "Ít hơn là Nhiều hơn"**
- Không nên cố gắng tối ưu **tất cả** ACs cùng lúc
- Chọn **3-5 Primary ACs** quan trọng nhất cho domain
- Quá nhiều ACs → kiến trúc phức tạp, khó maintain

### 2. **Traceability: ACs phải đến từ Requirements**
- Mỗi AC phải có **lý do nghiệp vụ rõ ràng**
- Liên kết với **User Stories**, **Functional Requirements**, **NFRs**
- Tránh "architecture astronaut" (thiết kế quá phức tạp không cần thiết)

### 3. **Trade-offs là Không Thể Tránh Khỏi**
- Tối ưu một AC thường **ảnh hưởng tiêu cực** đến AC khác
- Ví dụ:
  - **Scalability ↑** → **Simplicity ↓** (Microservices phức tạp hơn Monolith)
  - **Security ↑** → **Performance ↓** (Encryption tốn CPU)
  - **Extensibility ↑** → **Performance ↓** (Abstraction layers thêm overhead)

### 4. **Context-Driven: Phụ thuộc vào Domain**
- ITS (Intelligent Tutoring System) → **Modularity, Scalability, Performance** là top priority
- Banking System → **Security, Reliability, Consistency** là top priority
- Startup MVP → **Simplicity, Time-to-Market** là top priority

# Trade-off Analysis (Phân tích Đánh đổi)

## Nguyên tắc

> **"Mọi thứ đều là sự đánh đổi"** - Nguyên tắc vàng trong kiến trúc phần mềm

Mục tiêu của bước này là đưa ra quyết định **chấp nhận rủi ro ở đâu** và **tối ưu hóa ở đâu** để đạt được **"kiến trúc ít tệ nhất"** (_the least worst architecture_) phù hợp với ngữ cảnh ITS.

---

## Các Xung đột Architecture Characteristics

### 1. Scalability (AC2) & Modularity (AC1) ↔ Simplicity

| **Phân tích Đánh đổi** | **Quyết định và Lý do** |
|------------------------|-------------------------|
| Các phong cách kiến trúc tối ưu hóa Scalability (ví dụ: **Microservices**) có độ phức tạp cao hơn về phát triển, vận hành, và giám sát so với **Monolithic**. Độ phức tạp cao làm giảm Simplicity. | ✅ **Ưu tiên: Scalability & Modularity**<br><br>Chúng ta chấp nhận **Độ phức tạp Vận hành** (Complexity) cao hơn vì **Modularity** là bắt buộc để thực hiện **Live AI Model Swapping** (FR12), vốn là yêu cầu cốt lõi. |

---

### 2. Performance (AC3) ↔ Modularity/Coupling

| **Phân tích Đánh đổi** | **Quyết định và Lý do** |
|------------------------|-------------------------|
| Việc phân chia quá mịn (**Too Fine-Grained Components**) để đạt Modularity có thể dẫn đến quá nhiều lời gọi mạng (network calls), làm tăng độ trễ (Latency) và ảnh hưởng tiêu cực đến Performance. | ⚖️ **Ưu tiên: Cân bằng (Balanced Granularity)**<br><br>Phải đảm bảo các chức năng nghiệp vụ **liên quan chặt chẽ** (**Functional Cohesion**) được đóng gói trong cùng một Service (ví dụ: Logic Chấm điểm và Cập nhật Model) để tránh giao tiếp mạng không cần thiết, duy trì **Latency ≤ 500ms**. |

---

### 3. Security (AC6) ↔ Performance (AC3)

| **Phân tích Đánh đổi** | **Quyết định và Lý do** |
|------------------------|-------------------------|
| Tăng cường bảo mật (ví dụ: mã hóa dữ liệu người học (PII) khi truyền tải (**in transit**) và khi lưu trữ (**at rest**), hoặc thêm lớp xác thực giữa các service) sẽ làm tăng chi phí xử lý và độ trễ. | ✅ **Ưu tiên: Security (Data Protection)**<br><br>Chúng ta chấp nhận độ trễ nhỏ **có thể chấp nhận được** do mã hóa (TLS/HTTPS) để bảo vệ **Learner Model** (FR2) và PII. **Bảo mật là không thể thương lượng**. |

---

### 4. Testability (AC4) ↔ Development Cost

| **Phân tích Đánh đổi** | **Quyết định và Lý do** |
|------------------------|-------------------------|
| Áp dụng **Clean/Hexagonal Architecture** để đạt Testability cao (tuân thủ **DIP/SRP**) yêu cầu cấu trúc mã nguồn phức tạp hơn, dẫn đến chi phí phát triển ban đầu cao hơn và đường cong học tập khó hơn cho đội ngũ. | ✅ **Ưu tiên: Testability**<br><br>Chi phí ban đầu cao hơn được đánh đổi với **chi phí bảo trì** (Maintainability) thấp hơn về lâu dài và **độ tin cậy cao hơn** cho các thuật toán AI (FR7). Đây là một **quyết định chiến lược** cho tuổi thọ hệ thống. |

---

## Bảng Tổng hợp Trade-offs

| **Xung đột ACs** | **Quyết định** | **AC được Ưu tiên** | **AC bị Hy sinh** | **Lý do Chính** |
|------------------|----------------|---------------------|-------------------|-----------------|
| Scalability & Modularity ↔ Simplicity | Ưu tiên Scalability & Modularity | AC1, AC2 | Simplicity | Live AI Model Swapping (FR12) là yêu cầu cốt lõi |
| Performance ↔ Modularity | Cân bằng (Balanced Granularity) | AC3, AC1 | - | Duy trì latency ≤ 500ms bằng Functional Cohesion |
| Security ↔ Performance | Ưu tiên Security | AC6 | AC3 (chấp nhận +50-100ms) | Bảo vệ PII và LearnerModel không thể thương lượng |
| Testability ↔ Development Cost | Ưu tiên Testability | AC4 | - (chi phí ban đầu cao) | Độ tin cậy thuật toán AI và chi phí bảo trì dài hạn |
