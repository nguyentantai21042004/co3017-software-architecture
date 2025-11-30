# Capability: System Integration Testing

## ADDED Requirements

### Requirement: The entire system MUST function correctly as a whole.

#### Scenario: End-to-End Workflow Test
- **Given** the complete microservices suite is deployed and running in a test environment.
- **When** an automated test script simulates a user completing a full learning loop:
  1.  Requesting a lesson.
  2.  Submitting a correct answer to an easy question.
  3.  Requesting another lesson.
  4.  Submitting an incorrect answer to a hard question.
  5.  Requesting a final lesson.
- **Then** the system behaves as expected at each step.
- **And** the final lesson recommended is a remedial one, reflecting the updated, lower mastery score.

#### Scenario: System Resilience Test
- **Given** the complete microservices suite is running.
- **When** a non-critical service (e.g., `content-service`) is temporarily unavailable.
- **When** the `adaptive-engine` attempts to retrieve content.
- **Then** it fails gracefully and returns an appropriate error message to the client, without crashing the entire system.
