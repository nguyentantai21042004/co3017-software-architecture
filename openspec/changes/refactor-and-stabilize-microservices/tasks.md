# Tasks: Refactor and Stabilize Microservices

This work is broken down by service. The backend services should be stabilized before the client application.

1.  **Content Service (`content-service`)**
    - [x] Review core logic in the `Content Service` for bugs and correctness.
    - [x] Analyze and improve unit test coverage, focusing on content recommendation logic.
    - [x] Write integration tests for API endpoints.
    - [x] Refactor code to address any identified issues.

2.  **Scoring Service (`scoring-service`)**
    - [x] Review core logic in the `Scoring Service`, especially answer evaluation and event publishing.
    - [x] Analyze and improve unit test coverage.
    - [x] Write integration tests for the `/api/scoring/submit` endpoint and its interaction with RabbitMQ.
    - [x] Refactor code to address any identified issues.

3.  **Learner Model Service (`learner-model`)**
    - [ ] Review the logic for both the API and the consumer in the `Learner Model Service`.
    - [ ] Analyze and improve unit test coverage for mastery score calculations.
    - [ ] Write integration tests for the API and the RabbitMQ consumer.
    - [ ] Refactor code to address any identified issues.

4.  **Adaptive Engine Service (`adaptive-engine`)**
    - [ ] Review the orchestration logic in the `Adaptive Engine`.
    - [ ] Analyze and improve unit test coverage for the content recommendation algorithm.
    - [ ] Write integration tests that mock its dependencies (`content-service`, `learner-model-api`).
    - [ ] Refactor code to address any identified issues.

5.  **Client Application (`client`)**
    - [ ] Review all backend service changes (Content, Scoring, Learner Model, Adaptive Engine) to identify new APIs, updated endpoints, and interface changes.
    - [ ] Update client API service layer and TypeScript types/interfaces to match new backend contracts and ensure proper integration.
    - [ ] Review the frontend code for bugs, especially in state management and API service integration.
    - [ ] Add component tests for critical UI components.
    - [ ] Add end-to-end tests that simulate user interaction flows.
    - [ ] Refactor code to improve stability and user experience.

6.  **System-Wide Integration Testing**
    - [ ] Develop a suite of end-to-end tests that cover the primary user workflow:
        - User logs in.
        - User is presented with a question.
        - User submits an answer.
        - User's mastery score is updated.
        - User is presented with a new, adaptively chosen question.
    - [ ] Run these tests against a fully deployed system in a test environment.
