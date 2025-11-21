-- =============================================================================
-- DATABASE: learner_db
-- SERVICE: Learner Model Service (Golang) - Port 8083
-- PURPOSE: Store learner skill mastery scores (AI state)
-- =============================================================================

-- Drop database if exists (for clean re-initialization)
DROP DATABASE IF EXISTS learner_db;

-- Create database
CREATE DATABASE learner_db;

-- Connect to the new database
\c learner_db

-- =============================================================================
-- TABLE: skill_mastery
-- Stores the current mastery level for each user-skill combination
-- This is the core of the Adaptive Learning Engine
-- =============================================================================
CREATE TABLE skill_mastery (
    user_id VARCHAR(50) NOT NULL,
    skill_tag VARCHAR(50) NOT NULL,
    current_score INT DEFAULT 0,           -- Mastery score (0-100)
    last_updated TIMESTAMP DEFAULT NOW(),
    PRIMARY KEY (user_id, skill_tag)
);

-- Create index for efficient querying
CREATE INDEX idx_skill_mastery_user_id ON skill_mastery(user_id);
CREATE INDEX idx_skill_mastery_skill_tag ON skill_mastery(skill_tag);
CREATE INDEX idx_skill_mastery_last_updated ON skill_mastery(last_updated DESC);

-- =============================================================================
-- SEED DATA: Initial learner states for testing
-- =============================================================================

-- User 'user_01': Low mastery in algebra (to trigger remedial recommendation)
INSERT INTO skill_mastery (user_id, skill_tag, current_score)
VALUES ('user_01', 'math_algebra', 10);

-- User 'user_01': No mastery in geometry yet (to test new skill initialization)
INSERT INTO skill_mastery (user_id, skill_tag, current_score)
VALUES ('user_01', 'math_geometry', 0);

-- User 'user_02': High mastery in algebra (to test standard path)
INSERT INTO skill_mastery (user_id, skill_tag, current_score)
VALUES ('user_02', 'math_algebra', 80);

-- User 'user_02': Medium mastery in geometry
INSERT INTO skill_mastery (user_id, skill_tag, current_score)
VALUES ('user_02', 'math_geometry', 55);

-- =============================================================================
-- VERIFICATION
-- =============================================================================
SELECT 'learner_db initialized successfully!' AS status;
SELECT 'Total skill mastery records inserted: ' || COUNT(*) AS summary FROM skill_mastery;
SELECT 'Users with low mastery (score < 50): ' || COUNT(DISTINCT user_id) AS low_mastery_users
FROM skill_mastery WHERE current_score < 50;
