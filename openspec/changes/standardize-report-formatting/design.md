# Design: Standardizing Report Formatting & Writing Quality

## Context

Report Software Architecture cho hệ thống ITS đã hoàn thành về nội dung (95 trang, expected score 100%). Tuy nhiên, để đạt tiêu chuẩn học thuật cao nhất, cần chuẩn hoá về:

- Văn phong và ngôi xưng
- Quy tắc viết hoa
- Định dạng caption và tên file
- Loại bỏ các yếu tố phi học thuật (icons, emojis)

## Goals

1. **Consistency**: Đảm bảo tính nhất quán trong toàn bộ document
2. **Professionalism**: Văn phong học thuật, chuyên nghiệp
3. **Readability**: Dễ đọc, dễ theo dõi
4. **Maintainability**: Dễ dàng cập nhật và mở rộng trong tương lai

## Non-Goals

- Không thay đổi nội dung kỹ thuật
- Không thêm/bớt sections
- Không thay đổi cấu trúc document
- Không rename image files (chỉ document nếu cần)

## Decisions

### Decision 1: Capitalization Rules

**Rule**: Chỉ viết hoa trong các trường hợp sau:

1. Tên riêng (người, tổ chức)
2. Tên công nghệ/framework: Spring Boot, PostgreSQL, RabbitMQ, Go, Java
3. Tên hệ thống/module: Content Service, Scoring Service, Adaptive Engine
4. Tên diagram chính thức khi reference: "Figure 3.1", "Table 2.1"
5. Chữ cái đầu câu
6. Acronyms: API, REST, SOLID, DDD, CQRS

**Không viết hoa**:

- Danh từ chung: architecture, design, module, service (khi không phải tên riêng)
- Cụm từ mô tả: "the system architecture", "database design"

### Decision 2: Pronoun Replacement Strategy

| Original     | Replacement      | Example                   |
| ------------ | ---------------- | ------------------------- |
| "tôi"        | "nhóm"           | "Nhóm đã thiết kế..."     |
| "chúng tôi"  | "nhóm"           | "Nhóm thực hiện..."       |
| "thầy cô"    | (remove/passive) | "Được đánh giá..."        |
| "người đọc"  | (remove/passive) | "Có thể thấy rằng..."     |
| "giảng viên" | (remove/passive) | "Theo yêu cầu môn học..." |

**Preferred constructions**:

- Passive voice: "Hệ thống được thiết kế để..."
- Impersonal: "Báo cáo trình bày..."
- Group reference: "Nhóm đánh giá rằng..."

### Decision 3: Caption Positioning (LaTeX)

**Tables** - Caption ABOVE:

```latex
\begin{table}[htbp]
\caption{Comparison of Architecture Styles}
\label{tab:arch-comparison}
\begin{tabular}{...}
...
\end{tabular}
\end{table}
```

**Figures** - Caption BELOW:

```latex
\begin{figure}[htbp]
\centering
\includegraphics[width=0.8\textwidth]{images/system_architecture.png}
\caption{System Architecture Overview}
\label{fig:system-arch}
\end{figure}
```

### Decision 4: Icon/Emoji Replacement

| Icon | Text Replacement               |
| ---- | ------------------------------ |
| ✅   | "Implemented" or "Complete"    |
| ❌   | "Not implemented" or "Planned" |
| ⚠️   | "Partial" or "Warning"         |
| 📌   | (remove, use bold text)        |
| ✔    | "Yes" or checkmark in table    |
| ✨   | (remove entirely)              |
| 🔥   | (remove entirely)              |

**In tables**: Use text or LaTeX symbols:

- `\checkmark` for checkmarks
- `$\times$` for X marks
- Plain text for status

### Decision 5: File Naming Convention (Documentation Only)

Current naming is mostly consistent (snake_case). Document the standard:

**Images**:

- Pattern: `<type>_<description>.png`
- Examples: `erd_content_service.png`, `sequence_user_registration.png`

**Tables** (if exported as images):

- Pattern: `table_<topic>.png`

**Note**: No actual file renaming needed - current names are acceptable.

## Implementation Approach

### Pass 1: Text Cleanup (Automated + Manual)

1. Use grep/search to find icons and pronouns
2. Replace systematically file by file
3. Manual review for capitalization

### Pass 2: Formatting (Manual)

1. Review each file's caption positions
2. Verify LaTeX structure
3. Check spacing consistency

### Pass 3: Consistency (Semi-automated)

1. Compile LaTeX to check references
2. Review terminology usage
3. Grammar check

### Pass 4: Final Review (Manual)

1. Full read-through
2. Final compilation
3. Documentation

## Risks & Mitigations

| Risk                       | Mitigation                      |
| -------------------------- | ------------------------------- |
| Breaking LaTeX compilation | Compile after each file change  |
| Changing technical meaning | Review changes in context       |
| Missing some instances     | Use systematic search patterns  |
| Cross-reference breaks     | Verify all \ref{} after changes |

## Search Patterns for Implementation

### Icons/Emojis

```bash
grep -r "✅\|❌\|⚠️\|📌\|✔\|✨\|🔥" report/contents/
```

### Pronouns

```bash
grep -ri "thầy cô\|người đọc\|giảng viên" report/contents/
grep -ri "\btôi\b" report/contents/
```

### Capitalization Issues (manual review needed)

- Look for mid-sentence capitals
- Check consistency of term usage

## Quality Metrics

| Metric                   | Target |
| ------------------------ | ------ |
| Icons remaining          | 0      |
| Improper pronouns        | 0      |
| Capitalization issues    | 0      |
| Caption position errors  | 0      |
| LaTeX compilation errors | 0      |
| Broken references        | 0      |
