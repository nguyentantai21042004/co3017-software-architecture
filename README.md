# CO3017 - Software Architecture

## Overview

This repository consolidates every artifact produced for the **CO3017 Software Architecture** course at HCMUT. It mixes lecture notes, assignment documents, architectural deliverables, LaTeX sources, and the full Intelligent Tutoring System (ITS) microservices implementation hosted under `src/`. Treat the root as the canonical workspace for both documentation and code.

## Repository Map

| Path | Contents | Notes |
| --- | --- | --- |
| `assignment/` | Assignment briefs, architectural reports, roadmap, and diagrams used in reviews. | Start here when preparing submissions or presentations. |
| `claude/`, `gemini/` | AI-generated supporting material (summaries, rubrics, meeting minutes). | Helpful for quick refreshers before critiques. |
| `references/` | Official course handouts, slides, and the assignment PDF. | Keep immutable; cite these sources in reports. |
| `report/` | LaTeX project for the final architecture report (all sections plus assets). | Build with `latexmk -pdf main.tex`. |
| `src/` | Complete ITS microservices codebase plus testing guide and Postman collection. | See the section below for details. |

> Tip: use `find . -maxdepth 2 -type d` if you need a quick textual snapshot of subfolders.

## Intelligent Tutoring System (ITS)

The `src/` directory contains a production-ready adaptive learning platform composed of four services:

| Service | Tech Stack | Port | Database | Purpose |
| --- | --- | --- | --- | --- |
| Content | Java 17, Spring Boot | 8081 | `content_db` | Manages questions, skills, and learning objects. |
| Scoring | Go 1.25, Gin | 8082 | `scoring_db` | Scores learner submissions and emits events. |
| Learner Model | Go 1.25, Gin | 8083 | `learner_db` | Maintains per-skill mastery curves. |
| Adaptive Engine | Go 1.25, Gin | 8084 | Stateless | Orchestrates recommendations using content + mastery data. |

RabbitMQ bridges the scoring and learner-model services through asynchronous events, enabling eventual consistency without tight coupling.

### Prerequisites

- Java 17 or later
- Go 1.25.4 or later
- PostgreSQL 15
- RabbitMQ 3.x with management UI
- Maven 3.8+

### Database Initialization

```
psql -U postgres -h localhost -p 5432 -f init-scripts/01-init-content-db.sql
psql -U postgres -h localhost -p 5432 -f init-scripts/02-init-scoring-db.sql
psql -U postgres -h localhost -p 5432 -f init-scripts/03-init-learner-db.sql
```

### Running the Services

```
cd src/content        && mvn spring-boot:run
cd src/scoring        && go run cmd/api/main.go
cd src/learner-model  && go run cmd/api/main.go
cd src/adaptive-engine && go run cmd/api/main.go
```

Health checks:

```
curl http://localhost:8081/actuator/health
curl http://localhost:8082/health
curl http://localhost:8083/health
curl http://localhost:8084/health
```

For end-to-end verification or regression suites, rely on `src/TESTING_GUIDE.md` and the `ITS_Microservices.postman_collection.json`.

## Documentation & Study Assets

- `src/README.md`: In-depth system description, architecture diagrams, troubleshooting guidance, and demo scenario.
- `src/TESTING_GUIDE.md`: Step-by-step manual and automated testing procedures.
- `assignment/report/*.md`: Narrative responses for each grading rubric item (characteristics, styles, ADRs, SOLID application, reflection).
- `assignment/diagrams/*.md` plus `report/images/*.png`: Architectural diagrams (module, component, deployment, sequence, and domain views).
- `report/latex-formatting-requirements.md`: Publishing checklist before exporting PDFs from LaTeX.

## Working Practices

- Each assignment or sub-deliverable should include its own `README.md` outlining requirements, assumptions, and execution steps.
- Keep LaTeX templates documented so teammates can rebuild the report on any machine.
- Store Architecture Decision Records (ADRs) and supporting notes in `references/` or under `assignment/report/5-architecture-decisions.md` for traceability.
- Organize study material chronologically or by topic to keep weekly sync meetings efficient.

## Key Milestones (Semester 251)

- Practical Assignment 1: Week 7 (06 Oct 2025)
- Practical Assignment 2: Week 8 (13 Oct 2025)
- Final Assignment Submission: 07 Dec 2025 at 23:59 (hard deadline)
- Project Presentation: Week 15 (08 Dec 2025)

---

This repository serves as the single source of truth for coursework, documentation, and the ITS reference implementation. Keep it tidy and up to date before each studio or review session.
