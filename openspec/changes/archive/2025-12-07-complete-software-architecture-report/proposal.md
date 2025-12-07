# Change: Complete Software Architecture Report

## Why

The Software Architecture report for the Intelligent Tutoring System (ITS) project is currently at 87/100 (B+) according to the scoring rubric analysis. While the report has strong foundations in several areas (SOLID principles, ADRs, functional requirements), there are critical gaps that prevent it from achieving excellence:

1. **Missing Critical Content** (Current Score Impact: -13 points)
   - Reflection & Evaluation section is too brief (1 page vs. required 3-4 pages) - losing 3 points
   - Architecture Views lack complete diagrams (missing sequence diagrams, incomplete component diagrams) - losing 4 points
   - Requirements Analysis missing domain diagram and stakeholder matrix - losing 1.5 points
   - Architecture Design missing risk matrix and fitness functions - losing 3 points
   - Documentation Quality issues (no executive summary, inconsistent formatting) - losing 1.5 points

2. **Misalignment with Implementation**
   - Report describes architectural decisions but needs verification against actual codebase
   - ERD diagrams exist but need confirmation they match database schemas in microservices
   - Sequence diagrams need validation against actual service interactions

3. **Template Compliance Issues**
   - Current content doesn't fully follow `template-format.md` structure
   - LaTeX formatting needs to comply with `latex-formatting-requirements.md`
   - Diagram placement doesn't match recommendations in `missmatch-erd.md`

4. **Academic Requirements**
   - Scoring rubric (`scoring_rubic.md`) identifies specific gaps for achieving A+ grade
   - Missing quantitative metrics and evaluation methodology (ATAM)
   - Insufficient cross-referencing between report sections and implementation

This proposal creates a systematic approach to:
- Identify all gaps between current report and requirements
- Verify alignment between report claims and actual implementation
- Create missing artifacts (diagrams, sections, documentation)
- Ensure template and formatting compliance
- Achieve target score of 95-100/100 (A to A+)

## What Changes

### 1. Gap Analysis and Issue Tracking
- Create comprehensive issue files documenting every discrepancy between report and requirements
- Map each report section to implementation code for verification
- Identify missing diagrams, sections, and content based on scoring rubric
- Document all questions requiring clarification or user input

### 2. Content Completion
- **Executive Summary** (NEW) - 1-2 page overview of entire report
- **Reflection & Evaluation** (EXPAND) - Expand from 1 page to 3-4 pages with:
  - Quantitative metrics (before/after measurements)
  - ATAM evaluation methodology
  - Technical debt analysis
  - Lessons learned with depth
- **Architecture Views** (ENHANCE) - Add missing diagrams:
  - 5 comprehensive sequence diagrams for key scenarios
  - Complete component diagram showing all services with interfaces
  - Enhanced deployment diagram with infrastructure details
  - Data flow diagrams for AI pipeline
- **Requirements Analysis** (ENHANCE) - Add:
  - Domain Model UML diagram
  - Stakeholder influence/interest matrix
  - Acceptance criteria for user stories
- **Architecture Design** (ENHANCE) - Add:
  - Risk matrix with probability/impact analysis
  - Fitness functions for architecture characteristics
  - Migration strategy from MVP to microservices
  - Cost-benefit analysis (TCO)

### 3. Implementation Verification
- Create mapping document linking report sections to actual code
- Verify ERD diagrams match database schemas in:
  - User Management Service (PostgreSQL)
  - Content Service (PostgreSQL + JSONB)
  - Learner Model Service (PostgreSQL/Time-scale)
- Validate sequence diagrams against actual service interactions
- Confirm ADRs reflect actual implementation decisions

### 4. Template and Format Compliance
- Restructure report sections to match `template-format.md`
- Apply LaTeX formatting rules from `latex-formatting-requirements.md`
- Reorganize diagrams according to `missmatch-erd.md` recommendations
- Ensure consistent cross-referencing and labeling

### 5. Documentation and Tracking
- Create changelog files for all significant updates
- Maintain mapping.md showing report ↔ code ↔ artifacts relationships
- Document all assumptions and confirmations
- Create placeholder files for missing artifacts with clear specifications

## Impact

- **Affected Specs:** None (this is a documentation-only change)
- **Affected Code:** None (verification only, no code changes)
- **Affected Documentation:**
  - `report/contents/*.tex` - All 16 existing files will be reviewed and enhanced
  - `report/images/` - New diagrams will be added (5+ sequence diagrams, domain model, etc.)
  - NEW: `report/issues/` - Issue tracking files for gaps and confirmations
  - NEW: `report/diagrams/` - Organized diagram sources and specifications
  - NEW: `report/changelog/` - Update logs for all changes
  - NEW: `report/mapping.md` - Report-to-implementation mapping
- **Breaking Changes:** None
- **Dependencies:** 
  - Requires access to all microservice codebases for verification
  - May require user input for missing information or clarifications

## Implementation Approach

### Phase 1: Analysis and Planning (Quick Wins - 1 day)
1. Create issue tracking structure (`report/issues/`)
2. Analyze each report section against template and rubric
3. Create `mapping.md` with initial report ↔ code mappings
4. Document all questions and missing information
5. Identify quick wins (executive summary, formatting fixes)

### Phase 2: Content Gap Filling (High Impact - 2-3 days)
1. Write Executive Summary (1-2 pages)
2. Expand Reflection & Evaluation section (3-4 pages with metrics)
3. Create missing sequence diagrams (5 diagrams)
4. Add domain model UML diagram
5. Create risk matrix and mitigation strategies
6. Define fitness functions for architecture characteristics

### Phase 3: Implementation Verification (Medium Effort - 2 days)
1. Verify ERD diagrams against actual database schemas
2. Validate sequence diagrams against service code
3. Confirm ADRs reflect actual decisions
4. Update `mapping.md` with verified links
5. Document any discrepancies in changelog

### Phase 4: Template Compliance (Final Polish - 1 day)
1. Restructure sections to match template
2. Apply LaTeX formatting requirements
3. Reorganize diagram placement
4. Add cross-references throughout
5. Final consistency check

### Phase 5: Quality Assurance (Validation - 1 day)
1. Compile LaTeX and fix any errors
2. Verify all diagrams render correctly
3. Check all cross-references work
4. Validate against scoring rubric
5. Create final walkthrough document

## Success Criteria

1. ✅ All sections from `template-format.md` are present and complete
2. ✅ Report score improves from 87/100 to 95-100/100 based on `scoring_rubic.md`
3. ✅ All diagrams from `missmatch-erd.md` recommendations are created and properly placed
4. ✅ LaTeX compiles without errors and follows `latex-formatting-requirements.md`
5. ✅ `mapping.md` provides complete traceability between report and implementation
6. ✅ All issues documented in `report/issues/` are resolved or marked for user input
7. ✅ Changelog files document all significant changes with dates and rationale
8. ✅ Executive Summary provides clear 1-2 page overview
9. ✅ Reflection section is 3-4 pages with quantitative metrics and ATAM evaluation
10. ✅ All 5 key sequence diagrams are created and validated against implementation

## Risks and Mitigations

- **Risk:** Report claims may not match actual implementation
  - **Mitigation:** Systematic verification against codebase, document discrepancies in issues/
- **Risk:** Missing information may require user input, blocking progress
  - **Mitigation:** Create comprehensive issue files early, batch questions for user
- **Risk:** Diagram creation may be time-consuming
  - **Mitigation:** Prioritize high-impact diagrams first, use existing tools and templates
- **Risk:** LaTeX compilation errors from formatting changes
  - **Mitigation:** Make incremental changes, test compilation frequently
- **Risk:** Scope creep - trying to achieve perfect 100/100 may be inefficient
  - **Mitigation:** Focus on high-impact improvements first (Quick Wins + Medium Efforts target 95+)

## Open Questions

1. **Verification Scope:** Should we verify ALL claims in the report against implementation, or focus on critical architectural decisions?
2. **Diagram Tools:** What tool should be used for creating new diagrams (Mermaid, draw.io, PlantUML)?
3. **Missing Information:** For gaps in knowledge (e.g., cost analysis, specific metrics), should we:
   - Make reasonable assumptions and document them?
   - Request specific information from user?
   - Mark as "future work" if not critical?
4. **Implementation Discrepancies:** If report claims don't match implementation, should we:
   - Update report to match reality?
   - Flag as technical debt?
   - Both?
5. **Reflection Content:** What specific metrics should be included in the expanded Reflection section?
   - Code coverage percentages?
   - Performance benchmarks?
   - Development timeline?

## Related Work

- This builds on the existing report structure in `report/contents/` (16 .tex files)
- Uses existing diagrams in `report/images/` (20 PNG files) as foundation
- References the ITS microservices implementation in `sources/` directory
- Aligns with OpenSpec project documentation in `openspec/project.md`
- Follows academic requirements from `report/SA_hk251_Assignment.pdf`
