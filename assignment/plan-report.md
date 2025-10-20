# Nội dung cấu trúc báo cáo (phân theo các file *.tex) - Định dạng Markdown

## 1. File: `sections/01_overview.tex` (Mục 1 - Tổng quan)
| **Section**             | **Nội dung cần trình bày**                                                                                     | **Câu hỏi cốt lõi**                                                                                     |
|------------------------ |---------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------|
| 1 Overview              | Tổng quan về HMS, vai trò, chi tiết kỹ thuật (tính mô-đun, bảo mật).                                         | HMS là gì? Mục tiêu chính và những đặc điểm thiết kế quan trọng (VD: mô-đun, bảo mật) là gì?            |
| 1.1 System details      | Mô tả cách HMS tối ưu/quy trình (đăng ký, hồ sơ, thanh toán, ...).                                            | Lợi ích cốt lõi của thiết kế mô-đun, các biện pháp bảo mật?                                             |
| 1.2 Project Objectives  | Mục tiêu thiết kế và triển khai HMS tuân thủ nguyên tắc SOLID.                                                | Kiến trúc phần mềm HMS tuân thủ nguyên tắc nào? Đặc tính gì đạt được?                                   |
| 1.3 Project Scopes      | Xác định phạm vi (Bệnh nhân, Lịch hẹn, Thanh toán, Nhân sự).                                                  | 4 lĩnh vực chức năng cốt lõi trong HMS là gì?                                                           |
| 1.4 System requirements | Phân tích yêu cầu chức năng & phi chức năng (NFRs).                                                           | NFRs then chốt nào ảnh hưởng đến thiết kế kiến trúc (Scalability, Security,...)?                        |
| 1.5 Usecase diagram     | Sơ đồ Use Case tổng thể và chi tiết từng nghiệp vụ.                                                          | Use case mô tả quan hệ các Actor - chức năng thế nào?                                                   |

---

## 2. File: `sections/02_architecture_selection.tex` (Mục 2 - Lựa chọn Kiến trúc)
| **Section**                   | **Nội dung cần trình bày**                                                                        | **Câu hỏi cốt lõi**                                                         |
|------------------------------ |--------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------|
| 2 Comparison and selection    | Quy trình chọn kiến trúc tối ưu dựa trên Architectural Characteristics (ACs).                    | Phong cách kiến trúc nào, vì sao tối ưu (Layered, Microservices, SBA)?        |
| 2.1 Identifying ACs           | Trích xuất ACs (tường minh, ngầm định); chọn Top-3 ACs (Mô-đun, Mở rộng, Chịu lỗi).              | ACs ưu tiên hàng đầu là gì, ảnh hưởng tới sự sống lâu dài của HMS thế nào?    |
| 2.2 Considered Arch Styles    | So sánh Layered, Microservices, SBA dựa trên Top-3 ACs; Loại trừ dựa trên nhược điểm cụ thể.     | Vì sao loại Layered/Microservices?                                            |
| 2.3 Comparison and Selection  | Đánh giá (API Gateway vs. BFF), nêu quyết định thiết kế cuối cùng.                               | API Gateway giải quyết vấn đề gì? Vì sao chọn SBA + API Gateway?              |

---

## 3. File: `sections/03_architecture_design.tex` (Mục 3 - Thiết kế kiến trúc)
| **Section**                   | **Nội dung cần trình bày**                                                                      | **Câu hỏi cốt lõi**                                                                   |
|------------------------------ |--------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------|
| 3 Create the software architecture | Quy trình thiết kế, phân tách domain, component.                                          | Kiến trúc cuối cùng gồm các Domain/Service nào?                                      |
| 3.1 Component Identification Flow  | Quy trình xác định & tái cấu trúc component, thực hiện SRP.                             | Việc tách/bổ sung component thực thi SRP & giải quyết NFR như thế nào?               |
| 3.4 Final Architecture            | Chia thành 4 Domain chính (Patient, Appointment, Billing, Doctor), dịch vụ con đặc thù.  | Mỗi Domain có dịch vụ gì, chức năng gì?                                              |
| 3.5 Component Interactions        | Mô tả luồng giao tiếp: intra-domain/inter-domain/cross-cutting (API Gateway, Auth,...).   | Giao tiếp dịch vụ (đồng bộ/bất đồng bộ) hoàn thành nghiệp vụ như thế nào?            |

---

## 4. File: `sections/04_technologies.tex` (Mục 4 - Công nghệ sử dụng)
| **Section**          | **Nội dung cần trình bày**                                                                              | **Câu hỏi cốt lõi**                                               |
|--------------------- |--------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------|
| 4 Technologies       | Liệt kê, giải thích vai trò công nghệ cốt lõi.                                                         | Giải thích công nghệ hỗ trợ các ACs.                              |
| 4.4 Eureka           | Service Discovery: Service Registry - hỗ trợ Horizontal Scalability.                                   | Eureka giải quyết vấn đề nào trong kiến trúc SBA/Microservices?   |
| 4.6 Rabbit MQ        | Message Broker: giao tiếp bất đồng bộ (AMQP), tăng Concurrency & Reliability; cơ chế DLQ xử lý lỗi.    | RabbitMQ giải quyết NFR nào, cơ chế DLQ dùng ra sao?              |
| 4.1-4.3,4.5          | Vai trò Spring Boot (Framework), JWT (Authentication), MySQL (DB).                                    | JWT đảm bảo bảo mật, xác thực stateless như thế nào?               |

---

## 5. File: `sections/05_uml_diagrams.tex` (Mục 5 - UML Diagrams)
| **Section**              | **Nội dung cần trình bày**                                                         | **Câu hỏi cốt lõi**                                                                                 |
|------------------------  |-----------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------|
| 5 Create UML Diagrams    | Chứng minh cấu trúc/phân lớp qua sơ đồ UML.                                        | Các sơ đồ UML nào dùng mô tả kiến trúc?                                                             |
| 5.1 UML Components       | Sơ đồ phụ thuộc cấp cao: Service Domain và Cross-Cutting Components (API Gateway). | Mối quan hệ Service Domain và thành phần xuyên suốt thể hiện thế nào?                              |
| 5.2 UML Class Diagrams   | Sơ đồ lớp chi tiết cho core services (Billing, Patient, Doctor, API Gateway).      | Billing Service sử dụng OCP, DIP (Payment Gateway, Adapter) minh họa ra sao?                       |

---

## 6. File: `sections/06_solid_principles.tex` (Mục 6 - SOLID Principles)
| **Section**    | **Nội dung cần trình bày**                                                        | **Câu hỏi cốt lõi**                                                                    |
|--------------- |----------------------------------------------------------------------------------|----------------------------------------------------------------------------------------|
| 6 Apply SOLID  | Chứng minh thiết kế tuân thủ SOLID, liên hệ từng nguyên tắc.                      | Thiết kế tuân thủ 5 nguyên tắc SOLID như thế nào?                                      |
| 6.1 SRP        | Phân tách ranh giới nghiệp vụ rõ ràng (tách AuthService, MedicalService, v.v.).   | Các service tách rời đảm bảo mỗi module có lý do thay đổi duy nhất ra sao?             |
| 6.2 OCP        | Interface-based & Strategy Pattern (vd: Payment Adapter trong Billing Service).   | Thêm mới cổng thanh toán có cần sửa code cũ không (minh họa OCP)?                      |
| 6.5 DIP        | Module cấp cao phụ thuộc Abstraction (Interface); dùng Dependency Injection (DI). | DI dùng đảm bảo linh hoạt, tăng khả năng kiểm thử module như thế nào (DIP)?             |

---

## 8. File: `sections/08_future_extensibility.tex` (Mục 8 - Mở rộng tương lai)
| **Section**     | **Nội dung cần trình bày**                                                       | **Câu hỏi cốt lõi**                                                                             |
|---------------- |---------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------|
| 8 Future Extensibility   | Mô hình quản lý thay đổi để tích hợp chức năng mới (vd: quản lý xét nghiệm).   | Quy trình chuẩn để thêm chức năng mới là gì?                                                    |
| 8.1 Analyzing the Domain | Phân tích dựa trên: Độc lập Nghiệp vụ, Độ kết dính, Mở rộng riêng biệt.        | Khi nào tạo Module độc lập mới vs mở rộng Module hiện có?                                       |
| 8.2 Evaluating Using SOLID | Checklist SOLID kiểm tra chéo thiết kế tính năng mới.                        | Nếu vi phạm OCP (phải sửa code cũ), nên áp dụng mẫu thiết kế nào cho mở rộng?                   |
| 8.3 Identifying Abstractions| Xác định abstraction (Interface, DTO, Handler) để giảm coupling.            | Abstraction giúp tạo điểm mở rộng, giảm sửa đổi code cũ như thế nào?                            |
| 8.5 Using DI              | Đảm bảo dùng DI để giảm phụ thuộc cứng giữa module mới và cũ.                 | DI quan trọng ra sao với khả năng kiểm thử/mở rộng?                                             |