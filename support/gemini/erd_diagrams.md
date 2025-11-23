# Entity Relationship Diagrams (ERDs) for ITS

Dưới đây là mã Mermaid cho 3 ERD riêng biệt tương ứng với 3 Microservice chính, tuân thủ nguyên tắc Database-per-service.

## 1. User Management Service ERD (PostgreSQL)

Quản lý người dùng, phân quyền (RBAC) và thông tin cá nhân (PII).

```mermaid
erDiagram
    USERS {
        uuid id PK
        string email UK
        string password_hash
        string status
        timestamp created_at
    }
    ROLES {
        int id PK
        string name UK
        string description
    }
    PERMISSIONS {
        int id PK
        string resource
        string action
    }
    LEARNER_PROFILES {
        uuid user_id PK, FK
        string full_name
        string pii_data_encrypted
        jsonb preferences
    }
    
    USERS ||--o{ USERS_ROLES : has
    ROLES ||--o{ USERS_ROLES : assigned_to
    ROLES ||--o{ ROLES_PERMISSIONS : has
    PERMISSIONS ||--o{ ROLES_PERMISSIONS : granted_to
    USERS ||--|| LEARNER_PROFILES : has_profile
```

## 2. Content Service ERD (PostgreSQL + JSONB)

Quản lý nội dung khóa học, chương, bài học và ngân hàng câu hỏi. Sử dụng JSONB cho cấu trúc linh hoạt.

```mermaid
erDiagram
    COURSES {
        int id PK
        string title
        string description
        uuid instructor_id
        string status
    }
    CHAPTERS {
        int id PK
        int course_id FK
        string title
        int order_index
    }
    CONTENT_UNITS {
        int id PK
        int chapter_id FK
        string title
        enum type "VIDEO, TEXT, QUIZ"
        jsonb content_data "URL, text, or question_ids"
    }
    QUESTIONS {
        int id PK
        string content
        jsonb options
        string correct_answer
        string skill_tag
        int difficulty_level
        boolean is_remedial
    }
    TAGS {
        int id PK
        string name
        string type
    }

    COURSES ||--o{ CHAPTERS : contains
    CHAPTERS ||--o{ CONTENT_UNITS : contains
    CONTENT_UNITS }|--|{ TAGS : tagged_with
    QUESTIONS }|--|{ TAGS : tagged_with
```

## 3. Learner Model Service ERD (PostgreSQL / TimescaleDB)

Theo dõi tiến độ học tập, điểm số kỹ năng và lịch sử làm bài.

```mermaid
erDiagram
    SKILL_MASTERY {
        uuid learner_id PK
        string skill_tag PK
        float mastery_score
        timestamp last_updated
    }
    LEARNING_HISTORY {
        uuid id PK
        uuid learner_id FK
        int content_unit_id FK
        float score
        int time_spent_seconds
        timestamp completed_at
    }
    DIAGNOSTIC_RESULTS {
        uuid id PK
        uuid learner_id FK
        jsonb result_data
        timestamp created_at
    }

    SKILL_MASTERY }|--|| LEARNING_HISTORY : updated_by
```
