UC-09 – Làm Bài tập và Assessment

Mục tiêu: Học viên thực hiện bài tập hoặc bài kiểm tra; hệ thống lưu kết quả để phục vụ bước chấm điểm.

```mermaid
graph TD
    subgraph ITS
        UC05["UC-05: Tạo Bài tập"]
        UC09["UC-09: Làm Bài tập & Assessment"]
        UC10["UC-10: Chấm điểm & Phản hồi Tức thì"]
    end
    Learner((Learner Học sinh))
    Instructor((Instructor Giảng viên))
    Instructor --> UC05
    UC05 -- "includes" --> UC09
    Learner --> UC09
    UC09 -- "extends" --> UC10
```

Giải thích:
- Instructor tạo nội dung (UC-05).
- Learner thực hiện bài tập (UC-09).
- Khi nộp bài, hệ thống mở rộng (extends) đến UC-10 để chấm điểm và phản hồi.

---

UC-10 – Chấm điểm và Phản hồi Tức thì

Mục tiêu: Tự động chấm điểm và trả phản hồi nhanh cho người học (< 1 s), đồng thời cập nhật Learner Model.

```mermaid
graph TD
    subgraph ITS
        UC09["UC-09: Làm Bài tập"]
        UC10["UC-10: Chấm điểm & Phản hồi Tức thì"]
        FR5["Auto-Grading Service"]
        FR6["Feedback Generator"]
        UC11["UC-11: Gợi ý Bài học Bù"]
    end
    Learner
    Instructor
    Learner --> UC09
    UC09 -- "extends" --> UC10
    UC10 -- "includes" --> FR5
    UC10 -- "includes" --> FR6
    UC10 -- "extends" --> UC11
    Instructor -. monitor .-> UC10
```

Giải thích:
- UC-10 bao gồm hai chức năng chính:
  - Auto-Grading Service: tính điểm.
  - Feedback Generator: tạo gợi ý/hints.
- Nếu kết quả kém, hệ thống mở rộng đến UC-11 để đề xuất bài học bù.

---

UC-11 – Gợi ý Bài học Bù (Remediation)

Mục tiêu: Khi học viên yếu kỹ năng nào đó, hệ thống đề xuất nội dung phù hợp để củng cố.

```mermaid
graph TD
    subgraph ITS
        UC10["UC-10: Chấm điểm & Phản hồi"]
        UC11["UC-11: Gợi ý Bài học Bù"]
        FR6["Feedback & Remediation Service"]
        FR7["Adaptive Learning Engine"]
    end
    Learner
    Instructor
    UC10 -- "extends" --> UC11
    UC11 -- "includes" --> FR6
    UC11 -- "includes" --> FR7
    Learner --> UC11
    Instructor -. review .-> UC11
```
- Kích hoạt khi người học sai liên tục hoặc mastery score < 0.6.
- Adaptive Learning Engine và Feedback Service cùng xử lý để sinh danh sách bài bù.
- Instructor có thể xem và xác nhận lộ trình ôn tập.

Tổng quan chuỗi Use Case 09 → 10 → 11

```mermaid
graph LR
    UC09["UC-09: Làm Bài tập"] --> UC10["UC-10: Chấm điểm & Phản hồi"]
    UC10 --> UC11["UC-11: Gợi ý Bài học Bù"]
```

Đây là vòng lặp học tập thích ứng của ITS:
1. Người học làm bài (UC-09) →
2. Hệ thống chấm điểm và phản hồi (UC-10) →
3. Nếu cần, đề xuất bài học bù (UC-11) → quay về lộ trình mới.

```mermaid
classDiagram
    %% User Management
    class User {
        +String UserID
        +String Email
        +String PasswordHash
    }
    class Role {
        +String RoleID
        +String RoleName
    }
    class Permission {
        +String PermissionID
        +String PermissionName
    }
    User "1" --> "1..*" Role
    Role "1" --> "1..*" Permission

    %% Learner
    class Learner {
        +String LearnerID
        +String UserID
        +String FullName
    }
    class LearnerProfile {
        +String ProfileID
        +String Goals
        +String LearningStyle
    }
    class ProgressRecord {
        +String ProgressID
        +String ContentUnitID
        +String Status
        +Number Score
    }
    Learner "1" --> "1" LearnerProfile
    Learner "1" --> "0..*" ProgressRecord

    %% Learner Model
    class LearnerModel {
        +String LearnerID
        +Map SkillMasteryScores
    }
    class SkillMasteryScore {
        +String SkillID
        +Float MasteryScore
        +Date LastPracticed
    }
    LearnerModel "1" --> "1..*" SkillMasteryScore

    %% Content
    class Course {
        +String CourseID
        +String Title
        +String InstructorID
    }
    class Chapter {
        +String ChapterID
        +String Title
    }
    class ContentUnit {
        +String ContentUnitID
        +String Type
        +String Data
    }
    class MetadataTag {
        +String TagID
        +String TagName
        +String TagType
    }
    class Assessment {
        +String AssessmentID
        +String Type
    }
    Course "1" --> "1..*" Chapter
    Chapter "1" --> "1..*" ContentUnit
    ContentUnit "1" --> "0..*" MetadataTag
    ContentUnit "1" --> "0..1" Assessment

    %% Adaptive Path
    class AdaptivePath {
        +String PathID
        +String LearnerID
        +Date GeneratedAt
    }
    class PathNode {
        +String NodeID
        +String ContentUnitID
        +Float RecommendationScore
    }
    AdaptivePath "1" --> "1..*" PathNode

    %% Cross-Aggregate references
    LearnerModel ..> Learner : references
    AdaptivePath ..> Learner : references
    ProgressRecord ..> ContentUnit : references
```