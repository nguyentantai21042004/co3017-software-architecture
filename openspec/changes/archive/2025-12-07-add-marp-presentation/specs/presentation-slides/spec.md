# Presentation Slides Capability

## Overview

Capability này cung cấp bài thuyết trình (presentation) cho dự án Intelligent Tutoring System (ITS) sử dụng Marp - một công cụ tạo slides từ Markdown.

## ADDED Requirements

### Requirement: Marp Presentation Structure

The presentation system SHALL have a clear folder structure with necessary files to build and customize slides.

#### Scenario: Folder structure exists

- **WHEN** user navigates to `presentation/` folder
- **THEN** folder SHALL contain files: `slides.md`, `theme.css`, `README.md`
- **AND** folder `images/` SHALL exist for diagrams and screenshots

#### Scenario: README provides usage instructions

- **WHEN** user reads `presentation/README.md`
- **THEN** file SHALL contain Marp CLI installation instructions
- **AND** file SHALL contain commands to build slides (PDF, HTML, PPTX)
- **AND** file SHALL contain preview instructions

---

### Requirement: Slide Content Coverage

The presentation SHALL cover all major sections of the software architecture report according to the defined structure.

#### Scenario: All major sections included

- **WHEN** presentation is built
- **THEN** slides SHALL include these sections:
  - Title & Executive Summary (3-4 slides)
  - Requirements Analysis (4-5 slides)
  - Architecture Design (4-5 slides)
  - Architecture Decision Records (2-3 slides)
  - Architecture Views (3-4 slides)
  - SOLID Principles (2-3 slides)
  - Implementation (2-3 slides)
  - Demo & Code Showcase (4-6 slides)
  - Reflection & Future (2-3 slides)
  - Conclusion & Q&A (2 slides)

#### Scenario: Total slide count appropriate

- **WHEN** presentation is fully built
- **THEN** total slide count SHALL be between 28-38 slides
- **AND** presentation duration SHALL fit within 15-20 minutes

---

### Requirement: Visual Design Consistency

The presentation SHALL follow design guidelines and color palette as defined.

#### Scenario: Color palette applied

- **WHEN** slides are rendered
- **THEN** colors SHALL follow the palette:
  - Primary: Navy Blue (#1E3A8A)
  - Secondary: Blue (#3B82F6)
  - Accent: Green (#10B981)
  - Warning: Orange (#F59E0B)
  - Background: Light Gray (#F8FAFC) or White (#FFFFFF)
  - Text: Dark Gray (#1F2937)

#### Scenario: Typography consistent

- **WHEN** slides are rendered
- **THEN** headings SHALL use sans-serif bold font
- **AND** body text SHALL have size 18-24pt
- **AND** code blocks SHALL use monospace font

#### Scenario: Layout readable

- **WHEN** slides are displayed
- **THEN** each slide SHALL have maximum 5-6 bullet points
- **AND** white space SHALL be sufficient for readability from distance
- **AND** visual hierarchy SHALL be clear

---

### Requirement: Marp Syntax Compliance

The `slides.md` file SHALL use correct Marp syntax to ensure successful build.

#### Scenario: Valid Marp frontmatter

- **WHEN** file `slides.md` is parsed
- **THEN** file SHALL have frontmatter with `marp: true`
- **AND** frontmatter SHALL specify theme and pagination settings

#### Scenario: Slide separation correct

- **WHEN** slides are parsed
- **THEN** each slide SHALL be separated by `---`
- **AND** section dividers SHALL use `<!-- _class: lead -->` directive

#### Scenario: Images and diagrams referenced correctly

- **WHEN** slides contain images
- **THEN** image paths SHALL be relative to `presentation/` folder
- **AND** Marp image syntax SHALL be used correctly (e.g., `![bg](image.jpg)`)

---

### Requirement: Build and Export Support

The presentation SHALL be buildable into multiple output formats.

#### Scenario: PDF export works

- **WHEN** user runs `marp slides.md -o slides.pdf`
- **THEN** PDF file SHALL be created successfully
- **AND** PDF SHALL preserve formatting and images

#### Scenario: HTML export works

- **WHEN** user runs `marp slides.md -o slides.html`
- **THEN** HTML file SHALL be created successfully
- **AND** HTML SHALL be viewable in browser

#### Scenario: Preview mode works

- **WHEN** user runs `marp -p slides.md`
- **THEN** preview server SHALL start successfully
- **AND** slides SHALL display correctly in browser
