# Project Context

## Purpose
This project is an Intelligent Tutoring System (ITS) built on a microservices architecture. Its primary goal is to provide a personalized learning experience by dynamically adapting to each student's knowledge level and performance. The system recommends learning content (questions) based on a student's skill mastery, scores their submissions, and updates their learning profile in real-time.

## Tech Stack
The system is composed of several services, each with its own technology stack:
- **Frontend:**
  - **Client:** Next.js, React, TypeScript, Tailwind CSS
- **Backend:**
  - **Adaptive Engine:** Go, Gin
  - **Content Service:** Java 17, Spring Boot, Maven
  - **Scoring Service:** Go, Gin
  - **Learner Model:** Go, Gin
- **Infrastructure:**
  - **Database:** PostgreSQL
  - **Messaging:** RabbitMQ
  - **Containerization:** Docker

## Project Documentation
In addition to the source code, this repository contains a `report` directory with a detailed project report written in LaTeX.
- **`report/main.tex`**: The main LaTeX file that compiles the entire report.
- **`report/contents/`**: This directory contains individual `.tex` files for each section of the report, covering topics such as project scope, stakeholder analysis, architectural decisions, design principles, and system implementation.
- **`report/images/`**: This directory contains all the images and diagrams used in the report.

## Project Conventions

### Code Style
- **General:** Follow the idiomatic conventions of each service's language and framework (e.g., Go standard practices, Spring Boot conventions, Next.js/React best practices).
- **Naming:** Use `snake_case` for database columns and `camelCase` or `PascalCase` for code symbols as is conventional in the respective languages.

### Architecture Patterns
- **Microservices:** The system is decomposed into single-purpose services, each with its own data store and API.
  - **`client`**: The user-facing web application.
  - **`content-service`**: Manages learning materials and questions.
  - **`scoring-service`**: Evaluates student answers.
  - **`learner-model`**: Tracks student knowledge and skill mastery.
  - **`adaptive-engine`**: Orchestrates the learning experience by recommending content.
- **Event-Driven:** Services communicate asynchronously using a message broker (RabbitMQ). For example, the `scoring-service` publishes a `submission.scored` event, which the `learner-model-consumer` processes to update mastery scores.
- **Orchestration:** The `adaptive-engine` acts as an orchestrator, communicating synchronously with the `learner-model` and `content-service` to fetch data and provide a response to the client.

### Testing Strategy
- **Health Checks:** Each service exposes a `/health` endpoint for basic availability monitoring.
- **Unit & Integration Testing:** Each service should have its own suite of unit and integration tests to ensure correctness and reliability.

### Git Workflow
[This section is to be filled by the project team. Describe your branching strategy (e.g., GitFlow, Trunk-Based Development) and commit message conventions.]

## Domain Context
The core domain of the ITS revolves around the following concepts:
- **Question:** A learning item stored in the `content_db`. Each question has content, options, a correct answer, a `skill_tag`, a `difficulty_level`, and a flag for `is_remedial`.
- **Skill:** A specific topic or area of knowledge, identified by a `skill_tag` (e.g., `math_algebra`).
- **Submission:** A record of a student's answer to a question, stored in `scoring_db`. It includes the score and whether the answer was correct.
- **Skill Mastery:** A score (0-100) representing a student's proficiency in a specific skill. This is stored in `learner_db` and is the primary input for the adaptive algorithm.
- **Adaptive Flow:**
  1. The `client` requests the next lesson from the `adaptive-engine`.
  2. The `adaptive-engine` queries the `learner-model` to get the student's `current_score` for a given `skill_tag`.
  3. Based on the score, the `adaptive-engine` requests an appropriate question (e.g., remedial, standard, or difficult) from the `content-service`.
  4. The student submits an answer to the `scoring-service`.
  5. The `scoring-service` evaluates the answer, stores the result, and publishes an event to RabbitMQ.
  6. The `learner-model-consumer` receives the event and updates the student's `skill_mastery` score.

## Important Constraints
- **Containerized Environment:** The entire system is designed to run within Docker containers. Service-to-service communication relies on the Docker network (`its-network`).
- **Service Dependencies:** Services have explicit startup dependencies defined in `docker-compose.yml`. For example, the `adaptive-engine` depends on the `content-service` and `learner-model-api`.
- **Configuration:** Each service is configured via environment variables, which are injected by Docker Compose. These include database connection strings, service URLs, and credentials.

## External Dependencies
- **Docker Images:** The system relies on official public Docker images for its infrastructure components:
  - `postgres:15-alpine`
  - `rabbitmq:3.13-management-alpine`
- **Service Communication:** Services communicate with each other via REST APIs over the internal Docker network. The client application communicates with the backend via public-facing ports mapped to the host.