# Changelog: AI Pipeline Data Flow Diagram

**Date:** 2025-12-07
**Author:** AI Assistant
**Type:** Addition

## Summary

Created comprehensive AI Pipeline Data Flow diagram showing the complete flow from student submission to adaptive content recommendation.

## Rationale

Task 2.13 required creating a data flow diagram for the AI/ML pipeline showing:

- Student submission → Scoring service
- Score → Learner model update
- Mastery score → Adaptive engine
- Adaptive engine → Content recommendation

## Changes Made

### 1. Created PlantUML Source File

- **File:** `report/images/ai_pipeline_dataflow.puml`
- **Content:** Detailed data flow diagram showing:
  - **Input Layer:** Student Submission, Question Metadata, Current Mastery
  - **Scoring Pipeline:** Answer Validation, Score Calculation, Performance Metrics
  - **Event Bus:** scoring.completed, learner.updated events
  - **Learner Model Pipeline:** Mastery Update (BKT/IRT), Skill Decay, Learning Velocity
  - **Adaptive Engine Pipeline:** Content Filtering, Difficulty Selection (ZPD), Path Optimization
  - **Output Layer:** Next Content, Learning Path, Progress Report
  - **Data Storage:** questions, skill_mastery, submissions tables

### 2. Updated LaTeX File

- **File:** `report/contents/4.2_component_connector_view.tex`
- **Changes:**
  - Added new subsubsection "Luồng Dữ liệu AI Pipeline"
  - Added figure reference to AI pipeline diagram
  - Added detailed description of each pipeline stage
  - Added table of AI/ML algorithms used (BKT, IRT, ZPD)

### 3. AI/ML Algorithms Documented

- **Bayesian Knowledge Tracing (BKT):** Probabilistic model for estimating student knowledge
- **Item Response Theory (IRT):** Statistical model for assessing student ability and question difficulty
- **Zone of Proximal Development (ZPD):** Educational theory for optimal learning zone

### 4. Event Payload Structure

Documented the event payload structure for `scoring.completed`:

```json
{
  "user_id": "uuid",
  "question_id": "uuid",
  "score": 0.85,
  "skill_tags": ["algebra", "equations"],
  "time_spent_ms": 45000,
  "accuracy": 0.85,
  "timestamp": "ISO8601"
}
```

## Verification

- PlantUML source file created successfully
- LaTeX file updated with new content
- Pending: Generate PNG from PlantUML and compile LaTeX

## Related Issues

- Task 2.13: Create Data Flow Diagram for AI Pipeline
- Addresses gap identified in `report/issues/architecture-views-gaps.md` (-1 point for missing AI pipeline data flow)

## Note

The PNG image needs to be generated from the PlantUML source file using:

```bash
plantuml report/images/ai_pipeline_dataflow.puml
```

Or using an online PlantUML renderer.
