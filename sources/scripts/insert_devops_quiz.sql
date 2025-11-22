-- =====================================================================
-- DevOps Junior Quiz - 10 Questions
-- Converted from test_data.json
-- Topic: DevOps fundamentals for junior developers
-- =====================================================================

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

-- Verification query
SELECT
    id,
    content,
    correct_answer,
    skill_tag,
    difficulty_level,
    is_remedial
FROM questions
WHERE skill_tag = 'devops'
ORDER BY id DESC
LIMIT 10;

-- Summary
SELECT
    'DevOps questions inserted!' AS status,
    COUNT(*) AS total_devops_questions
FROM questions
WHERE skill_tag = 'devops';
