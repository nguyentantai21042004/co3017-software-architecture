# Chương 1: Thiết kế Kiến trúc Hệ thống

> Đây là chương **quan trọng nhất**, chiếm phần lớn báo cáo, đi từ **phân tích yêu cầu** → **thiết kế chi tiết**.

---

## 1. Phân tích Yêu cầu (Pages 5-11)

### Mô tả cái gì
Phần này **định nghĩa "vấn đề"** một cách chi tiết. Nó bao gồm:

### Nội dung chi tiết

**Phạm vi Dự án (Project Scope)**
- Các module chính sẽ làm: Quản lý Bệnh nhân, Đặt lịch hẹn, ...
- Các yếu tố loại trừ: Không tích hợp phần cứng y tế

**Stakeholders (Các bên liên quan)**
- Quản trị viên
- Bác sĩ
- Bệnh nhân

**Yêu cầu Chức năng (Functional Requirements)**
- Các user story cụ thể
- Ví dụ: *"Là một bệnh nhân, tôi muốn..."*

**Yêu cầu Phi Chức năng (Non-Functional Requirements)**
- Hiệu suất: Phản hồi < 1 giây
- Độ tin cậy: Uptime 99.9%
- Bảo mật: Mã hóa AES-256
- V.v.

### Câu hỏi chính
- Hệ thống phải **làm gì** (chức năng)?
- Hệ thống phải **làm tốt như thế nào** (phi chức năng)?
- **Ai sẽ sử dụng** nó?

### Ý nghĩa
Đây là **nền tảng của toàn bộ thiết kế**. Mọi quyết định kiến trúc sau này đều phải được đưa ra để **đáp ứng các yêu cầu này**.

### Tổng nội dung
- Định nghĩa rõ ràng những gì hệ thống cần làm
- Các mục tiêu hiệu suất cụ thể:
  - Hỗ trợ **10.000 bệnh nhân**
  - **100.000 giao dịch/ngày**
- Các yêu cầu bảo mật

---

## 2. Chọn Kiểu Kiến trúc & Mẫu Kiến trúc (Pages 12-18)

### Mô tả cái gì
Đây là phần **"lập luận"** cho quyết định kiến trúc.

### Nội dung chi tiết

**So sánh các kiểu kiến trúc**
- So sánh: **Microservices** vs **Monolithic** vs **Layered**
- Dựa trên các yêu cầu phi chức năng đã xác định
- Tập trung vào:
  - Khả năng mở rộng (Scalability)
  - Tính sẵn sàng (Availability)
  - Khả năng bảo trì (Maintainability)

**Kết luận**
- Biện minh **tại sao Microservices** là lựa chọn vượt trội

**Mẫu kiến trúc**
- So sánh và chọn mẫu **API Gateway** để quản lý giao tiếp

### Câu hỏi chính
- Với các yêu cầu đã đặt ra, **cấu trúc cấp cao nào** (kiểu kiến trúc) là tốt nhất cho hệ thống này?
- **Tại sao?**

### Ý nghĩa
Đây là **quyết định thiết kế trung tâm** của báo cáo. Nó **kết nối**:
- **"Vấn đề"** (yêu cầu)  →  **"Giải pháp"** (kiến trúc)

### Tổng nội dung
- Lập luận rằng kiến trúc **Monolithic**:
  - **KHÔNG** thể đáp ứng yêu cầu **99.9% uptime**
  - **KHÔNG** thể đáp ứng yêu cầu **khả năng mở rộng**
- Do đó, kiến trúc **Microservices** là lựa chọn lý tưởng
- Chọn **API Gateway** làm mẫu điều phối trung tâm

---

## 3. Thiết kế Kiến trúc Hệ thống (Pages 19-25)

### Mô tả cái gì
Chuyển từ **"tại sao"** (lựa chọn) sang **"cái gì"** (thiết kế).

### Nội dung chi tiết

**Xác định Actors & Actions**
- **Actors (Diễn viên):** Bệnh nhân, Bác sĩ, ...
- **Actions (Hành động):** Đặt lịch hẹn, Thêm hồ sơ, ...

**Nhóm thành Components**
- Nhóm các actions thành các microservices
- Ví dụ:
  - AppointmentScheduling Service
  - PatientManagement Service
  - ...

**Trực quan hóa thiết kế**
- Sơ đồ UML:
  - Package Diagram
  - Component Diagram
  - Deployment Diagram

### Câu hỏi chính
- Kiến trúc Microservices đã chọn sẽ được **cấu trúc cụ thể** như thế nào?
- Có những **service nào**?
- Chúng **tương tác và được triển khai** ra sao?

### Ý nghĩa
**Cung cấp bản thiết kế (blueprint) kỹ thuật cấp cao** của hệ thống.

### Tổng nội dung
- Bản thiết kế phác thảo hệ thống
- Xác định **8 service chính:**
  - Patient
  - Staff
  - Billing
  - Authentication
  - ...
- Sơ đồ UML cho thấy:
  - Các service được **đóng gói** (Package Diagram)
  - Các service **tương tác với nhau** qua API Gateway (Component Diagram)
  - Được **triển khai** trong Docker container trên Ubuntu (Deployment Diagram)

---

## 4. Lược đồ Lớp UML - Chi tiết (Pages 26-55)

### Mô tả cái gì
Một phần **rất chi tiết**, đi sâu vào **bên trong từng microservice** đã xác định ở phần 3.

### Nội dung chi tiết

**Cho mỗi service:**
- API Gateway
- Authentication
- Patient
- Staff
- ...

**Trình bày:**
1. **Sơ đồ Lớp (Class Diagram)** chi tiết
2. **Phân tích SOLID** - tuân thủ 5 nguyên tắc:
   - **Single Responsibility Principle (SRP)**
   - **Open/Closed Principle (OCP)**
   - **Liskov Substitution Principle (LSP)**
   - **Interface Segregation Principle (ISP)**
   - **Dependency Inversion Principle (DIP)**
3. **Thảo luận** về "Khả năng mở rộng" (Extensibility)

### Câu hỏi chính
- Cấu trúc **mã nguồn bên trong** của mỗi microservice là gì?
- Làm thế nào thiết kế nội bộ **đảm bảo các nguyên tắc** kỹ thuật phần mềm tốt (SOLID)?

### Ý nghĩa
**Thể hiện sự hiểu biết sâu sắc** về:
- Thiết kế hướng đối tượng
- **Liên kết** kiến trúc vĩ mô (Microservices) ↔️ thiết kế vi mô (SOLID, class design)

### Tổng nội dung - Ví dụ Patient Service

**Các lớp:**
- `PatientController`
- `PatientService` (interface)
- `PatientServiceImpl`
- `PatientRepository`

**Phân tích SOLID:**
- **Single Responsibility**: Controller chỉ xử lý HTTP
- **Dependency Inversion**: Controller phụ thuộc vào interface `PatientService`
  - Không phụ thuộc vào class `PatientServiceImpl`

---

## Tóm tắt Cấu trúc Chương 1

| Phần | Trang | Tập trung | Output |
|------|-------|----------|--------|
| **1. Phân tích Yêu cầu** | 5-11 | Định nghĩa "vấn đề" | User Stories + NFR |
| **2. Chọn Kiến trúc** | 12-18 | So sánh & lập luận | Quyết định: Microservices |
| **3. Thiết kế Hệ thống** | 19-25 | Xác định service & diagram | UML diagrams (Package, Component, Deployment) |
| **4. Lược đồ Lớp** | 26-55 | Chi tiết từng service | Class Diagram + SOLID analysis |

---

## Nhận xét chung

**Cấu trúc logic dạng Funnel:**
```
Phần 1: Vấn đề (cái gì & tại sao)
   ↓
Phần 2: Lựa chọn kiến trúc (tại sao lựa chọn này)
   ↓
Phần 3: Thiết kế tổng quan (cái gì & cách nó hoạt động)
   ↓
Phần 4: Chi tiết triển khai (làm thế nào chi tiết)
```

**Mỗi phần đều trả lời:**
- **Câu hỏi chính**
- **Ý nghĩa**
- **Tổng nội dung cụ thể**
