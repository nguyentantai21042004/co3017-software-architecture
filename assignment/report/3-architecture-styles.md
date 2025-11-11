# So Sánh và Lựa Chọn Architecture Styles

## Mục Tiêu

So sánh các kiểu kiến trúc phổ biến và chọn ra phong cách phù hợp nhất để hiện thực hóa các Architecture Characteristics (ACs) đã được ưu tiên cao nhất cho ITS:
- **AC1: Modularity**
- **AC2: Scalability**  
- **AC3: Performance**
- **AC4: Testability**

---

## 1. Ma Trận Đặc Điểm Kiến Trúc (Architecture Characteristics Matrix)

### 1.1. Bảng So Sánh Tổng Quát

Bảng dưới đây so sánh 8 phong cách kiến trúc phổ biến dựa trên 13 đặc điểm quan trọng:

| **Đặc Điểm / Tiêu Chí** | **Layered** | **Modular Monolith** | **Microkernel** | **Microservices** | **Service-based** | **Service-oriented (SOA)** | **Event-driven** | **Space-based** |
|-------------------------|-------------|----------------------|-----------------|-------------------|-------------------|----------------------------|------------------|-----------------|
| **Partitioning** | Technical | Domain | Technical | Domain | Domain | Technical | Technical | Technical |
| **Cost** | $ | $ | $ | $$$$$ | $$ | $$$$$ | $$ | $$$$$ |
| **Maintainability** | ⭐ | ⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐ | ⭐⭐⭐ | ⭐⭐⭐⭐ |
| **Testability** | ⭐⭐ | ⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐ | ⭐ | ⭐⭐ | ⭐ |
| **Deployability** | ⭐ | ⭐⭐⭐ | ⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐ | ⭐ | ⭐⭐⭐ | ⭐⭐⭐⭐ |
| **Simplicity** | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐ | ⭐⭐ | ⭐⭐ | ⭐ | ⭐⭐ | ⭐ |
| **Scalability** | ⭐ | ⭐⭐ | ⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ |
| **Elasticity** | ⭐ | ⭐⭐ | ⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐ | ⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ |
| **Responsiveness** | ⭐⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐ | ⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ |
| **Fault-tolerance** | ⭐ | ⭐⭐ | ⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐ | ⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ |
| **Evolvability** | ⭐ | ⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ |
| **Abstraction** | ⭐ | ⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐ | ⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐ |
| **Interoperability** | ⭐ | ⭐⭐ | ⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐ |

**Chú thích:**
- **Partitioning:** Technical (phân chia theo kỹ thuật) vs Domain (phân chia theo nghiệp vụ)
- **Cost:** $ (thấp) → $$$$$ (rất cao)
- **Đặc điểm khác:** ⭐ (kém) → ⭐⭐⭐⭐⭐ (xuất sắc)

### 1.2. Giải Thích Các Đặc Điểm Quan Trọng

#### **Partitioning (Phương Pháp Phân Chia)**
- **Technical:** Phân chia theo các tầng kỹ thuật (Presentation, Business, Data)
  - *Ví dụ:* Layered, Event-driven, Space-based
- **Domain:** Phân chia theo các miền nghiệp vụ (User Management, Content Delivery, Assessment)
  - *Ví dụ:* Microservices, Service-based, Modular Monolith

#### **Maintainability (Khả Năng Bảo Trì)**
- **Cao (⭐⭐⭐⭐⭐):** Microservices - mỗi service độc lập, dễ sửa đổi
- **Thấp (⭐):** Layered, SOA - tight coupling, thay đổi lan tỏa

#### **Testability (Khả Năng Kiểm Thử)**
- **Cao (⭐⭐⭐⭐⭐):** Microservices - test độc lập từng service
- **Thấp (⭐):** SOA, Space-based - phụ thuộc nhiều, khó mock

#### **Scalability (Khả Năng Mở Rộng)**
- **Cao (⭐⭐⭐⭐⭐):** Microservices, Event-driven, Space-based - scale từng component
- **Thấp (⭐):** Layered - phải scale toàn bộ application

---

## 2. Ánh Xạ Architecture Characteristics Của ITS

Dựa trên các **Architecture Characteristics** đã xác định cho ITS, chúng ta ánh xạ chúng với bảng so sánh:

| **ITS Architecture Characteristics** | **Tương Ứng Trong Bảng** | **Ưu Tiên** |
|--------------------------------------|--------------------------|-------------|
| AC1: Modularity | Maintainability, Evolvability | ⭐⭐⭐ (Cao) |
| AC2: Scalability | Scalability, Elasticity | ⭐⭐⭐ (Cao) |
| AC3: Performance | Responsiveness | ⭐⭐⭐ (Cao) |
| AC4: Testability | Testability | ⭐⭐⭐ (Cao) |
| AC5: Deployability | Deployability | ⭐⭐⭐ (Cao) |
| AC6: Security | *(Cần đánh giá riêng)* | ⭐⭐⭐ (Cao) |

---

## 3. So Sánh Các Phong Cách Kiến Trúc Cho ITS

### 3.1. Đánh Giá Chi Tiết

| **Architecture Style** | **Phù Hợp Với ITS?** | **Ưu Điểm Cho ITS** | **Nhược Điểm Cho ITS** | **Điểm Tổng** |
|------------------------|----------------------|---------------------|------------------------|---------------|
| **Layered** | ❌ Không | - Đơn giản (⭐⭐⭐⭐⭐)<br>- Chi phí thấp ($) | - Scalability thấp (⭐)<br>- Testability thấp (⭐⭐)<br>- Không hỗ trợ Modularity | **2/10** |
| **Modular Monolith** | ⚠️ Có thể (cho MVP) | - Chi phí thấp ($)<br>- Maintainability tốt (⭐⭐⭐)<br>- Đơn giản vừa phải (⭐⭐⭐⭐) | - Scalability hạn chế (⭐⭐)<br>- Testability trung bình (⭐⭐⭐)<br>- Không hỗ trợ live model swapping | **5/10** |
| **Microkernel** | ⚠️ Có thể | - Evolvability cao (⭐⭐⭐⭐⭐)<br>- Abstraction tốt (⭐⭐⭐⭐⭐)<br>- Hỗ trợ plugin architecture | - Scalability hạn chế (⭐⭐)<br>- Responsiveness thấp (⭐⭐)<br>- Không phù hợp với distributed system | **6/10** |
| **Microservices** | ✅ **Rất Phù Hợp** | - Maintainability xuất sắc (⭐⭐⭐⭐⭐)<br>- Testability xuất sắc (⭐⭐⭐⭐⭐)<br>- Scalability xuất sắc (⭐⭐⭐⭐⭐)<br>- Deployability xuất sắc (⭐⭐⭐⭐⭐)<br>- **Hỗ trợ live model swapping** | - Chi phí cao ($$$$$)<br>- Độ phức tạp cao (⭐⭐)<br>- Cần DevOps expertise | **9/10** |
| **Service-based** | ⚠️ Có thể | - Chi phí vừa phải ($$)<br>- Scalability tốt (⭐⭐⭐⭐)<br>- Interoperability tốt (⭐⭐⭐⭐) | - Testability thấp (⭐⭐)<br>- Deployability thấp (⭐⭐)<br>- Không linh hoạt như Microservices | **6/10** |
| **SOA** | ❌ Không | - Interoperability cao (⭐⭐⭐⭐⭐) | - Chi phí rất cao ($$$$$)<br>- Simplicity rất thấp (⭐)<br>- Testability thấp (⭐)<br>- **Quá phức tạp cho ITS** | **3/10** |
| **Event-driven** | ✅ **Phù Hợp (Kết Hợp)** | - Scalability xuất sắc (⭐⭐⭐⭐⭐)<br>- Responsiveness xuất sắc (⭐⭐⭐⭐⭐)<br>- Fault-tolerance xuất sắc (⭐⭐⭐⭐⭐)<br>- **Phù hợp với real-time feedback** | - Chi phí vừa phải ($$)<br>- Testability thấp (⭐⭐)<br>- Cần message broker infrastructure | **8/10** |
| **Space-based** | ❌ Không | - Scalability xuất sắc (⭐⭐⭐⭐⭐)<br>- Elasticity xuất sắc (⭐⭐⭐⭐⭐) | - Chi phí rất cao ($$$$$)<br>- Simplicity rất thấp (⭐)<br>- Testability rất thấp (⭐)<br>- **Quá phức tạp, không cần thiết** | **4/10** |

### 3.2. So Sánh Chi Tiết: Monolithic vs Microservices

| **Tiêu chí** | **Monolithic Architecture** | **Microservices Architecture** |
|--------------|----------------------------|-------------------------------|
| **AC1: Modularity** | Thấp. Mặc dù có thể tổ chức module lô-gíc (package), việc triển khai vật lý là đơn lẻ, dẫn đến khớp nối cứng nhắc giữa các module. | Rất Cao. Các dịch vụ AI/Domain (ví dụ: Adaptive Engine, Scoring Engine) được cô lập thành các đơn vị triển khai độc lập (Architecture Quantum). |
| **AC2: Scalability** | Thấp. Phải nhân bản toàn bộ ứng dụng (kể cả phần không cần mở rộng). Khó mở rộng các module tính toán nặng (AI) độc lập. | Rất Cao. Có thể Scale Horizontally chỉ các service cần thiết (ví dụ: Scoring Engine) để xử lý tải nặng. |
| **AC3: Performance** | Trung bình. Không có network overhead, nhưng khó tối ưu từng phần. | Cao. Có thể tối ưu từng service, nhưng có network latency. |
| **AC4: Testability** | Trung bình. Khó kiểm thử Logic Nghiệp vụ vì nó thường phụ thuộc vào Data Access Layer, làm giảm tính độc lập. | Rất Cao. Mỗi service là một khối mã nhỏ, dễ dàng kiểm thử đơn vị độc lập. Logic cốt lõi (Domain Service) được tách khỏi I/O (tuân thủ DIP). |
| **AC5: Deployability** | Thấp. Thay đổi nhỏ (bug fix) yêu cầu triển khai lại toàn bộ ứng dụng, không thể hỗ trợ Live AI Model Swapping (FR9). | Rất Cao. Cho phép triển khai độc lập và liên tục từng dịch vụ (ví dụ: hoán đổi phiên bản Adaptive Engine V2 mà không cần downtime). |
| **Simplicity** | Cao. Đơn giản khi bắt đầu, vận hành và gỡ lỗi (debugging) dễ hơn. | Thấp. Độ phức tạp cao (network, distributed tracing, service mesh). |

---

## 4. Quyết Định Kiến Trúc

### 4.1. Lựa Chọn: Hybrid Microservices + Event-Driven Architecture

**Quyết định:** Chúng ta sẽ sử dụng **Microservices + Event-Driven (Hybrid Architecture)**

**Lý do:**

1. **Microservices làm nền tảng:**
   - Đáp ứng tất cả yêu cầu về Modularity (AC1), Scalability (AC2), Testability (AC4)
   - Hỗ trợ live AI model swapping (FR9, FR12)
   - Cho phép independent deployment (AC5: Deployability)
   - Mỗi service có thể sử dụng công nghệ phù hợp nhất

2. **Event-Driven cho real-time components:**
   - Xử lý real-time feedback (FR6)
   - Xử lý adaptive learning (FR4)
   - Giảm coupling giữa các services
   - Tăng responsiveness và fault-tolerance

### 4.2. Justification (Biện Minh Chi Tiết)

#### **Bắt buộc về Modularity & Deployability**
ITS có yêu cầu cốt lõi là **Live AI Model Swapping (FR9, FR12)**. Chỉ Microservices mới cho phép:
- Triển khai và hoán đổi các phiên bản Mô hình AI (ví dụ: AdaptivePathGenerator) độc lập
- Không gây downtime khi cập nhật
- Củng cố tính Modularity (AC1) và Deployability (AC5)

#### **Yêu cầu về Scalability**
Tải tính toán nặng và biến động của các thuật toán AI/ML (FR7) yêu cầu:
- Mở rộng độc lập từng service
- Auto-scaling theo nhu cầu thực tế
- Điều mà Monolithic không thể đáp ứng hiệu quả (AC2)

#### **Hỗ trợ Testability**
Microservices, khi kết hợp với Clean/Hexagonal Architecture bên trong mỗi service:
- Logic AI/Domain có tính Testability (AC4) rất cao
- Cần thiết cho tính đúng đắn của thuật toán
- Dễ dàng unit test và integration test

---

## 5. Lựa Chọn Mẫu Kiến Trúc Nội Bộ (Internal Architecture Pattern)

### 5.1. Clean Architecture / Hexagonal Architecture

Để đảm bảo **Testability (AC4)** và **Maintainability (AC7)** bên trong mỗi Microservice, ta sẽ áp dụng:

**Clean Architecture (hoặc Hexagonal/Onion Architecture):**
- Bảo vệ Logic Nghiệp vụ Cốt lõi (Domain Services) khỏi các chi tiết bên ngoài
- Tách biệt khỏi Database, Framework, và UI

**Thực thi Dependency Inversion Principle (DIP):**
- Các Interactor/Use Cases (Policy Modules, như AdaptivePathGenerator) chỉ phụ thuộc vào Interfaces (Abstraction)
- Không phụ thuộc vào các lớp triển khai (Concretion)

**Thực thi Single Responsibility Principle (SRP):**
- Đảm bảo các lớp bên trong mỗi tầng có độ Functional Cohesion cao nhất
- Mỗi class chỉ có một lý do để thay đổi

### 5.2. Cấu Trúc Bên Trong Mỗi Service

```
Microservice (ví dụ: Learner Model Service)
├── Domain Layer (Entities, Value Objects, Domain Services)
│   └── Logic nghiệp vụ cốt lõi, độc lập với framework
├── Application Layer (Use Cases, Interactors)
│   └── Orchestrate domain logic, implement business flows
├── Infrastructure Layer (Repositories, External Services)
│   └── Database access, HTTP clients, message queues
└── Presentation Layer (Controllers, DTOs)
    └── REST API, GraphQL, gRPC endpoints
```

---

## 6. Trade-offs Được Chấp Nhận

### 6.1. Complexity vs Scalability & Modularity

| **Trade-off** | **Quyết Định** | **Mitigation** |
|---------------|---------------|----------------|
| **Simplicity (⭐⭐) ↔ Scalability (⭐⭐⭐⭐⭐)** | Chấp nhận độ phức tạp cao để đạt Scalability tối ưu | - Sử dụng managed Kubernetes (GKE, EKS)<br>- Infrastructure as Code (Terraform)<br>- Training DevOps team |
| **Development Cost (+30%) ↔ Testability (⭐⭐⭐⭐⭐)** | Chấp nhận chi phí cao hơn để đạt Testability cao | - Invest in testing frameworks<br>- TDD practices<br>- Automated testing |
| **Network Latency ↔ Modularity** | Chấp nhận latency để đạt tính module hóa | - Implement caching (Redis)<br>- Use gRPC cho internal calls<br>- Optimize service boundaries |

### 6.2. Tổng Kết Trade-offs

**Chấp nhận:**
- ❌ High complexity (Simplicity ⭐⭐)
- ❌ High cost ($$$$$)
- ❌ Steep learning curve
- ❌ Network latency overhead

**Đạt được:**
- ✅ Excellent Modularity (⭐⭐⭐⭐⭐)
- ✅ Excellent Scalability (⭐⭐⭐⭐⭐)
- ✅ Excellent Testability (⭐⭐⭐⭐⭐)
- ✅ Excellent Deployability (⭐⭐⭐⭐⭐)
- ✅ Support for Live AI Model Swapping

---

## 7. Kết Luận

### 7.1. Architecture Style Cuối Cùng

**Hybrid Microservices + Event-Driven Architecture**

**Justification:**
1. ✅ Đáp ứng tất cả Architecture Characteristics quan trọng của ITS
2. ✅ Hỗ trợ tất cả Functional Requirements (đặc biệt FR9, FR12: Live AI Model Swapping)
3. ✅ Đáp ứng Non-Functional Requirements (scalability, performance, availability)
4. ⚠️ Trade-offs chấp nhận được với chiến lược mitigation rõ ràng

### 7.2. Internal Architecture Pattern

**Clean Architecture** cho mỗi Microservice:
- High testability
- Technology independence
- Long-term maintainability
- Tuân thủ SOLID principles

---

**Tài liệu tham khảo:**
- Software Architecture: The Hard Parts (Neal Ford et al.)
- Building Microservices (Sam Newman)
- Fundamentals of Software Architecture (Mark Richards, Neal Ford)
- Clean Architecture (Robert C. Martin)
