# ERD Directory

Entity-Relationship Diagrams for microservice databases.

## Purpose

Store ERD source files for the three main microservices with persistent data stores.

## Diagrams

1. **User Service ERD** (`user-service.drawio`)
   - Tables: Users, Roles, Permissions, Users_Roles, Roles_Permissions, Learner_Profiles
   - Exported to: `report/images/erd_user_service.png`
   - Referenced in: `4.1_module_view.tex` Section 3.1.4

2. **Content Service ERD** (`content-service.drawio`)
   - Tables: Courses, Chapters, Content_Units, Metadata_Tags, Content_Tags
   - Exported to: `report/images/erd_content_service.png`
   - Referenced in: `4.1_module_view.tex` Section 3.1.4

3. **Learner Model Service ERD** (`learner-model-service.drawio`)
   - Tables: Skill_Mastery, Learning_History, Diagnostic_Results
   - Exported to: `report/images/erd_learner_model_service.png`
   - Referenced in: `4.1_module_view.tex` Section 3.1.4

## Important Notes

- **Microservices Approach:** Each service has its own database - do NOT create a monolithic ERD
- **Verification:** ERDs must match actual database schemas in implementation
- **Updates:** If schema changes, update both source file and exported image
