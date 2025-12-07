# Capability: Software Architecture Report Documentation - Delta

## ADDED Requirements

### Requirement: Skywork AI Presentation Export

The architecture report MUST be exportable to a Markdown format optimized for Skywork AI automatic slide generation.

**Rationale:** Presentation slides are essential for communicating architecture decisions to stakeholders. Skywork AI can automatically generate professional slides from properly formatted Markdown, reducing manual effort and ensuring consistency.

**Acceptance Criteria:**

- `report/presentation/slides.md` contains complete slide content
- Markdown follows Skywork AI tag syntax for visual elements
- Content is condensed to presentation-appropriate length (≤2 sentences per bullet)
- All major report sections are represented in slides
- Visual placeholders are included for diagrams and charts
- README.md provides usage instructions

#### Scenario: Complete Slide Deck Generation

**Given** the LaTeX report is complete  
**When** the conversion process is executed  
**Then** `report/presentation/slides.md` contains:

- Title slide with project name and team info
- Section dividers for each major chapter
- Content slides for each subsection
- Visual placeholders using correct tag syntax
- Conclusion and Q&A slides

**And** the slide count is between 25-35 slides  
**And** all content is in Vietnamese with English technical terms

#### Scenario: Structure Mapping Compliance

**Given** a LaTeX section needs to be converted  
**When** applying the structure mapping rules  
**Then** the conversion follows:

- `\title{...}` → `# [SLIDE DECK]: ...`
- `\section{...}` → `## ...` (Section Divider)
- `\subsection{...}` → `### ...` (Content Slide)
- Paragraphs → Bullet points (summarized)
- Tables → `[CHART_TYPE: ...]` or Markdown tables
- Figures → `[PLACEHOLDER: ...]` with descriptions

**And** header hierarchy is maintained correctly

#### Scenario: Content Condensation Rules Applied

**Given** a paragraph of academic text needs conversion  
**When** applying content condensation rules  
**Then** the output:

- Contains only key points as bullet points
- Each bullet point has ≤ 2 sentences
- Focuses on keywords, results, and numbers
- Removes complex mathematical formulas
- Preserves meaning and context

**And** no important information is lost

#### Scenario: Visual Placeholder Tags Correct

**Given** a diagram or chart needs to be represented  
**When** inserting visual placeholders  
**Then** the correct tag syntax is used:

- `[PLACEHOLDER: description]` for standard images
- `[FULL_SLIDE_IMAGE_PLACEHOLDER: description]` for hero/background images
- `[CHART_TYPE: type, data summary]` for charts and graphs

**And** descriptions are detailed enough for AI image generation  
**And** chart types are appropriate (Bar/Pie/Line/Radar)

#### Scenario: Section Coverage Verification

**Given** the slide deck is complete  
**When** verifying section coverage  
**Then** all major sections are represented:

- Executive Summary (1-2 slides)
- Project Scope and Objectives (2-3 slides)
- Stakeholder Analysis (1-2 slides)
- Functional Requirements (2-3 slides)
- Non-Functional Requirements (1-2 slides)
- Architecture Characteristics (2-3 slides)
- Architecture Style Selection (2-3 slides)
- Architecture Decision Records (3-4 slides)
- Design Principles (1-2 slides)
- Module View (1-2 slides)
- Component & Connector View (2-3 slides)
- Allocation View (1-2 slides)
- Behavior View (2-3 slides)
- SOLID Principles (5 slides, one per principle)
- System Implementation (2-3 slides)
- Reflection and Evaluation (2-3 slides)
- Conclusion and Q&A (2 slides)

**And** no section is missing or empty

#### Scenario: Skywork AI Compatibility Test

**Given** the slides.md file is complete  
**When** the content is pasted into Skywork AI  
**Then** Skywork AI:

- Recognizes the slide structure from headers
- Processes all placeholder tags correctly
- Generates appropriate visual elements
- Creates a cohesive presentation flow

**And** no syntax errors are reported  
**And** all slides render correctly

### Requirement: Presentation Documentation

The presentation export MUST include documentation explaining how to use the generated Markdown with Skywork AI.

**Rationale:** Users need clear instructions to effectively use the generated Markdown file with Skywork AI. Documentation ensures consistent results and reduces support burden.

**Acceptance Criteria:**

- `report/presentation/README.md` exists
- README explains the purpose of slides.md
- README provides step-by-step Skywork AI usage instructions
- README documents the tag syntax used
- README includes troubleshooting tips

#### Scenario: User Follows README Instructions

**Given** a user wants to create slides from the Markdown  
**When** they follow the README.md instructions  
**Then** they can:

- Understand the purpose of each file
- Copy the slides.md content correctly
- Paste into Skywork AI interface
- Configure Skywork AI settings appropriately
- Generate the final presentation

**And** the process completes successfully  
**And** the generated slides match expectations
