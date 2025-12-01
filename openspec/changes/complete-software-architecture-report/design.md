# Design: Complete Software Architecture Report

## Overview

This change is a documentation-only effort to complete the Software Architecture report for the ITS project. The design focuses on creating a systematic, traceable approach to identify gaps, verify claims, and enhance content to achieve academic excellence (95-100/100 score).

## Key Design Decisions

### 1. Issue-Driven Approach

**Decision:** Use issue tracking files to document every gap, question, and discrepancy.

**Rationale:**
- Provides clear audit trail of what was found and how it was resolved
- Enables batching of user questions to minimize interruptions
- Creates reusable documentation for future report updates
- Separates concerns: analysis → resolution → validation

**Structure:**
```
report/issues/
├── README.md                          # Overview of issue tracking
├── requirements-gaps.md               # Chapter 2 gaps
├── architecture-design-gaps.md        # Chapter 3 gaps
├── architecture-views-gaps.md         # Chapter 4 gaps
├── solid-gaps.md                      # Chapter 5 gaps
├── implementation-gaps.md             # Chapter 6 gaps
├── reflection-gaps.md                 # Missing reflection section
├── user-questions.md                  # Questions requiring user input
├── quick-wins.md                      # Immediate actionable items
├── erd-verification.md                # ERD vs schema verification results
├── sequence-verification.md           # Sequence diagram verification
├── adr-verification.md                # ADR implementation verification
├── solid-verification.md              # SOLID example verification
├── verification-summary.md            # Overall verification statistics
└── rubric-validation.md               # Final score validation
```

### 2. Mapping-Based Traceability

**Decision:** Create a comprehensive mapping document linking report → code → artifacts.

**Rationale:**
- Ensures every claim in report can be traced to implementation
- Facilitates verification process
- Enables future maintenance (when code changes, know what report sections to update)
- Provides transparency for academic review

**Mapping Structure:**
```markdown
## Chapter 2: Requirements Analysis

| Report Section | Code Location | Artifacts | Status |
|----------------|---------------|-----------|--------|
| 2.3.1 User Stories | N/A (requirements) | Table 2.1 | [VERIFIED] |
| 2.3.2 Use Cases | sources/client/e2e/*.spec.ts | Figures 2.1-2.3 | [VERIFIED] |
| 2.3.3 Domain Model | sources/*/models/ | Figure 2.4 | [VERIFY] |
| 2.4 NFRs | openspec/project.md | Table 2.2 | [VERIFIED] |

## Chapter 3: Architecture Design

| Report Section | Code Location | Artifacts | Status |
|----------------|---------------|-----------|--------|
| 3.3 ADR-1 Microservices | sources/* (5 services) | N/A | [VERIFIED] |
| 3.3 ADR-2 Event-Driven | sources/*/rabbitmq/ | Figure 3.5 | [VERIFIED] |
...
```

### 3. Phased Execution Strategy

**Decision:** Execute in 5 distinct phases: Analysis → Content → Verification → Compliance → QA.

**Rationale:**
- **Phase 1 (Analysis):** Understand the problem completely before attempting solutions
- **Phase 2 (Content):** Create missing content based on analysis
- **Phase 3 (Verification):** Ensure report claims match reality
- **Phase 4 (Compliance):** Apply formatting and structural requirements
- **Phase 5 (QA):** Validate final quality

**Dependencies:**
```
Phase 1 (Analysis)
    ↓
Phase 2 (Content) ← Can start some tasks in parallel
    ↓
Phase 3 (Verification) ← Depends on Phase 2 completion
    ↓
Phase 4 (Compliance) ← Can start in parallel with Phase 3
    ↓
Phase 5 (QA) ← Depends on all previous phases
```

### 4. Diagram Organization Strategy

**Decision:** Organize diagrams by type in `report/diagrams/` with source files, separate from final outputs in `report/images/`.

**Rationale:**
- Source files (`.drawio`, `.mmd`, `.puml`) are editable and version-controllable
- Exported images (`.png`) are what LaTeX includes
- Clear separation enables easy updates (edit source → export → LaTeX auto-updates)
- README files in each directory document diagram purpose and maintenance

**Structure:**
```
report/diagrams/
├── README.md
├── erd/
│   ├── README.md
│   ├── user-service.drawio
│   ├── content-service.drawio
│   └── learner-model-service.drawio
├── sequence/
│   ├── README.md
│   ├── user-registration.mmd
│   ├── adaptive-delivery.mmd
│   └── ...
├── uml/
│   ├── README.md
│   └── domain-model.drawio
└── architecture/
    ├── README.md
    ├── system-decomposition.drawio
    ├── service-architecture.drawio
    └── ...

report/images/
├── erd_user_service.png              # Exported from diagrams/erd/
├── erd_content_service.png
├── domain_model_uml.png               # Exported from diagrams/uml/
└── ...
```

### 5. Changelog-Based Documentation

**Decision:** Create individual changelog files for each significant update.

**Rationale:**
- Provides detailed history of what changed and why
- Enables rollback if needed
- Documents decision rationale for future reference
- Supports academic integrity (shows work progression)

**Naming Convention:**
```
report/changelog/
├── README.md
├── executive-summary-20251201.md
├── reflection-expansion-20251201.md
├── risk-matrix-20251201.md
├── sequence-diagrams-20251201.md
└── ...
```

**Changelog Template:**
```markdown
# Changelog: [Component Name]

**Date:** YYYY-MM-DD
**Author:** [Name]
**Type:** [Addition|Modification|Deletion]

## Summary
Brief description of what changed.

## Rationale
Why this change was made.

## Changes Made
- Detailed list of changes
- File locations
- Line numbers if applicable

## Verification
How the change was verified (compilation, review, etc.)

## Related Issues
Links to issue files or user questions resolved.
```

### 6. Verification Strategy

**Decision:** Systematic verification of report claims against implementation using a three-step process.

**Process:**
1. **Extract Claims:** Read report section, identify all factual claims about implementation
2. **Locate Evidence:** Find corresponding code/config/documentation
3. **Compare & Document:** 
   - Match → Mark [VERIFIED] in mapping.md
   - Mismatch → Document in issues/[component]-verification.md
   - Missing → Mark [MISSING] and create issue

**Example Verification:**
```markdown
## ADR-2: Event-Driven Architecture

**Claim:** "Services communicate asynchronously via RabbitMQ for scoring updates"

**Evidence:**
- File: sources/scoring-service/src/publisher.go
- Lines: 45-67
- Code: PublishScoringEvent() publishes to "scoring.completed" exchange

**Status:** [VERIFIED]

---

**Claim:** "Learner model consumes events to update mastery scores"

**Evidence:**
- File: sources/learner-model/src/consumer.go
- Lines: 23-45
- Code: ConsumesScoringEvents() subscribes to "scoring.completed"

**Status:** [VERIFIED]
```

### 7. Quick Wins Prioritization

**Decision:** Identify and execute high-impact, low-effort tasks first.

**Rationale:**
- Builds momentum
- Shows progress quickly
- Reduces overall risk (if time runs short, high-impact items are done)
- Enables early user feedback on direction

**Quick Win Criteria:**
- Can be completed without user input
- Takes < 2 hours
- Improves rubric score by ≥ 0.5 points
- No dependencies on other tasks

**Examples:**
- Executive Summary (2 hours, +1 point)
- LaTeX formatting fixes (1 hour, +0.5 points)
- Cross-reference additions (1 hour, +0.5 points)
- Diagram reorganization (1 hour, +0.5 points)

## Technical Approach

### LaTeX Compilation Workflow

```bash
# Standard compilation
cd report/
pdflatex main.tex
pdflatex main.tex  # Run twice for references

# With bibliography (if added)
pdflatex main.tex
bibtex main
pdflatex main.tex
pdflatex main.tex
```

### Diagram Export Workflow

```bash
# For draw.io diagrams
# 1. Edit in draw.io desktop or web
# 2. Export as PNG with:
#    - Transparent background: No
#    - Border width: 10px
#    - Scale: 200% (for high resolution)
# 3. Save to report/images/

# For Mermaid diagrams
# 1. Edit .mmd file
# 2. Use mermaid-cli or online editor
# 3. Export as PNG
# 4. Save to report/images/

# For PlantUML diagrams
# 1. Edit .puml file
# 2. Run: plantuml diagram.puml
# 3. Move generated PNG to report/images/
```

### Verification Automation (Optional)

```bash
# Script to check if all referenced images exist
#!/bin/bash
cd report/
grep -r "includegraphics" contents/*.tex | \
  sed 's/.*{\(.*\)}.*/\1/' | \
  while read img; do
    if [ ! -f "$img" ]; then
      echo "Missing: $img"
    fi
  done

# Script to check if all labels are referenced
#!/bin/bash
cd report/
# Extract all labels
grep -r "\\label{" contents/*.tex | \
  sed 's/.*\\label{\(.*\)}.*/\1/' > /tmp/labels.txt

# Extract all references
grep -r "\\ref{" contents/*.tex | \
  sed 's/.*\\ref{\(.*\)}.*/\1/' > /tmp/refs.txt

# Find unreferenced labels
comm -23 <(sort /tmp/labels.txt) <(sort /tmp/refs.txt)
```

## Quality Metrics

### Success Metrics

1. **Completeness:** All sections from template-format.md present
2. **Accuracy:** ≥ 95% of claims verified against implementation
3. **Quality:** Rubric score ≥ 95/100
4. **Consistency:** 100% LaTeX formatting compliance
5. **Traceability:** 100% of report sections mapped to code/artifacts

### Validation Checklist

Before marking complete:
- [ ] LaTeX compiles without errors
- [ ] All diagrams render correctly in PDF
- [ ] All cross-references work
- [ ] All tables and figures have captions and labels
- [ ] Mapping.md has 100% coverage
- [ ] All issue files are resolved or marked for user input
- [ ] Rubric validation shows ≥ 95/100
- [ ] Walkthrough document is complete

## Risk Mitigation

### Risk: Report claims don't match implementation

**Mitigation:**
- Systematic verification in Phase 3
- Document all discrepancies in issue files
- Provide options: update report to match reality, or flag as technical debt

### Risk: Missing information blocks progress

**Mitigation:**
- Batch all questions in user-questions.md
- Prioritize questions (blocking vs. nice-to-have)
- Make reasonable assumptions where possible, document them
- Continue with tasks that don't depend on missing info

### Risk: Scope creep (trying to achieve perfect 100/100)

**Mitigation:**
- Target 95/100 as "good enough"
- Focus on high-impact improvements first (Quick Wins + Medium Efforts)
- Time-box each phase
- Final polish (95→100) is optional based on time available

### Risk: LaTeX compilation errors from changes

**Mitigation:**
- Make incremental changes
- Compile after each significant change
- Keep backup of working version
- Use version control (git) to enable rollback

## Future Enhancements

After completing this change, consider:

1. **Automated Report Generation:** Extract metrics from codebase automatically
2. **Living Documentation:** Keep report in sync with code via CI/CD
3. **Interactive Diagrams:** Use web-based diagram viewers
4. **Multi-Format Output:** Generate HTML/Markdown versions alongside PDF
5. **Template Reusability:** Create template for future projects
