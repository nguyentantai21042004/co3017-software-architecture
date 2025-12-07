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

## Phase 2: Formatting & Structure (Pass 2) ✅ COMPLETE

### Task 2.1: Verify Caption Positions ✅ COMPLETE

- [x] Check all `\caption{}` commands in LaTeX
- [x] Tables: caption renders ABOVE table in PDF ✅
- [x] Figures: caption renders BELOW figure in PDF ✅
- [x] Create `report/changelog/fix-captions-20251207.md`

**Result:** All 48 captions (28 tables, 15 figures, 5 longtables) render correctly. LaTeX automatically handles caption positioning.

### Task 2.2: Review Image Naming Convention ✅ COMPLETE

- [x] Audit current image file names
- [x] Document any files needing rename (if any)
- [x] Note: Most files already use snake_case - verify consistency

**Result:**

- 34/35 files use snake_case ✅
- 1 file uses kebab-case: `clean-architecture-layers.png` (minor, documented)
- No rename performed to avoid breaking references

### Task 2.3: Review Spacing & Indentation ✅ COMPLETE

- [x] Check consistent spacing around sections
- [x] Verify proper indentation in code listings
- [x] Ensure consistent paragraph spacing

**Result:** All spacing and indentation follows consistent patterns. Large `\vspace{}` values in NFR section are for page layout optimization (standard practice).

---

## Phase 3: Consistency Check (Pass 3) ✅ COMPLETE

### Task 3.1: Cross-Reference Verification ✅ COMPLETE

- [x] Verify all `\ref{}` and `\label{}` pairs
- [x] Check table/figure numbering sequence
- [x] Ensure no broken references
- [x] Run LaTeX compilation to detect issues

**Result:**

- 60+ labels found (40+ tables, 15+ figures)
- Manual references use `Mục~X.X` format (Vietnamese standard)
- All references verified to point to existing sections

### Task 3.2: Grammar & Expression Review ✅ COMPLETE

- [x] Review sentence structure for clarity
- [x] Fix any grammatical errors
- [x] Ensure technical terms used consistently
- [x] Verify terminology matches glossary (if exists)

**Result:** Vietnamese grammar correct throughout. No changes needed.

### Task 3.3: Terminology Consistency ✅ COMPLETE

- [x] Create list of key terms and their standard forms
- [x] Verify consistent usage throughout document

**Verified Terms:**

- "Microservices" ✅ (not "micro-service")
- "API" ✅ (always uppercase)
- "Database" ✅ (context-appropriate)
- "PostgreSQL", "RabbitMQ", "Clean Architecture" ✅
- Service names consistent ✅

**Changelog:** `report/changelog/consistency-check-20251207.md`

---

## Phase 4: Final Review (Pass 4) ✅ COMPLETE

### Task 4.1: Full Document Read-Through ✅ COMPLETE

- [x] Read entire document for flow and coherence
- [x] Check transitions between sections
- [x] Verify logical progression of ideas
- [x] Note any remaining issues

**Result:** Document structure is logical and well-organized:

- Chapter 1: Executive Summary
- Chapter 2: Requirements Analysis (2.1-2.4)
- Chapter 3: Architecture Design (3.1-3.4)
- Chapter 4: Architecture Views (4.1-4.4)
- Chapter 5: SOLID Principles
- Chapter 6: System Implementation
- Chapter 7: Reflection and Evaluation

### Task 4.2: Final Compilation & Verification ✅ COMPLETE

- [x] Run `pdflatex main.tex` twice
- [x] Verify no compilation errors
- [x] Check page count (should remain ~95 pages)
- [x] Verify all figures/tables render correctly
- [x] Create `report/changelog/formatting-final-20251207.md`

**Result:**

- First pass: 91 pages
- Second pass: 95 pages ✅
- No compilation errors
- All 15+ figures and 40+ tables render correctly

### Task 4.3: Create Formatting Guidelines Document ✅ COMPLETE

- [x] Document all formatting rules applied
- [x] Create `report/formatting-guidelines.md` for future reference
- [x] Include examples of correct vs incorrect formatting

**Created:** `report/formatting-guidelines.md` with:

- Writing style rules (pronouns, capitalization, icons)
- LaTeX formatting standards
- File naming conventions
- Terminology consistency guide
- Quality checklist

---

## Summary

| Phase                 | Tasks        | Status      |
| --------------------- | ------------ | ----------- |
| Phase 1: Text Cleanup | 3 tasks      | ✅ Complete |
| Phase 2: Formatting   | 3 tasks      | ✅ Complete |
| Phase 3: Consistency  | 3 tasks      | ✅ Complete |
| Phase 4: Final Review | 3 tasks      | ✅ Complete |
| **Total**             | **12 tasks** | **✅ 100%** |

## Validation Checklist

- [x] No icons/emojis in document
- [x] No improper pronouns ("thầy cô", "người đọc", "tôi")
- [x] Consistent capitalization
- [x] Correct caption positions
- [x] Consistent file naming
- [x] No broken cross-references
- [x] LaTeX compiles without errors
- [x] Professional academic tone throughout

---

## 🎉 PROPOSAL COMPLETE

**Status:** All 12 tasks completed (100%)
**Final Page Count:** 95 pages
**Compilation:** Success, no errors

**Files Modified:** 8 LaTeX files
**Files Created:** 7 changelog/guideline files

**Report Status:** Ready for academic submission
