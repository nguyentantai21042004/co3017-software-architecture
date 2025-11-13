3.4 Behavioral View (Góc nhìn Hành vi)Góc nhìn Hành vi mô tả cách các thành phần kiến trúc (từ Mục 3.2 và 3.3) tương tác với nhau theo thời gian để thực hiện các kịch bản nghiệp vụ (use cases) cụ thể. Góc nhìn này sử dụng Sơ đồ Tuần tự (Sequence Diagrams) để làm rõ luồng thông điệp và sự cộng tác giữa các microservice.Góc nhìn này rất quan trọng để xác thực rằng kiến trúc đã chọn có thể đáp ứng các yêu cầu chức năng (FRs) và phi chức năng (ACs) phức tạp, đặc biệt là AC3: Performance (thông qua các luồng đồng bộ/bất đồng bộ) và Độ tin cậy.3.4.1 Các Kịch bản Chính (Key Scenarios)Dưới đây là các sơ đồ tuần tự cho 5 kịch bản quan trọng nhất của hệ thống ITS.1. Đăng ký Người dùng và Nhập môn (User Registration and Onboarding)Mô tả: Luồng này (UC-01) mô tả cách một người dùng mới (Learner) tạo tài khoản. Nó sử dụng một mô hình bất đồng bộ (event-driven) để tách rời việc tạo thông tin xác thực (AuthN) khỏi việc tạo hồ sơ người dùng (PII), tuân thủ ADR-6 và ADR-7.Các thành phần tham gia: Client, API Gateway, Auth Service (Java), User Management Service (Java), RabbitMQ.sequenceDiagram
    actor Learner
    participant Client
    participant APIGW as API Gateway (Go)
    participant AuthSvc as Auth Service (Java)
    participant Broker as RabbitMQ
    participant UserSvc as User Mgmt Svc (Java)

    Learner ->> Client: 1. Điền và Gửi Form Đăng ký
    Client ->> APIGW: 2. POST /auth/register
    APIGW ->> AuthSvc: 3. POST /v1/register
    
    AuthSvc ->> AuthSvc: 4. Tạo entry xác thực (AuthN)
    AuthSvc -->> Broker: 5. Publish Event [UserRegistered]<br>(Gửi UserID, Email)
    
    AuthSvc ->> APIGW: 6. Trả về "Success 201 Created"
    APIGW ->> Client: 7. Trả về "Success"
    Client ->> Learner: 8. Thông báo Đăng ký Thành công
    
    activate UserSvc
    Broker -->> UserSvc: 9. Consume Event [UserRegistered]<br>(Async)
    UserSvc ->> UserSvc: 10. Tạo Hồ sơ Người dùng (Profile)<br>(Lưu PII theo ADR-7)
    deactivate UserSvc
2. Cung cấp Nội dung Thích ứng (Adaptive Content Delivery)Mô tả: Luồng này (UC-08) đã được mô tả trong Mục 3.2.3. Nó xảy ra khi người học yêu cầu bài học tiếp theo. Đây là một luồng đồng bộ (synchronous) phức tạp, đòi hỏi sự tương tác nhanh giữa Adaptive Engine và Learner Model Service (cả hai đều bằng Go, theo ADR-1).Các thành phần tham gia: Learner, Client, API Gateway, Adaptive Engine, Learner Model Service, Content Service.sequenceDiagram
    actor Learner
    participant Client as Browser/Client
    participant APIGW as API Gateway (Go)
    participant AdaptiveSvc as Adaptive Engine (Go)
    participant ModelSvc as Learner Model (Go)
    participant ContentSvc as Content Service (Java)

    Learner ->> Client: 1. Nhấn "Bài học Tiếp theo"
    Client ->> APIGW: 2. POST /api/adaptive/next<br>(Gửi JWT)
    
    APIGW ->> APIGW: 3. Xác thực JWT (ADR-6)
    APIGW ->> AdaptiveSvc: 4. POST /v1/next<br>(Gửi X-User-ID)

    AdaptiveSvc ->> ModelSvc: 5. GET /v1/model/{user-id}<br>(Query đồng bộ qua gRPC/REST)
    ModelSvc ->> AdaptiveSvc: 6. Trả về SkillMasteryScores (JSON)

    AdaptiveSvc ->> AdaptiveSvc: 7. Chạy thuật toán đề xuất<br>(Dựa trên scores)
    AdaptiveSvc ->> APIGW: 8. Trả về { contentId: "unit-123" }

    APIGW ->> Client: 9. Trả về { contentId: "unit-123" }

    Client ->> APIGW: 10. GET /api/content/unit-123<br>(Yêu cầu tải nội dung)
    APIGW ->> ContentSvc: 11. GET /v1/content/unit-123

    ContentSvc ->> APIGW: 12. Trả về ContentData (JSON)
    APIGW ->> Client: 13. Trả về ContentData

    Client ->> Learner: 14. Hiển thị nội dung bài học
3. Tạo Phản hồi Tức thì (Real-time Feedback Generation)Mô tả: Kịch bản này (một phần của UC-10) là luồng đồng bộ nhanh, tập trung vào việc cung cấp gợi ý (hints) ngay lập tức cho người học mà không cần nộp bài. Điều này yêu cầu hiệu năng cực cao (AC3: < 500ms), là lý do Scoring & Feedback Service được viết bằng Go (ADR-1).Các thành phần tham gia: Learner, Client, API Gateway, Scoring & Feedback Service.sequenceDiagram
    actor Learner
    participant Client
    participant APIGW as API Gateway (Go)
    participant ScoringSvc as Scoring Svc (Go)

    Learner ->> Client: 1. Đang làm bài tập, nhấn "Gợi ý" (Hint)
    Client ->> APIGW: 2. POST /api/scoring/hint<br>(Gửi questionId, currentState)
    
    APIGW ->> APIGW: 3. Xác thực JWT (ADR-6)
    APIGW ->> ScoringSvc: 4. POST /v1/hint

    ScoringSvc ->> ScoringSvc: 5. Chạy logic tạo gợi ý nhanh<br>(Rule-based/Simple AI)
    ScoringSvc ->> APIGW: 6. Trả về { hint: "Hãy kiểm tra lại..." }
    
    APIGW ->> Client: 7. Trả về JSON chứa Gợi ý
    Client ->> Learner: 8. Hiển thị Gợi ý cho Người học
4. Nộp và Chấm điểm Bài tập (Assessment Submission and Scoring)Mô tả: Đây là luồng (UC-10) quan trọng nhất, đã được mô tả trong Mục 3.2.3. Nó sử dụng mô hình lai (hybrid):Đồng bộ (Sync): Phản hồi điểm số cơ bản về client ngay lập tức (<500ms).Bất đồng bộ (Async): Publish sự kiện SubmissionCompleted lên RabbitMQ để Learner Model Service cập nhật kỹ năng (skill mastery) trong nền.Các thành phần tham gia: Learner, Client, API Gateway, Scoring Service, RabbitMQ, Learner Model Service.sequenceDiagram
    actor Learner
    participant Client
    participant APIGW as API Gateway
    participant ScoringSvc as Scoring Svc (Go)
    participant Broker as RabbitMQ
    participant ModelSvc as Learner Model Svc (Go)

    Learner ->> Client: 1. Nhấn "Nộp bài" (Submit)
    Client ->> APIGW: 2. POST /api/scoring/submit
    APIGW ->> ScoringSvc: 3. Gửi dữ liệu bài làm

    %% Sync Path (Tức thì)
    ScoringSvc ->> ScoringSvc: 4. Chấm điểm nhanh (Ví dụ: 8/10)
    ScoringSvc ->> APIGW: 5. Trả về phản hồi tức thì (<500ms)<br>{ score: 8, feedbackId: "f123" }
    APIGW ->> Client: 6. Hiển thị "Bạn đạt 8/10"
    Client ->> Learner: 7. Nhận điểm
    
    %% Async Path (Cập nhật nền)
    ScoringSvc -->> Broker: 8. Publish Event [SubmissionCompleted]<br>(Gửi kết quả chi tiết, UserID)
    
    activate ModelSvc
    Broker -->> ModelSvc: 9. Consume Event [SubmissionCompleted]
    ModelSvc ->> ModelSvc: 10. Cập nhật SkillMasteryScores<br>(Lưu vào MongoDB)
    deactivate ModelSvc
5. Tạo Báo cáo Giảng viên (Instructor Report Generation)Mô tả: Luồng này (UC-13/14) là một kịch bản "đọc" (read) phức tạp. Nó yêu cầu tổng hợp (orchestration) dữ liệu từ nhiều microservice khác nhau. Chúng ta giả định Content Service (Java, theo ADR-1) là dịch vụ thực hiện việc tổng hợp này, vì nó đòi hỏi logic nghiệp vụ phức tạp.Các thành phần tham gia: Instructor, Client, API Gateway, Content Service (Orchestrator), User Management Service, Learner Model Service.sequenceDiagram
    actor Instructor
    participant Client
    participant APIGW as API Gateway (Go)
    participant ContentSvc as Content Svc (Java)
    participant UserSvc as User Mgmt Svc (Java)
    participant ModelSvc as Learner Model Svc (Go)

    Instructor ->> Client: 1. Yêu cầu xem Báo cáo Lớp học
    Client ->> APIGW: 2. GET /api/reports/class/{classId}
    APIGW ->> ContentSvc: 3. GET /v1/reports/class/{classId}

    activate ContentSvc
    ContentSvc ->> UserSvc: 4. GET /v1/class/{classId}/students<br>(Lấy danh sách SV)
    UserSvc -->> ContentSvc: 5. Trả về [studentId1, studentId2, ...]

    ContentSvc ->> ModelSvc: 6. GET /v1/models/bulk?ids=...<br>(Lấy Skill Scores của cả lớp)
    ModelSvc -->> ContentSvc: 7. Trả về [SkillModel1, SkillModel2, ...]

    ContentSvc ->> ContentSvc: 8. Tổng hợp dữ liệu<br>(Tính toán điểm yếu chung, trung bình...)
    
    ContentSvc ->> APIGW: 9. Trả về Báo cáo (JSON)
    deactivate ContentSvc
    
    APIGW ->> Client: 10. Trả về Báo cáo (JSON)
    Client ->> Instructor: 11. Hiển thị Báo cáo Tổng hợp
