# Requirements Document

## Introduction

Dự án này là một Hệ thống Gia sư Thông minh (Intelligent Tutoring System - ITS) được phát triển cho môn học CO3017 - Kỹ Thuật Phần Mềm tại HCMUT. Hiện tại, repository chứa đầy đủ source code, báo cáo LaTeX, và tài liệu phân tích kiến trúc. Tuy nhiên, các file README từ root đến từng service đều thiếu nội dung hoặc không có ý nghĩa, gây khó khăn cho việc hiểu và sử dụng hệ thống.

Mục tiêu của spec này là tạo ra một bộ tài liệu README hoàn chỉnh, chuyên nghiệp, và có ý nghĩa bằng tiếng Việt, giúp người đọc (giảng viên, sinh viên, hoặc developer khác) có thể:

- Hiểu rõ mục đích và kiến trúc của hệ thống
- Cài đặt và chạy hệ thống một cách dễ dàng
- Nắm được cấu trúc và vai trò của từng component
- Tham khảo API và tài liệu kỹ thuật chi tiết

## Glossary

- **ITS (Intelligent Tutoring System)**: Hệ thống Gia sư Thông minh - ứng dụng sử dụng AI để cung cấp trải nghiệm học tập cá nhân hóa
- **Repository**: Kho lưu trữ mã nguồn và tài liệu trên GitHub
- **README**: File tài liệu markdown mô tả về một thư mục hoặc dự án
- **Service**: Một microservice trong kiến trúc hệ thống
- **Root README**: File README.md ở thư mục gốc của repository
- **Content Service**: Service quản lý câu hỏi và tài liệu học tập
- **Scoring Service**: Service đánh giá câu trả lời của học sinh
- **Learner Model Service**: Service theo dõi mức độ thành thạo của học sinh
- **Adaptive Engine**: Service đề xuất nội dung học tập cá nhân hóa
- **Client**: Ứng dụng web frontend cho học sinh
- **Docker**: Nền tảng containerization để triển khai ứng dụng
- **Microservices Architecture**: Kiến trúc phần mềm chia hệ thống thành các service độc lập

## Requirements

### Requirement 1

**User Story:** Là một giảng viên hoặc người đánh giá, tôi muốn đọc README tổng quan ở root để hiểu toàn bộ dự án, mục đích, và cấu trúc repository, để có thể nhanh chóng nắm bắt được nội dung bài nộp.

#### Acceptance Criteria

1. WHEN a reader opens the root README.md THEN the System SHALL display a comprehensive overview in Vietnamese including project purpose, course context, and repository structure
2. WHEN a reader views the root README THEN the System SHALL present a clear table of contents with links to major sections
3. WHEN a reader needs to understand the architecture THEN the System SHALL provide a high-level architecture diagram with Vietnamese labels
4. WHEN a reader wants to run the system THEN the System SHALL include quick start instructions with Docker commands
5. WHEN a reader explores the repository structure THEN the System SHALL document all major directories with their purposes in Vietnamese

### Requirement 2

**User Story:** Là một developer, tôi muốn đọc README của thư mục sources để hiểu kiến trúc microservices và cách chạy toàn bộ hệ thống, để có thể setup môi trường development.

#### Acceptance Criteria

1. WHEN a developer opens sources/README.md THEN the System SHALL explain the microservices architecture with service dependencies
2. WHEN a developer needs to start services THEN the System SHALL provide step-by-step Docker Compose instructions
3. WHEN a developer wants to verify the system THEN the System SHALL include health check commands for all services
4. WHEN a developer encounters issues THEN the System SHALL document common problems and solutions
5. WHEN a developer needs technical details THEN the System SHALL list all services with ports, databases, and technologies

### Requirement 3

**User Story:** Là một developer, tôi muốn đọc README của từng service (Content, Scoring, Learner Model, Adaptive Engine, Client) để hiểu chức năng, API, và cách chạy service đó độc lập.

#### Acceptance Criteria

1. WHEN a developer opens a service README THEN the System SHALL describe the service purpose and responsibilities in Vietnamese
2. WHEN a developer needs to run a service locally THEN the System SHALL provide local development setup instructions
3. WHEN a developer wants to test APIs THEN the System SHALL document all API endpoints with curl examples
4. WHEN a developer needs configuration details THEN the System SHALL list all environment variables with descriptions
5. WHEN a developer explores the codebase THEN the System SHALL explain the internal architecture and key components

### Requirement 4

**User Story:** Là một sinh viên hoặc người đọc, tôi muốn các README được viết bằng tiếng Việt chuyên nghiệp, để dễ hiểu và phù hợp với bối cảnh môn học tại HCMUT.

#### Acceptance Criteria

1. WHEN a reader views any README THEN the System SHALL use professional Vietnamese technical terminology
2. WHEN a reader encounters technical terms THEN the System SHALL provide Vietnamese translations with English terms in parentheses
3. WHEN a reader follows instructions THEN the System SHALL use clear, imperative Vietnamese commands
4. WHEN a reader needs context THEN the System SHALL reference the course assignment requirements appropriately
5. WHEN a reader views code examples THEN the System SHALL include Vietnamese comments explaining key concepts

### Requirement 5

**User Story:** Là một người đánh giá, tôi muốn README tổng quan liên kết rõ ràng với các artifacts khác (báo cáo LaTeX, markdown analysis, presentation), để có thể điều hướng toàn bộ submission.

#### Acceptance Criteria

1. WHEN a reviewer opens the root README THEN the System SHALL include a section linking to the LaTeX report with build instructions
2. WHEN a reviewer needs analysis documents THEN the System SHALL provide links to markdown analysis files with descriptions
3. WHEN a reviewer wants to view diagrams THEN the System SHALL reference the diagrams directory with explanations
4. WHEN a reviewer explores the presentation THEN the System SHALL link to presentation materials with viewing instructions
5. WHEN a reviewer needs assignment context THEN the System SHALL reference markdown/assignment.md with a summary

### Requirement 6

**User Story:** Là một developer, tôi muốn README của sources/ bao gồm hướng dẫn testing và troubleshooting, để có thể kiểm tra và debug hệ thống hiệu quả.

#### Acceptance Criteria

1. WHEN a developer needs to test the system THEN the System SHALL document integration test procedures
2. WHEN a developer encounters errors THEN the System SHALL provide a troubleshooting section with common issues
3. WHEN a developer wants to verify data THEN the System SHALL explain how to access databases and message queues
4. WHEN a developer needs logs THEN the System SHALL document how to view and interpret service logs
5. WHEN a developer runs tests THEN the System SHALL explain the test data setup and cleanup procedures

### Requirement 7

**User Story:** Là một người đọc, tôi muốn mỗi README có cấu trúc nhất quán và dễ scan, để có thể nhanh chóng tìm thông tin cần thiết.

#### Acceptance Criteria

1. WHEN a reader opens any README THEN the System SHALL follow a consistent structure with standard sections
2. WHEN a reader scans a README THEN the System SHALL use clear headings and subheadings in Vietnamese
3. WHEN a reader needs quick reference THEN the System SHALL include tables for configuration and API documentation
4. WHEN a reader views code blocks THEN the System SHALL use proper markdown syntax highlighting
5. WHEN a reader navigates sections THEN the System SHALL provide a table of contents for READMEs longer than 200 lines

### Requirement 8

**User Story:** Là một giảng viên, tôi muốn README root giải thích rõ cách dự án đáp ứng yêu cầu assignment, để có thể đánh giá submission một cách chính xác.

#### Acceptance Criteria

1. WHEN a reviewer reads the root README THEN the System SHALL include a section mapping implementation to assignment requirements
2. WHEN a reviewer evaluates architecture THEN the System SHALL reference where architecture decisions are documented
3. WHEN a reviewer checks SOLID principles THEN the System SHALL link to SOLID principles documentation with examples
4. WHEN a reviewer assesses implementation THEN the System SHALL list which core modules have been implemented
5. WHEN a reviewer needs evidence THEN the System SHALL provide links to relevant code, diagrams, and reports

### Requirement 9

**User Story:** Là một developer mới, tôi muốn README có prerequisites và system requirements rõ ràng, để biết cần cài đặt gì trước khi bắt đầu.

#### Acceptance Criteria

1. WHEN a new developer starts THEN the System SHALL list all required software with minimum versions
2. WHEN a developer checks compatibility THEN the System SHALL specify supported operating systems
3. WHEN a developer needs tools THEN the System SHALL document required development tools with installation links
4. WHEN a developer prepares environment THEN the System SHALL explain hardware requirements for running Docker
5. WHEN a developer verifies setup THEN the System SHALL provide commands to check installed prerequisites

### Requirement 10

**User Story:** Là một contributor, tôi muốn README của mỗi service giải thích project structure và coding conventions, để có thể đóng góp code một cách nhất quán.

#### Acceptance Criteria

1. WHEN a contributor explores a service THEN the System SHALL document the directory structure with explanations
2. WHEN a contributor writes code THEN the System SHALL reference coding style guidelines
3. WHEN a contributor adds features THEN the System SHALL explain the layered architecture pattern used
4. WHEN a contributor needs examples THEN the System SHALL point to representative code files
5. WHEN a contributor submits changes THEN the System SHALL document the testing requirements
