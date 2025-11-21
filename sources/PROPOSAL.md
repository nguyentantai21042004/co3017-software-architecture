### PHẦN 1: HẠ TẦNG & CƠ SỞ DỮ LIỆU (INFRASTRUCTURE LAYER)

#### Task 1.1: Thiết lập Docker Environment

  * **Mục tiêu:** Có môi trường chạy DB và Queue mà không cần cài cắm rác máy.
  * **File:** `docker-compose.yml`
  * **Nội dung code:**
    ```yaml
    version: '3.8'
    services:
      postgres:
        image: postgres:15-alpine
        environment:
          POSTGRES_USER: admin
          POSTGRES_PASSWORD: password
          POSTGRES_DB: its_main_db # Ta sẽ tạo schema riêng cho từng service trong này
        ports:
          - "5432:5432"
        volumes:
          - ./init-scripts:/docker-entrypoint-initdb.d # Để tự chạy script tạo bảng

      rabbitmq:
        image: rabbitmq:3-management-alpine
        ports:
          - "5672:5672"   # App connect
          - "15672:15672" # Admin UI
        environment:
          RABBITMQ_DEFAULT_USER: guest
          RABBITMQ_DEFAULT_PASS: guest
    ```
  * **Đầu ra:** Lệnh `docker-compose up -d` chạy thành công. Vào `localhost:15672` login được.

#### Task 1.2: Khởi tạo Database riêng biệt cho từng Service (chuẩn Microservices, mỗi service 1 database vật lý)

  * **Mục tiêu:** Mỗi service phải có 1 database vật lý RIÊNG, không dùng chung schema trong 1 database. Việc này đáp ứng đúng design microservices (mỗi service tự chủ dữ liệu, không share schema). Đồng thời, vẫn bảo đảm đủ các cột mới (`is_remedial`, `difficulty_level`) để hỗ trợ Adaptive.
  * **File:** `init-scripts/01-init.sql`
  * **Nội dung SQL:** Ví dụ dành cho PostgreSQL, hãy chạy trên 3 connections hoặc chỉ định lại database khi init (nếu dùng docker-compose thì thêm 3 service Postgres, mỗi service 1 DB cho 1 service).
    ```sql
    -- 1. Database cho CONTENT SERVICE (Java)
    CREATE DATABASE content_db;
    \c content_db

    CREATE TABLE questions (
        id SERIAL PRIMARY KEY,
        content TEXT NOT NULL,
        options JSONB,                     -- VD: ["A", "B", "C", "D"]
        correct_answer VARCHAR(255),
        skill_tag VARCHAR(50) NOT NULL,    -- VD: "math_algebra"
        difficulty_level INT DEFAULT 1,    -- 1: Dễ, 2: TB, 3: Khó
        is_remedial BOOLEAN DEFAULT FALSE, -- TRUE: Bài này dùng để dạy kèm/ôn tập
        created_at TIMESTAMP DEFAULT NOW()
    );

    -- Seed data để test Adaptive Flow
    INSERT INTO questions (content, correct_answer, skill_tag, difficulty_level, is_remedial) VALUES
    ('Bài toán khó: Giải phương trình X...', 'A', 'math_algebra', 3, FALSE), -- Bài chính (ID 1)
    ('Bài ôn tập: Nhắc lại quy tắc chuyển vế...', 'B', 'math_algebra', 1, TRUE); -- Bài chữa cháy (ID 2)

    -- 2. Database cho SCORING SERVICE (Go)
    CREATE DATABASE scoring_db;
    \c scoring_db

    CREATE TABLE submissions (
        id SERIAL PRIMARY KEY,
        user_id VARCHAR(50),
        question_id INT,
        submitted_answer VARCHAR(255),
        score_awarded INT,
        is_passed BOOLEAN, -- True nếu score >= 50%
        created_at TIMESTAMP DEFAULT NOW()
    );

    -- 3. Database cho LEARNER MODEL SERVICE (Go)
    CREATE DATABASE learner_db;
    \c learner_db

    CREATE TABLE skill_mastery (
        user_id VARCHAR(50),
        skill_tag VARCHAR(50),
        current_score INT DEFAULT 0, -- Điểm thông thạo hiện tại (0-100)
        last_updated TIMESTAMP DEFAULT NOW(),
        PRIMARY KEY (user_id, skill_tag)
    );
    -- Seed user chưa biết gì
    INSERT INTO skill_mastery (user_id, skill_tag, current_score) VALUES ('user_01', 'math_algebra', 10);
    ```
  * **Đầu ra:** Có đủ 3 database vật lý độc lập, mỗi DB tương ứng 1 service, mỗi bảng dữ liệu nằm đúng DB của service mình (KHÔNG dùng schema chung trong 1 DB to). Có dữ liệu mẫu sẵn để test Adaptive Engine.

-----

### PHẦN 2: CONTENT SERVICE (JAVA / SPRING BOOT)

#### Task 2.1: API Lấy bài học (Get & Recommendation)

  * **Công nghệ:** Java 17+, Spring Boot 3, Spring Data JPA.
  * **Logic cần code:**
    1.  Entity `Question` map với bảng `content_svc.questions`.
    2.  Repository có method:
          * `findById(Long id)`
          * `findFirstBySkillTagAndIsRemedial(String skill, Boolean isRemedial)` (Tìm bài theo skill và loại bài).
    3.  Controller `QuestionController`:
          * `GET /api/content/{id}`: Trả về chi tiết câu hỏi.
          * `GET /api/content/recommend`: Param `?skill=math&type=remedial`.
  * **Đầu ra API:**
      * Request: `GET /recommend?skill=math_algebra&type=remedial`
      * Response: JSON chứa câu hỏi ID 2 (Bài ôn tập).

-----

### PHẦN 3: SCORING SERVICE (GOLANG) - *THE PRODUCER*

#### Task 3.1: Logic Chấm điểm (Sync)

  * **Công nghệ:** Golang, Gin Framework.
  * **Logic:**
      * Endpoint: `POST /api/scoring/submit`.
      * Payload: `{ "user_id": "user_01", "question_id": 1, "answer": "C" }`.
      * Process:
          * Lấy đáp án đúng (tạm thời hardcode hoặc query DB content).
          * So sánh: Nếu sai -\> Score = 0. Nếu đúng -\> Score = 100.
          * Lưu vào bảng `scoring_svc.submissions`.
  * **Đầu ra:** Response JSON `< 500ms`: `{ "correct": false, "score": 0, "feedback": "Sai rồi" }`.

#### Task 3.2: Logic Bắn Event (Async)

  * **Công nghệ:** RabbitMQ (`amqp` library).
  * **Logic:**
      * Ngay sau khi lưu DB xong (ở Task 3.1), tạo một Goroutine (để không block response).
      * Publish message vào Exchange `its.events`.
      * Routing Key: `event.submission`.
      * Body:
        ```json
        {
          "event": "SubmissionCompleted",
          "user_id": "user_01",
          "skill_tag": "math_algebra",
          "score_obtained": 0,
          "timestamp": "2025-..."
        }
        ```

-----

### PHẦN 4: LEARNER MODEL SERVICE (GOLANG) - *THE CONSUMER*

#### Task 4.1: Xử lý Event (Worker)

  * **Công nghệ:** Golang, Goroutines.
  * **Logic:**
      * Khởi động Consumer lắng nghe Queue `learner.updates`.
      * Khi nhận message `SubmissionCompleted`:
          * Parse JSON.
          * Tính toán lại Mastery Score: `NewScore = (OldScore + ScoreObtained) / 2` (Công thức giả định).
          * Update bảng `learner_svc.skill_mastery`.
          * Log: "Updated user\_01 math\_algebra to 5".

#### Task 4.2: Internal API (Cho Adaptive Engine gọi)

  * **Mục tiêu:** Để thằng Adaptive Engine biết user đang giỏi hay dốt.
  * **Endpoint:** `GET /internal/learner/{user_id}/mastery`.
  * **Logic:** Query DB trả về điểm hiện tại.
  * **Đầu ra:** `{"skill": "math_algebra", "mastery_score": 5}`.

-----

### PHẦN 5: ADAPTIVE ENGINE (GOLANG) - *THE BRAIN*

*(Đây là phần WOW feature)*

#### Task 5.1: Logic Điều phối (Orchestration Logic)

  * **Công nghệ:** Golang.
  * **Endpoint:** `POST /api/adaptive/next-lesson`.
  * **Payload:** `{ "user_id": "user_01", "current_skill": "math_algebra" }`.
  * **Pseudocode Logic (Copy vào code):**
    ```go
    // 1. Gọi sang Learner Model Service để lấy điểm hiện tại
    mastery := httpGet("http://learner-service/internal/learner/user_01/mastery")

    var nextContentID int

    // 2. Logic thích ứng (Adaptive Rule)
    if mastery.Score < 50 {
        // User đang yếu -> Cần học Remedial (Bổ trợ)
        // Gọi sang Content Service lấy bài Remedial
        content := httpGet("http://content-service/api/content/recommend?skill=math_algebra&type=remedial")
        nextContentID = content.ID
    } else {
        // User giỏi -> Học bài tiếp theo (Standard)
        content := httpGet("http://content-service/api/content/recommend?skill=math_algebra&type=standard")
        nextContentID = content.ID
    }

    // 3. Trả về cho Client
    return JSON({ "next_lesson_id": nextContentID, "reason": "Based on your recent performance..." })
    ```

-----

### PHẦN 6: INTEGRATION & TESTING (GIAI ĐOẠN KHỚP LỆNH)

#### Task 6.1: API Gateway (Golang Reverse Proxy)

  * **Mục tiêu:** Client chỉ cần biết 1 địa chỉ `localhost:8080`.
  * **Cấu hình Route:**
      * `/content/*` -\> `http://localhost:8081` (Java)
      * `/scoring/*` -\> `http://localhost:8082` (Go Scoring)
      * `/adaptive/*` -\> `http://localhost:8083` (Go Adaptive)

#### Task 6.2: Kịch bản Test Demo (End-to-End)

Bạn thực hiện đúng trình tự này để quay video demo hoặc báo cáo:

1.  **Check trạng thái đầu:** User 01 có điểm Math = 10 (Rất thấp).
2.  **Hành động:** User làm bài tập khó (ID 1).
      * Call `POST /scoring/submit` với đáp án SAI.
3.  **Phản ứng hệ thống:**
      * Scoring trả về `score: 0`.
      * (Ngầm) RabbitMQ đẩy tin -\> Learner Model update điểm xuống còn 5.
4.  **Hành động tiếp theo (Adaptive):**
      * User bấm "Tiếp tục".
      * Call `POST /adaptive/next-lesson`.
5.  **Kết quả WOW:**
      * Adaptive Engine thấy điểm = 5 (\< 50).
      * Nó tự động trả về **Bài ID 2** (Là bài có `is_remedial = TRUE` mà ta đã seed ở Task 1.2).
      * *Client hiển thị:* "Chúng tôi thấy bạn gặp khó khăn. Hãy làm bài ôn tập này nhé."

-----

### TỔNG KẾT

Bạn có 6 phần việc rõ ràng. Nếu bạn làm một mình, tôi khuyên nên làm theo thứ tự:

1.  **Task 1.1 & 1.2** (Có DB để chứa dữ liệu).
2.  **Task 2.1 & 3.1** (Để có cái mà gọi API: Lấy đề & Chấm điểm).
3.  **Task 3.2 & 4.1** (Kết nối RabbitMQ - Phần khó nhất nhưng hay nhất).
4.  **Task 5.1** (Lắp não cho hệ thống - Phần Adaptive).

Bạn muốn tôi viết code chi tiết (full file) cho Task nào trước?