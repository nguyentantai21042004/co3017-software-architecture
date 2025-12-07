# Báo cáo Kiến trúc Phần mềm - ITS

## Mục lục

- [Tổng quan](#tổng-quan)
- [Cấu trúc Thư mục](#cấu-trúc-thư-mục)
- [Yêu cầu Hệ thống](#yêu-cầu-hệ-thống)
- [Biên dịch Báo cáo](#biên-dịch-báo-cáo)
- [Xem Báo cáo](#xem-báo-cáo)
- [Hình ảnh và Sơ đồ](#hình-ảnh-và-sơ-đồ)
- [Khắc phục Sự cố](#khắc-phục-sự-cố)

---

## Tổng quan

Thư mục này chứa báo cáo Kiến trúc Phần mềm (Software Architecture) cho dự án Hệ thống Gia sư Thông minh (Intelligent Tutoring System - ITS), được viết bằng LaTeX theo chuẩn của Đại học Bách Khoa TP.HCM.

**Thông tin môn học:**

- **Mã môn:** CO3017 - Kiến Trúc Phần Mềm
- **Học kỳ:** HK251 (2024-2025)
- **Giảng viên hướng dẫn:** ThS. Trần Trương Tuấn Phát

**Nội dung báo cáo:**

- Tóm tắt điều hành (Executive Summary)
- Phân tích yêu cầu (Requirements Analysis)
- Quyết định kiến trúc (Architecture Decisions)
- Các góc nhìn kiến trúc (Architecture Views)
- Áp dụng nguyên tắc SOLID
- Triển khai hệ thống (System Implementation)
- Đánh giá và phản ánh (Reflection & Evaluation)

---

## Cấu trúc Thư mục

```text
report/
├── main.tex              # File LaTeX chính
├── main.pdf              # Báo cáo đã biên dịch (PDF)
├── SA_hk251_Assignment.pdf  # Đề bài assignment
├── contents/             # Nội dung các chương
│   ├── 1_executive_summary.tex
│   ├── 2.1_project_scope_and_objectives.tex
│   ├── 2.2_stakeholder_analysis.tex
│   ├── 2.3_functional_requirements.tex
│   ├── 2.4_non_functional_requirements.tex
│   ├── 3.1_architecture_characteristics_prioritization.tex
│   ├── 3.2_architecture_style_selection.tex
│   ├── 3.3_architecture_decision_records.tex
│   ├── 3.4_design_principles.tex
│   ├── 4.1_module_view.tex
│   ├── 4.2_component_connector_view.tex
│   ├── 4.3_allocation_view.tex
│   ├── 4.4_behavior_view.tex
│   ├── 5_apply_SOLID_principle.tex
│   ├── 6_system_implementation.tex
│   └── 7_reflection_and_evaluation.tex
├── images/               # Hình ảnh và sơ đồ (PNG)
├── puml/                 # Source files PlantUML
└── changelogs/           # Lịch sử thay đổi
```

---

## Yêu cầu Hệ thống

### Phần mềm Bắt buộc

| Phần mềm             | Phiên bản Tối thiểu | Mục đích                 |
| -------------------- | ------------------- | ------------------------ |
| TeX Live hoặc MiKTeX | 2022+               | Biên dịch LaTeX          |
| latexmk              | 4.70+               | Tự động hóa biên dịch    |
| PlantUML             | 1.2023+             | Tạo sơ đồ UML (tùy chọn) |
| Java JRE             | 11+                 | Chạy PlantUML (tùy chọn) |

### Kiểm tra Cài đặt

Chạy các lệnh sau để xác nhận các công cụ đã được cài đặt đúng:

```bash
# Kiểm tra LaTeX (yêu cầu phiên bản 2022+)
pdflatex --version
# Output mong đợi: pdfTeX 3.x hoặc cao hơn

# Kiểm tra latexmk (yêu cầu phiên bản 4.70+)
latexmk --version
# Output mong đợi: Latexmk, John Collins, ...

# Kiểm tra PlantUML (tùy chọn, yêu cầu Java 11+)
java -version
# Output mong đợi: openjdk version "11.x.x" hoặc cao hơn

plantuml -version
# Output mong đợi: PlantUML version 1.2023.x hoặc cao hơn
```

**Lưu ý:** Nếu bất kỳ lệnh nào báo lỗi "command not found", hãy cài đặt theo hướng dẫn bên dưới.

### Cài đặt trên macOS

```bash
# Sử dụng Homebrew
brew install --cask mactex
brew install latexmk

# PlantUML (tùy chọn)
brew install plantuml
```

### Cài đặt trên Ubuntu/Debian

```bash
# TeX Live đầy đủ
sudo apt-get install texlive-full latexmk

# PlantUML (tùy chọn)
sudo apt-get install plantuml
```

### Cài đặt trên Windows

1. Tải và cài đặt [MiKTeX](https://miktex.org/download)
2. Cài đặt latexmk qua MiKTeX Console
3. (Tùy chọn) Tải [PlantUML](https://plantuml.com/download)

---

## Biên dịch Báo cáo

### Sử dụng latexmk (Khuyến nghị)

```bash
cd report

# Biên dịch PDF
latexmk -pdf main.tex

# Biên dịch và tự động mở PDF
latexmk -pdf -pv main.tex

# Chế độ theo dõi (tự động biên dịch khi có thay đổi)
latexmk -pdf -pvc main.tex
```

### Sử dụng pdflatex Trực tiếp

```bash
cd report

# Biên dịch nhiều lần để cập nhật references
pdflatex main.tex
pdflatex main.tex
pdflatex main.tex
```

### Dọn dẹp File Tạm

```bash
cd report

# Xóa tất cả file tạm
latexmk -C

# Hoặc xóa thủ công
rm -f *.aux *.log *.out *.toc *.lof *.lot *.fls *.fdb_latexmk *.synctex.gz
```

---

## Xem Báo cáo

### File PDF Đã Biên dịch

Báo cáo đã được biên dịch sẵn tại `report/main.pdf`. Mở bằng bất kỳ PDF viewer nào:

```bash
# macOS
open report/main.pdf

# Linux
xdg-open report/main.pdf

# Windows
start report/main.pdf
```

### Đề bài Assignment

Đề bài gốc có tại `report/SA_hk251_Assignment.pdf`.

---

## Hình ảnh và Sơ đồ

### Thư mục Images

Thư mục `images/` chứa tất cả hình ảnh được sử dụng trong báo cáo:

| Loại Sơ đồ        | Files                                                  | Mô tả                       |
| ----------------- | ------------------------------------------------------ | --------------------------- |
| Sequence Diagrams | `*_sequence.png`                                       | Sơ đồ tuần tự các use cases |
| ERD               | `erd_*.png`                                            | Sơ đồ quan hệ thực thể      |
| Deployment        | `deployment_*.png`                                     | Sơ đồ triển khai            |
| Architecture      | `service_architecture.png`, `system_decomposition.png` | Kiến trúc hệ thống          |
| Screenshots       | `dashboard.png`, `homepage.png`, `learnpage.png`       | Giao diện ứng dụng          |

### Tạo Sơ đồ từ PlantUML

Thư mục `puml/` chứa source files PlantUML. Để tạo lại hình ảnh:

```bash
cd report

# Tạo một sơ đồ cụ thể
plantuml -tpng puml/erd_content_service.puml -o ../images/

# Tạo tất cả sơ đồ
plantuml -tpng puml/*.puml -o ../images/
```

### Danh sách PlantUML Files

| File                             | Mô tả                         |
| -------------------------------- | ----------------------------- |
| `erd_content_service.puml`       | ERD của Content Service       |
| `erd_learner_model_service.puml` | ERD của Learner Model Service |
| `erd_user_service.puml`          | ERD của User Service          |
| `erd_mvp_overview.puml`          | Tổng quan ERD cho MVP         |
| `domain_model.puml`              | Mô hình miền nghiệp vụ        |
| `deployment_diagram.puml`        | Sơ đồ triển khai              |
| `enhanced_deployment.puml`       | Sơ đồ triển khai mở rộng      |
| `ai_pipeline_dataflow.puml`      | Luồng dữ liệu AI pipeline     |
| `*_sequence.puml`                | Các sơ đồ tuần tự             |

---

## Khắc phục Sự cố

### Các Lỗi Thường gặp

#### 1. Thiếu Gói LaTeX (Missing Package)

**Triệu chứng:**

```text
! LaTeX Error: File `vntex.sty' not found.
```

**Giải pháp:**

```bash
# TeX Live
tlmgr install vntex

# MiKTeX: Mở MiKTeX Console và cài đặt package
```

#### 2. Lỗi Phông chữ Tiếng Việt (Vietnamese Font Error)

**Triệu chứng:**

```text
! Package inputenc Error: Unicode character ... not set up for use with LaTeX.
```

**Giải pháp:**
Đảm bảo đã cài đặt package `vntex` và sử dụng encoding UTF-8:

```bash
tlmgr install vntex collection-langvietnamese
```

#### 3. Hình ảnh Không Tìm thấy (Image Not Found)

**Triệu chứng:**

```text
! LaTeX Error: File `images/diagram.png' not found.
```

**Giải pháp:**

- Kiểm tra file tồn tại trong thư mục `images/`
- Tạo lại từ PlantUML nếu cần:
  ```bash
  plantuml -tpng puml/diagram.puml -o ../images/
  ```

#### 4. Lỗi Tràn Bộ nhớ (Memory Overflow)

**Triệu chứng:**

```text
! TeX capacity exceeded, sorry [main memory size=...]
```

**Giải pháp:**

```bash
# Tăng memory trong texmf.cnf hoặc sử dụng LuaLaTeX
lualatex main.tex
```

#### 5. Tham chiếu Không Cập nhật (References Not Updated)

**Triệu chứng:**

- Mục lục hiển thị "??"
- Số trang không đúng

**Giải pháp:**

```bash
# Biên dịch nhiều lần
latexmk -pdf main.tex

# Hoặc xóa cache và biên dịch lại
latexmk -C
latexmk -pdf main.tex
```

### Các Gói LaTeX Cần thiết (Required LaTeX Packages)

Báo cáo sử dụng các gói (packages) sau (đã được cài sẵn trong TeX Live Full):

| Gói (Package) | Mục đích                       |
| ------------- | ------------------------------ |
| `vntex`       | Hỗ trợ tiếng Việt              |
| `mathptmx`    | Phông chữ Times (Font Times)   |
| `graphicx`    | Chèn hình ảnh                  |
| `hyperref`    | Liên kết trong PDF             |
| `fancyhdr`    | Đầu trang/Chân trang tùy chỉnh |
| `tikz`        | Vẽ đồ họa                      |
| `listings`    | Hiển thị mã nguồn (code)       |
| `booktabs`    | Bảng đẹp                       |
| `geometry`    | Căn lề trang                   |
| `caption`     | Chú thích hình ảnh và bảng     |

---

## Liên kết Liên quan

### Tài liệu Dự án

| Tài liệu                | Đường dẫn                                              | Mô tả                                  |
| ----------------------- | ------------------------------------------------------ | -------------------------------------- |
| **Root README**         | [../README.md](../README.md)                           | Tổng quan dự án, cấu trúc repository   |
| **Sources README**      | [../sources/README.md](../sources/README.md)           | Hướng dẫn microservices và Docker      |
| **Presentation README** | [../presentation/README.md](../presentation/README.md) | Slide thuyết trình và hướng dẫn export |
| **Assignment PDF**      | [./SA_hk251_Assignment.pdf](./SA_hk251_Assignment.pdf) | Đề bài gốc của môn học                 |

### Tài liệu Phân tích Kiến trúc

| Tài liệu                         | Đường dẫn                                                                                                    | Nội dung                            |
| -------------------------------- | ------------------------------------------------------------------------------------------------------------ | ----------------------------------- |
| **Stakeholder Analysis**         | [../markdown/report/1-analyst.md](../markdown/report/1-analyst.md)                                           | Phân tích stakeholder, user stories |
| **Architecture Characteristics** | [../markdown/report/2-architecture-characteristics.md](../markdown/report/2-architecture-characteristics.md) | Đặc tính kiến trúc                  |
| **Architecture Styles**          | [../markdown/report/3-architecture-styles.md](../markdown/report/3-architecture-styles.md)                   | So sánh kiểu kiến trúc              |
| **Architecture Decisions**       | [../markdown/report/5-architecture-decisions.md](../markdown/report/5-architecture-decisions.md)             | ADRs                                |
| **SOLID Principles**             | [../markdown/report/6-SOLID-principles.md](../markdown/report/6-SOLID-principles.md)                         | Ví dụ áp dụng SOLID                 |
| **Reflection Report**            | [../markdown/report/7-reflection-report.md](../markdown/report/7-reflection-report.md)                       | Đánh giá và bài học                 |
| **Microservices Analysis**       | [../markdown/microservices.md](../markdown/microservices.md)                                                 | Chi tiết domain model               |

### Service READMEs

| Service                   | Đường dẫn                                                                    | Mô tả                    |
| ------------------------- | ---------------------------------------------------------------------------- | ------------------------ |
| **Content Service**       | [../sources/content/README.md](../sources/content/README.md)                 | Quản lý nội dung học tập |
| **Scoring Service**       | [../sources/scoring/README.md](../sources/scoring/README.md)                 | Chấm điểm                |
| **Learner Model Service** | [../sources/learner-model/README.md](../sources/learner-model/README.md)     | Skill mastery tracking   |
| **Adaptive Engine**       | [../sources/adaptive-engine/README.md](../sources/adaptive-engine/README.md) | Recommendation           |
| **Client**                | [../sources/client/README.md](../sources/client/README.md)                   | Frontend UI/UX           |

### Diagrams

| Diagram                       | Đường dẫn                                                                                                                | Loại               |
| ----------------------------- | ------------------------------------------------------------------------------------------------------------------------ | ------------------ |
| **Domain Model**              | [../markdown/diagrams/domain_model_class_diagram.md](../markdown/diagrams/domain_model_class_diagram.md)                 | Class Diagram      |
| **Deployment Architecture**   | [../markdown/diagrams/deployment_architecture_onprem.md](../markdown/diagrams/deployment_architecture_onprem.md)         | Deployment Diagram |
| **Adaptive Content Delivery** | [../markdown/diagrams/adaptive_content_delivery_sequence.md](../markdown/diagrams/adaptive_content_delivery_sequence.md) | Sequence Diagram   |

---

## Lịch sử Thay đổi

Xem thư mục `changelogs/` để biết chi tiết các thay đổi đã thực hiện trên báo cáo.

```bash
ls -la report/changelogs/
```

### Các Thay đổi Gần đây

| Ngày       | File                              | Mô tả                       |
| ---------- | --------------------------------- | --------------------------- |
| 2025-12-07 | `acceptance-criteria-20251207.md` | Cập nhật tiêu chí chấp nhận |
| 2025-12-07 | `fitness-functions-20251207.md`   | Thêm hàm đánh giá kiến trúc |
| 2025-12-07 | `component-diagram-20251207.md`   | Cập nhật sơ đồ thành phần   |

Xem [changelogs/README.md](./changelogs/README.md) để biết quy ước đặt tên và template changelog.

---

## Đóng góp

Khi chỉnh sửa báo cáo:

1. **Tạo changelog** - Ghi lại thay đổi trong `changelogs/`
2. **Biên dịch kiểm tra** - Chạy `latexmk -pdf main.tex` để đảm bảo không có lỗi
3. **Cập nhật hình ảnh** - Nếu thay đổi PlantUML, tạo lại PNG
4. **Commit message** - Sử dụng format: `docs(report): <mô tả ngắn gọn>`

---

_Cập nhật lần cuối: Tháng 12, 2025_
