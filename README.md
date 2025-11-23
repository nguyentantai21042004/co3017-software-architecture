# CO3017 - Software Architecture

## Overview

This repository contains all artifacts for the **CO3017 Software Architecture** course at HCMUT, including architectural analysis, implementation, and documentation for an **Intelligent Tutoring System (ITS)** built with microservices architecture.

## Repository Structure

```
co3017-software-architecture/
├── analysis/              # Architectural Analysis & Research
│   ├── report/           # Analysis reports (Markdown)
│   ├── diagrams/         # Architecture diagrams (Mermaid)
│   ├── assignment.md     # Assignment requirements
│   └── ...
│
├── sources/              # Source Code & Implementation
│   ├── services/         # Microservices (Content, Scoring, Learner, Adaptive)
│   ├── client/           # Frontend (Next.js)
│   ├── infrastructure/   # Docker, scripts, configs
│   └── tests/            # Integration & system tests
│
├── report/               # LaTeX Report
│   ├── contents/         # LaTeX source files
│   ├── images/           # Diagrams & screenshots
│   └── main.tex          # Main LaTeX file
│
└── support/              # AI Assistance & References
    ├── claude/           # Claude AI outputs
    └── gemini/           # Gemini AI outputs
```

## Directory Details

### `analysis/` - Architectural Analysis & Research

Contains all architectural analysis, requirements, and design decisions:

- **`report/`**: Detailed analysis reports in Markdown format
  - `1-analyst.md` - Stakeholder analysis & requirements
  - `2-architecture-characteristics.md` - Architecture characteristics prioritization
  - `3-architecture-styles.md` - Architecture style comparison
  - `5-architecture-decisions.md` - Architecture Decision Records (ADRs)
  - `6-SOLID-principles.md` - SOLID principles application
  - `7-reflection-report.md` - Project reflection

- **`diagrams/`**: Architecture diagrams (Mermaid/PlantUML)
  - Sequence diagrams (use cases, workflows)
  - Deployment diagrams
  - Domain model diagrams

- **Other files**: `assignment.md`, `microservices.md`, `roadmap.md`, `system-comparison.md`

### `sources/` - Source Code & Implementation

Complete ITS microservices implementation with Docker support:

- **Services**:
  - `content-service/` - Java 17 + Spring Boot (Port 8081)
  - `scoring-service/` - Go 1.23 + Gin (Port 8082)
  - `learner-model-service/` - Go 1.23 + Gin (Port 8083)
  - `adaptive-engine/` - Go 1.23 + Gin (Port 8084)

- **Client**: `client/` - Next.js 15 + React 19 (Port 3000)

- **Infrastructure**:
  - `docker-compose.infra.yml` - Databases, RabbitMQ, MinIO
  - `docker-compose.yml` - Application services
  - `scripts/` - Database initialization & utilities

- **Documentation**: See `sources/README.md` for detailed setup instructions

### `report/` - LaTeX Report

Final architecture report in LaTeX format:

- **`contents/`**: LaTeX source files for each section
- **`images/`**: Diagrams and screenshots for the report
- **`main.tex`**: Main LaTeX file

**Build report**:
```bash
cd report
latexmk -pdf main.tex
```

### `support/` - AI Assistance & References

Supporting materials from AI tools:
- `claude/` - Claude AI generated content
- `gemini/` - Gemini AI generated content

## Quick Start

### Prerequisites

- **Docker Desktop** 4.0+ (recommended)
- **Java 17+**, **Go 1.23+**, **Node.js 20+** (for local development)
- **PostgreSQL 15**, **RabbitMQ 3.x**, **Maven 3.8+**

### Running the System

**Option 1: Docker (Recommended)**
```bash
cd sources
make setup          # Build, start infrastructure & services, init DBs
make health         # Check all services
```

**Option 2: Local Development**
```bash
cd sources
make dev            # Start only infrastructure (DB, RabbitMQ, MinIO)
# Then run each service locally in separate terminals
```

See `sources/README.md` for detailed instructions.

## Documentation

### Main Documentation

- **`sources/README.md`**: Complete system documentation, architecture, setup guide
- **`analysis/report/`**: Architectural analysis and design decisions
- **`report/latex-formatting-requirements.md`**: LaTeX formatting guidelines

### Service Documentation

Each service has its own README:
- `sources/content-service/README.md`
- `sources/scoring-service/README.md`
- `sources/learner-model-service/README.md`
- `sources/adaptive-engine/README.md`
- `sources/client/README.md`

## System Architecture

The ITS is built as a microservices architecture:

```
┌─────────────┐
│   Client    │ (Next.js)
└──────┬──────┘
       │
       ▼
┌─────────────────┐      ┌────────────────┐
│ Adaptive Engine │─────▶│ Learner Model  │
│   (Port 8084)   │      │  (Port 8083)   │
└────────┬────────┘      └────────────────┘
         │
         ▼
┌─────────────────┐
│ Content Service │
│   (Port 8081)   │
└─────────────────┘
         │
         ▼
┌─────────────────┐      ┌──────────────┐
│ Scoring Service │─────▶│  RabbitMQ    │
│   (Port 8082)   │      │  (Port 5672) │
└─────────────────┘      └──────────────┘
```

**Key Features**:
- **Event-Driven**: RabbitMQ for asynchronous communication
- **Database-per-Service**: Each service has its own database
- **Containerized**: Full Docker support with docker-compose
- **Scalable**: Microservices can be scaled independently

## Workflow

### 1. Analysis → `analysis/`
- Write all analysis in Markdown
- Create diagrams in `analysis/diagrams/`
- Document decisions in `analysis/report/`

### 2. Implementation → `sources/`
- Develop services in `sources/services/`
- Write technical docs in `sources/document/`
- Test in `sources/tests/`

### 3. Report → `report/`
- Convert analysis to LaTeX
- Build PDF: `latexmk -pdf main.tex`

## Useful Commands

### Docker Management (from `sources/`)
```bash
make help          # Show all available commands
make setup         # Complete setup (build, start, init, check)
make infra         # Start infrastructure only
make services      # Start application services only
make logs          # View all logs
make health        # Check service health
make test          # Run end-to-end test
```

### Database Operations
```bash
make db-init       # Initialize databases
make db-content    # Connect to content database
make db-backup     # Backup all databases
```

## Key Milestones (Semester 251)

- **Practical Assignment 1**: Week 7 (06 Oct 2025)
- **Practical Assignment 2**: Week 8 (13 Oct 2025)
- **Final Assignment Submission**: 07 Dec 2025 at 23:59
- **Project Presentation**: Week 15 (08 Dec 2025)

## Quick Links

- [System Documentation](./sources/README.md)
- [Architecture Analysis](./analysis/report/)
- [Docker Setup Guide](./sources/README.md#quick-start)
- [LaTeX Report](./report/)

## License

This repository is for academic purposes as part of CO3017 Software Architecture course at HCMUT.

---

**Last Updated**: 23 November 2025  
**Course**: CO3017 - Software Architecture  
**Institution**: Ho Chi Minh City University of Technology (HCMUT)
