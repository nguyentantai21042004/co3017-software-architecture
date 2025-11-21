-- Content Units Table
-- Đơn vị nội dung (bài học/quiz/assignment)

CREATE TABLE IF NOT EXISTS content_units (
    unit_id UUID PRIMARY KEY,
    chapter_id UUID NOT NULL,
    unit_type VARCHAR(20) NOT NULL CHECK (unit_type IN ('TEXT', 'VIDEO', 'QUIZ', 'CODING_TASK')),
    metadata_config JSONB NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_content_units_chapter_id ON content_units(chapter_id);
CREATE INDEX IF NOT EXISTS idx_content_units_unit_type ON content_units(unit_type);
CREATE INDEX IF NOT EXISTS idx_content_units_created_at ON content_units(created_at);
CREATE INDEX IF NOT EXISTS idx_content_units_updated_at ON content_units(updated_at);

-- Foreign key constraint (nếu có bảng chapters)
-- ALTER TABLE content_units ADD CONSTRAINT fk_content_units_chapter_id 
--     FOREIGN KEY (chapter_id) REFERENCES chapters(chapter_id) ON DELETE CASCADE;

-- Trigger để tự động cập nhật updated_at
CREATE OR REPLACE FUNCTION update_content_units_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER trigger_update_content_units_updated_at
    BEFORE UPDATE ON content_units
    FOR EACH ROW
    EXECUTE FUNCTION update_content_units_updated_at();

-- Sample data
INSERT INTO content_units (unit_id, chapter_id, unit_type, metadata_config, created_at, updated_at) VALUES
(
    '550e8400-e29b-41d4-a716-446655440001',
    '550e8400-e29b-41d4-a716-446655440010',
    'TEXT',
    '{"title": "Giới thiệu về Java", "content": "Java là một ngôn ngữ lập trình...", "estimated_time": 30}',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
),
(
    '550e8400-e29b-41d4-a716-446655440002',
    '550e8400-e29b-41d4-a716-446655440010',
    'VIDEO',
    '{"title": "Video hướng dẫn Java cơ bản", "video_url": "https://example.com/java-basic.mp4", "duration": 1800}',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
),
(
    '550e8400-e29b-41d4-a716-446655440003',
    '550e8400-e29b-41d4-a716-446655440011',
    'QUIZ',
    '{"title": "Quiz về Java cơ bản", "questions": [{"id": 1, "question": "Java là gì?", "options": ["Ngôn ngữ lập trình", "Hệ điều hành", "Database"], "correct_answer": 0}], "time_limit": 900}',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
),
(
    '550e8400-e29b-41d4-a716-446655440004',
    '550e8400-e29b-41d4-a716-446655440011',
    'CODING_TASK',
    '{"title": "Viết chương trình Hello World", "description": "Viết chương trình Java in ra Hello World", "test_cases": [{"input": "", "expected_output": "Hello World"}], "language": "java"}',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
);
