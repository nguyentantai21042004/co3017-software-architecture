# Change: Add Marp Presentation Slides for ITS Project

## Why

Dự án cần một bài thuyết trình (presentation) cho buổi báo cáo cuối kỳ môn CO3017 - Software Architecture. Hiện tại đã có file `GAMMA_PROMPT.md` định nghĩa cấu trúc slides nhưng chưa có folder và nội dung slides thực tế. Cần tạo presentation sử dụng Marp (Markdown-based presentation) để dễ dàng maintain và version control.

## What Changes

- Tạo folder `presentation/` tại root của repository
- Tạo file `presentation/slides.md` - nội dung chính của presentation (Marp format)
- Tạo file `presentation/theme.css` - custom theme cho slides
- Tạo file `presentation/README.md` - hướng dẫn sử dụng và build slides
- Di chuyển `GAMMA_PROMPT.md` vào folder `presentation/` (nếu cần)

## Cấu Trúc Slides Dự Kiến

Dựa trên nội dung report và `GAMMA_PROMPT.md`:

| Section                   | Số slides | Nội dung                                             |
| ------------------------- | --------- | ---------------------------------------------------- |
| Title + Executive Summary | 3-4       | Giới thiệu dự án, team, mục tiêu                     |
| Requirements Analysis     | 4-5       | Stakeholders, Functional/Non-functional requirements |
| Architecture Design       | 4-5       | Microservices, Event-Driven, Clean Architecture      |
| ADRs                      | 2-3       | Architecture Decision Records quan trọng             |
| Architecture Views        | 3-4       | Module, C&C, Allocation, Behavior views              |
| SOLID Principles          | 2-3       | Áp dụng SOLID trong dự án                            |
| Implementation            | 2-3       | Tech stack, cấu trúc code                            |
| Demo & Code Showcase      | 4-6       | API endpoints, Docker setup, code examples           |
| Reflection & Future       | 2-3       | Lessons learned, future improvements                 |
| Conclusion & Q&A          | 2         | Tổng kết, Q&A                                        |
| **Tổng cộng**             | **28-38** |                                                      |

## Impact

- Affected specs: Không ảnh hưởng specs hiện tại (capability mới)
- Affected code: Không ảnh hưởng source code
- New files:
  - `presentation/slides.md`
  - `presentation/theme.css`
  - `presentation/README.md`
  - `presentation/images/` (folder cho diagrams/screenshots)
