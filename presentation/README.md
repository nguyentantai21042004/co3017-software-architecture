# Bài Thuyết Trình - Intelligent Tutoring System

## Tổng Quan

Thư mục này chứa bài thuyết trình (presentation) cho dự án Hệ thống Gia sư Thông minh (Intelligent Tutoring System - ITS), được xây dựng bằng [Marp](https://marp.app/) - công cụ tạo slide từ Markdown.

**Thông tin bài thuyết trình:**

| Thuộc tính | Giá trị                     |
| ---------- | --------------------------- |
| Môn học    | CO3017 - Kiến Trúc Phần Mềm |
| Trường     | Đại học Bách Khoa TP.HCM    |
| Sinh viên  | Nguyễn Tấn Tài - 2212990    |
| Số slides  | 57 slides                   |
| Thời lượng | ~30 phút                    |

## Mục Lục

- [Xem Bài Thuyết Trình](#xem-bài-thuyết-trình)
- [Cấu Trúc Thư Mục](#cấu-trúc-thư-mục)
- [Yêu Cầu Hệ Thống](#yêu-cầu-hệ-thống)
- [Build từ Source](#build-từ-source)
- [Export Định Dạng Khác](#export-định-dạng-khác)
- [Hình Ảnh và Diagrams](#hình-ảnh-và-diagrams)
- [Tùy Chỉnh Theme](#tùy-chỉnh-theme)
- [Khắc Phục Sự Cố](#khắc-phục-sự-cố)

## Xem Bài Thuyết Trình

### Cách Nhanh Nhất

Mở file PDF đã được build sẵn:

```bash
# macOS
open slides.pdf

# Linux
xdg-open slides.pdf

# Windows
start slides.pdf
```

### Xem Dạng HTML

Mở file `slides.html` trong trình duyệt để xem với các tính năng tương tác:

```bash
# macOS
open slides.html

# Linux
xdg-open slides.html

# Windows
start slides.html
```

### Xem Trực Tiếp (Live Preview)

Nếu đã cài đặt Marp CLI, chạy preview server với live reload:

```bash
# Sử dụng Makefile
make preview

# Hoặc chạy trực tiếp
marp -p slides.md --allow-local-files
```

Mở trình duyệt tại `http://localhost:8080` để xem slides.

## Cấu Trúc Thư Mục

```text
presentation/
├── slides.md           # Source file Markdown
├── slides.pdf          # PDF đã build
├── slides.html         # HTML đã build
├── theme.css           # Custom theme cho slides
├── Makefile            # Automation commands
├── export-diagram.sh   # Script export Mermaid diagrams
├── images/             # Hình ảnh sử dụng trong slides
│   ├── hcmut.png                              # Logo HCMUT
│   ├── header.png                             # Header image
│   ├── homepage.png                           # Screenshot trang chủ
│   ├── dashboard.png                          # Screenshot dashboard
│   ├── learnpage.png                          # Screenshot trang học
│   ├── service_architecture.png               # Sơ đồ kiến trúc services
│   ├── domain_model_class_diagram.png         # Class diagram
│   ├── deployment_architecture_onprem.png     # Deployment diagram
│   ├── adaptive_content_delivery_sequence.png # Sequence diagram
│   ├── assessment_submission_and_scoring_sequence.png
│   ├── architectural-characteristics-radar.png
│   ├── clean-architecture-layers.png
│   └── ...
└── README.md           # File này
```

## Yêu Cầu Hệ Thống

### Để Xem Slides

- Trình đọc PDF (Adobe Reader, Preview, Foxit, etc.)
- Hoặc trình duyệt web hiện đại (Chrome, Firefox, Safari, Edge)

### Để Build từ Source

| Phần mềm | Phiên bản | Mục đích             |
| -------- | --------- | -------------------- |
| Node.js  | >= 18.x   | Runtime cho Marp CLI |
| npm      | >= 9.x    | Package manager      |
| Marp CLI | >= 3.x    | Build slides         |

### Kiểm Tra Cài Đặt

```bash
# Kiểm tra Node.js
node --version
# Output: v18.x.x hoặc cao hơn

# Kiểm tra npm
npm --version
# Output: 9.x.x hoặc cao hơn

# Kiểm tra Marp CLI
marp --version
# Output: @marp-team/marp-cli 3.x.x
```

### Cài Đặt Marp CLI

```bash
# Cài đặt global
npm install -g @marp-team/marp-cli

# Hoặc sử dụng Makefile
make install
```

## Build từ Source

### Sử Dụng Makefile (Khuyến nghị)

```bash
# Build PDF
make pdf

# Build HTML
make html

# Build cả PDF và HTML
make all

# Build tất cả định dạng (PDF, HTML, PPTX)
make build-all
```

### Sử Dụng Marp CLI Trực Tiếp

```bash
# Build PDF
marp slides.md -o slides.pdf --allow-local-files

# Build HTML
marp slides.md -o slides.html --allow-local-files

# Build PowerPoint
marp slides.md -o slides.pptx --allow-local-files
```

### Build với Custom Theme

```bash
# PDF với theme
make pdf-theme

# HTML với theme
make html-theme

# Hoặc trực tiếp
marp slides.md -o slides.pdf --theme theme.css --allow-local-files
```

### Watch Mode (Auto-rebuild)

```bash
# Tự động rebuild khi file thay đổi
make watch

# Hoặc
marp -w slides.md -o slides.pdf --allow-local-files
```

## Export Định Dạng Khác

### PDF (Mặc định)

```bash
make pdf
# Output: slides.pdf
```

### HTML

```bash
make html
# Output: slides.html
```

HTML output có các tính năng:

- Navigation bằng phím mũi tên
- Fullscreen mode (phím F)
- Overview mode (phím O)
- Speaker notes (nếu có)

### PowerPoint (PPTX)

```bash
make pptx
# Output: slides.pptx
```

**Lưu ý:** PPTX export có thể mất một số formatting so với PDF/HTML.

### Mở File Sau Khi Build

```bash
# Build và mở PDF
make open-pdf

# Build và mở HTML
make open-html
```

## Hình Ảnh và Diagrams

### Danh Sách Hình Ảnh

| File                                             | Mô tả                              |
| ------------------------------------------------ | ---------------------------------- |
| `hcmut.png`                                      | Logo Đại học Bách Khoa TP.HCM      |
| `header.png`                                     | Header image cho title slide       |
| `homepage.png`                                   | Screenshot trang chủ ứng dụng      |
| `dashboard.png`                                  | Screenshot dashboard học sinh      |
| `learnpage.png`                                  | Screenshot trang học tập           |
| `service_architecture.png`                       | Sơ đồ kiến trúc microservices      |
| `domain_model_class_diagram.png`                 | Class diagram domain model         |
| `deployment_architecture_onprem.png`             | Sơ đồ triển khai on-premise        |
| `adaptive_content_delivery_sequence.png`         | Sequence diagram adaptive delivery |
| `assessment_submission_and_scoring_sequence.png` | Sequence diagram scoring flow      |
| `architectural-characteristics-radar.png`        | Radar chart đặc tính kiến trúc     |
| `clean-architecture-layers.png`                  | Sơ đồ Clean Architecture           |
| `system_decomposition.png`                       | Sơ đồ phân rã hệ thống             |
| `skateholder-analysis.png`                       | Ma trận phân tích stakeholder      |
| `usecase_9.png`                                  | Use case diagram                   |
| `mvp-status.png`                                 | Trạng thái MVP                     |
| `dev-roadmap.png`                                | Lộ trình phát triển                |

### Export Mermaid Diagrams

Nếu cần tạo diagram mới từ Mermaid:

```bash
# Cài đặt Mermaid CLI
make install-mermaid
# Hoặc: npm install -g @mermaid-js/mermaid-cli

# Export diagram
make diagram
# Hoặc: ./export-diagram.sh input.md output.png
```

### Thêm Hình Ảnh Mới

1. Đặt file hình ảnh vào thư mục `images/`
2. Tham chiếu trong `slides.md`:

```markdown
<!-- Hình ảnh bên phải -->

![bg right:40% fit](images/your-image.png)

<!-- Hình ảnh centered -->

![center](images/your-image.png)

<!-- Hình ảnh với kích thước cụ thể -->

![width:600px](images/your-image.png)
```

## Tùy Chỉnh Theme

### Cấu Trúc Chủ đề (Theme Structure)

File `theme.css` định nghĩa chủ đề tùy chỉnh `its-theme`:

```css
/* @theme its-theme */

@import "default";

:root {
  --color-primary: #1e3a8a; /* Xanh đậm */
  --color-secondary: #3b82f6; /* Xanh nhạt */
  --color-accent: #10b981; /* Xanh lá */
  --color-warning: #f59e0b; /* Vàng */
  --color-background: #f8fafc; /* Nền sáng */
  --color-text: #1f2937; /* Text đậm */
}
```

### Các Lớp Đặc Biệt (Special Classes)

| Lớp (Class)   | Mô tả                               | Sử dụng                        |
| ------------- | ----------------------------------- | ------------------------------ |
| `title-slide` | Slide tiêu đề với nền HCMUT         | `<!-- _class: title-slide -->` |
| `lead`        | Slide đầu mục với hiệu ứng gradient | `<!-- _class: lead -->`        |

### Thay Đổi Màu Sắc (Change Colors)

Chỉnh sửa các biến CSS trong `theme.css`:

```css
:root {
  --color-primary: #YOUR_COLOR;
  --color-secondary: #YOUR_COLOR;
  /* ... */
}
```

### Thay Đổi Phông chữ (Change Font)

```css
:root {
  --font-heading: "Your Font", sans-serif;
  --font-body: "Your Font", sans-serif;
  --font-code: "Your Mono Font", monospace;
}
```

### Thêm Lớp Tùy chỉnh (Add Custom Class)

```css
/* Trong theme.css */
section.your-class {
  background-color: #custom;
  /* ... */
}
```

Sử dụng trong slides:

```markdown
<!-- _class: your-class -->

# Your Slide Content
```

## Khắc Phục Sự Cố

### Marp CLI Không Tìm Thấy

```bash
# Kiểm tra cài đặt
which marp

# Cài đặt lại
npm install -g @marp-team/marp-cli

# Hoặc sử dụng npx
npx @marp-team/marp-cli slides.md -o slides.pdf
```

### Hình Ảnh Không Hiển Thị

```bash
# Đảm bảo sử dụng flag --allow-local-files
marp slides.md -o slides.pdf --allow-local-files

# Kiểm tra đường dẫn hình ảnh
ls -la images/
```

### Theme Không Áp Dụng

```bash
# Kiểm tra tên theme trong slides.md
# Dòng đầu tiên phải có:
# ---
# marp: true
# theme: its-theme
# ---

# Build với theme
marp slides.md -o slides.pdf --theme theme.css --allow-local-files
```

### PDF Bị Cắt Nội Dung

Điều chỉnh font size hoặc layout trong `slides.md`:

```markdown
<!-- Giảm font size cho slide cụ thể -->
<style scoped>
section { font-size: 20px; }
</style>
```

### Preview Server Không Chạy

```bash
# Kiểm tra port 8080
lsof -i :8080

# Sử dụng port khác
marp -p slides.md --allow-local-files --server.port 3000
```

## Lệnh Makefile Đầy Đủ

```bash
make help           # Hiển thị tất cả lệnh có sẵn
make pdf            # Build PDF
make html           # Build HTML
make pptx           # Build PowerPoint
make all            # Build PDF + HTML
make build-all      # Build tất cả định dạng
make preview        # Live preview server
make preview-theme  # Preview với custom theme
make watch          # Auto-rebuild khi thay đổi
make clean          # Xóa files đã build
make install        # Cài đặt Marp CLI
make check          # Kiểm tra Marp đã cài đặt
make count          # Đếm số slides
make info           # Thông tin presentation
make validate       # Kiểm tra markdown syntax
make diagram        # Export Mermaid diagram
make install-mermaid # Cài đặt Mermaid CLI
```

## Nội Dung Bài Thuyết Trình

Bài thuyết trình bao gồm các phần chính:

1. **Tóm Tắt Tổng Quan** - Executive Summary
2. **Phân Tích Bối Cảnh** - Context và Requirements
3. **Phân Tích Stakeholder** - Ma trận các bên liên quan
4. **Yêu Cầu Chức Năng** - User Stories và Use Cases
5. **Yêu Cầu Phi Chức Năng** - Architecture Characteristics
6. **Thiết Kế Kiến Trúc** - So sánh và quyết định
7. **Architecture Decision Records** - ADRs
8. **Architecture Views** - Module, C&C, Allocation, Behavior
9. **Áp Dụng SOLID** - Ví dụ trong code
10. **Triển Khai Hệ Thống** - MVP Status
11. **Đánh Giá và Phản Hồi** - Lessons Learned
12. **Demo & Code Showcase** - Screenshots và code structure
13. **Kết Luận** - Key Takeaways

## Liên kết Liên quan

### Tài liệu Dự án

| Tài liệu           | Đường dẫn                                    | Mô tả                                |
| ------------------ | -------------------------------------------- | ------------------------------------ |
| **Root README**    | [../README.md](../README.md)                 | Tổng quan dự án, cấu trúc repository |
| **Sources README** | [../sources/README.md](../sources/README.md) | Hướng dẫn microservices và Docker    |
| **Report README**  | [../report/README.md](../report/README.md)   | Hướng dẫn build báo cáo LaTeX        |
| **Báo cáo PDF**    | [../report/main.pdf](../report/main.pdf)     | Báo cáo kiến trúc phần mềm           |

### Tài liệu Phân tích

| Tài liệu                         | Đường dẫn                                                                                                    | Nội dung              |
| -------------------------------- | ------------------------------------------------------------------------------------------------------------ | --------------------- |
| **Assignment**                   | [../markdown/assignment.md](../markdown/assignment.md)                                                       | Yêu cầu bài tập gốc   |
| **Stakeholder Analysis**         | [../markdown/report/1-analyst.md](../markdown/report/1-analyst.md)                                           | Phân tích stakeholder |
| **Architecture Characteristics** | [../markdown/report/2-architecture-characteristics.md](../markdown/report/2-architecture-characteristics.md) | Đặc tính kiến trúc    |
| **Architecture Decisions**       | [../markdown/report/5-architecture-decisions.md](../markdown/report/5-architecture-decisions.md)             | ADRs                  |
| **SOLID Principles**             | [../markdown/report/6-SOLID-principles.md](../markdown/report/6-SOLID-principles.md)                         | Ví dụ áp dụng SOLID   |

### Service READMEs

| Service                   | Đường dẫn                                                                    | Mô tả                    |
| ------------------------- | ---------------------------------------------------------------------------- | ------------------------ |
| **Content Service**       | [../sources/content/README.md](../sources/content/README.md)                 | Quản lý nội dung học tập |
| **Scoring Service**       | [../sources/scoring/README.md](../sources/scoring/README.md)                 | Chấm điểm                |
| **Learner Model Service** | [../sources/learner-model/README.md](../sources/learner-model/README.md)     | Skill mastery tracking   |
| **Adaptive Engine**       | [../sources/adaptive-engine/README.md](../sources/adaptive-engine/README.md) | Recommendation           |
| **Client**                | [../sources/client/README.md](../sources/client/README.md)                   | Frontend UI/UX           |

### External Documentation

- [Marp Documentation](https://marp.app/) - Tài liệu Marp chính thức
- [Mermaid Documentation](https://mermaid.js.org/) - Tài liệu Mermaid diagrams

---

**Lần cập nhật cuối:** 2024-12-07
