-- Insert Test Data for Integration Tests
-- This script inserts math and science questions for testing the adaptive learning flow

-- Math Remedial Questions (5 questions)
INSERT INTO questions (content, correct_answer, skill_tag, is_remedial, difficulty_level, created_at)
VALUES
('What is 2 + 2?', '4', 'math', true, 1, NOW()),
('What is 5 - 3?', '2', 'math', true, 1, NOW()),
('What is 3 × 4?', '12', 'math', true, 1, NOW()),
('What is 10 ÷ 2?', '5', 'math', true, 1, NOW()),
('What is 7 + 8?', '15', 'math', true, 1, NOW());

-- Math Standard Questions (5 questions)
INSERT INTO questions (content, correct_answer, skill_tag, is_remedial, difficulty_level, created_at)
VALUES
('Solve: 2x + 5 = 13. What is x?', '4', 'math', false, 2, NOW()),
('What is the square root of 144?', '12', 'math', false, 2, NOW()),
('Calculate: (3 + 5) × 2 - 4', '12', 'math', false, 2, NOW()),
('What is 15% of 200?', '30', 'math', false, 2, NOW()),
('If a triangle has angles 60° and 70°, what is the third angle?', '50', 'math', false, 2, NOW());

-- Science Remedial Questions (5 questions)
INSERT INTO questions (content, correct_answer, skill_tag, is_remedial, difficulty_level, created_at)
VALUES
('What is the chemical symbol for water?', 'H2O', 'science', true, 1, NOW()),
('How many planets are in our solar system?', '8', 'science', true, 1, NOW()),
('What gas do plants absorb from the atmosphere?', 'CO2', 'science', true, 1, NOW()),
('What is the freezing point of water in Celsius?', '0', 'science', true, 1, NOW()),
('What force pulls objects toward Earth?', 'gravity', 'science', true, 1, NOW());

-- Science Standard Questions (5 questions)
INSERT INTO questions (content, correct_answer, skill_tag, is_remedial, difficulty_level, created_at)
VALUES
('What is the powerhouse of the cell?', 'mitochondria', 'science', false, 2, NOW()),
('What is the speed of light in vacuum (in km/s)?', '300000', 'science', false, 2, NOW()),
('What is the atomic number of Carbon?', '6', 'science', false, 2, NOW()),
('What is the process by which plants make food?', 'photosynthesis', 'science', false, 2, NOW()),
('What is the smallest unit of life?', 'cell', 'science', false, 2, NOW());

-- Verify inserted data
SELECT 
    skill_tag, 
    is_remedial, 
    COUNT(*) as question_count 
FROM questions 
WHERE skill_tag IN ('math', 'science')
GROUP BY skill_tag, is_remedial 
ORDER BY skill_tag, is_remedial;

