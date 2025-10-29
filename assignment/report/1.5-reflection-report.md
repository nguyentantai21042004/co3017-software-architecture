# 1.7 Báo Cáo Phản Ánh (Reflection Report)

## 1. SOLID đã cải thiện thiết kế như thế nào?

### SRP (Single Responsibility Principle)
Việc tách riêng **LearnerAggregate** (quản lý thông tin cá nhân của người học) khỏi **LearnerModelAggregate** (quản lý trạng thái và dữ liệu AI về tiến trình học tập) giúp đảm bảo mỗi service chỉ tập trung vào một trách nhiệm duy nhất. Khi cần hoán đổi hoặc nâng cấp mô hình AI (US8/FR12), chúng ta chỉ cần thay đổi tại LearnerModel Service mà không ảnh hưởng đến dịch vụ quản lý người dùng. Điều này giúp đơn giản hóa bảo trì và tận dụng khả năng mở rộng của microservices.

### DIP (Dependency Inversion Principle)
Trong thiết kế Clean Architecture (Sơ đồ 2), việc phụ thuộc vào Interface (ví dụ: `LearnerModelRepository`) giúp các thành phần nghiệp vụ như **AdaptivePathGenerator** không phụ thuộc trực tiếp vào lớp cài đặt cụ thể (Postgres, Redis, v.v). Nhờ đó, có thể dễ dàng thay thế hoặc mock tầng lưu trữ trong quá trình kiểm thử (Testability - AC4) mà không cần thay đổi logic nghiệp vụ. Điều này giúp hệ thống linh hoạt hơn và giảm sự phụ thuộc chặt chẽ giữa các lớp.

## 2. Những thách thức & đánh đổi trong thiết kế

### Thách thức (Challenges)
- **Quản lý độ phức tạp:** Kiến trúc microservices phức tạp vượt trội so với mô hình monolithic truyền thống, đòi hỏi nhiều công sức cho vận hành (vận hành hạ tầng, CI/CD, logging, monitoring, v.v).
- **Giao tiếp giữa các dịch vụ:** Việc phân tách các chức năng thành nhiều service khiến khả năng kiểm soát luồng dữ liệu/transaction trở nên khó khăn hơn, dễ phát sinh lỗi phân tán.

### Đánh đổi (Trade-offs)
- **Eventual Consistency:** Việc sử dụng Kafka (kiến trúc bất đồng bộ) để đạt yêu cầu mở rộng (AC2 - Scalability) đồng nghĩa với phải chấp nhận tính nhất quán cuối cùng thay vì nhất quán ngay lập tức. Ví dụ: Khi học sinh nộp bài, điểm số có thể mất vài giây mới cập nhật vào hệ thống (LearnerModel Service), không phản ánh ngay lập tức trên dashboard.
- **Tăng overhead vận hành:** Hệ thống nhiều dịch vụ đòi hỏi giải pháp quản lý trạng thái, đồng bộ dữ liệu, bảo mật liên dịch vụ, trong khi monolithic đơn giản hóa các vấn đề này nhưng lại thiếu tính linh hoạt khi phát triển mở rộng.

---

**Kết luận:** Việc bám sát các nguyên tắc SOLID, đặc biệt là SRP và DIP, đã giúp kiến trúc của hệ thống có tính modular cao, dễ mở rộng, bảo trì và kiểm thử. Dù phải đánh đổi một số yếu tố (độ phức tạp, consistency), thiết kế này đặt nền móng vững chắc để phát triển hệ thống ITS ở quy mô lớn và đáp ứng được các mục tiêu đã đề ra ở Phase 1.

Sau khi hoàn thiện bản phản ánh này, tài liệu cho Phase 1 (Thiết kế) đã đầy đủ và sẵn sàng tiếp bước sang Phase 2 (Implementation).