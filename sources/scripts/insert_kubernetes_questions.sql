-- =============================================================================
-- Kubernetes Questions - Complete Set
-- Purpose: Insert all Kubernetes-related questions into content_db
-- Includes: Standard questions and remedial questions
-- =============================================================================

-- Connect to content_db (assumes database already exists)
\c content_db

-- =============================================================================
-- KUBERNETES STANDARD QUESTIONS (10 questions)
-- =============================================================================

-- Question 1: Controller Manager (kubernetes_architecture, difficulty 2, standard)
INSERT INTO questions (content, options, correct_answer, skill_tag, difficulty_level, is_remedial, created_at)
VALUES (
    'Thành phần chính chịu trách nhiệm quản lý trạng thái mong muốn của một cluster (như số replica Pod) trong Kubernetes là gì?',
    '["A. Pod", "B. Deployment", "C. Controller Manager", "D. etcd"]',
    'C',
    'kubernetes',
    2,
    FALSE,
    NOW()
);

-- Question 2: Deployment for stateless apps (kubernetes_workloads, difficulty 1, standard)
INSERT INTO questions (content, options, correct_answer, skill_tag, difficulty_level, is_remedial, created_at)
VALUES (
    'Đối tượng Kubernetes nào thường được sử dụng nhất để triển khai và quản lý vòng đời của một ứng dụng stateless?',
    '["A. StatefulSet", "B. DaemonSet", "C. Job", "D. Deployment"]',
    'D',
    'kubernetes',
    1,
    FALSE,
    NOW()
);

-- Question 3: kubectl apply command (kubectl_commands, difficulty 1, standard)
INSERT INTO questions (content, options, correct_answer, skill_tag, difficulty_level, is_remedial, created_at)
VALUES (
    'Lệnh kubectl nào được sử dụng để tạo một Pod từ file cấu hình YAML?',
    '["A. kubectl run", "B. kubectl create", "C. kubectl apply -f", "D. kubectl exec"]',
    'C',
    'kubernetes',
    1,
    FALSE,
    NOW()
);

-- Question 4: Scheduler component (kubernetes_architecture, difficulty 2, standard)
INSERT INTO questions (content, options, correct_answer, skill_tag, difficulty_level, is_remedial, created_at)
VALUES (
    'Thành phần nào trong Kubernetes chịu trách nhiệm lập lịch (schedule) các Pod lên các Node?',
    '["A. Kubelet", "B. Kube-proxy", "C. Scheduler", "D. Controller Manager"]',
    'C',
    'kubernetes',
    2,
    FALSE,
    NOW()
);

-- Question 5: LoadBalancer Service (kubernetes_networking, difficulty 2, standard)
INSERT INTO questions (content, options, correct_answer, skill_tag, difficulty_level, is_remedial, created_at)
VALUES (
    'Loại Service nào trong Kubernetes hiển thị một địa chỉ IP cố định bên ngoài cluster và ánh xạ traffic đến các Pod?',
    '["A. ClusterIP", "B. NodePort", "C. LoadBalancer", "D. Headless Service"]',
    'C',
    'kubernetes',
    2,
    FALSE,
    NOW()
);

-- Question 6: ConfigMap usage (kubernetes_configuration, difficulty 2, standard)
INSERT INTO questions (content, options, correct_answer, skill_tag, difficulty_level, is_remedial, created_at)
VALUES (
    'ConfigMap trong Kubernetes chủ yếu được sử dụng để làm gì?',
    '["A. Lưu trữ thông tin nhạy cảm như mật khẩu", "B. Tách dữ liệu cấu hình khỏi container image", "C. Cấp quyền truy cập cho Pod", "D. Định nghĩa các policy mạng"]',
    'B',
    'kubernetes',
    2,
    FALSE,
    NOW()
);

-- Question 7: Pod as smallest unit (kubernetes_core_concepts, difficulty 1, standard)
INSERT INTO questions (content, options, correct_answer, skill_tag, difficulty_level, is_remedial, created_at)
VALUES (
    'Đối tượng nào sau đây ĐÚNG khi nói về đơn vị triển khai nhỏ nhất và có thể tạo được trong Kubernetes?',
    '["A. Node", "B. Container", "C. Pod", "D. Cluster"]',
    'C',
    'kubernetes',
    1,
    FALSE,
    NOW()
);

-- Question 8: kubectl create vs apply (kubectl_commands, difficulty 3, standard)
-- Note: This is a short answer question - no options
INSERT INTO questions (content, options, correct_answer, skill_tag, difficulty_level, is_remedial, created_at)
VALUES (
    'Hãy giải thích ngắn gọn sự khác biệt chính giữa `kubectl create` và `kubectl apply`.',
    NULL,
    'kubectl create được dùng để tạo mới một tài nguyên. kubectl apply được dùng để tạo mới HOẶC cập nhật một tài nguyên dựa trên file cấu hình, và nó lưu lại cấu hình đã áp dụng để quản lý future updates.',
    'kubernetes',
    3,
    FALSE,
    NOW()
);

-- Question 9: LoadBalancer on cloud (kubernetes_networking, difficulty 3, standard)
-- Note: This is a short answer question - no options
INSERT INTO questions (content, options, correct_answer, skill_tag, difficulty_level, is_remedial, created_at)
VALUES (
    'Một lập trình viên triển khai ứng dụng web và cần đảm bảo nó có thể truy cập được từ internet. Họ đã tạo một Deployment và một Service. Dịch vụ nên được cấu hình với kiểu (type) nào để đáp ứng yêu cầu này, và họ cần làm gì thêm trên môi trường cloud (như AWS, GCP)?',
    NULL,
    'Kiểu Service nên dùng là LoadBalancer. Trên môi trường cloud, khi tạo Service với type: LoadBalancer, Kubernetes sẽ tự động tạo một Load Balancer bên ngoài (ví dụ: Application Load Balancer trên AWS) mà cấp một IP/DNS public cho ứng dụng. Người dùng thường không cần thao tác thủ công nào khác, cloud provider sẽ xử lý.',
    'kubernetes',
    3,
    FALSE,
    NOW()
);

-- Question 10: Kubelet responsibilities (kubernetes_architecture, difficulty 2, REMEDIAL)
-- Note: This is a short answer question - no options
INSERT INTO questions (content, options, correct_answer, skill_tag, difficulty_level, is_remedial, created_at)
VALUES (
    'Hãy mô tả ngắn gọn nhiệm vụ chính của Kubelet trên mỗi Node worker trong cụm Kubernetes.',
    NULL,
    'Kubelet là một tác nhân (agent) chạy trên mỗi node. Nhiệm vụ chính của nó là đảm bảo các container được mô tả trong PodSpec (thông từ API server) đang chạy và khỏe mạnh (healthy) trên node đó.',
    'kubernetes',
    2,
    TRUE,
    NOW()
);

-- =============================================================================
-- VERIFICATION & SUMMARY
-- =============================================================================

-- Check all inserted Kubernetes questions
SELECT
    id,
    LEFT(content, 60) AS content_preview,
    correct_answer,
    skill_tag,
    difficulty_level,
    is_remedial
FROM questions
WHERE skill_tag = 'kubernetes'
ORDER BY id DESC
LIMIT 20;

-- Summary by remedial status
SELECT
    skill_tag,
    COUNT(*) AS total_questions,
    SUM(CASE WHEN is_remedial THEN 1 ELSE 0 END) AS remedial_count,
    SUM(CASE WHEN NOT is_remedial THEN 1 ELSE 0 END) AS standard_count,
    MIN(difficulty_level) AS min_difficulty,
    MAX(difficulty_level) AS max_difficulty
FROM questions
WHERE skill_tag = 'kubernetes'
GROUP BY skill_tag
ORDER BY skill_tag;

-- Overall summary
SELECT
    'Kubernetes questions inserted successfully!' AS status,
    COUNT(*) AS total_kubernetes_questions,
    SUM(CASE WHEN is_remedial THEN 1 ELSE 0 END) AS remedial_questions,
    SUM(CASE WHEN NOT is_remedial THEN 1 ELSE 0 END) AS standard_questions
FROM questions
WHERE skill_tag = 'kubernetes';

