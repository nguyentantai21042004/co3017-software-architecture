# UML Diagrams Directory

UML diagrams for domain modeling and class structures.

## Purpose

Store UML diagram source files for domain models and class diagrams.

## Diagrams

1. **Domain Model UML** (`domain-model.drawio`)
   - Shows: Entities, Value Objects, Aggregates, Relationships
   - Entities: User, Learner, Course, Chapter, ContentUnit, Question, Submission, SkillMastery
   - Exported to: `report/images/domain_model_uml.png`
   - Referenced in: `2.3_functional_requirements.tex` Section 2.3.3 (or 1.3.3)

2. **SOLID Principle Diagrams** (optional, if created)
   - Class diagrams illustrating each SOLID principle
   - Exported to: `report/images/solid_*.png`
   - Referenced in: `5_apply_SOLID_principle.tex`

## Notes

- Domain Model is conceptual (business logic), not physical (database)
- Shows aggregate boundaries and domain relationships
- Different from ERDs which show database tables
