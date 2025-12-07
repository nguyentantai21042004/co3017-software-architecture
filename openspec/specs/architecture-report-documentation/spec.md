# Capability: Software Architecture Report Documentation

## Overview

This capability defines the requirements for maintaining comprehensive, accurate, and academically rigorous software architecture documentation for the ITS project. The documentation serves as the authoritative source for understanding architectural decisions, design principles, and system structure.

## MODIFIED Requirements

### Requirement: Complete Architecture Report Structure

The software architecture report MUST follow the established template format and include all required sections with appropriate depth and detail.

**Rationale:** Academic requirements and professional standards demand comprehensive documentation that covers all aspects of system architecture. Incomplete or superficial documentation fails to communicate design decisions effectively and limits the report's value for stakeholders.

**Acceptance Criteria:**
- Report includes all sections defined in template-format.md
- Each section meets minimum content requirements from scoring rubric
- Executive summary provides 1-2 page overview
- Reflection section is 3-4 pages with quantitative metrics
- All architecture views include appropriate diagrams

#### Scenario: Student Reviews Complete Report

**Given** a student needs to understand the ITS architecture  
**When** they read the software architecture report  
**Then** they can find:
- Clear project objectives and scope
- Complete functional and non-functional requirements
- Justified architecture decisions with alternatives considered
- Multiple architecture views (module, C&C, allocation, behavioral)
- SOLID principles application with code examples
- Reflection on design process with lessons learned

**And** all sections are present and complete per template requirements

#### Scenario: Professor Evaluates Report Against Rubric

**Given** a professor is grading the report using the scoring rubric  
**When** they evaluate each section  
**Then** the report achieves:
- Requirements Analysis: ≥ 14/15 points (includes domain model, stakeholder matrix, acceptance criteria)
- Architecture Design: ≥ 24/25 points (includes risk matrix, fitness functions, migration strategy)
- Architecture Views: ≥ 19/20 points (includes all required diagrams)
- SOLID Application: ≥ 19/20 points (includes UML diagrams, test examples)
- Reflection & Evaluation: ≥ 9/10 points (3-4 pages with metrics and ATAM)
- Documentation Quality: ≥ 9/10 points (executive summary, consistent formatting)

**And** overall score is ≥ 95/100

### Requirement: Implementation Verification and Traceability

All claims in the architecture report MUST be verifiable against the actual implementation, with clear traceability between documentation and code.

**Rationale:** Documentation that doesn't match reality is worse than no documentation. Traceability ensures the report accurately reflects the system as built and enables future maintenance.

**Acceptance Criteria:**
- mapping.md provides complete report → code → artifacts links
- All ERD diagrams match actual database schemas
- All sequence diagrams reflect actual service interactions
- All ADRs correspond to implemented decisions
- All SOLID examples reference actual code
- Verification status is documented for all claims

#### Scenario: Developer Verifies ERD Against Database Schema

**Given** a developer needs to verify the User Service ERD  
**When** they compare the ERD diagram to the actual database schema  
**Then** all tables shown in ERD exist in the database:
- Users table with columns: id, email, password_hash, status
- Roles table with columns: id, name, description
- Permissions table with columns: id, resource, action
- Users_Roles junction table
- Roles_Permissions junction table
- Learner_Profiles table with encrypted PII

**And** all relationships and cardinalities match the implementation  
**And** verification status is marked [VERIFIED] in mapping.md

#### Scenario: Reviewer Traces ADR to Implementation

**Given** a reviewer reads ADR-2 about event-driven architecture  
**When** they follow the mapping to the codebase  
**Then** they can find:
- RabbitMQ publisher in scoring-service at specified file path
- RabbitMQ consumer in learner-model at specified file path
- Event schema definitions
- Message queue configuration

**And** the implementation matches the ADR description  
**And** mapping.md links ADR-2 to specific code locations

### Requirement: LaTeX Formatting Compliance

All LaTeX source files MUST comply with the established formatting requirements to ensure consistent, professional presentation.

**Rationale:** Consistent formatting improves readability, ensures proper PDF rendering, and meets academic submission standards. Inconsistent formatting appears unprofessional and may cause compilation errors.

**Acceptance Criteria:**
- All .tex files follow latex-formatting-requirements.md
- No manual section numbering in headings
- All special characters properly escaped
- All tables use consistent formatting (vertical centering, justified text)
- All figures use consistent formatting (captions, labels, FloatBarrier)
- LaTeX compiles without errors
- PDF renders correctly with no overflow or clipping

#### Scenario: LaTeX Compilation Succeeds

**Given** all report content is complete  
**When** running `pdflatex main.tex`  
**Then** compilation completes without errors  
**And** PDF is generated successfully  
**And** all diagrams render at appropriate sizes  
**And** all cross-references resolve correctly  
**And** table of contents is complete and accurate

#### Scenario: Table Formatting is Consistent

**Given** a report section contains multiple tables  
**When** reviewing the LaTeX source  
**Then** all tables:
- Use `tabularx` or `longtable` environments
- Have vertically centered cells using `m{width}`
- Have justified text using `\justifying`
- Include `\caption{}` and `\label{}`
- Are followed by `\FloatBarrier` at section end

**And** all tables render consistently in the PDF

### Requirement: Diagram Completeness and Organization

All required architecture diagrams MUST be created, properly organized, and correctly placed in the report according to the diagram placement guidelines.

**Rationale:** Diagrams are essential for communicating architecture visually. Missing or poorly placed diagrams reduce report effectiveness. Organized source files enable future updates.

**Acceptance Criteria:**
- All diagrams from missmatch-erd.md recommendations exist
- Diagram source files organized in report/diagrams/ by type
- Exported images in report/images/ for LaTeX inclusion
- Each diagram has README.md explaining purpose and maintenance
- Diagrams placed in correct report sections
- All diagrams referenced in text with proper labels

#### Scenario: All Required Sequence Diagrams Exist

**Given** the behavior view section requires 5 key sequence diagrams  
**When** reviewing report/images/ directory  
**Then** the following diagrams exist:
- user_registration_sequence.png
- adaptive_content_delivery_sequence.png
- assessment_submission_and_scoring_sequence.png
- real_time_feedback_sequence.png
- instructor_report_generation_sequence.png

**And** each diagram has a source file in report/diagrams/sequence/  
**And** each diagram is referenced in 4.4_behavior_view.tex  
**And** each diagram renders correctly in the PDF

#### Scenario: ERD Diagrams are Properly Placed

**Given** the report follows missmatch-erd.md recommendations  
**When** reviewing the Module View section  
**Then** a new subsection "3.1.4 Data Persistence Design" exists  
**And** it contains three ERD diagrams:
- User Management Service ERD
- Content Service ERD
- Learner Model Service ERD

**And** each ERD is properly captioned and labeled  
**And** ERDs are not in Chapter 2 (requirements)  
**And** ERDs show microservice-specific schemas, not a monolithic database

### Requirement: Issue Tracking and Change Documentation

All gaps, discrepancies, and changes MUST be documented in structured issue files and changelogs to maintain transparency and traceability.

**Rationale:** Systematic issue tracking ensures no gaps are overlooked. Changelogs provide audit trail for academic integrity and enable understanding of how the report evolved.

**Acceptance Criteria:**
- report/issues/ contains gap analysis for each chapter
- report/issues/user-questions.md lists all questions requiring input
- report/changelog/ contains dated entries for all significant changes
- Each changelog entry includes: date, author, summary, rationale, changes made, verification
- All issues are either resolved or marked for user input
- Verification summary documents overall statistics

#### Scenario: Gap Analysis is Complete

**Given** Phase 1 analysis is complete  
**When** reviewing report/issues/ directory  
**Then** the following files exist:
- requirements-gaps.md
- architecture-design-gaps.md
- architecture-views-gaps.md
- solid-gaps.md
- implementation-gaps.md
- reflection-gaps.md

**And** each file documents:
- What is currently present
- What is missing per template and rubric
- Specific actions needed to close gaps
- Priority/impact of each gap

#### Scenario: Changes are Documented in Changelog

**Given** the executive summary was created  
**When** reviewing report/changelog/ directory  
**Then** a file executive-summary-YYYYMMDD.md exists  
**And** it contains:
- Date and author
- Summary of what was created
- Rationale for the content
- File location and structure
- Verification method (compilation, review)

**And** similar changelog entries exist for all other significant changes
