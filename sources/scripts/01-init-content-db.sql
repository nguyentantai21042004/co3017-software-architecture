-- =============================================================================
-- DATABASE: content_db
-- SERVICE: Content Service (Java/Spring Boot) - Port 8081
-- PURPOSE: Store questions, skills, and learning content
-- NOTE: Database is created by init-multiple-postgresql-databases.sh
-- This script only creates tables and inserts data
-- =============================================================================

-- =============================================================================
-- TABLE: questions
-- Stores all learning questions with metadata for adaptive learning
-- =============================================================================
CREATE TABLE IF NOT EXISTS questions (
    id BIGSERIAL PRIMARY KEY,              -- Changed from SERIAL to BIGSERIAL for Long compatibility
    content TEXT NOT NULL,
    options JSONB,                         -- Example: ["A. Option1", "B. Option2", "C. Option3", "D. Option4"]
    correct_answer TEXT NOT NULL,          -- Changed from VARCHAR(255) to TEXT for long-form answers
    skill_tag VARCHAR(100) NOT NULL,       -- Changed from VARCHAR(50) to VARCHAR(100) to match JPA entity
    difficulty_level INTEGER DEFAULT 1,    -- 1: Easy, 2: Medium, 3: Hard
    is_remedial BOOLEAN DEFAULT FALSE,     -- TRUE: Remedial/review question
    created_at TIMESTAMP DEFAULT NOW()
);

-- Create index for efficient querying by skill and type
CREATE INDEX IF NOT EXISTS idx_questions_skill_remedial ON questions(skill_tag, is_remedial);
CREATE INDEX IF NOT EXISTS idx_questions_difficulty ON questions(difficulty_level);

-- =============================================================================
-- SEED DATA: Sample questions for testing Adaptive Flow
-- =============================================================================

-- Clear existing data (optional, for clean re-initialization)
TRUNCATE TABLE questions RESTART IDENTITY CASCADE;

-- Question 1: Hard main question (for testing failure scenario)
INSERT INTO questions (content, options, correct_answer, skill_tag, difficulty_level, is_remedial)
VALUES (
    'Bài toán khó: Giải phương trình bậc hai x² + 5x + 6 = 0',
    '["A. x = -2 và x = -3", "B. x = 2 và x = 3", "C. x = -1 và x = -6", "D. Vô nghiệm"]',
    'A',
    'math_algebra',
    3,
    FALSE
);

-- Question 2: Easy remedial question (for adaptive recommendation)
INSERT INTO questions (content, options, correct_answer, skill_tag, difficulty_level, is_remedial)
VALUES (
    'Bài ôn tập: Nhắc lại quy tắc chuyển vế trong phương trình. Với phương trình x + 3 = 7, x bằng bao nhiêu?',
    '["A. 4", "B. 10", "C. -4", "D. 3"]',
    'A',
    'math_algebra',
    1,
    TRUE
);

-- Question 3: Medium standard question
INSERT INTO questions (content, options, correct_answer, skill_tag, difficulty_level, is_remedial)
VALUES (
    'Bài tập chuẩn: Giải phương trình 2x - 4 = 10',
    '["A. x = 3", "B. x = 7", "C. x = 5", "D. x = 14"]',
    'B',
    'math_algebra',
    2,
    FALSE
);

-- Question 4: Another remedial for geometry (different skill)
INSERT INTO questions (content, options, correct_answer, skill_tag, difficulty_level, is_remedial)
VALUES (
    'Bài ôn tập: Diện tích hình chữ nhật có chiều dài 5cm, chiều rộng 3cm là bao nhiêu?',
    '["A. 8 cm²", "B. 15 cm²", "C. 16 cm²", "D. 30 cm²"]',
    'B',
    'math_geometry',
    1,
    TRUE
);

-- Question 5: Hard geometry question
INSERT INTO questions (content, options, correct_answer, skill_tag, difficulty_level, is_remedial)
VALUES (
    'Bài toán khó: Tính thể tích hình cầu có bán kính 3cm (lấy π ≈ 3.14)',
    '["A. 113.04 cm³", "B. 28.26 cm³", "C. 37.68 cm³", "D. 84.78 cm³"]',
    'A',
    'math_geometry',
    3,
    FALSE
);

-- =============================================================================
-- VERIFICATION
-- =============================================================================
SELECT 'content_db initialized successfully!' AS status;
SELECT 'Total questions inserted: ' || COUNT(*) AS summary FROM questions;
SELECT 'Remedial questions: ' || COUNT(*) AS remedial_count FROM questions WHERE is_remedial = TRUE;
SELECT 'Standard questions: ' || COUNT(*) AS standard_count FROM questions WHERE is_remedial = FALSE;
