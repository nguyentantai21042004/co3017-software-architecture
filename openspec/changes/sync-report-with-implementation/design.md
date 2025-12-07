# Design: Sync Report with Implementation

## Context

Phase 3 verification revealed that the Software Architecture report describes both MVP implementation and Target Architecture without clear distinction. This creates potential confusion for readers about what is actually implemented vs. planned.

**Verification Statistics:**
| Category | MVP | Target | Match Rate |
|----------|-----|--------|------------|
| Database Tables | 3 | 11 | 21% |
| Sequence Diagrams | 2 | 3 | 40% |
| ADRs | 5 | 5 | 50% |
| SOLID Examples | 13 | 2 | 87% |
| **Overall** | **23** | **21** | **52%** |

## Goals / Non-Goals

### Goals

- Add clear MVP vs Target Architecture labels to all diagrams
- Create implementation status section documenting what's built
- Update ADR table with implementation status column
- Create MVP-specific ERD showing only implemented tables
- Maintain academic quality (99.5/100 score)

### Non-Goals

- Changing the actual implementation
- Removing Target Architecture documentation
- Rewriting existing content
- Adding new features to the system

## Decisions

### Decision 1: Visual Labeling Strategy

**What:** Use consistent visual labels throughout the report

**Labels:**

- `[MVP]` or `✅ MVP Implementation` - For verified, implemented features
- `[Target]` or `⚠️ Target Architecture` - For planned features
- `[Partial]` or `⚠️ Partial Implementation` - For partially implemented

**Why:** Clear, consistent labeling helps readers understand implementation status at a glance

**Alternatives Considered:**

1. Color coding only → Rejected (not accessible, doesn't work in B&W print)
2. Separate MVP and Target sections → Rejected (too much restructuring)
3. Footnotes only → Rejected (easy to miss)

### Decision 2: Implementation Status Table Format

**What:** Add standardized status table to relevant sections

**Format:**

```latex
\begin{table}[ht]
\centering
\begin{tabularx}{\textwidth}{|l|c|X|}
\hline
\textbf{Component} & \textbf{Status} & \textbf{Notes} \\
\hline
Content Service & ✅ MVP & questions table implemented \\
User Service & ❌ Target & Planned for Phase 3 \\
\hline
\end{tabularx}
\caption{Implementation Status}
\end{table}
```

**Why:** Tables provide quick reference for implementation status

### Decision 3: MVP ERD Approach

**What:** Create new MVP-specific ERD diagram alongside existing Target ERDs

**Approach:**

1. Keep existing ERDs (they show Target Architecture)
2. Add new `erd_mvp_overview.png` showing only 3 implemented tables
3. Add legend explaining MVP vs Target distinction

**Why:** Shows both current reality and future vision

**Alternatives Considered:**

1. Replace Target ERDs with MVP ERDs → Rejected (loses architectural vision)
2. Modify existing ERDs with overlays → Rejected (complex, hard to maintain)

### Decision 4: ADR Status Column

**What:** Add "Implementation Status" column to ADR summary table

**Status Values:**

- ✅ Implemented - Decision fully implemented in MVP
- ⚠️ Partial - Decision partially implemented
- ❌ Planned - Decision documented but not yet implemented

**Why:** ADRs document decisions, not implementation. Adding status clarifies what's built.

## Implementation Approach

### Phase 1: Diagram Labeling (2 hours)

1. **ERD Labels** (`4.1_module_view.tex`)

   - Add `[Target Architecture]` to User Service ERD caption
   - Add `[MVP + Target]` to Content Service ERD caption
   - Add `[MVP + Target]` to Learner Model ERD caption

2. **Sequence Diagram Labels** (`4.4_behavior_view.tex`)

   - Add `[MVP ✅]` to Adaptive Content Delivery
   - Add `[MVP ✅]` to Assessment Submission
   - Add `[Target]` to User Registration, Real-time Feedback, Instructor Report

3. **Component Labels** (`4.2_component_connector_view.tex`)
   - Add status column to component table
   - Mark 4 services as MVP, 3 as Target

### Phase 2: Implementation Status Section (2 hours)

1. **New Section** (`6_system_implementation.tex`)

   ```latex
   \subsection{Trạng Thái Triển Khai}

   \indentpar Báo cáo này mô tả cả kiến trúc MVP (đã triển khai)
   và Target Architecture (kế hoạch). Bảng sau tóm tắt trạng thái
   triển khai hiện tại:

   [Verification Statistics Table]

   \subsubsection{MVP Implementation}
   - 4 microservices hoạt động
   - 3 databases với 3 tables
   - RabbitMQ async communication
   - Adaptive learning flow verified

   \subsubsection{Target Architecture}
   - User Management Service
   - Authentication/Authorization
   - Reporting & Analytics
   - Real-time WebSocket feedback
   ```

2. **ADR Status Update** (`3.3_architecture_decision_records.tex`)
   - Add "Trạng Thái" column to ADR summary table
   - Update each ADR with implementation status

### Phase 3: MVP ERD Creation (1-2 hours)

1. **Create PlantUML Source** (`erd_mvp_overview.puml`)

   ```plantuml
   @startuml
   title MVP Database Schema (3 Tables)

   entity "questions" as q {
     * id : BIGSERIAL
     --
     content : TEXT
     options : JSONB
     correct_answer : TEXT
     skill_tag : VARCHAR(100)
     difficulty_level : INTEGER
     is_remedial : BOOLEAN
   }

   entity "submissions" as s {
     * id : BIGSERIAL
     --
     user_id : VARCHAR(50)
     question_id : BIGINT
     submitted_answer : VARCHAR(255)
     score_awarded : INTEGER
     is_passed : BOOLEAN
   }

   entity "skill_mastery" as sm {
     * user_id : VARCHAR(50)
     * skill_tag : VARCHAR(100)
     --
     current_score : INTEGER
     last_updated : TIMESTAMP
   }

   q ||--o{ s : "question_id"
   sm }o--|| q : "skill_tag"
   @enduml
   ```

2. **Generate PNG and Add to Report**

### Phase 4: Final Verification (30 minutes)

1. Compile LaTeX twice
2. Verify all labels render correctly
3. Check page count
4. Update verification documents

## Risks / Trade-offs

### Risk 1: Increased Page Count

- **Risk:** Adding status tables and labels may increase page count
- **Mitigation:** Use compact table formatting, inline labels where possible
- **Acceptable:** Up to 5 additional pages (93 → 98 pages)

### Risk 2: Visual Clutter

- **Risk:** Too many labels may distract from content
- **Mitigation:** Use subtle, consistent styling; avoid excessive emoji
- **Approach:** Use text labels `[MVP]` rather than emoji in formal sections

### Risk 3: Inconsistent Labeling

- **Risk:** Missing labels in some sections
- **Mitigation:** Create checklist of all diagrams/tables to label
- **Verification:** Review each chapter systematically

## Quality Metrics

### Success Criteria

1. ✅ All 3 ERD diagrams have MVP/Target labels
2. ✅ All 5 sequence diagrams have implementation status
3. ✅ Component table has status column
4. ✅ ADR table has status column
5. ✅ New "Implementation Status" section exists
6. ✅ MVP ERD diagram created
7. ✅ LaTeX compiles without errors
8. ✅ Page count: 93-98 pages

### Verification Checklist

- [ ] `4.1_module_view.tex` - ERD labels added
- [ ] `4.2_component_connector_view.tex` - Component status added
- [ ] `4.4_behavior_view.tex` - Sequence diagram labels added
- [ ] `3.3_architecture_decision_records.tex` - ADR status column added
- [ ] `6_system_implementation.tex` - Implementation status section added
- [ ] `erd_mvp_overview.puml` - MVP ERD created
- [ ] All changelog files created
- [ ] LaTeX compiles successfully

## Open Questions

1. **Label Language:** Should labels be in Vietnamese or English?

   - Recommendation: Vietnamese for consistency with report (`[MVP]`, `[Kiến trúc Mục tiêu]`)

2. **Emoji Usage:** Should we use emoji (✅, ❌) in formal report?

   - Recommendation: Use sparingly, prefer text labels in main content

3. **Legend Placement:** Where should the MVP/Target legend appear?
   - Recommendation: At the beginning of Chapter 4 (Architecture Views)

## Rubric Traceability Design (NEW)

### Course Rubric Overview

Based on `report/proposal/rubic.md`, the assignment is graded as follows:

| Component                       | Weight | Sub-components                                                                    |
| ------------------------------- | ------ | --------------------------------------------------------------------------------- |
| **Task 1: Architecture Design** | 55%    | Context (5%), Style (3%), Design (20%), UML (7%), SOLID (15%), Extensibility (5%) |
| **Task 2: Code Implementation** | 30%    | Core Functions (15%), SOLID in Code (15%), Bonus >1 module (+10%)                 |
| **Task 3: Documentation**       | 5%     | Reflection (3%), Division of Work (2%)                                            |
| **Presentation**                | 10%    | Demo (2%), Skills (3%), Slides (3%), Q&A (2%)                                     |

### Decision 5: Rubric Traceability Matrix

**What:** Create comprehensive mapping from rubric criteria to report sections

**Format:**

```markdown
| Rubric Criteria      | Weight | Report Section | Evidence            | Status |
| -------------------- | ------ | -------------- | ------------------- | ------ |
| 1.1 ITS Context      | 5%     | Ch.1, 2.1      | Context description | ✅     |
| 1.2 Style Comparison | 3%     | 3.2            | Comparison table    | ✅     |
| ...                  | ...    | ...            | ...                 | ...    |
```

**Why:**

- Ensures all grading criteria are addressed
- Provides quick reference for reviewers
- Identifies any gaps before submission

### Decision 6: SOLID Dual Coverage

**What:** Verify SOLID principles are covered in both:

1. Task 1 (15%): Documentation in Chapter 5
2. Task 2 (15%): Code examples with evidence

**Mapping:**
| Principle | Task 1 (Report) | Task 2 (Code) |
|-----------|-----------------|---------------|
| SRP (3%+3%) | Section 5.1 | Service separation |
| OCP (3%+3%) | Section 5.2 | Interface design |
| LSP (3%+3%) | Section 5.3 | Contract compliance |
| ISP (3%+3%) | Section 5.4 | Focused interfaces |
| DIP (3%+3%) | Section 5.5 | Constructor injection |

**Why:** SOLID accounts for 30% of total grade (15% + 15%)

### Rubric Traceability Implementation

#### Task 1: Software Architecture Design (55%)

| Criteria                              | Weight | Report Location        | Evidence                                     |
| ------------------------------------- | ------ | ---------------------- | -------------------------------------------- |
| **1.1 ITS Context**                   | 5%     | Chapter 1, Section 2.1 | Project scope, objectives, stakeholders      |
| **1.2 Architecture Style Comparison** | 3%     | Section 3.2            | Monolith vs Microservices comparison table   |
| **1.3 Overall Architecture Design**   | 20%    | Chapter 3, 4           | ADRs, Module View, C&C View, Allocation View |
| **1.4 UML Class Diagram**             | 7%     | Section 4.1            | Domain Model, ERD diagrams                   |
| **1.5 SOLID Principles**              | 15%    | Chapter 5              | 5 principles × 3% each                       |
| **1.6 Future Extensibility**          | 5%     | Section 3.2, Chapter 7 | Migration strategy, technical debt           |

#### Task 2: Code Implementation (30%)

| Criteria                     | Weight | Evidence Location  | Status                  |
| ---------------------------- | ------ | ------------------ | ----------------------- |
| **2.1 Core Functionalities** | 15%    | `sources/*/`       | 4 microservices working |
| **2.1b Bonus: >1 module**    | +10%   | `sources/*/`       | 4 modules implemented   |
| **2.2 SOLID in Code**        | 15%    | Chapter 5 examples | Code snippets verified  |

#### Task 3: Documentation (5%)

| Criteria                  | Weight | Report Location | Status                           |
| ------------------------- | ------ | --------------- | -------------------------------- |
| **3.1 Reflection Report** | 3%     | Chapter 7       | ATAM evaluation, lessons learned |
| **3.2 Division of Work**  | 2%     | Appendix/README | Team contributions               |

### Rubric Compliance Checklist

- [ ] Task 1.1: ITS Context (5%) - Chapter 1, 2.1
- [ ] Task 1.2: Style Comparison (3%) - Section 3.2
- [ ] Task 1.3: Architecture Design (20%) - Chapter 3, 4
- [ ] Task 1.4: UML Class Diagram (7%) - Section 4.1
- [ ] Task 1.5: SOLID Principles (15%) - Chapter 5
- [ ] Task 1.6: Future Extensibility (5%) - Section 3.2, Chapter 7
- [ ] Task 2.1: Core Functionalities (15%) - sources/\*/
- [ ] Task 2.1b: Bonus >1 module (+10%) - 4 microservices
- [ ] Task 2.2: SOLID in Code (15%) - Chapter 5 examples
- [ ] Task 3.1: Reflection Report (3%) - Chapter 7
- [ ] Task 3.2: Division of Work (2%) - Appendix

### Expected Score Calculation

| Task      | Max Score | Expected | Notes                  |
| --------- | --------- | -------- | ---------------------- |
| Task 1    | 55%       | 55%      | All criteria addressed |
| Task 2    | 30% + 10% | 40%      | Bonus achieved         |
| Task 3    | 5%        | 5%       | Complete               |
| **Total** | **90%**   | **100%** | With bonus             |

**Note:** Presentation (10%) is separate and not covered in this proposal.
