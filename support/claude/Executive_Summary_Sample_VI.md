# TỔNG QUAN DỰ ÁN
## Hệ Thống Hướng Dẫn Thông Minh (ITS) - Thiết Kế Kiến Trúc Phần Mềm

---

## Tổng Quan Dự Án

**Tầm Nhìn:** Xây dựng một hệ thống học tập thông minh có khả năng cá nhân hóa trải nghiệm học tập cho từng người dùng, giúp mọi học viên đều có thể tiếp cận giáo dục chất lượng theo cách phù hợp nhất với năng lực và tiến độ riêng, đồng thời hệ thống phải mở rộng và phục vụ hàng nghìn người dùng song song.

**Những Thách Thức Chính:**
- Đảm bảo phục vụ trên 5.000 người dùng đồng thời với thời gian phản hồi dưới 500ms
- Duy trì cân bằng giữa hiệu năng kỹ thuật và tốc độ triển khai
- Mở rộng theo chiều ngang mà vẫn giữ được hiệu quả học tập và chất lượng AI
- Hỗ trợ thử nghiệm và cập nhật mô hình AI mà không ảnh hưởng đến hệ thống vận hành

**Phương Pháp Kiến Trúc Được Chọn:**  
Áp dụng kiến trúc microservices lai (Hybrid Microservices) kết hợp với mô hình hướng sự kiện (Event-Driven Architecture) nhằm đạt được tính mô-đun, khả năng mở rộng, hiệu suất cao và khả năng kiểm thử toàn diện.

---

## Các Quyết Định Kiến Trúc Chính

1. **Microservices Lai Ghép + Event-Driven Architecture**  
   Phân hủy thành 5 dịch vụ (Quản Lý Người Dùng, Nội Dung, Mô Hình Học Viên, Công Cụ Thích Ứng, Đánh Giá) kết hợp xử lý sự kiện không đồng bộ cho phân tích thời gian thực.

2. **Đa Ngôn Ngữ Lập Trình (Polyglot Programming):**
   - **Java/Spring Boot:** Xử lý logic nghiệp vụ cốt lõi.
   - **Golang:** Phục vụ các tác vụ tính toán hiệu năng cao, đặc biệt cho AI/ML.
   - **PostgreSQL:** Lưu trữ dữ liệu.
   - **RabbitMQ:** Truyền phát và xử lý sự kiện.

3. **Kiến Trúc Sạch (Clean Architecture):**  
   Tách biệt rõ ràng giữa tầng nghiệp vụ và tầng hạ tầng, giúp dễ bảo trì, kiểm thử và mở rộng. Đảm bảo độ bao phủ kiểm thử lớn hơn 85%.

4. **Điều Phối Bằng Kubernetes:**  
   Hỗ trợ tự động mở rộng, triển khai xanh–xanh (blue-green) để tránh downtime, cùng cơ chế tự phục hồi pod nhằm tối ưu hóa hiệu suất hệ thống.

5. **Khả Quan Sát (Observability-first):**  

---

## Kết Quả Dự Kiến

**Hiệu Suất & Khả Năng Mở Rộng:**
- Hỗ trợ 5.000 trở lên người dùng đồng thời (mở rộng tới 9 000)
- Thời gian phản ứng nhỏ hơn năm trăm mili giây (phân vị thứ chín mươi lăm) dưới tải cao
- Tính khả dụng chín mươi chín phẩy chín phần trăm (ít hơn 9 giờ ngừng hoạt động mỗi năm)

**Sự Xuất Sắc Kỹ Thuật:**
- Tính mô đun: Phát triển và triển khai dịch vụ độc lập
- Khả năng phục hồi: Ngắt mạch và cơ chế dự phòng ngăn chặn lỗi xếp tầng
- Tính linh hoạt: Hỗ trợ kiểm tra A/B mô hình AI mà không thay đổi hệ thống
- Khả năng bảo trì: Ranh giới rõ ràng giảm tải nhận thức cho nhà phát triển
