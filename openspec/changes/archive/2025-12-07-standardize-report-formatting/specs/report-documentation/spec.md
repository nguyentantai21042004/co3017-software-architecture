# Report Documentation Spec Delta

## ADDED Requirements

### Requirement: Academic Writing Standards

The report SHALL follow academic writing standards for Vietnamese university submissions.

#### Scenario: Capitalization compliance

- **WHEN** reviewing any text in the report
- **THEN** only proper nouns, technology names, system names, and sentence beginnings are capitalized

#### Scenario: Pronoun usage compliance

- **WHEN** referring to the project team
- **THEN** the report uses "nhóm" (first person plural) or passive voice constructions
- **AND** does not use "tôi", "thầy cô", "người đọc", or "giảng viên"

#### Scenario: Icon-free content

- **WHEN** displaying status or emphasis
- **THEN** the report uses text descriptions or LaTeX symbols instead of Unicode icons/emojis

### Requirement: Caption Positioning Standards

The report SHALL follow standard academic conventions for caption placement.

#### Scenario: Table caption position

- **WHEN** a table is included in the report
- **THEN** the caption appears ABOVE the table content

#### Scenario: Figure caption position

- **WHEN** a figure is included in the report
- **THEN** the caption appears BELOW the figure content

### Requirement: Consistent Terminology

The report SHALL use consistent terminology throughout all sections.

#### Scenario: Technology name consistency

- **WHEN** referring to technologies and frameworks
- **THEN** the same capitalization and spelling is used throughout
- **AND** examples include: "Spring Boot" (not "SpringBoot"), "PostgreSQL" (not "Postgresql"), "RabbitMQ" (not "RabbitMq")

#### Scenario: Service name consistency

- **WHEN** referring to system services
- **THEN** the same naming convention is used throughout
- **AND** examples include: "Content Service", "Scoring Service", "Learner Model Service", "Adaptive Engine"

## MODIFIED Requirements

### Requirement: Report Quality Standards

The report SHALL maintain professional quality suitable for academic submission.

#### Scenario: LaTeX compilation success

- **WHEN** the report is compiled with pdflatex
- **THEN** no errors are produced
- **AND** all cross-references resolve correctly

#### Scenario: Formatting consistency

- **WHEN** reviewing the compiled PDF
- **THEN** spacing, indentation, and layout are consistent throughout
- **AND** no formatting anomalies are present

#### Scenario: Academic tone

- **WHEN** reading any section of the report
- **THEN** the writing maintains a professional, objective tone
- **AND** avoids informal language or colloquialisms
