-- =====================================================================
-- DevOps Remedial Questions - 5 Basic Questions
-- These are simpler questions for learners with low mastery scores
-- =====================================================================

-- Remedial Question 1: Very basic DevOps definition
INSERT INTO questions (content, options, correct_answer, skill_tag, difficulty_level, is_remedial, created_at)
VALUES (
    'DevOps là gì?',
    '["A. Một ngôn ngữ lập trình mới", "B. Một phương pháp kết hợp Development (phát triển) và Operations (vận hành)", "C. Một loại database", "D. Một hệ điều hành"]',
    'B',
    'devops',
    1,
    TRUE,
    NOW()
);

-- Remedial Question 2: Basic version control
INSERT INTO questions (content, options, correct_answer, skill_tag, difficulty_level, is_remedial, created_at)
VALUES (
    'Git được sử dụng để làm gì?',
    '["A. Vẽ đồ họa", "B. Quản lý phiên bản code (version control)", "C. Chỉnh sửa ảnh", "D. Gửi email"]',
    'B',
    'devops',
    1,
    TRUE,
    NOW()
);

-- Remedial Question 3: Simple CI/CD concept
INSERT INTO questions (content, options, correct_answer, skill_tag, difficulty_level, is_remedial, created_at)
VALUES (
    'CI/CD giúp làm gì?',
    '["A. Tự động hóa việc build, test và deploy code", "B. Viết code nhanh hơn", "C. Thiết kế giao diện đẹp hơn", "D. Lưu trữ database"]',
    'A',
    'devops',
    1,
    TRUE,
    NOW()
);

-- Remedial Question 4: Basic container concept
INSERT INTO questions (content, options, correct_answer, skill_tag, difficulty_level, is_remedial, created_at)
VALUES (
    'Container trong DevOps có tác dụng gì?',
    '["A. Đóng gói ứng dụng cùng các dependencies để chạy ở mọi môi trường", "B. Chỉ để lưu trữ dữ liệu", "C. Thay thế hoàn toàn database", "D. Chỉ dùng để test"]',
    'A',
    'devops',
    1,
    TRUE,
    NOW()
);

-- Remedial Question 5: Basic deployment
INSERT INTO questions (content, options, correct_answer, skill_tag, difficulty_level, is_remedial, created_at)
VALUES (
    'Deploy code nghĩa là gì?',
    '["A. Xóa code cũ", "B. Viết code mới", "C. Đưa code lên môi trường production để người dùng sử dụng", "D. Test code"]',
    'C',
    'devops',
    1,
    TRUE,
    NOW()
);

-- Verification query
SELECT
    id,
    content,
    correct_answer,
    skill_tag,
    difficulty_level,
    is_remedial
FROM questions
WHERE skill_tag = 'devops' AND is_remedial = TRUE
ORDER BY id DESC;

-- Summary
SELECT
    'Remedial DevOps questions inserted!' AS status,
    COUNT(*) AS total_remedial_questions
FROM questions
WHERE skill_tag = 'devops' AND is_remedial = TRUE;
