-- Unit Tags Table
-- Liên kết N:M giữa ContentUnit và MetadataTag

CREATE TABLE IF NOT EXISTS unit_tags (
    unit_id UUID NOT NULL,
    tag_id INTEGER NOT NULL,
    relevance_score FLOAT NOT NULL CHECK (relevance_score >= 0.0 AND relevance_score <= 1.0),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (unit_id, tag_id)
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_unit_tags_unit_id ON unit_tags(unit_id);
CREATE INDEX IF NOT EXISTS idx_unit_tags_tag_id ON unit_tags(tag_id);
CREATE INDEX IF NOT EXISTS idx_unit_tags_relevance_score ON unit_tags(relevance_score);
CREATE INDEX IF NOT EXISTS idx_unit_tags_created_at ON unit_tags(created_at);
CREATE INDEX IF NOT EXISTS idx_unit_tags_updated_at ON unit_tags(updated_at);

-- Foreign key constraints (nếu có bảng content_units và metadata_tags)
-- ALTER TABLE unit_tags ADD CONSTRAINT fk_unit_tags_unit_id 
--     FOREIGN KEY (unit_id) REFERENCES content_units(unit_id) ON DELETE CASCADE;
-- ALTER TABLE unit_tags ADD CONSTRAINT fk_unit_tags_tag_id 
--     FOREIGN KEY (tag_id) REFERENCES metadata_tags(tag_id) ON DELETE CASCADE;

-- Trigger để tự động cập nhật updated_at
CREATE OR REPLACE FUNCTION update_unit_tags_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER trigger_update_unit_tags_updated_at
    BEFORE UPDATE ON unit_tags
    FOR EACH ROW
    EXECUTE FUNCTION update_unit_tags_updated_at();

-- Sample data
INSERT INTO unit_tags (unit_id, tag_id, relevance_score, created_at, updated_at) VALUES
('550e8400-e29b-41d4-a716-446655440001', 1, 0.9, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('550e8400-e29b-41d4-a716-446655440001', 2, 0.8, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('550e8400-e29b-41d4-a716-446655440001', 9, 1.0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('550e8400-e29b-41d4-a716-446655440002', 1, 0.7, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('550e8400-e29b-41d4-a716-446655440002', 9, 1.0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('550e8400-e29b-41d4-a716-446655440003', 1, 0.8, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('550e8400-e29b-41d4-a716-446655440003', 3, 0.6, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('550e8400-e29b-41d4-a716-446655440003', 9, 1.0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('550e8400-e29b-41d4-a716-446655440004', 1, 0.9, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('550e8400-e29b-41d4-a716-446655440004', 4, 0.7, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('550e8400-e29b-41d4-a716-446655440004', 10, 0.8, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
