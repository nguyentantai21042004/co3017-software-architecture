Dưới đây là nội dung chi tiết cho Chương 2, Mục 2.1, được tổng hợp chính xác từ các tài liệu bạn đã cung cấp.

---

## CHAPTER 2: ARCHITECTURE DESIGN

### 2.1 Architecture Characteristics Prioritization

Việc ưu tiên hóa các Đặc tính Kiến trúc (Architecture Characteristics - ACs) là bước nền tảng, đóng vai trò là cầu nối chiến lược giữa yêu cầu nghiệp vụ và các quyết định kỹ thuật. Các quyết định này định hình cấu trúc hệ thống và quản lý các đánh đổi (trade-offs) không thể tránh khỏi.

#### 2.1.1 Characteristics Matrix

Bảng ma trận dưới đây tổng hợp 9 đặc tính kiến trúc đã được xác định cho hệ thống ITS, phân tích tác động nghiệp vụ, rủi ro/độ phức tạp kỹ thuật, và mức độ ưu tiên tương ứng.

| Characteristic | Business Impact | Technical Risk / Effort | Priority | Notes (Vai trò chính) |
| :--- | :--- | :--- | :--- | :--- |
| **Modularity (AC1)** | ⭐⭐⭐⭐⭐ **High** | **High** | **1** (Cao nhất) | Định hình kiến trúc Microservices; cô lập logic AI; hỗ trợ "Live AI Model Swapping" (FR12). |
| **Scalability (AC2)** | ⭐⭐⭐⭐⭐ **High** | **High** | **1** (Cao nhất) | Xử lý tải người dùng đồng thời (≥ 5,000) và các tác vụ tính toán AI nặng. |
| **Performance (AC3)** | ⭐⭐⭐⭐⭐ **High** | **Low** | **1** (Cao nhất) | Đảm bảo trải nghiệm học tập mượt mà; phản hồi thời gian thực (\<500ms). |
| **Testability (AC4)** | ⭐⭐⭐⭐ **High** | **Low** | **2** (Cao) | Đảm bảo tính đúng đắn của thuật toán AI; hỗ trợ bởi Clean Architecture (ADR-3). |
| **Security (AC6)** | ⭐⭐⭐⭐ **High** | **High** | **2** (Cao) | Bảo vệ dữ liệu nhạy cảm (PII) của người học và nội dung (FR11); Tuân thủ GDPR/FERPA. |
| **Maintainability (AC7)** | ⭐⭐⭐ **Medium** | **Low** | **2** (Cao) | Giảm chi phí vòng đời; dễ dàng sửa lỗi và cải tiến hệ thống. |
| **Deployability (AC5)** | ⭐⭐⭐ **Medium** | **Medium** | **3** (Trung bình) | Hỗ trợ triển khai độc lập từng service và "Live AI Model Swapping". |
| **Observability (AC9)** | ⭐⭐ **Medium** | **Low** | **3** (Trung bình) | Debug, monitor, và phát hiện sự cố trong hệ thống microservices phân tán. |
| **Extensibility (AC8)** | ⭐ **Low** | **Medium** | **4** (Thấp) | Hỗ trợ thêm tính năng mới (ví dụ: loại câu hỏi mới) mà không sửa code lõi (tuân thủ OCP). |

---

#### 2.1.2 Trade-off Analysis

Mọi quyết định kiến trúc đều là sự đánh đổi. Dưới đây là các phân tích đánh đổi quan trọng nhất được chấp nhận để tối ưu cho các đặc tính đã ưu tiên.

**Trade-off: (AC1) Modularity & (AC2) Scalability vs. Simplicity (Độ đơn giản)**

* **Scenario:** Khi lựa chọn kiểu kiến trúc (ví dụ: Microservices) để hỗ trợ khả năng mở rộng cho \>5,000 người dùng và yêu cầu nghiệp vụ cốt lõi "Live AI Model Swapping" (FR12).
* **Decision:** Chọn **Modularity & Scalability** hơn Simplicity.
* **Rationale:** Chấp nhận độ phức tạp vận hành (complexity) cao hơn của kiến trúc Microservices. Đây là yêu cầu bắt buộc để đáp ứng tính mô-đun (cô lập logic AI) và khả năng triển khai độc lập (FR12), vốn là yêu cầu cốt lõi của hệ thống ITS.
* **Mitigation (Giảm thiểu):**
    1.  **Lộ trình Phát triển (Evolution Path):** Bắt đầu với **Modular Monolith** trong giai đoạn MVP (Giai đoạn 1) để giảm độ phức tạp ban đầu, sau đó di trú dần sang Microservices bằng **Strangler Fig Pattern** khi thực sự cần (Giai đoạn 2).
    2.  **Sử dụng Dịch vụ Managed:** Giảm tải vận hành bằng cách sử dụng các dịch vụ được quản lý (ví dụ: GKE, EKS cho Kubernetes).
    3.  **Đào tạo & Tiêu chuẩn hóa:** Đào tạo team về DevOps, chuẩn hóa cấu trúc dự án và chia sẻ các mẫu CI/CD.

**Trade-off: (AC4) Testability vs. Development Cost (Chi phí Phát triển)**

* **Scenario:** Khi quyết định cấu trúc code bên trong mỗi service để đảm bảo logic thuật toán AI (FR7) có thể được kiểm thử độc lập và chính xác.
* **Decision:** Chọn **Testability** hơn chi phí phát triển ban đầu.
* **Rationale:** Áp dụng bắt buộc **Clean/Hexagonal Architecture** (ADR-3) và **Repository Pattern** (ADR-4). Mặc dù việc này đòi hỏi cấu trúc code phức tạp hơn (nhiều boilerplate, interfaces) và đường cong học tập (learning curve) khó hơn, nó mang lại lợi ích chiến lược dài hạn: chi phí bảo trì (AC7) thấp hơn và độ tin cậy của thuật toán AI cao hơn.
* **Mitigation (Giảm thiểu):**
    1.  **Cung cấp Mẫu (Templates):** Tạo các mẫu code (code templates) và ví dụ rõ ràng cho cả Java và Golang để đảm bảo tính nhất quán (như đã định nghĩa trong ADR-3 và ADR-4).
    2.  **Chiến lược Kiểm thử rõ ràng:** Định nghĩa rõ "Testing Pyramid" (ADR-5) để team tập trung vào Unit Test (mục tiêu \>80% coverage) cho logic nghiệp vụ, giảm sự phụ thuộc vào E2E test chậm và đắt đỏ.

**Trade-off: (AC6) Security vs. (AC3) Performance**

* **Scenario:** Khi triển khai các biện pháp bảo mật bắt buộc, như mã hóa PII và xác thực token, vốn tiêu tốn thêm tài nguyên CPU và tăng độ trễ (latency).
* **Decision:** Chọn **Security** hơn Performance (trong giới hạn chấp nhận được).
* **Rationale:** Việc bảo vệ dữ liệu nhạy cảm PII của người học (tuân thủ GDPR/FERPA) và đảm bảo tính toàn vẹn của `LearnerModel` là **không thể thương lượng**. Chúng ta chấp nhận một độ trễ nhỏ (ví dụ: 50-100ms) do mã hóa (TLS, pgcrypto) và xác thực JWT tại API Gateway (ADR-6).
* **Mitigation (Giảm thiểu):**
    1.  **Tối ưu Hóa tại Cổng:** Thực hiện xác thực JWT tập trung tại API Gateway (ADR-6), thay vì để mỗi service tự validate, giúp giảm gánh nặng cho các service nghiệp vụ.
    2.  **Mã hóa Chọn lọc:** Chỉ mã hóa các cột PII nhạy cảm nhất (ADR-7), thay vì mã hóa toàn bộ cơ sở dữ liệu (Full Disk Encryption), để cân bằng giữa an ninh và hiệu năng query.

**Trade-off: (AC3) Performance vs. (AC1) Modularity/Coupling**

* **Scenario:** Khi phân rã hệ thống thành các Microservices. Việc phân chia quá mịn (fine-grained) để tối ưu Modularity có thể dẫn đến quá nhiều lời gọi mạng (network calls), làm tăng độ trễ và ảnh hưởng tiêu cực đến Performance.
* **Decision:** Chọn **Cân bằng (Balanced Granularity)**, ưu tiên Performance cho các luồng nghiệp vụ thời gian thực.
* **Rationale:** Phải đảm bảo các chức năng nghiệp vụ liên quan chặt chẽ (Functional Cohesion) được đóng gói trong cùng một Service (Architecture Quantum). Ví dụ: logic chấm điểm và cập nhật model ban đầu có thể nằm cùng nhau để tránh giao tiếp mạng không cần thiết, duy trì mục tiêu latency ≤ 500ms.
* **Mitigation (Giảm thiểu):**
    1.  **Tuân thủ DDD:** Phân rã service dựa trên Bounded Context của Domain-Driven Design (DDD), thay vì phân rã theo kỹ thuật, để đảm bảo tính gắn kết nghiệp vụ.
    2.  **Giao tiếp Nội bộ Hiệu quả:** Sử dụng gRPC hoặc các giao thức nhị phân hiệu quả (thay vì REST/JSON) cho giao tiếp service-to-service trong nội bộ (nếu cần).
    3.  **Circuit Breaker & Retry:** Áp dụng các mẫu (pattern) như Circuit Breaker và Retry Logic (với Exponential Backoff) để xử lý lỗi mạng và latency cao một cách linh hoạt.

    Chào bạn, tôi đã chuẩn bị nội dung chi tiết cho Mục 2.2, được trích xuất và tổng hợp chính xác từ các tài liệu phân tích (`2-architecture-characteristics.md` và `3-architecture-styles.md`) mà bạn đã cung cấp.

---

### 2.2 Architecture Style Selection

Việc lựa chọn kiểu kiến trúc (Architecture Style) là quyết định quan trọng nhất, định hình cấu trúc tổng thể của hệ thống. Quyết định này được đưa ra dựa trên một bộ tiêu chí đánh giá nghiêm ngặt, nhằm tìm ra kiến trúc "ít tệ nhất" (*the least worst architecture*), tối ưu hóa các đặc tính kiến trúc (ACs) đã được ưu tiên [cite: 2-architecture-characteristics.md, 3-architecture-styles.md].

#### 2.2.1 Evaluation Criteria

Các kiểu kiến trúc được đánh giá dựa trên các tiêu chí sau, bắt nguồn trực tiếp từ các yêu cầu nghiệp vụ và bối cảnh kỹ thuật của dự án ITS:

1.  **Alignment with Primary ACs (Mức độ đáp ứng các ACs chính):** Đây là tiêu chí quan trọng nhất. Kiến trúc được chọn phải tối ưu hóa 4 ACs chính đã ưu tiên:
    * **AC1: Modularity** (Hỗ trợ "Live AI Model Swapping" FR9, FR12)
    * **AC2: Scalability** (Hỗ trợ \>5,000 người dùng đồng thời)
    * **AC3: Performance** (Hỗ trợ phản hồi thời gian thực FR6)
    * **AC4: Testability** (Đảm bảo tính đúng đắn của thuật toán AI) [cite: 2-architecture-characteristics.md, 3-architecture-styles.md].
2.  **Technical Complexity (Độ phức tạp Kỹ thuật):** Đánh giá độ phức tạp khi phát triển, vận hành và giám sát (Observability) kiến trúc đó [cite: 3-architecture-styles.md]. Đây là một phần của trade-off (đánh đổi) với Scalability [cite: 2-architecture-characteristics.md].
3.  **Team Expertise (Kỹ năng của Đội ngũ):** Xem xét rủi ro về kỹ năng. Ví dụ, một kiến trúc Microservices đòi hỏi kỹ năng DevOps chuyên biệt để quản lý Kubernetes và Service Mesh, đây là một rủi ro phải được quản lý [cite: 3-architecture-styles.md].
4.  **Cost Implications (Tác động Chi phí):** Phân tích Tổng Chi phí Sở hữu (TCO), bao gồm chi phí hạ tầng ban đầu (ví dụ: Kubernetes, Message Broker) và chi phí bảo trì, mở rộng dài hạn [cite: 3-architecture-styles.md].

#### 2.2.2 Architecture Styles Comparison

Dựa trên các tiêu chí trên, chúng ta đã so sánh 8 phong cách kiến trúc phổ biến. Bảng dưới đây tóm tắt các lựa chọn chính, đánh giá chúng dựa trên các ACs ưu tiên và các tiêu chí phụ.

* **Đánh giá:** ⭐ (Kém) đến ⭐⭐⭐⭐⭐ (Xuất sắc).
* **Complexity (Độ phức tạp):** Đánh giá đảo ngược của "Simplicity" (Độ đơn giản) trong tệp phân tích [cite: 3-architecture-styles.md].
* **Cost (Chi phí):** $ (Thấp) đến $$$$$ (Rất cao) [cite: 3-architecture-styles.md].
* **Score (Điểm):** Điểm trung bình của 4 cột ACs (Modularity, Scalability, Performance, Complexity).



| Style | AC1: Modularity | AC2: Scalability | AC3: Performance | Complexity | Cost | **Score (Avg)** |
| :--- | :---: | :---: | :---: | :---: | :---: | :---: |
| Layered (Monolith) | ⭐ | ⭐ | ⭐⭐⭐⭐ | ⭐ | $ | **1.75 / 5.0** |
| Modular Monolith | ⭐⭐⭐ | ⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐ | $ | **2.75 / 5.0** |
| Microkernel | ⭐⭐⭐⭐ | ⭐⭐ | ⭐⭐ | ⭐⭐ | $ | **2.5 / 5.0** |
| **Microservices** | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐⭐ | $$$$$ | **4.5 / 5.0** |
| Service-based | ⭐⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐ | ⭐⭐⭐ | $$ | **3.25 / 5.0** |
| Service-oriented (SOA) | ⭐ | ⭐⭐⭐ | ⭐⭐ | ⭐⭐⭐⭐⭐ | $$$$$ | **2.75 / 5.0** |
| **Event-driven** | ⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ | $$ | **4.25 / 5.0** |
| Space-based | ⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ | $$$$$ | **4.5 / 5.0** |

**Phân tích Kết quả:**

* **Layered (Monolith):** Bị loại bỏ hoàn toàn do thất bại trong việc đáp ứng các ACs quan trọng nhất là Modularity (⭐) và Scalability (⭐) [cite: 3-architecture-styles.md].
* **Modular Monolith:** Là một phương án dự phòng tốt (Fallback Option) hoặc chiến lược MVP (Giai đoạn 1), vì nó cân bằng giữa chi phí ($) và Modularity/Testability (⭐⭐⭐) [cite: 3-architecture-styles.md].
* **Microservices:** Đạt điểm cao nhất (4.5) về khả năng đáp ứng các ACs, nhưng cũng đi kèm với Độ phức tạp (⭐⭐⭐⭐) và Chi phí ($$$$$) cao nhất [cite: 3-architecture-styles.md].
* **Event-driven:** Đạt điểm rất cao (4.25), đặc biệt xuất sắc ở Scalability (⭐⭐⭐⭐⭐) và Performance/Responsiveness (⭐⭐⭐⭐⭐). Đây là một thành phần bổ trợ hoàn hảo [cite: 3-architecture-styles.md].
* **Space-based:** Mặc dù đạt điểm cao (4.5), kiến trúc này được đánh giá là quá phức tạp (Complexity ⭐⭐⭐⭐⭐) và không cần thiết cho bối cảnh của ITS [cite: 3-architecture-styles.md].

---

#### 2.2.3 Final Architecture Decision

Dựa trên phân tích đánh giá và so sánh ở trên, quyết định kiến trúc cuối cùng được đưa ra.

**Selected Architecture: Hybrid Microservices + Event-Driven Architecture** [cite: 3-architecture-styles.md]

Chúng ta lựa chọn một kiến trúc lai (hybrid), kết hợp sức mạnh của hai phong cách:

1.  **Microservices** được sử dụng làm cấu trúc phân chia (partitioning) chính, dựa trên miền nghiệp vụ (Domain-Driven Design) [cite: 3-architecture-styles.md, main.pdf].
2.  **Event-Driven Architecture (EDA)** được sử dụng làm mẫu giao tiếp chính (communication pattern) cho các tương tác bất đồng bộ, xử lý thời gian thực [cite: 3-architecture-styles.md, main.pdf].

**Justification (Lý do):**

1.  **Tối ưu hóa các ACs Cốt lõi:** Kiến trúc Microservices đáp ứng trực tiếp các yêu cầu cao nhất:
    * **AC1: Modularity & AC5: Deployability:** Cho phép triển khai và hoán đổi các phiên bản Mô hình AI (ví dụ: `AdaptivePathGenerator`) một cách độc lập mà không gây downtime, đáp ứng yêu cầu **"Live AI Model Swapping" (FR9, FR12)** [cite: 3-architecture-styles.md].
    * **AC2: Scalability:** Cho phép mở rộng (scale) độc lập từng service (ví dụ: `ScoringService`) khi tải tính toán AI tăng cao, thay vì phải mở rộng toàn bộ ứng dụng [cite: 3-architecture-styles.md].
    * **AC4: Testability:** Khi kết hợp với Clean Architecture (quyết định trong mục 4.2 của file 3), mỗi service trở thành một khối mã nhỏ, dễ dàng kiểm thử độc lập, đảm bảo tính đúng đắn của thuật toán [cite: 3-architecture-styles.md].

2.  **Đáp ứng Yêu cầu Nghiệp vụ Thời gian thực:** Thành phần Event-Driven (EDA) là bắt buộc để:
    * Xử lý **"Real-time Feedback" (FR6)**: Khi một học sinh nộp bài, sự kiện `SubmissionCompleted` được bắn đi, cho phép nhiều service (Scoring, LearnerModel) xử lý song song và bất đồng bộ.
    * Hỗ trợ **"Adaptive Learning" (FR4, FR7):** Các cập nhật về mô hình người học (LearnerModel) có thể được xử lý qua hàng đợi (message queue), giúp tăng **AC3: Performance/Responsiveness** và **Fault-tolerance** (khả năng chịu lỗi) cho hệ thống [cite: 3-architecture-styles.md, main.pdf].

3.  **Quản lý Rủi ro và Đánh đổi (Trade-off):** Chúng ta nhận diện rõ rủi ro lớn nhất của lựa chọn này là **Độ phức tạp Vận hành Cao** [cite: 3-architecture-styles.md]. Quyết định này được đưa ra kèm theo một chiến lược giảm thiểu rủi ro (Mitigation Strategy) rõ ràng thông qua Lộ trình Phát triển theo từng giai đoạn [cite: 3-architecture-styles.md].

**Implementation Strategy (Chiến lược Triển khai):**

Để giảm thiểu rủi ro về độ phức tạp ban đầu, chúng ta sẽ **KHÔNG** xây dựng Microservices từ ngày đầu tiên. Thay vào đó, chúng ta áp dụng một lộ trình phát triển (Evolution Path) linh hoạt [cite: 3-architecture-styles.md]:

* **Phase 1: Modular Monolith (MVP)**
    * Xây dựng một ứng dụng Monolith duy nhất.
    * Bên trong monolith, code được tổ chức nghiêm ngặt thành các **modules nghiệp vụ** (ví dụ: `learner`, `content`, `scoring`) và bắt buộc tuân thủ **Clean Architecture** (ADR-3) và **DIP** (ADR-4) [cite: 3-architecture-styles.md, 5-architecture-decisions.md].
    * *Mục tiêu:* Time-to-market nhanh, chi phí thấp, và tập trung vào nghiệp vụ, trong khi vẫn đảm bảo AC4 (Testability) [cite: 3-architecture-styles.md].

* **Phase 2: Extract Critical Services (Strangler Fig Pattern)**
    * Khi hệ thống phát triển và có yêu cầu thực tế về Scalability (ví dụ: `ScoringService` bị quá tải), chúng ta sẽ áp dụng **Strangler Fig Pattern** [cite: 3-architecture-styles.md].
    * Tách module `Scoring` ra thành Microservice đầu tiên.
    * Một API Gateway (hoặc Proxy) sẽ được thiết lập để điều hướng traffic `/api/score` đến service mới, trong khi các traffic khác vẫn đi vào Monolith [cite: 3-architecture-styles.md].

* **Phase 3: Full Microservices Ecosystem**
    * Lặp lại quy trình ở Giai đoạn 2, tiếp tục "bóc tách" (extract) các service khác (như `AdaptiveEngine`, `LearnerModelService`) ra khỏi Monolith khi nhu cầu nghiệp vụ và tải hệ thống yêu cầu.
    * *Mục tiêu:* Đạt được một hệ thống Microservices phân tán hoàn chỉnh, tối ưu cho AC1, AC2, và AC5 [cite: 3-architecture-styles.md].