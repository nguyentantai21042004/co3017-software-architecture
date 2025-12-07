# Design Document: LaTeX to Skywork AI Markdown Conversion

## Context

### Background

Dự án ITS (Intelligent Tutoring System) có một báo cáo LaTeX hoàn chỉnh trong `report/` directory. Báo cáo này cần được chuyển đổi thành định dạng Markdown tối ưu để sử dụng với Skywork AI - một công cụ tạo slide tự động.

### Stakeholders

- **Nhóm sinh viên**: Cần slide để thuyết trình báo cáo
- **Giảng viên**: Đánh giá nội dung thuyết trình
- **Skywork AI**: Công cụ nhận Markdown và tạo slide

### Constraints

- Phải tuân thủ cú pháp Tags đặc biệt của Skywork AI
- Nội dung phải cô đọng, phù hợp với slide (không quá 2 câu/bullet)
- Giữ nguyên cấu trúc logic của báo cáo gốc

## Goals / Non-Goals

### Goals

- Tạo file Markdown có thể sử dụng trực tiếp với Skywork AI
- Cô đọng nội dung học thuật thành văn phong thuyết trình
- Tận dụng hệ thống Tags để tạo visual elements tự động
- Đảm bảo coverage đầy đủ các section quan trọng của báo cáo

### Non-Goals

- Không tạo slide trực tiếp (việc này do Skywork AI thực hiện)
- Không thay đổi nội dung báo cáo LaTeX gốc
- Không tạo hình ảnh mới (chỉ tạo placeholders)

## Decisions

### Decision 1: Cấu trúc Output Directory

**What**: Tạo `report/presentation/` directory riêng biệt
**Why**:

- Tách biệt với báo cáo LaTeX gốc
- Dễ quản lý và version control
- Có thể regenerate mà không ảnh hưởng source

### Decision 2: Single File Approach

**What**: Tất cả slides trong một file `slides.md`
**Why**:

- Skywork AI hoạt động tốt nhất với single input file
- Dễ dàng copy/paste vào tool
- Giữ nguyên flow của presentation

### Decision 3: Tag System

**What**: Sử dụng 3 loại Tags chính

```markdown
[PLACEHOLDER: description] # Ảnh minh họa thường
[FULL_SLIDE_IMAGE_PLACEHOLDER: ...] # Ảnh nền/hero slide
[CHART_TYPE: type, data] # Biểu đồ/đồ thị
```

**Why**: Đây là cú pháp được định nghĩa trong term.md, tối ưu cho Skywork AI

### Decision 4: Content Condensation Rules

**What**: Áp dụng quy tắc cô đọng nội dung

- Mỗi bullet point ≤ 2 câu
- Tập trung vào keywords, kết quả, con số
- Loại bỏ công thức toán học phức tạp
  **Why**: Slide cần ngắn gọn, dễ đọc trong thời gian ngắn

### Alternatives Considered

| Option                 | Pros                | Cons                         | Decision |
| ---------------------- | ------------------- | ---------------------------- | -------- |
| Multiple MD files      | Modular, dễ edit    | Khó sử dụng với Skywork AI   | Rejected |
| Include LaTeX formulas | Giữ nguyên chi tiết | Skywork AI không render được | Rejected |
| Auto-convert images    | Tự động hóa         | Phức tạp, cần thêm tools     | Rejected |

## Mapping Rules

### Structure Mapping (LaTeX → Markdown)

| LaTeX                 | Markdown              | Slide Type            |
| --------------------- | --------------------- | --------------------- |
| `\title{...}`         | `# [SLIDE DECK]: ...` | Title Slide           |
| `\section{...}`       | `## ...`              | Section Divider       |
| `\subsection{...}`    | `### ...`             | Content Slide         |
| `\subsubsection{...}` | `#### ...`            | Sub-content (nếu cần) |

### Content Mapping

| LaTeX Element       | Markdown Output                         |
| ------------------- | --------------------------------------- |
| Paragraph           | Bullet points (tóm tắt)                 |
| `\begin{itemize}`   | `* item`                                |
| `\begin{enumerate}` | `1. item`                               |
| `\begin{table}`     | `[CHART_TYPE: ...]` hoặc Markdown table |
| `\begin{figure}`    | `[PLACEHOLDER: ...]`                    |
| `\includegraphics`  | `[PLACEHOLDER: ...]` với mô tả          |
| Math equations      | Text description của ý nghĩa            |

### Section-Specific Guidelines

#### Executive Summary

- 1 slide overview
- 3-5 key points
- Hero image placeholder

#### Requirements (Section 2)

- 1-2 slides per subsection
- Focus on key use cases
- Stakeholder matrix as chart

#### Architecture (Section 3)

- ADRs: Decision → Rationale format
- Characteristics: Radar chart placeholder
- Style comparison: Table hoặc bullet points

#### Views (Section 4)

- 1 slide per view type
- Diagram placeholder cho mỗi view
- Brief description

#### Implementation (Section 5-6)

- SOLID: 1 slide per principle
- Implementation: Feature highlights
- Screenshots placeholders

#### Reflection (Section 7)

- Lessons learned
- Challenges & solutions
- Future work

## Risks / Trade-offs

### Risk 1: Information Loss

**Risk**: Cô đọng quá mức có thể mất thông tin quan trọng
**Mitigation**:

- Review với báo cáo gốc
- Giữ lại key metrics và decisions
- Link đến báo cáo đầy đủ trong README

### Risk 2: Skywork AI Compatibility

**Risk**: Tags có thể không được Skywork AI xử lý đúng
**Mitigation**:

- Tuân thủ chính xác cú pháp trong term.md
- Test với một section nhỏ trước

### Risk 3: Visual Consistency

**Risk**: Placeholders có thể tạo ra hình ảnh không phù hợp
**Mitigation**:

- Mô tả chi tiết trong placeholder
- Specify style (professional, technical, etc.)

## Output Structure

```
report/presentation/
├── README.md           # Hướng dẫn sử dụng
└── slides.md           # Nội dung slides chính
```

### slides.md Structure

```markdown
# [SLIDE DECK]: Intelligent Tutoring System - ITS

## 1. Executive Summary

### Overview

- Key point 1
- Key point 2
  [FULL_SLIDE_IMAGE_PLACEHOLDER: ...]

## 2. Requirements Analysis

### Project Scope

- ...
  [PLACEHOLDER: ...]

### Stakeholders

- ...
  [CHART_TYPE: ...]

... (tiếp tục cho các sections khác)

## 7. Conclusion

### Key Takeaways

- ...
  [FULL_SLIDE_IMAGE_PLACEHOLDER: ...]

### Q&A

- Thank you
- Contact information
```

## Open Questions

1. **Số lượng slides tối ưu?**

   - Đề xuất: 25-35 slides cho presentation 15-20 phút

2. **Ngôn ngữ output?**

   - Đề xuất: Tiếng Việt (theo báo cáo gốc) với technical terms giữ nguyên tiếng Anh

3. **Mức độ chi tiết cho diagrams?**
   - Đề xuất: Mô tả đủ để Skywork AI hiểu context, không quá chi tiết
