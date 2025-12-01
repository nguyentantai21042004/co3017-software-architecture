# Diagrams Directory

This directory contains source files for all diagrams used in the Software Architecture report.

## Purpose

Maintain editable source files separate from exported images. This enables easy updates and version control of diagram sources.

## Directory Structure

- `erd/` - Entity-Relationship Diagrams (database schemas)
- `sequence/` - Sequence Diagrams (behavioral flows)
- `uml/` - UML Diagrams (domain model, class diagrams)
- `architecture/` - Architecture Diagrams (system decomposition, deployment, etc.)

## Workflow

1. **Create/Edit:** Work with source files (.drawio, .mmd, .puml) in appropriate subdirectory
2. **Export:** Export to PNG with appropriate resolution (200% scale, 10px border)
3. **Save:** Place exported PNG in `report/images/` for LaTeX inclusion
4. **Reference:** Update LaTeX files to reference the image in `report/images/`

## Tools

- **draw.io:** For ERDs, UML diagrams, architecture diagrams (.drawio files)
- **Mermaid:** For sequence diagrams, flowcharts (.mmd files)
- **PlantUML:** Alternative for UML diagrams (.puml files)

## Naming Conventions

- Source files: `descriptive-name.{drawio|mmd|puml}`
- Exported images: `descriptive_name.png` (use underscores, match LaTeX naming)
