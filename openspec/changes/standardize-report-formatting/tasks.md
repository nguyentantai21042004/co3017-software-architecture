# Tasks: Standardizing Report Formatting & Writing Quality

## Phase 1: Text Cleanup (Pass 1) ✅ COMPLETE

### Task 1.1: Remove Icons & Emojis ✅ COMPLETE

- [x] Scan all 17 LaTeX files for icons/emojis
- [x] Remove: ✔ ✨ 📌 🔥 ✅ ❌ ⚠️ and similar symbols
- [x] Replace with text equivalents where needed (e.g., "Verified", "Warning")
- [x] Create `report/changelog/remove-icons-20251207.md`

**Changes Made:**

- Replaced `→` with `$\rightarrow$` in `2.3_functional_requirements.tex` (event chain)
- Replaced `→` with `:` in `4.3_allocation_view.tex` (section references, network layers)

### Task 1.2: Fix Pronoun Usage ✅ COMPLETE

- [x] Search for "thầy cô", "người đọc", "giảng viên", "tôi"
- [x] Replace with "nhóm" or passive voice constructions
- [x] Create `report/changelog/fix-pronouns-20251207.md`

**Changes Made:**

- `7_reflection_and_evaluation.tex`: "chúng tôi" → "nhóm" (2 instances)
- `4.1_module_view.tex`: "chúng tôi" → "nhóm" (1 instance)
- `5_apply_SOLID_principle.tex`: "Chúng tôi" → "Nhóm" (1 instance)
- `4.2_component_connector_view.tex`: "chúng tôi" → "nhóm" (2 instances)
- `6_system_implementation.tex`: "chúng tôi" → "báo cáo/nhóm" (2 instances)
- `3.3_architecture_decision_records.tex`: "của tôi" → removed (1 instance)

**Note:** User Stories retain "tôi" as it's standard format.

### Task 1.3: Standardize Capitalization ✅ COMPLETE

- [x] Review each file for improper capitalization
- [x] Keep capitalized: proper nouns, technology names, system names, diagram titles
- [x] Lowercase: common nouns, descriptive phrases mid-sentence
- [x] Create `report/changelog/fix-capitalization-20251207.md`

**Result:** All 17 files already follow proper academic capitalization standards. No changes needed.

---

## Phase 2: Formatting & Structure (Pass 2)

### Task 2.1: Verify Caption Positions

- [ ] Check all `\caption{}` commands in LaTeX
- [ ] Tables: caption BEFORE `\begin{tabular}`
- [ ] Figures: caption AFTER `\includegraphics`
- [ ] Fix any incorrect positions
- [ ] Create `report/changelog/fix-captions-20251207.md`

### Task 2.2: Review Image Naming Convention

- [ ] Audit current image file names
- [ ] Document any files needing rename (if any)
- [ ] Note: Most files already use snake_case - verify consistency
- [ ] Update `\includegraphics{}` references if files renamed

**Current naming analysis:**

- ✅ `adaptive_content_delivery_sequence.png` - correct
- ✅ `erd_content_service.png` - correct
- ✅ `system_decomposition.png` - correct
- Review: `usecase_9.png`, `usecase_10.png`, `usecase_11.png`

### Task 2.3: Review Spacing & Indentation

- [ ] Check consistent spacing around sections
- [ ] Verify proper indentation in code listings
- [ ] Ensure consistent paragraph spacing
- [ ] Fix any formatting inconsistencies

---

## Phase 3: Consistency Check (Pass 3)

### Task 3.1: Cross-Reference Verification

- [ ] Verify all `\ref{}` and `\label{}` pairs
- [ ] Check table/figure numbering sequence
- [ ] Ensure no broken references
- [ ] Run LaTeX compilation to detect issues

### Task 3.2: Grammar & Expression Review

- [ ] Review sentence structure for clarity
- [ ] Fix any grammatical errors
- [ ] Ensure technical terms used consistently
- [ ] Verify terminology matches glossary (if exists)

### Task 3.3: Terminology Consistency

- [ ] Create list of key terms and their standard forms
- [ ] Verify consistent usage throughout document
- [ ] Examples:
  - "microservice" vs "micro-service" vs "Microservice"
  - "API" vs "Api" vs "api"
  - "database" vs "Database"

---

## Phase 4: Final Review (Pass 4)

### Task 4.1: Full Document Read-Through

- [ ] Read entire document for flow and coherence
- [ ] Check transitions between sections
- [ ] Verify logical progression of ideas
- [ ] Note any remaining issues

### Task 4.2: Final Compilation & Verification

- [ ] Run `pdflatex main.tex` twice
- [ ] Verify no compilation errors
- [ ] Check page count (should remain ~95 pages)
- [ ] Verify all figures/tables render correctly
- [ ] Create `report/changelog/formatting-final-20251207.md`

### Task 4.3: Create Formatting Guidelines Document

- [ ] Document all formatting rules applied
- [ ] Create `report/formatting-guidelines.md` for future reference
- [ ] Include examples of correct vs incorrect formatting

---

## Summary

| Phase                 | Tasks        | Estimated Time |
| --------------------- | ------------ | -------------- |
| Phase 1: Text Cleanup | 3 tasks      | 1.5-2 hours    |
| Phase 2: Formatting   | 3 tasks      | 1-1.5 hours    |
| Phase 3: Consistency  | 3 tasks      | 1-1.5 hours    |
| Phase 4: Final Review | 3 tasks      | 0.5-1 hour     |
| **Total**             | **12 tasks** | **4-6 hours**  |

## Validation Checklist

- [ ] No icons/emojis in document
- [ ] No improper pronouns ("thầy cô", "người đọc", "tôi")
- [ ] Consistent capitalization
- [ ] Correct caption positions
- [ ] Consistent file naming
- [ ] No broken cross-references
- [ ] LaTeX compiles without errors
- [ ] Professional academic tone throughout
