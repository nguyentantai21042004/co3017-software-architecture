-- Insert E2E Test Data
-- This script inserts test data for E2E testing:
-- 1. Questions for Content Service (content_db)
-- 2. Test user mastery data for Learner Model Service (learner_db)

-- ============================================================================
-- CONTENT SERVICE TEST DATA (content_db)
-- ============================================================================

-- Math Remedial Questions (for E2E testing)
INSERT INTO questions (content, options, correct_answer, skill_tag, is_remedial, difficulty_level, created_at)
VALUES
('What is 2 + 2?', '["2", "3", "4", "5"]', '4', 'math', true, 1, NOW()),
('What is 5 - 3?', '["1", "2", "3", "4"]', '2', 'math', true, 1, NOW()),
('What is 3 × 4?', '["9", "10", "12", "15"]', '12', 'math', true, 1, NOW()),
('What is 10 ÷ 2?', '["2", "5", "10", "20"]', '5', 'math', true, 1, NOW()),
('What is 7 + 8?', '["14", "15", "16", "17"]', '15', 'math', true, 1, NOW())
ON CONFLICT DO NOTHING;

-- Math Standard Questions
INSERT INTO questions (content, options, correct_answer, skill_tag, is_remedial, difficulty_level, created_at)
VALUES
('Solve: 2x + 5 = 13. What is x?', '["2", "4", "6", "8"]', '4', 'math', false, 2, NOW()),
('What is the square root of 144?', '["10", "11", "12", "13"]', '12', 'math', false, 2, NOW()),
('Calculate: (3 + 5) × 2 - 4', '["8", "10", "12", "16"]', '12', 'math', false, 2, NOW())
ON CONFLICT DO NOTHING;

-- Science Remedial Questions
INSERT INTO questions (content, options, correct_answer, skill_tag, is_remedial, difficulty_level, created_at)
VALUES
('What is the chemical symbol for water?', '["H2O", "CO2", "O2", "H2"]', 'H2O', 'science', true, 1, NOW()),
('How many planets are in our solar system?', '["6", "7", "8", "9"]', '8', 'science', true, 1, NOW()),
('What gas do plants absorb from the atmosphere?', '["O2", "N2", "CO2", "H2"]', 'CO2', 'science', true, 1, NOW()),
('What is the freezing point of water in Celsius?', '["-10", "0", "10", "100"]', '0', 'science', true, 1, NOW()),
('What force pulls objects toward Earth?', '["magnetism", "gravity", "friction", "tension"]', 'gravity', 'science', true, 1, NOW())
ON CONFLICT DO NOTHING;

-- Science Standard Questions
INSERT INTO questions (content, options, correct_answer, skill_tag, is_remedial, difficulty_level, created_at)
VALUES
('What is the powerhouse of the cell?', '["nucleus", "mitochondria", "ribosome", "chloroplast"]', 'mitochondria', 'science', false, 2, NOW()),
('What is the atomic number of Carbon?', '["4", "6", "8", "12"]', '6', 'science', false, 2, NOW()),
('What is the smallest unit of life?', '["atom", "molecule", "cell", "tissue"]', 'cell', 'science', false, 2, NOW())
ON CONFLICT DO NOTHING;

