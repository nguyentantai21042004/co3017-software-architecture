-- Metadata Tags Table
-- Danh sách kỹ năng/chủ đề cốt lõi (Source of Truth)

CREATE TABLE IF NOT EXISTS metadata_tags (
    tag_id SERIAL PRIMARY KEY,
    tag_name VARCHAR(100) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_metadata_tags_tag_name ON metadata_tags(tag_name);
CREATE INDEX IF NOT EXISTS idx_metadata_tags_created_at ON metadata_tags(created_at);
CREATE INDEX IF NOT EXISTS idx_metadata_tags_updated_at ON metadata_tags(updated_at);

-- Trigger để tự động cập nhật updated_at
CREATE OR REPLACE FUNCTION update_metadata_tags_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER trigger_update_metadata_tags_updated_at
    BEFORE UPDATE ON metadata_tags
    FOR EACH ROW
    EXECUTE FUNCTION update_metadata_tags_updated_at();

-- Sample data
INSERT INTO metadata_tags (tag_name, created_at, updated_at) VALUES
('Java Programming', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Object-Oriented Programming', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Data Structures', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Algorithms', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Database Design', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Web Development', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Spring Framework', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('RESTful APIs', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Beginner Level', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Intermediate Level', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Advanced Level', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Frontend Development', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Backend Development', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Full-Stack Development', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Software Testing', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
