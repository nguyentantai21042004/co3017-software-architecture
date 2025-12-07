# Change: Standardizing Report Formatting & Writing Quality

## Why

Report hiện tại đã hoàn chỉnh về mặt nội dung (95 trang, expected score 100%), tuy nhiên cần được chuẩn hoá về câu chữ, cách trình bày và phong cách học thuật để đảm bảo tính chuyên nghiệp, nhất quán và phù hợp với tiêu chuẩn báo cáo của môn học.

## What Changes

### 1. Capitalization Rules (Quy tắc viết hoa)

- Chỉ viết hoa: tên riêng, tên công nghệ/framework, tên hệ thống/module, tên diagram chính thức, chữ cái đầu câu
- Loại bỏ viết hoa tuỳ tiện giữa câu với danh từ chung

### 2. File Naming & Caption Conventions

- Ảnh: `figure_<module>_<index>.png` (snake_case)
- Bảng: `table_<topic>_<index>.png`
- Caption: Bảng đặt phía trên, Hình đặt phía dưới

### 3. Pronoun Usage (Ngôi xưng)

- Loại bỏ: "thầy cô", "người đọc", "giảng viên", "tôi"
- Sử dụng: "nhóm" (ngôi thứ nhất số nhiều) hoặc câu bị động trung lập
- Ví dụ: "Nhóm thực hiện...", "Hệ thống được thiết kế để..."

### 4. Remove Icons & Emojis

- Xóa toàn bộ: ✔ ✨ 📌 🔥 ✅ ❌ ⚠️ và các biểu tượng tương tự
- Giữ format academic-style, tối giản

## Impact

### Affected Files (17 LaTeX files)

- `report/contents/1_executive_summary.tex`
- `report/contents/2.1_project_scope_and_objectives.tex`
- `report/contents/2.2_stakeholder_analysis.tex`
- `report/contents/2.3_functional_requirements.tex`
- `report/contents/2.4_non_functional_requirements.tex`
- `report/contents/2.5_constraints_and_assumptions.tex`
- `report/contents/3.1_architecture_characteristics_prioritization.tex`
- `report/contents/3.2_architecture_style_selection.tex`
- `report/contents/3.3_architecture_decision_records.tex`
- `report/contents/3.4_design_principles.tex`
- `report/contents/4.1_module_view.tex`
- `report/contents/4.2_component_connector_view.tex`
- `report/contents/4.3_allocation_view.tex`
- `report/contents/4.4_behavior_view.tex`
- `report/contents/5_apply_SOLID_principle.tex`
- `report/contents/6_system_implementation.tex`
- `report/contents/7_reflection_and_evaluation.tex`

### Image Files (35 files)

- Current naming: Mixed conventions (snake_case, some inconsistent)
- Most already follow snake_case - minimal changes needed

## Success Criteria

1. Không còn viết hoa tuỳ tiện trong văn bản
2. Tất cả caption đúng vị trí (bảng: trên, hình: dưới)
3. Không còn ngôi xưng "thầy cô", "người đọc", "tôi"
4. Không còn icons/emojis trong report
5. LaTeX compile thành công, không lỗi
6. Văn phong nhất quán, chuyên nghiệp

## Risk Assessment

| Risk                     | Likelihood | Impact | Mitigation                       |
| ------------------------ | ---------- | ------ | -------------------------------- |
| LaTeX compilation errors | Low        | Medium | Compile after each file change   |
| Cross-reference breaks   | Low        | Medium | Verify all \ref{} after renaming |
| Content meaning change   | Low        | High   | Review each change carefully     |

## Timeline

- Estimated effort: 4-6 hours
- 4 passes as defined in implementation plan
