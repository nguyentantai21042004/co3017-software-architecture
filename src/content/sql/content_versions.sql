-- Content Versions Table
-- Phiên bản nội dung với semantic versioning

CREATE TABLE IF NOT EXISTS content_versions (
    version_id BIGSERIAL PRIMARY KEY,
    unit_id UUID NOT NULL,
    version_number VARCHAR(20) NOT NULL,
    content_data JSONB NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_content_versions_unit_id ON content_versions(unit_id);
CREATE INDEX IF NOT EXISTS idx_content_versions_version_number ON content_versions(version_number);
CREATE INDEX IF NOT EXISTS idx_content_versions_is_active ON content_versions(is_active);
CREATE INDEX IF NOT EXISTS idx_content_versions_created_at ON content_versions(created_at);
CREATE INDEX IF NOT EXISTS idx_content_versions_updated_at ON content_versions(updated_at);

-- Unique constraint để đảm bảo mỗi unit chỉ có một version active
CREATE UNIQUE INDEX IF NOT EXISTS idx_content_versions_unit_active 
ON content_versions(unit_id) WHERE is_active = TRUE;

-- Foreign key constraint (nếu có bảng content_units)
-- ALTER TABLE content_versions ADD CONSTRAINT fk_content_versions_unit_id 
--     FOREIGN KEY (unit_id) REFERENCES content_units(unit_id) ON DELETE CASCADE;

-- Trigger để tự động cập nhật updated_at
CREATE OR REPLACE FUNCTION update_content_versions_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER trigger_update_content_versions_updated_at
    BEFORE UPDATE ON content_versions
    FOR EACH ROW
    EXECUTE FUNCTION update_content_versions_updated_at();

-- Function để deactivate tất cả versions của một unit
CREATE OR REPLACE FUNCTION deactivate_unit_versions(p_unit_id UUID)
RETURNS VOID AS $$
BEGIN
    UPDATE content_versions 
    SET is_active = FALSE, updated_at = CURRENT_TIMESTAMP
    WHERE unit_id = p_unit_id;
END;
$$ language 'plpgsql';

-- Sample data
INSERT INTO content_versions (unit_id, version_number, content_data, is_active, created_at, updated_at) VALUES
(
    '550e8400-e29b-41d4-a716-446655440001',
    '1.0.0',
    '{"title": "Giới thiệu về Java - Version 1.0", "content": "Java là một ngôn ngữ lập trình hướng đối tượng...", "estimated_time": 30, "sections": [{"title": "Khái niệm cơ bản", "content": "Java là gì?"}]}',
    TRUE,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
),
(
    '550e8400-e29b-41d4-a716-446655440001',
    '1.1.0',
    '{"title": "Giới thiệu về Java - Version 1.1", "content": "Java là một ngôn ngữ lập trình hướng đối tượng mạnh mẽ...", "estimated_time": 35, "sections": [{"title": "Khái niệm cơ bản", "content": "Java là gì?"}, {"title": "Lịch sử Java", "content": "Java được phát triển bởi Sun Microsystems..."}]}',
    FALSE,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
),
(
    '550e8400-e29b-41d4-a716-446655440002',
    '1.0.0',
    '{"title": "Video hướng dẫn Java cơ bản - Version 1.0", "video_url": "https://example.com/java-basic-v1.mp4", "duration": 1800, "thumbnail_url": "https://example.com/thumb-v1.jpg", "quality": "HD"}',
    TRUE,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
),
(
    '550e8400-e29b-41d4-a716-446655440003',
    '1.0.0',
    '{"title": "Quiz về Java cơ bản - Version 1.0", "questions": [{"id": 1, "question": "Java là gì?", "options": ["Ngôn ngữ lập trình", "Hệ điều hành", "Database"], "correct_answer": 0, "points": 10}], "time_limit": 900, "passing_score": 70}',
    TRUE,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
),
(
    '550e8400-e29b-41d4-a716-446655440003',
    '1.1.0',
    '{"title": "Quiz về Java cơ bản - Version 1.1", "questions": [{"id": 1, "question": "Java là gì?", "options": ["Ngôn ngữ lập trình", "Hệ điều hành", "Database"], "correct_answer": 0, "points": 10}, {"id": 2, "question": "JVM là gì?", "options": ["Java Virtual Machine", "Java Version Manager", "Java Virtual Memory"], "correct_answer": 0, "points": 10}], "time_limit": 1200, "passing_score": 70}',
    FALSE,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
),
(
    '550e8400-e29b-41d4-a716-446655440004',
    '1.0.0',
    '{"title": "Viết chương trình Hello World - Version 1.0", "description": "Viết chương trình Java in ra Hello World", "instructions": "Sử dụng System.out.println để in ra Hello World", "test_cases": [{"input": "", "expected_output": "Hello World", "points": 100}], "language": "java", "time_limit": 300}',
    TRUE,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
);
