# Report Documentation Specification

## ADDED Requirements

### Requirement: MVP vs Target Architecture Labeling

The report SHALL clearly distinguish between MVP (implemented) and Target Architecture (planned) features throughout all diagrams and tables.

#### Scenario: ERD Diagram Labeling

- **WHEN** a reader views an ERD diagram
- **THEN** the diagram caption SHALL indicate whether it shows MVP implementation, Target Architecture, or both
- **AND** the reader SHALL understand which tables are currently implemented

#### Scenario: Sequence Diagram Labeling

- **WHEN** a reader views a sequence diagram
- **THEN** the diagram caption SHALL indicate whether the flow is implemented in MVP or planned for Target Architecture
- **AND** verified MVP diagrams SHALL be marked with implementation status

#### Scenario: Component Status Indication

- **WHEN** a reader views the component/service table
- **THEN** each component SHALL have an implementation status column
- **AND** the status SHALL be one of: ✅ MVP, ⚠️ Partial, ❌ Target

### Requirement: Implementation Status Section

The report SHALL include a dedicated section documenting the current implementation status with verification statistics.

#### Scenario: Verification Statistics Display

- **WHEN** a reader views the Implementation Status section
- **THEN** the section SHALL display verification statistics table showing:
  - Database Tables match rate (3/14 = 21%)
  - Sequence Diagrams match rate (2/5 = 40%)
  - ADRs match rate (5/10 = 50%)
  - SOLID Examples match rate (13/15 = 87%)
  - Overall match rate (23/44 = 52%)

#### Scenario: MVP Capabilities Documentation

- **WHEN** a reader wants to understand current MVP capabilities
- **THEN** the section SHALL list all implemented services and their functionality
- **AND** the section SHALL explain the MVP vs Target Architecture approach

### Requirement: ADR Implementation Status

The ADR summary table SHALL include an implementation status column indicating whether each architectural decision has been implemented.

#### Scenario: ADR Status Column

- **WHEN** a reader views the ADR summary table
- **THEN** each ADR SHALL have a status indicator:
  - ✅ Implemented - Decision fully implemented in MVP
  - ⚠️ Partial - Decision partially implemented
  - ❌ Planned - Decision documented but not yet implemented

#### Scenario: ADR Status Accuracy

- **WHEN** the ADR status is displayed
- **THEN** the status SHALL accurately reflect the Phase 3 verification results
- **AND** ADR-1, ADR-2, ADR-3, ADR-4, ADR-8 SHALL be marked as Implemented
- **AND** ADR-5 SHALL be marked as Partial
- **AND** ADR-6, ADR-7, ADR-9, ADR-10 SHALL be marked as Planned

### Requirement: MVP ERD Overview Diagram

The report SHALL include an MVP-specific ERD diagram showing only the implemented database tables.

#### Scenario: MVP ERD Content

- **WHEN** a reader views the MVP ERD diagram
- **THEN** the diagram SHALL show only the 3 implemented tables:
  - questions (Content Service)
  - submissions (Scoring Service)
  - skill_mastery (Learner Model Service)
- **AND** the diagram SHALL show relationships between tables

#### Scenario: MVP ERD Placement

- **WHEN** the MVP ERD is added to the report
- **THEN** it SHALL be placed in the Module View section (Chapter 4.1)
- **AND** it SHALL have a caption indicating "MVP Database Schema"

## MODIFIED Requirements

### Requirement: Report Accuracy

The report SHALL accurately represent the current state of implementation while also documenting the target architecture vision.

#### Scenario: Honest Documentation

- **WHEN** a reader reviews the report
- **THEN** they SHALL be able to distinguish between what is implemented and what is planned
- **AND** the report SHALL not mislead readers about current capabilities

#### Scenario: Maintained Academic Quality

- **WHEN** MVP/Target labels are added
- **THEN** the report SHALL maintain its academic quality (99.5/100 score)
- **AND** the report SHALL compile without LaTeX errors
- **AND** the page count SHALL remain between 93-100 pages

### Requirement: Rubric Traceability Matrix

The report SHALL include a comprehensive traceability matrix mapping all course rubric criteria to specific report sections and evidence.

#### Scenario: Task 1 Traceability (55%)

- **WHEN** a reviewer checks Task 1 criteria
- **THEN** the traceability matrix SHALL map each criterion to report sections:
  - ITS Context (5%) → Chapter 1, Section 2.1
  - Architecture Style Comparison (3%) → Section 3.2
  - Overall Architecture Design (20%) → Chapter 3, 4
  - UML Class Diagram (7%) → Section 4.1
  - SOLID Principles (15%) → Chapter 5
  - Future Extensibility (5%) → Section 3.2, Chapter 7

#### Scenario: Task 2 Traceability (30%)

- **WHEN** a reviewer checks Task 2 criteria
- **THEN** the traceability matrix SHALL map each criterion to code evidence:
  - Core Functionalities (15%) → sources/\*/ (4 microservices)
  - Bonus >1 module (+10%) → 4 microservices implemented
  - SOLID in Code (15%) → Chapter 5 code examples

#### Scenario: Task 3 Traceability (5%)

- **WHEN** a reviewer checks Task 3 criteria
- **THEN** the traceability matrix SHALL map each criterion:
  - Reflection Report (3%) → Chapter 7
  - Division of Work (2%) → Appendix/README

### Requirement: SOLID Dual Coverage

The report SHALL demonstrate SOLID principles in both documentation (Task 1 - 15%) and code implementation (Task 2 - 15%).

#### Scenario: SOLID Documentation Coverage

- **WHEN** a reviewer checks SOLID in documentation
- **THEN** Chapter 5 SHALL contain explanations for all 5 principles:
  - Single Responsibility (3%)
  - Open/Closed (3%)
  - Liskov Substitution (3%)
  - Interface Segregation (3%)
  - Dependency Inversion (3%)

#### Scenario: SOLID Code Coverage

- **WHEN** a reviewer checks SOLID in code
- **THEN** the report SHALL provide code examples demonstrating each principle:
  - SRP: Service separation (scoring, content, learner-model)
  - OCP: Interface-based design (repository/interface.go)
  - LSP: Contract compliance in services
  - ISP: Focused Repository interfaces
  - DIP: Constructor injection in main.go

### Requirement: Rubric Compliance Verification

The report documentation SHALL be verifiable against the official course rubric.

#### Scenario: Complete Rubric Coverage

- **WHEN** the report is evaluated against the course rubric
- **THEN** all Task 1 criteria (55%) SHALL be addressed
- **AND** all Task 2 criteria (30%) SHALL be demonstrated
- **AND** all Task 3 criteria (5%) SHALL be documented
- **AND** the traceability matrix SHALL provide clear evidence for each criterion

#### Scenario: Score Calculation

- **WHEN** calculating the expected score
- **THEN** Task 1 (55%) + Task 2 (30% + 10% bonus) + Task 3 (5%) SHALL equal 100%
- **AND** the report SHALL achieve maximum score for all documented criteria
