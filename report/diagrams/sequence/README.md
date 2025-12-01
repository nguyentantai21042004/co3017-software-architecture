# Sequence Diagrams Directory

Sequence diagrams showing key behavioral scenarios in the ITS system.

## Purpose

Store sequence diagram source files for the 5 key user workflows.

## Diagrams

1. **User Registration** (`user-registration.mmd`)
   - Flow: User → Auth Service → Database
   - Exported to: `report/images/user_registration_sequence.png`
   - Referenced in: `4.4_behavior_view.tex` Section 3.4.1

2. **Adaptive Content Delivery** (`adaptive-delivery.mmd`)
   - Flow: Client → Adaptive Engine → Learner Model → Content Service
   - Exported to: `report/images/adaptive_content_delivery_sequence.png`
   - Referenced in: `4.4_behavior_view.tex` Section 3.4.1

3. **Assessment Submission and Scoring** (`assessment-scoring.mmd`)
   - Flow: Client → Scoring Service → RabbitMQ → Learner Model Consumer
   - Exported to: `report/images/assessment_submission_and_scoring_sequence.png`
   - Referenced in: `4.4_behavior_view.tex` Section 3.4.1

4. **Real-time Feedback** (`real-time-feedback.mmd`)
   - Flow: Scoring Service → Event Bus → Client (WebSocket)
   - Exported to: `report/images/real_time_feedback_sequence.png`
   - Referenced in: `4.4_behavior_view.tex` Section 3.4.1

5. **Instructor Report Generation** (`instructor-report.mmd`)
   - Flow: Instructor → Reporting Service → Learner Model → Content Service
   - Exported to: `report/images/instructor_report_generation_sequence.png`
   - Referenced in: `4.4_behavior_view.tex` Section 3.4.1

## Verification

- Each sequence diagram must be verified against actual service code
- Check that lifelines match actual service names
- Verify message flows match actual API calls and events
