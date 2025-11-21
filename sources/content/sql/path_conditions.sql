-- Path Conditions Table
-- Hỗ trợ cấu hình lộ trình học tập

CREATE TABLE IF NOT EXISTS path_conditions (
    condition_id UUID PRIMARY KEY,
    source_unit_id UUID NOT NULL,
    target_unit_id UUID NOT NULL,
    required_score_pct INTEGER NOT NULL CHECK (required_score_pct >= 0 AND required_score_pct <= 100),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT check_no_self_reference CHECK (source_unit_id != target_unit_id)
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_path_conditions_source_unit_id ON path_conditions(source_unit_id);
CREATE INDEX IF NOT EXISTS idx_path_conditions_target_unit_id ON path_conditions(target_unit_id);
CREATE INDEX IF NOT EXISTS idx_path_conditions_required_score_pct ON path_conditions(required_score_pct);
CREATE INDEX IF NOT EXISTS idx_path_conditions_created_at ON path_conditions(created_at);
CREATE INDEX IF NOT EXISTS idx_path_conditions_updated_at ON path_conditions(updated_at);

-- Unique constraint để tránh duplicate path conditions
CREATE UNIQUE INDEX IF NOT EXISTS idx_path_conditions_unique_path 
ON path_conditions(source_unit_id, target_unit_id);

-- Foreign key constraints (nếu có bảng content_units)
-- ALTER TABLE path_conditions ADD CONSTRAINT fk_path_conditions_source_unit_id 
--     FOREIGN KEY (source_unit_id) REFERENCES content_units(unit_id) ON DELETE CASCADE;
-- ALTER TABLE path_conditions ADD CONSTRAINT fk_path_conditions_target_unit_id 
--     FOREIGN KEY (target_unit_id) REFERENCES content_units(unit_id) ON DELETE CASCADE;

-- Trigger để tự động cập nhật updated_at
CREATE OR REPLACE FUNCTION update_path_conditions_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER trigger_update_path_conditions_updated_at
    BEFORE UPDATE ON path_conditions
    FOR EACH ROW
    EXECUTE FUNCTION update_path_conditions_updated_at();

-- Sample data
INSERT INTO path_conditions (condition_id, source_unit_id, target_unit_id, required_score_pct, created_at, updated_at) VALUES
(
    '550e8400-e29b-41d4-a716-446655440101',
    '550e8400-e29b-41d4-a716-446655440001',
    '550e8400-e29b-41d4-a716-446655440002',
    70,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
),
(
    '550e8400-e29b-41d4-a716-446655440102',
    '550e8400-e29b-41d4-a716-446655440001',
    '550e8400-e29b-41d4-a716-446655440003',
    80,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
),
(
    '550e8400-e29b-41d4-a716-446655440103',
    '550e8400-e29b-41d4-a716-446655440002',
    '550e8400-e29b-41d4-a716-446655440004',
    75,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
),
(
    '550e8400-e29b-41d4-a716-446655440104',
    '550e8400-e29b-41d4-a716-446655440003',
    '550e8400-e29b-41d4-a716-446655440004',
    85,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
);
