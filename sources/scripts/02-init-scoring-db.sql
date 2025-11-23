-- =============================================================================
-- DATABASE: scoring_db
-- SERVICE: Scoring Service (Golang) - Port 8082
-- PURPOSE: Store submission records and scoring results
-- =============================================================================

-- Drop database if exists (for clean re-initialization)
DROP DATABASE IF EXISTS scoring_db;

-- Create database
CREATE DATABASE scoring_db;

-- Connect to the new database
\c scoring_db

-- =============================================================================
-- TABLE: submissions
-- Stores all user submissions and their scoring results
-- =============================================================================
CREATE TABLE submissions (
    id BIGSERIAL PRIMARY KEY,              -- Changed from SERIAL to BIGSERIAL for int64 compatibility
    user_id VARCHAR(50) NOT NULL,
    question_id BIGINT NOT NULL,           -- Changed from INT to BIGINT to match Go int64 type
    submitted_answer VARCHAR(255) NOT NULL,
    score_awarded INTEGER NOT NULL,        -- Score in percentage (0-100)
    is_passed BOOLEAN NOT NULL,            -- TRUE if score >= 50%
    created_at TIMESTAMP DEFAULT NOW()
);

-- Create indexes for efficient querying
CREATE INDEX idx_submissions_user_id ON submissions(user_id);
CREATE INDEX idx_submissions_question_id ON submissions(question_id);
CREATE INDEX idx_submissions_created_at ON submissions(created_at DESC);

-- Composite index for user performance tracking
CREATE INDEX idx_submissions_user_question ON submissions(user_id, question_id);

-- =============================================================================
-- SEED DATA: Sample submissions for testing (Optional)
-- =============================================================================

-- Example: User 'user_01' made some previous attempts
INSERT INTO submissions (user_id, question_id, submitted_answer, score_awarded, is_passed)
VALUES
    ('user_01', 3, 'B', 100, TRUE),  -- Correct answer on medium question
    ('user_01', 2, 'A', 100, TRUE);  -- Correct answer on remedial question

-- =============================================================================
-- VERIFICATION
-- =============================================================================
SELECT 'scoring_db initialized successfully!' AS status;
SELECT 'Total submissions inserted: ' || COUNT(*) AS summary FROM submissions;
