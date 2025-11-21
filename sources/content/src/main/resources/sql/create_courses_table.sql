-- Create courses table
CREATE TABLE IF NOT EXISTS courses (
    course_id UUID PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    instructor_id UUID NOT NULL,
    structure_type VARCHAR(20) NOT NULL CHECK (structure_type IN ('LINEAR', 'ADAPTIVE')),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for better performance
CREATE INDEX IF NOT EXISTS idx_courses_instructor_id ON courses(instructor_id);
CREATE INDEX IF NOT EXISTS idx_courses_structure_type ON courses(structure_type);
CREATE INDEX IF NOT EXISTS idx_courses_title ON courses(title);
CREATE INDEX IF NOT EXISTS idx_courses_created_at ON courses(created_at);

-- Create trigger to automatically update updated_at timestamp
CREATE OR REPLACE FUNCTION update_courses_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER trigger_courses_updated_at
    BEFORE UPDATE ON courses
    FOR EACH ROW
    EXECUTE FUNCTION update_courses_updated_at();

-- Insert sample data
INSERT INTO courses (course_id, title, description, instructor_id, structure_type) VALUES
    ('550e8400-e29b-41d4-a716-446655440001', 'Introduction to Programming', 'Learn the fundamentals of programming with Java', '660e8400-e29b-41d4-a716-446655440001', 'LINEAR'),
    ('550e8400-e29b-41d4-a716-446655440002', 'Advanced Algorithms', 'Master complex algorithms and data structures', '660e8400-e29b-41d4-a716-446655440002', 'ADAPTIVE'),
    ('550e8400-e29b-41d4-a716-446655440003', 'Web Development', 'Build modern web applications', '660e8400-e29b-41d4-a716-446655440001', 'LINEAR')
ON CONFLICT (course_id) DO NOTHING;
