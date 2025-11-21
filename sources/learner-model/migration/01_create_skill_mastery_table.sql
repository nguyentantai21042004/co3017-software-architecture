-- Create skill_mastery table
CREATE TABLE IF NOT EXISTS skill_mastery (
    user_id VARCHAR(255) NOT NULL,
    skill_tag VARCHAR(255) NOT NULL,
    current_score INTEGER NOT NULL DEFAULT 0 CHECK (current_score >= 0 AND current_score <= 100),
    last_updated TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (user_id, skill_tag)
);

-- Add indexes for better query performance
CREATE INDEX IF NOT EXISTS idx_skill_mastery_user_id ON skill_mastery(user_id);
CREATE INDEX IF NOT EXISTS idx_skill_mastery_skill_tag ON skill_mastery(skill_tag);
CREATE INDEX IF NOT EXISTS idx_skill_mastery_last_updated ON skill_mastery(last_updated);

-- Add comment
COMMENT ON TABLE skill_mastery IS 'Stores user mastery levels for different skills';
COMMENT ON COLUMN skill_mastery.current_score IS 'Mastery score from 0-100';
