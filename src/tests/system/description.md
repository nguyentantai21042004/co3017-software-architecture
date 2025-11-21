# BACKEND TEST SPECIFICATIONS: ADAPTIVE REMEDIATION SYSTEM

**System Scope:** Scoring Service (Go), Learner Model Service (Go), Adaptive Engine (Go), Content Service (Java).
**Architecture Style:** Hybrid Microservices + Event-Driven (RabbitMQ).

-----

## PART 1: SCORING SERVICE (Producer)

*Service chịu tải cao, xử lý logic chấm điểm đồng bộ và phát tin bất đồng bộ.*

### 1.1 Functional Cases (Happy Path)

#### TC\_SCORING\_01: Submit Correct Answer & Publish Event

  * **Description:** Người dùng nộp đúng, nhận điểm, event được bắn đi.
  * **Input (POST /api/submit):** `{"user_id": "u1", "question_id": 101, "answer": "A"}` (Giả sử A là đúng).
  * **Logic Check:**
    1.  Lookup đáp án đúng (CorrectAnswer="A").
    2.  Calculate Score = 100.
    3.  Insert vào bảng `submissions`.
    4.  Publish message lên RabbitMQ `scoring.exchange`.
  * **Expected Output (HTTP 200):** `{"correct": true, "score": 100}`.
  * **Backend Verification:**
      * **DB:** `SELECT count(*) FROM submissions WHERE user_id='u1' AND score_awarded=100` -\> Result: 1.
      * **RabbitMQ:** Check queue, message payload phải chứa `{"score_obtained": 100, "skill_tag": "..."}`.

#### TC\_SCORING\_02: Submit Incorrect Answer & Publish Event

  * **Description:** Người dùng nộp sai, nhận 0 điểm, event vẫn phải được bắn đi để Learner Model biết mà trừ điểm/giữ nguyên.
  * **Input:** `{"user_id": "u1", "question_id": 101, "answer": "B"}`.
  * **Expected Output (HTTP 200):** `{"correct": false, "score": 0}`.
  * **Backend Verification:** Event payload phải chứa `{"score_obtained": 0}`.

### 1.2 Edge Cases & Error Handling

#### TC\_SCORING\_03: Invalid Question ID

  * **Input:** `question_id`: 99999 (Không tồn tại).
  * **Logic:** Service gọi Content Service (hoặc tra cache) không thấy ID.
  * **Expected Output:** HTTP 404 Not Found hoặc HTTP 400 Bad Request.
  * **Side Effect:** **KHÔNG** được lưu DB, **KHÔNG** được bắn Event.

#### TC\_SCORING\_04: RabbitMQ Downtime Resilience

  * **Condition:** Stop RabbitMQ container (`docker stop rabbitmq`).
  * **Action:** Gọi API Submit.
  * **Expected Behavior:**
      * **Option A (Strong Consistency):** Trả về HTTP 500 (Nếu yêu cầu bắt buộc phải lưu được event).
      * **Option B (Availability prioritized - Recommended):** Trả về HTTP 200, nhưng log error "Failed to publish event" vào file log/Sentry để retry sau (Outbox Pattern).
      * *Test này đảm bảo hệ thống chấm điểm không "chết" theo Message Queue.*

-----

## PART 2: LEARNER MODEL SERVICE (Consumer)

*Worker chạy ngầm, chịu trách nhiệm cập nhật trạng thái kỹ năng của người dùng.*

### 2.1 Logic Calculation Cases

#### TC\_LEARNER\_01: Update Score - Weighted Average

  * **Pre-condition:** User u1 có `math_algebra` mastery = 50.
  * **Event Input:** `SubmissionCompleted` { "score\_obtained": 100 }.
  * **Logic:** Công thức cập nhật (Ví dụ: `(Old * 0.7) + (New * 0.3)`).
      * Calculation: `(50 * 0.7) + (100 * 0.3) = 35 + 30 = 65`.
  * **Verification:**
      * **DB:** `SELECT current_score FROM skill_mastery WHERE user_id='u1'` -\> Result: 65.

#### TC\_LEARNER\_02: Update Score - Penalize Incorrect Answer

  * **Pre-condition:** User u1 có `math_algebra` mastery = 50.
  * **Event Input:** `SubmissionCompleted` { "score\_obtained": 0 }.
  * **Logic:** `(50 * 0.7) + (0 * 0.3) = 35`.
  * **Verification:** DB `current_score` giảm xuống 35.

### 2.2 Integrity & Concurrency

#### TC\_LEARNER\_03: New User Cold Start

  * **Pre-condition:** User `u_new` chưa từng có record trong bảng `skill_mastery`.
  * **Event Input:** `SubmissionCompleted` { "user\_id": "u\_new", "score": 100, "skill": "math" }.
  * **Logic:** Check DB thấy null -\> Insert record mới thay vì Update.
  * **Verification:** Record mới được tạo với `current_score` = 100 (hoặc theo logic khởi tạo).

#### TC\_LEARNER\_04: Idempotency Check (Quan trọng cho Event-Driven)

  * **Description:** Message bị gửi lặp lại 2 lần do mạng lag (At-least-once delivery).
  * **Action:** Gửi cùng 1 message (cùng `event_id` hoặc `submission_id`) 2 lần vào Queue.
  * **Expected Behavior:**
      * Lần 1: Update điểm, Acknowledge message.
      * Lần 2: Phát hiện duplicate (dựa trên log xử lý hoặc logic nghiệp vụ), **KHÔNG** cộng điểm lần nữa, nhưng vẫn Acknowledge để xóa message khỏi queue.
  * **Verification:** Điểm số trong DB chỉ thay đổi 1 lần.

-----

## PART 3: CONTENT SERVICE (Java Data Provider)

*Service quản lý kho dữ liệu bài tập, hỗ trợ query theo metadata.*

### 3.1 Query Logic

#### TC\_CONTENT\_01: Filter Remedial Content

  * **Input:** `GET /recommend?skill=math&type=remedial`.
  * **Logic:** JPA/SQL Query `WHERE skill_tag='math' AND is_remedial=TRUE`.
  * **Verification:** JSON trả về phải có `id` của bài tập có flag `remedial=true`. Nếu trả về bài khó -\> Bug.

#### TC\_CONTENT\_02: Filter Standard Content

  * **Input:** `GET /recommend?skill=math&type=standard`.
  * **Verification:** JSON trả về bài có `is_remedial=FALSE`.

#### TC\_CONTENT\_03: No Content Available

  * **Input:** Skill không tồn tại hoặc hết bài tập.
  * **Expected Output:**
      * HTTP 200 với Body rỗng hoặc Default content (Fallback).
      * Hoặc HTTP 404 (Tùy design, nên dùng 200 + empty list).

-----

## PART 4: ADAPTIVE ENGINE (Orchestrator)

*Bộ não ra quyết định, kết nối dữ liệu từ Learner Model và Content.*

### 4.1 Decision Logic Cases

#### TC\_ADAPTIVE\_01: Trigger Remediation (Weak Student)

  * **Mock Data Setup:**
      * Mock `LearnerService.GetMastery("u1")` returns `{ score: 30 }`.
  * **Action:** `POST /adaptive/next` cho user `u1`.
  * **Logic:**
    1.  Check Score (30) \< Threshold (50).
    2.  Decision = "REMEDIAL".
    3.  Call `ContentService.GetRecommendation(type="remedial")`.
  * **Expected Output:** JSON trả về bài học Remedial.

#### TC\_ADAPTIVE\_02: Trigger Advancement (Strong Student)

  * **Mock Data Setup:**
      * Mock `LearnerService.GetMastery("u2")` returns `{ score: 85 }`.
  * **Action:** `POST /adaptive/next` cho user `u2`.
  * **Logic:**
    1.  Check Score (85) \>= Threshold (50).
    2.  Decision = "STANDARD" (hoặc ADVANCED).
    3.  Call `ContentService.GetRecommendation(type="standard")`.
  * **Expected Output:** JSON trả về bài học Standard/Advanced.

### 4.2 Resilience Cases

#### TC\_ADAPTIVE\_03: Learner Service Unavailable (Fallback)

  * **Condition:** `Learner Model Service` bị down.
  * **Action:** `POST /adaptive/next`.
  * **Logic:** Không lấy được điểm số hiện tại.
  * **Fallback Strategy:** Mặc định coi như user bình thường (Standard) hoặc an toàn hơn là cho làm bài Dễ (Remedial).
  * **Expected Output:** HTTP 200 (Không được crash 500), trả về bài học theo Default Strategy.

-----

## SUMMARY OF SQL VALIDATION SCRIPTS

*Dùng các câu lệnh này để verify dữ liệu sau khi chạy test.*

```sql
-- 1. Kiểm tra User u1 đã nộp bài chưa và điểm có lưu đúng không?
SELECT * FROM scoring_svc.submissions 
WHERE user_id = 'u1' ORDER BY created_at DESC LIMIT 1;

-- 2. Kiểm tra User u1 đã được cập nhật điểm thông thạo chưa?
-- (Chạy sau khi nộp bài khoảng 1-2s)
SELECT * FROM learner_svc.skill_mastery 
WHERE user_id = 'u1' AND skill_tag = 'math_algebra';

-- 3. Kiểm tra xem có bài tập Remedial nào trong kho không?
SELECT count(*) FROM content_svc.questions 
WHERE skill_tag = 'math_algebra' AND is_remedial = TRUE;
```