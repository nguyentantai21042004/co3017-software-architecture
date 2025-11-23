-- =============================================================================
-- DevOps Questions - Complete Set
-- Purpose: Insert all DevOps-related questions into content_db
-- Includes: Standard questions (10) and remedial questions (5)
-- =============================================================================

-- Connect to content_db (assumes database already exists)
\c content_db

-- =============================================================================
-- DEVOPS STANDARD QUESTIONS (10 questions)
-- =============================================================================

-- Question 1: DevOps basics
INSERT INTO questions (content, options, correct_answer, skill_tag, difficulty_level, is_remedial, created_at)
VALUES (
    'Mục tiêu chính của DevOps là gì?',
    '["A. Tách biệt hoàn toàn team Dev và Ops", "B. Tự động hoá và tăng khả năng cộng tác giữa Dev và Ops", "C. Chỉ tập trung vào viết code nhanh hơn", "D. Chỉ tập trung vào bảo mật hệ thống"]',
    'B',
    'devops',
    1,
    FALSE,
    NOW()
);

-- Question 2: CI/CD
INSERT INTO questions (content, options, correct_answer, skill_tag, difficulty_level, is_remedial, created_at)
VALUES (
    'CI trong DevOps là viết tắt của cụm từ nào?',
    '["A. Constant Integration", "B. Continuous Integration", "C. Central Integration", "D. Continuous Improvement"]',
    'B',
    'devops',
    1,
    FALSE,
    NOW()
);

-- Question 3: Container/Docker
INSERT INTO questions (content, options, correct_answer, skill_tag, difficulty_level, is_remedial, created_at)
VALUES (
    'Công cụ nào sau đây thường được dùng để containerize ứng dụng?',
    '["A. Git", "B. Docker", "C. Ansible", "D. Jenkins"]',
    'B',
    'devops',
    1,
    FALSE,
    NOW()
);

-- Question 4: CI/CD benefits
INSERT INTO questions (content, options, correct_answer, skill_tag, difficulty_level, is_remedial, created_at)
VALUES (
    'Lợi ích chính của CI/CD pipeline là gì?',
    '["A. Giảm số lượng test cần viết", "B. Tự động build, test và deploy, giảm lỗi do thao tác tay", "C. Không cần review code nữa", "D. Chỉ deploy được một lần mỗi tháng"]',
    'B',
    'devops',
    2,
    FALSE,
    NOW()
);

-- Question 5: Infrastructure as Code concept
INSERT INTO questions (content, options, correct_answer, skill_tag, difficulty_level, is_remedial, created_at)
VALUES (
    'Khái niệm Infrastructure as Code (IaC) mô tả điều gì?',
    '["A. Viết tài liệu hạ tầng bằng file Word", "B. Cấu hình hạ tầng bằng cách click chuột trên giao diện", "C. Định nghĩa và quản lý hạ tầng bằng mã (code/script) có thể version control", "D. Thuê nhà cung cấp bên ngoài quản lý hạ tầng"]',
    'C',
    'devops',
    2,
    FALSE,
    NOW()
);

-- Question 6: IaC tools
INSERT INTO questions (content, options, correct_answer, skill_tag, difficulty_level, is_remedial, created_at)
VALUES (
    'Công cụ nào sau đây là ví dụ điển hình cho Infrastructure as Code?',
    '["A. Terraform", "B. Docker", "C. Git", "D. Slack"]',
    'A',
    'devops',
    2,
    FALSE,
    NOW()
);

-- Question 7: Monitoring
INSERT INTO questions (content, options, correct_answer, skill_tag, difficulty_level, is_remedial, created_at)
VALUES (
    'Trong bối cảnh DevOps, mục đích chính của việc monitoring hệ thống là gì?',
    '["A. Chỉ để thu thập log cho đầy ổ đĩa", "B. Phát hiện sớm sự cố, theo dõi performance và độ ổn định của hệ thống", "C. Thay thế hoàn toàn việc test", "D. Tăng chi phí vận hành"]',
    'B',
    'devops',
    2,
    FALSE,
    NOW()
);

-- Question 8: Rollback
INSERT INTO questions (content, options, correct_answer, skill_tag, difficulty_level, is_remedial, created_at)
VALUES (
    'Khi deploy phiên bản mới gây lỗi, khái niệm ''rollback'' nghĩa là gì?',
    '["A. Dừng luôn hệ thống để điều tra", "B. Xoá toàn bộ dữ liệu người dùng", "C. Quay lại chạy phiên bản ổn định trước đó", "D. Tắt monitoring để không thấy lỗi"]',
    'C',
    'devops',
    2,
    FALSE,
    NOW()
);

-- Question 9: Git basics
INSERT INTO questions (content, options, correct_answer, skill_tag, difficulty_level, is_remedial, created_at)
VALUES (
    'Trong Git, lệnh nào dùng để đẩy (upload) commit từ local lên remote repository (như GitHub, GitLab)?',
    '["A. git pull", "B. git push", "C. git clone", "D. git status"]',
    'B',
    'devops',
    1,
    FALSE,
    NOW()
);

-- Question 10: Cloud basics
INSERT INTO questions (content, options, correct_answer, skill_tag, difficulty_level, is_remedial, created_at)
VALUES (
    'Trong các loại dịch vụ cloud sau, loại nào cung cấp máy ảo, mạng, lưu trữ để tự cài hệ điều hành và phần mềm?',
    '["A. SaaS (Software as a Service)", "B. PaaS (Platform as a Service)", "C. IaaS (Infrastructure as a Service)", "D. FaaS (Function as a Service)"]',
    'C',
    'devops',
    2,
    FALSE,
    NOW()
);

-- =============================================================================
-- DEVOPS REMEDIAL QUESTIONS (5 questions - Basic level)
-- =============================================================================

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

-- =============================================================================
-- VERIFICATION & SUMMARY
-- =============================================================================

-- Check all inserted DevOps questions
SELECT
    id,
    LEFT(content, 60) AS content_preview,
    correct_answer,
    skill_tag,
    difficulty_level,
    is_remedial
FROM questions
WHERE skill_tag = 'devops'
ORDER BY is_remedial, id DESC;

-- Summary by remedial status
SELECT
    skill_tag,
    COUNT(*) AS total_questions,
    SUM(CASE WHEN is_remedial THEN 1 ELSE 0 END) AS remedial_count,
    SUM(CASE WHEN NOT is_remedial THEN 1 ELSE 0 END) AS standard_count,
    MIN(difficulty_level) AS min_difficulty,
    MAX(difficulty_level) AS max_difficulty
FROM questions
WHERE skill_tag = 'devops'
GROUP BY skill_tag
ORDER BY skill_tag;

-- Overall summary
SELECT
    'DevOps questions inserted successfully!' AS status,
    COUNT(*) AS total_devops_questions,
    SUM(CASE WHEN is_remedial THEN 1 ELSE 0 END) AS remedial_questions,
    SUM(CASE WHEN NOT is_remedial THEN 1 ELSE 0 END) AS standard_questions
FROM questions
WHERE skill_tag = 'devops';

