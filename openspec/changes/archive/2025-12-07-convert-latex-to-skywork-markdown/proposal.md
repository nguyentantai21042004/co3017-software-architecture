# Change: Convert LaTeX Report to Skywork AI-Optimized Markdown

## Why

Báo cáo LaTeX hiện tại (`report/main.tex` và các file trong `report/contents/`) cần được chuyển đổi thành định dạng Markdown tối ưu để sử dụng với **Skywork AI** nhằm tự động tạo slide thuyết trình chuyên nghiệp. Việc chuyển đổi này giúp:

- Tạo slide presentation tự động từ nội dung báo cáo
- Tận dụng hệ thống Tags đặc biệt để điều hướng hình ảnh và biểu đồ
- Cô đọng nội dung học thuật thành văn phong thuyết trình dễ hiểu

## What Changes

### Tạo mới

- **`report/presentation/slides.md`**: File Markdown chính chứa toàn bộ nội dung slide
- **`report/presentation/README.md`**: Hướng dẫn sử dụng file Markdown với Skywork AI

### Quy trình chuyển đổi

1. **Structure Mapping**: Chuyển đổi cấu trúc LaTeX sang Markdown headers

   - `\title{...}` → `# [SLIDE DECK]: ...`
   - `\section{...}` → `## ...` (Slide Divider)
   - `\subsection{...}` → `### ...` (Content Slides)

2. **Content Condensation**: Cô đọng nội dung

   - Tóm tắt đoạn văn thành bullet points
   - Quy tắc 2 câu cho mỗi bullet point
   - Loại bỏ công thức phức tạp, giữ ý nghĩa

3. **Visual Placeholders**: Chèn Tags cho hình ảnh
   - `[PLACEHOLDER: ...]` cho ảnh minh họa
   - `[FULL_SLIDE_IMAGE_PLACEHOLDER: ...]` cho ảnh nền/hero
   - `[CHART_TYPE: ...]` cho biểu đồ/đồ thị

## Impact

### Affected Files

- **Input**: Tất cả file `.tex` trong `report/contents/`
- **Output**: Tạo mới `report/presentation/` directory

### Affected Specs

- `architecture-report-documentation` - Mở rộng khả năng xuất báo cáo

### Dependencies

- Không có breaking changes
- Không ảnh hưởng đến source code của hệ thống ITS
- Chỉ tạo thêm artifact mới từ báo cáo hiện có

## Scope

### Các section cần chuyển đổi (theo thứ tự trong main.tex)

1. Executive Summary
2. Project Scope and Objectives
3. Stakeholder Analysis
4. Functional Requirements
5. Non-Functional Requirements
6. Architecture Characteristics Prioritization
7. Architecture Style Selection
8. Architecture Decision Records
9. Design Principles
10. Module View
11. Component & Connector View
12. Allocation View
13. Behavior View
14. SOLID Principles Application
15. System Implementation
16. Reflection and Evaluation

### Hình ảnh cần xử lý

- Diagrams từ `report/images/` sẽ được tham chiếu hoặc thay thế bằng placeholders phù hợp
