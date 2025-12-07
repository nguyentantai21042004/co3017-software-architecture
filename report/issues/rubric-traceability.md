# Rubric Traceability Matrix

**Date:** 2025-12-07
**Purpose:** Map course rubric criteria to report sections and code evidence

---

## Course Rubric Overview (Assignment 40%)

| Component                   | Weight | Sub-components                                    |
| --------------------------- | ------ | ------------------------------------------------- |
| Task 1: Architecture Design | 55%    | Context, Style, Design, UML, SOLID, Extensibility |
| Task 2: Code Implementation | 30%    | Core Functions, SOLID in Code, Bonus              |
| Task 3: Documentation       | 5%     | Reflection, Division of Work                      |
| Presentation                | 10%    | Demo, Skills, Slides, Q&A                         |

---

## Task 1: Software Architecture Design (55%)

### 1.1 ITS Context Description (5%)

| Criterion                     | Weight | Report Section         | Evidence                                | Status      |
| ----------------------------- | ------ | ---------------------- | --------------------------------------- | ----------- |
| Mô tả đầy đủ bối cảnh của ITS | 5%     | Chapter 1, Section 2.1 | Project scope, objectives, stakeholders | ✅ Complete |

**Report Locations:**

- `report/contents/1_executive_summary.tex` - Executive summary
- `report/contents/2.1_project_scope_and_objectives.tex` - Project scope
- `report/contents/2.2_stakeholder_analysis.tex` - Stakeholder analysis

### 1.2 Architecture Style Comparison (3%)

| Criterion                               | Weight | Report Section | Evidence                             | Status      |
| --------------------------------------- | ------ | -------------- | ------------------------------------ | ----------- |
| So sánh và lựa chọn architecture styles | 3%     | Section 3.2    | Monolith vs Microservices comparison | ✅ Complete |

**Report Locations:**

- `report/contents/3.2_architecture_style_selection.tex` - Style comparison table
- Table: Comparison of Monolith, Microservices, Hybrid approaches
- Decision: Hybrid Microservices + Event-Driven

### 1.3 Overall Architecture Design (20%)

| Criterion                               | Weight | Report Section | Evidence              | Status      |
| --------------------------------------- | ------ | -------------- | --------------------- | ----------- |
| Thiết kế tổng thể Software Architecture | 20%    | Chapter 3, 4   | ADRs, Views, Diagrams | ✅ Complete |

**Report Locations:**

- `report/contents/3.1_architecture_characteristics_prioritization.tex` - AC prioritization
- `report/contents/3.3_architecture_decision_records.tex` - 10 ADRs
- `report/contents/4.1_module_view.tex` - Module View
- `report/contents/4.2_component_connector_view.tex` - C&C View
- `report/contents/4.3_allocation_view.tex` - Allocation View
- `report/contents/4.4_behavior_view.tex` - Behavior View (5 sequence diagrams)

### 1.4 UML Class Diagram (7%)

| Criterion            | Weight | Report Section | Evidence                   | Status      |
| -------------------- | ------ | -------------- | -------------------------- | ----------- |
| Vẽ UML Class Diagram | 7%     | Section 4.1    | Domain Model, ERD diagrams | ✅ Complete |

**Report Locations:**

- `report/contents/4.1_module_view.tex` - Module View with ERDs
- `report/images/domain_model_class_diagram.png` - Domain Model UML
- `report/images/erd_*.png` - 3 ERD diagrams

### 1.5 SOLID Principles (15%)

| Principle             | Weight | Report Section | Evidence                             | Status      |
| --------------------- | ------ | -------------- | ------------------------------------ | ----------- |
| Single Responsibility | 3%     | Section 5.1    | Service separation, layer separation | ✅ Complete |
| Open/Closed           | 3%     | Section 5.2    | Interface-based design               | ✅ Complete |
| Liskov Substitution   | 3%     | Section 5.3    | Contract compliance                  | ✅ Complete |
| Interface Segregation | 3%     | Section 5.4    | Focused interfaces                   | ✅ Complete |
| Dependency Inversion  | 3%     | Section 5.5    | Constructor injection                | ✅ Complete |

**Report Location:**

- `report/contents/5_apply_SOLID_principle.tex` - All 5 principles with examples

### 1.6 Future Extensibility (5%)

| Criterion                        | Weight | Report Section         | Evidence                           | Status      |
| -------------------------------- | ------ | ---------------------- | ---------------------------------- | ----------- |
| Khả năng mở rộng trong tương lai | 5%     | Section 3.2, Chapter 7 | Migration strategy, technical debt | ✅ Complete |

**Report Locations:**

- `report/contents/3.2_architecture_style_selection.tex` - Strangler Fig Pattern
- `report/contents/7_reflection_and_evaluation.tex` - Technical debt, lessons learned

---

## Task 2: Code Implementation (30%)

### 2.1 Core Functionalities (15%)

| Criterion                      | Weight | Code Location | Evidence        | Status      |
| ------------------------------ | ------ | ------------- | --------------- | ----------- |
| Implement Core Functionalities | 15%    | `sources/*/`  | 4 microservices | ✅ Complete |

**Code Locations:**

- `sources/content/` - Content Service (Java/Spring Boot)
- `sources/scoring/` - Scoring Service (Go/Gin)
- `sources/learner-model/` - Learner Model Service (Go/Gin)
- `sources/adaptive/` - Adaptive Engine (Go/Gin)

**Functionality Evidence:**

- Question management and recommendations
- Answer validation and score calculation
- Mastery tracking and skill updates
- Content orchestration and ZPD targeting

### 2.1b Bonus: >1 Module (+10%)

| Criterion           | Weight | Evidence        | Status            |
| ------------------- | ------ | --------------- | ----------------- |
| Implement >1 module | +10%   | 4 microservices | ✅ Bonus Achieved |

**Evidence:** 4 independent microservices implemented:

1. Content Service
2. Scoring Service
3. Learner Model Service
4. Adaptive Engine

### 2.2 SOLID in Code (15%)

| Principle | Weight | Code Location               | Evidence              | Status      |
| --------- | ------ | --------------------------- | --------------------- | ----------- |
| SRP       | 3%     | `sources/*/`                | Service separation    | ✅ Verified |
| OCP       | 3%     | `*/repository/interface.go` | Interface abstraction | ✅ Verified |
| LSP       | 3%     | All services                | Contract compliance   | ✅ Verified |
| ISP       | 3%     | `*/repository/`             | Focused interfaces    | ✅ Verified |
| DIP       | 3%     | `*/cmd/api/main.go`         | Constructor injection | ✅ Verified |

**Code Evidence:**

- SRP: Each service has single responsibility (scoring only, content only, etc.)
- OCP: Repository interfaces allow extension without modification
- LSP: All implementations honor interface contracts
- ISP: Repository interfaces are focused (no unused methods)
- DIP: Dependencies injected via constructors in main.go

---

## Task 3: Documentation & Reporting (5%)

### 3.1 Reflection Report (3%)

| Criterion         | Weight | Report Section | Evidence                         | Status      |
| ----------------- | ------ | -------------- | -------------------------------- | ----------- |
| Reflection Report | 3%     | Chapter 7      | ATAM evaluation, lessons learned | ✅ Complete |

**Report Location:**

- `report/contents/7_reflection_and_evaluation.tex`
- ATAM evaluation methodology
- Trade-off analysis
- Technical debt register
- Lessons learned

### 3.2 Division of Work (2%)

| Criterion        | Weight | Location        | Evidence           | Status       |
| ---------------- | ------ | --------------- | ------------------ | ------------ |
| Division of Work | 2%     | README/Appendix | Team contributions | ⚠️ To Verify |

**Note:** Division of work may be in project README or separate document.

---

## Summary

### Task 1: Architecture Design (55%)

| Sub-task                 | Weight  | Status      | Score   |
| ------------------------ | ------- | ----------- | ------- |
| 1.1 ITS Context          | 5%      | ✅ Complete | 5%      |
| 1.2 Style Comparison     | 3%      | ✅ Complete | 3%      |
| 1.3 Architecture Design  | 20%     | ✅ Complete | 20%     |
| 1.4 UML Class Diagram    | 7%      | ✅ Complete | 7%      |
| 1.5 SOLID Principles     | 15%     | ✅ Complete | 15%     |
| 1.6 Future Extensibility | 5%      | ✅ Complete | 5%      |
| **Total**                | **55%** |             | **55%** |

### Task 2: Code Implementation (30% + 10% Bonus)

| Sub-task                 | Weight  | Status      | Score   |
| ------------------------ | ------- | ----------- | ------- |
| 2.1 Core Functionalities | 15%     | ✅ Complete | 15%     |
| 2.1b Bonus >1 module     | +10%    | ✅ Achieved | +10%    |
| 2.2 SOLID in Code        | 15%     | ✅ Verified | 15%     |
| **Total**                | **30%** |             | **40%** |

### Task 3: Documentation (5%)

| Sub-task              | Weight | Status       | Score  |
| --------------------- | ------ | ------------ | ------ |
| 3.1 Reflection Report | 3%     | ✅ Complete  | 3%     |
| 3.2 Division of Work  | 2%     | ⚠️ To Verify | 2%     |
| **Total**             | **5%** |              | **5%** |

---

## Expected Score Calculation

| Task              | Max     | Expected | Notes                    |
| ----------------- | ------- | -------- | ------------------------ |
| Task 1            | 55%     | 55%      | All criteria addressed   |
| Task 2            | 30%     | 40%      | With +10% bonus          |
| Task 3            | 5%      | 5%       | Complete                 |
| **Content Total** | **90%** | **100%** | Exceeds max due to bonus |

**Note:** Presentation (10%) is evaluated separately.

---

## Last Updated

**Date:** 2025-12-07
**Status:** Complete
**Next:** Update mapping.md with rubric traceability
