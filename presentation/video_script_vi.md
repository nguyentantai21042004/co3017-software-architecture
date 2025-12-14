# Kịch Bản Chi Tiết: Intelligent Tutoring System

**Người trình bày:** Nguyễn Tấn Tài
**Thời lượng ước tính:** ~30 phút

---

## 0. Mở đầu

### Slide 1: Trang bìa
"Xin chào Thầy và các bạn. Em là Nguyễn Tấn Tài, đại diện cho nhóm 13. Hôm nay, em xin phép được trình bày báo cáo bài tập lớn môn Kiến trúc Phần mềm với đề tài là **'Intelligent Tutoring System'** - hay còn gọi là Hệ thống gia sư thông minh."

---

## 1. Tóm Tắt Tổng Quan (Executive Summary)

### Slide 2: Mục lục (Nếu có) hoặc Chuyển tiếp
"Bài thuyết trình của em sẽ đi qua **các nội dung chính** từ **tổng quan dự án**, **phân tích yêu cầu**, đến **thiết kế kiến trúc chi tiết** và cuối cùng là **demo sản phẩm**."

### Slide 3: Tầm Nhìn Dự Án
"Đầu tiên, về tầm nhìn của dự án. nhóm mong muốn xây dựng **một hệ thống học tập** có khả năng **cá nhân hóa trải nghiệm** NGHỈ **cho từng người dùng**, giống như việc có một gia sư riêng 1-kèm-1 vậy. Hệ thống không chỉ cung cấp bài học mà còn phải hiểu được năng lực của người học để điều chỉnh nội dung phù hợp. Mục tiêu kỹ thuật của nhóm là hệ thống phải phục vụ được hơn 5,000 người dùng đồng thời với thời gian phản hồi cực nhanh, dưới 500ms với mỗi thao tác với hệ thống"

### Slide 4: Thách Thức Chính
"Tuy nhiên, để đạt được tầm nhìn đó, nhóm phải đối mặt với những vấn đề cần xem xét.
Thứ nhất là áp lực về hiệu năng cao khi xem xét số lượng người dùng đồng thời lớn (Concurrent Users).
Thứ hai là bài toán mở rộng hệ thống (Scaling). Làm sao để khi người dùng tăng đột biến, hệ thống vẫn đứng vững, tốc độ phản vẫn vẫn đạt mức p80?
Thứ ba là khả năng cập nhật các mô hình AI mới nhất mà không được làm gián đoạn trải nghiệm của người dùng, hay còn gọi là Zero Downtime Deployment.
Và cuối cùng, code base phải sạch, tuân thủ nguyên tắc SOLID để dễ dàng bảo trì về sau."

### Slide 5: Giải Pháp Kiến Trúc
"Để giải quyết các thách thức trên, nhóm đã đề xuất giải pháp kiến trúc **Hybrid Microservices** kết hợp với **Event-Driven Architecture**.
Nhóm sử dụng **Polyglot Programming**, tức là đa ngôn ngữ lập trình. Cụ thể là kết hợp Java Spring Boot cho các xử lý nghiệp vụ phức tạp và Go (Gin framework) cho các tác vụ cần hiệu năng cao.
Hệ thống sử dụng PostgreSQL làm cơ sở dữ liệu chính và RabbitMQ để đảm bảo giao tiếp bất đồng bộ giữa các service."

---

## 2. Phân Tích Bối Cảnh (Context Analysis)

### Slide 6: Vision Statement
"Về Vision Statement. nhóm muốn tái hiện trải nghiệm học tập nơi mà nội dung được điều chỉnh theo đúng tốc độ và sở thích của từng cá nhân. Hệ thống phải có khả năng mở rộng để phục vụ hàng nghìn sinh viên cùng lúc mà không bị nghẽn cổ chai."

### Slide 7: Bối Cảnh Kinh Doanh
"Về bối cảnh kinh doanh, thị trường E-learning đang chuyển dịch mạnh mẽ sang hướng Cá nhân hóa (Personalization).
Khách hàng mục tiêu của nhóm bao gồm học sinh K-12, sinh viên đại học, và cả các giảng viên, quản trị viên hệ thống.
Tiêu chí thành công của dự án được đo lường bằng khả năng chịu tải 5,000 CCU, cam kết SLA 99.5% và mục tiêu cải thiện 40% tỷ lệ giữ chân người học (retention rate)."

### Slide 8: Bối Cảnh Kỹ Thuật
"Về mặt kỹ thuật, hệ thống này không đứng độc lập hoàn toàn mà cần tích hợp với các hệ thống Learning Management System (LMS) hiện có.
nhóm sử dụng Auth Service tập trung với chuẩn JWT và OAuth2 để bảo mật.
Toàn bộ hệ thống được container hóa và vận hành trên Kubernetes để tận dụng khả năng Auto-scaling.
Ngoài ra, để giám sát sức khỏe hệ thống, nhóm tích hợp bộ công cụ Prometheus, Grafana và Loki."

---

## 3. Phân Tích Stakeholders

### Slide 9: Ma Trận Stakeholder
"Tiếp theo là phân tích các bên liên quan (Stakeholders).
Quan trọng nhất là **Learner (Người học)** và **Instructor (Giảng viên)**. Đây là nhóm có ảnh hưởng cao và sự quan tâm cao, nên nhóm xếp vào nhóm 'Manage Closely'.
Gọi ý: Khi nói đến đây, hãy chỉ vào góc trên bên phải của ma trận.
**Admin** quan tâm đến chi phí và bảo mật, trong khi **AI Engineer** cần môi trường để thử nghiệm thuật toán mới.
Cuối cùng là **System Architect**, người đảm bảo tính toàn vẹn của kiến trúc Clean Architecture."

### Slide 10: Chiến Lược Quản Lý
"Dựa trên ma trận đó, nhóm đưa ra chiến lược quản lý cụ thể. Tập trung tối đa nguồn lực để thỏa mãn Learner và Instructor vì họ quyết định sự thành bại của sản phẩm. Nhóm Admin và Architect cần được giữ hài lòng (Keep Satisfied) thông qua các báo cáo và tuân thủ quy chuẩn kỹ thuật."

---

## 4. Yêu Cầu Chức Năng (Functional Requirements)

### Slide 11: User Stories Chính
"Hệ thống được xây dựng xoay quanh các User Stories cốt lõi.
Quan trọng nhất là khả năng 'Đánh giá kiến thức' để đề xuất lộ trình.
Hệ thống phải có khả năng gợi ý và giải thích ngay lập tức khi người học mắc lỗi sai (độ trễ dưới 500ms).
Người học cũng cần một Dashboard trực quan để theo dõi tiến độ của mình theo thời gian thực."

### Slide 12: Use Cases Quan Trọng
"Từ User Stories, nhóm chi tiết hóa thành các Use Cases.
Nổi bật là **UC-04: Xây dựng Learner Model**, nơi hệ thống tạo dựng hồ sơ năng lực người học từ bài kiểm tra đầu vào.
Và **UC-18: Live Swap AI Model**, một tính năng kỹ thuật cho phép thay nóng các mô hình gợi ý mà không cần tắt restart server."

---

## 5. Yêu Cầu Phi Chức Năng (Non-Functional Requirements)

### Slide 13: Các Đặc Tính Kiến Trúc Chính
"Về yêu cầu phi chức năng, nhóm xác định 4 trụ cột kiến trúc (Architecture Characteristics) quan trọng nhất:
1.  **Modularity (Tính mô-đun hóa):** Để đảm bảo Low Coupling, High Cohesion.
2.  **Scalability (Khả năng mở rộng):** Để đáp ứng tải 5,000 người dùng.
3.  **Performance (Hiệu năng):** Đảm bảo phản hồi nhanh.
4.  **Testability (Khả năng kiểm thử):** Đảm bảo độ tin cậy của phần mềm."

### Slide 14: Biểu đồ Radar Đặc Tinh Kiến Trúc
"Biểu đồ Radar trên màn hình thể hiện sự đánh đổi (Trade-offs) của nhóm. nhóm chấp nhận hy sinh một chút sự đơn giản (Simplicity) để đổi lấy Scalability và Modularity cao, phù hợp với tính chất của một hệ thống phân tán phức tạp."

---

## 6. Thiết Kế Kiến Trúc (Architecture Design)

### Slide 15: So Sánh Các Phương Án
"Trước khi chốt phương án cuối cùng, nhóm đã cân nhắc kỹ lưỡng.
Mô hình **Monolith** thì đơn giản lúc đầu nhưng về sau sẽ rất khó scale và deploy.
**Serverless** thì chi phí khó dự đoán và gặp vấn đề cold-start.
Do đó, **Microservices** là lựa chọn phù hợp nhất dù độ phức tạp ban đầu cao hơn."

### Slide 16: Quyết Định - Hybrid Microservices
"Chính xác hơn, nhóm chọn **Hybrid Microservices**. Tại sao là Hybrid? Vì nhóm không chia nhỏ mọi thứ một cách cực đoan. Mỗi service sẽ tương ứng với một Bounded Context lớn, ví dụ như Scoring hay Content. Các service này giao tiếp bất đồng bộ qua RabbitMQ để giảm sự phụ thuộc trực tiếp lẫn nhau (Decoupling)."

---

## 7. Các Quyết Định Kiến Trúc (ADRs)

### Slide 17: ADR-1 Polyglot Programming
"Em xin trình bày sâu hơn về các quyết định kiến trúc (ADR).
Đầu tiên là **Polyglot Programming**. nhóm dùng Java cho Content Service vì Java có hệ sinh thái thư viện doanh nghiệp rất mạnh. Còn với Scoring Service hay Adaptive Engine, nơi cần tốc độ tính toán ma trận nhanh, nhóm chọn Go (Golang). Đây là chiến lược 'Best tool for the job'."

### Slide 18: ADR-2 Database per Service
"Thứ hai là **Database per Service**. Mỗi Microservice sở hữu database riêng biệt. Content Service có DB riêng, Learner Model có DB riêng. Điều này ngăn chặn việc một service này chọc ngoáy vào dữ liệu của service kia, đảm bảo tính đóng gói."

### Slide 19: ADR-3 Clean Architecture
"Thứ ba là **Clean Architecture**. Trong từng service, code được tổ chức thành các lớp (Layers) rõ ràng: Domain, Application, Infrastructure. Quy tắc bất di bất dịch là lớp bên trong không được phụ thuộc vào lớp bên ngoài. Điều này giúp nhóm dễ dàng viết Unit Test cho Business Logic mà không cần quan tâm đến Database hay Web Framework."

---

## 8. Các Góc Nhìn Kiến Trúc (Architecture Views)

### Slide 20: Module View
"Ở góc nhìn Module, hệ thống gồm 4 khối chính:
**Content Service** quản lý ngân hàng câu hỏi.
**Scoring Service** đảm nhận việc chấm điểm.
**Learner Model** lưu trữ hồ sơ năng lực người học.
Và **Adaptive Engine** đóng vai trò bộ não, sử dụng dữ liệu từ Learner Model để quyết định bài học tiếp theo."

### Slide 21: Component & Connector View
"Ở Component View, các bạn có thể thấy rõ luồng giao tiếp. Các mũi tên liền là gọi REST API đồng bộ, các mũi tên đứt là bắn message qua RabbitMQ. Việc dùng Message Queue giúp Scoring Service sau khi chấm điểm xong có thể trả về kết quả ngay cho người dùng mà không cần đợi Learner Model cập nhật xong dữ liệu, giúp giảm độ trễ đáng kể."

### Slide 22: Allocation View
"Về triển khai (Allocation View), mỗi service được đóng gói trong một Docker Container riêng biệt. Chúng được orchestrate bởi Kubernetes, cho phép nhóm scale số lượng Pods của Scoring Service lên độc lập khi có nhiều người nộp bài cùng lúc."

---

## 9. Góc Nhìn Hành Vi (Behavior View)

### Slide 23: Luồng Adaptive Content Delivery
"Hãy cùng xem qua luồng hoạt động chính: **Giao bài học thích ứng**.
1. Khi Client yêu cầu bài học mới, request đến Adaptive Engine.
2. Engine gọi Learner Model để lấy chỉ số thành thạo hiện tại.
3. Dựa trên chỉ số đó, Engine tính toán và yêu cầu Content Service lấy ra câu hỏi phù hợp nhất (không quá dễ, không quá khó).
4. Kết quả được trả về Client.
Tất cả diễn ra gần như tức thời."

### Slide 24: Luồng Assessment & Scoring
"Với luồng **Chấm điểm**:
1. Sinh viên nộp bài -> Scoring Service chấm điểm và lưu kết quả.
2. Ngay lập tức, Scoring Service bắn sự kiện `SubmissionGraded` vào RabbitMQ.
3. Learner Model Service lắng nghe sự kiện này và cập nhật lại hồ sơ năng lực của sinh viên một cách âm thầm (background process)."

---

## 10. Áp Dụng SOLID

### Slide 25: Nguyên lý SOLID
"Kiến trúc này là hiện thân của SOLID.
**Single Responsibility (SRP):** Mỗi service chỉ làm một việc. Scoring chỉ chấm điểm, không quản lý nội dung.
**Open/Closed (OCP):** nhóm dùng Strategy Pattern cho các thuật toán chấm điểm. Muốn thêm thuật toán mới? Chỉ cần viết thêm class mới implement interface, không cần sửa code cũ.
**Dependency Inversion (DIP):** Tầng Use Case chỉ làm việc với Interface, không phụ thuộc vào Database cụ thể nào."

---

## 11. Triển Khai & Kết Quả (Implementation Status)

### Slide 26: Trạng Thái MVP
"Hiện tại, nhóm đã hoàn thành MVP.
4/7 Service core đã chạy ổn định.
Các luồng nghiệp vụ chính như Adaptive Delivery và Scoring đã được kiểm chứng (Verified).
Cơ sở dữ liệu đã thiết kế và migrate thành công."

### Slide 27: Kết Quả Đạt Được
"Các chỉ số code quality rất khả quan. Complexity được giữ ở mức thấp 7.2. Coupling giữa các module trung bình chỉ 3.8. Đặc biệt, tính năng Adaptive Learning đã hoạt động 100% theo thiết kế."

---

## 12. Đánh Giá (Evaluation)

### Slide 28: Điểm Mạnh & Thách Thức
"Nhìn lại quá trình, điểm sáng lớn nhất là sự kết hợp giữa **Polyglot** và **Clean Architecture**. Nó giúp code rất tường minh và dễ test.
Tuy nhiên, khó khăn là sự phức tạp trong quản lý Distributed Transaction và việc setup môi trường dev tốn nhiều RAM hơn so với Monolith."

---

## 13. Hướng Tương Lai (Future Roadmap)

### Slide 29: Roadmap
"Trong sprint tới, nhóm sẽ tập trung hoàn thiện User Management Service và tích hợp API Gateway để quản lý routing tập trung. Xa hơn nữa là tính năng Analytics Dashboard cho giảng viên."

---

## 14. Demo & Code Showcase

### Slide 30 - 35: Demo Ứng Dụng
**(Phần này bạn vừa thao tác trên màn hình vừa nói)**
"Bây giờ, em xin mời Thầy và các bạn xem demo thực tế.
Đây là **Trang chủ**, giao diện rất trực quan. Em sẽ đóng vai một học viên mới...
Em vào phần **Học tập**. Hệ thống đưa ra câu hỏi đầu tiên ở mức độ trung bình...
Em chọn đáp án sai. Hệ thống giải thích ngay tại sao sai... Và quan trọng hơn, câu hỏi tiếp theo hệ thống đưa ra đã được hạ độ khó xuống để giúp em củng cố kiến thức nền tảng...
Sau khi hoàn thành bài học, em quay lại **Dashboard**. Các bạn thấy đấy, biểu đồ năng lực của em đã thay đổi ngay lập tức."

### Slide 36 - 38: Code Tour
"Về code, đây là cấu trúc pres project Java (Content Service)... Các package được chia theo Clean Architecture: Domain, Infrastructure, Application...
Còn đây là project Go (Scoring Service)... Cấu trúc folder chuẩn Go Standard Layout...
File Docker Compose này giúp dựng toàn bộ hệ thống chỉ với một lệnh `docker-compose up`."

---

## 15. Kết Luận

### Slide 39: Kết Luận
"Tổng kết lại, ITS không chỉ là một ứng dụng web, mà là một minh chứng cho việc áp dụng các kiến trúc phần mềm hiện đại như Microservices, Event-Driven và Clean Architecture để giải quyết bài toán giáo dục phức tạp."

### Slide 40: Q&A
"Em xin kết thúc phần trình bày tại đây. Rất mong nhận được câu hỏi và góp ý từ Thầy và các bạn. Em xin cảm ơn!"

---
**[HẾT]**
