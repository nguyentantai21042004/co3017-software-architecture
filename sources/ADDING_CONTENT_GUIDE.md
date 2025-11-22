# Guide: Adding New Subjects and Questions

This guide explains how to add new learning subjects and questions to your Intelligent Tutoring System (ITS).

## Table of Contents
1. [Quick Start](#quick-start)
2. [Understanding the System](#understanding-the-system)
3. [Adding Questions to Existing Subjects](#adding-questions-to-existing-subjects)
4. [Adding a Completely New Subject](#adding-a-completely-new-subject)
5. [Testing Your Changes](#testing-your-changes)
6. [Troubleshooting](#troubleshooting)

---

## Quick Start

**TL;DR:** To add new content:
1. Insert questions into database using SQL
2. Add metadata to frontend (optional, for icons/descriptions)
3. Restart services (if needed)
4. Content appears automatically!

---

## Understanding the System

### How Content Discovery Works

Your system uses **dynamic content discovery**:

```
Database (questions table)
    ↓
Content Service (GET /api/content/skills)
    ↓
Frontend Dashboard (auto-fetches skills)
    ↓
User sees all skills automatically!
```

**Key Point:** You DON'T need to hardcode skills anywhere. The system automatically discovers them from the database.

### Database Schema

```sql
CREATE TABLE questions (
    id SERIAL PRIMARY KEY,
    content TEXT NOT NULL,
    options JSONB,              -- For multiple choice: ["A. Option1", "B. Option2", ...]
    correct_answer TEXT NOT NULL,
    skill_tag VARCHAR(100) NOT NULL,  -- ← This determines the subject!
    difficulty_level INTEGER,
    is_remedial BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW()
);
```

### Important Fields

| Field | Purpose | Example |
|-------|---------|---------|
| `skill_tag` | Identifies the subject | `"devops"`, `"kubernetes_architecture"` |
| `is_remedial` | For low-mastery students | `true` or `false` |
| `options` | Multiple choice options (JSONB) | `["A. Answer1", "B. Answer2"]` |
| `correct_answer` | The answer key | `"B"` or full text for short answer |

---

## Adding Questions to Existing Subjects

### Example: Adding More DevOps Questions

**Step 1: Create SQL File**

```sql
-- sources/scripts/insert_devops_advanced.sql
INSERT INTO questions (content, options, correct_answer, skill_tag, difficulty_level, is_remedial, created_at)
VALUES (
    'What is the purpose of a blue-green deployment?',
    '["A. To save costs", "B. To enable zero-downtime deployments", "C. To test in production", "D. To use two cloud providers"]',
    'B',
    'devops',  -- ← Same skill_tag as existing DevOps questions
    2,
    FALSE,
    NOW()
);

-- Add more questions...
```

**Step 2: Insert into Database**

```bash
PGPASSWORD=postgres psql -h localhost -U postgres -d content_db \
    -f sources/scripts/insert_devops_advanced.sql
```

**Step 3: Done!**

The questions are immediately available. No service restart needed for adding questions to existing subjects.

---

## Adding a Completely New Subject

Let's add a **"Cloud Computing"** subject as an example.

### Step 1: Design Your Skill Tags

Choose a consistent naming scheme:

**Option A: Single Skill Tag**
- `skill_tag = "cloud"`
- All cloud questions use the same tag

**Option B: Multiple Skill Tags (Recommended)**
- `skill_tag = "cloud_aws"`
- `skill_tag = "cloud_azure"`
- `skill_tag = "cloud_gcp"`
- Allows more granular progress tracking

For this example, we'll use **`"cloud"`**.

### Step 2: Create Questions (SQL File)

Create `sources/scripts/insert_cloud_quiz.sql`:

```sql
-- =====================================================================
-- Cloud Computing Quiz
-- =====================================================================

-- Standard Questions (for students with mastery >= 50)
INSERT INTO questions (content, options, correct_answer, skill_tag, difficulty_level, is_remedial, created_at)
VALUES
(
    'What does IaaS stand for in cloud computing?',
    '["A. Internet as a Service", "B. Infrastructure as a Service", "C. Integration as a Service", "D. Intelligence as a Service"]',
    'B',
    'cloud',
    1,
    FALSE,  -- ← Not remedial
    NOW()
),
(
    'Which cloud service model gives you the most control over the infrastructure?',
    '["A. SaaS", "B. PaaS", "C. IaaS", "D. FaaS"]',
    'C',
    'cloud',
    2,
    FALSE,
    NOW()
);

-- Remedial Questions (for students with mastery < 50)
INSERT INTO questions (content, options, correct_answer, skill_tag, difficulty_level, is_remedial, created_at)
VALUES
(
    'What is cloud computing?',
    '["A. Storing data on your local computer", "B. Delivering computing services over the internet", "C. A type of weather prediction", "D. A programming language"]',
    'B',
    'cloud',
    1,
    TRUE,  -- ← Remedial question
    NOW()
),
(
    'What is the main benefit of cloud computing?',
    '["A. Faster internet", "B. Pay only for what you use", "C. Better graphics", "D. Offline access"]',
    'B',
    'cloud',
    1,
    TRUE,
    NOW()
);

-- Verification
SELECT skill_tag, COUNT(*),
       SUM(CASE WHEN is_remedial THEN 1 ELSE 0 END) as remedial_count
FROM questions
WHERE skill_tag = 'cloud'
GROUP BY skill_tag;
```

### Step 3: Insert into Database

```bash
PGPASSWORD=postgres psql -h localhost -U postgres -d content_db \
    -f sources/scripts/insert_cloud_quiz.sql
```

### Step 4: Add Frontend Metadata (Optional but Recommended)

Edit `sources/client/app/dashboard/page.tsx`:

```typescript
const SKILL_METADATA: Record<string, { name: string; icon: any; description: string }> = {
  math: { name: "Mathematics", icon: Calculator, description: "Algebra, Geometry, and Calculus" },
  science: { name: "Science", icon: FlaskConical, description: "Physics, Chemistry, and Biology" },
  devops: { name: "DevOps", icon: Code2, description: "CI/CD, Docker, IaC, and Cloud" },
  cloud: { name: "Cloud Computing", icon: Cloud, description: "AWS, Azure, GCP, and Cloud Architecture" },  // ← Add this
  // ... other skills
}
```

**Don't forget to import the icon:**

```typescript
import { Calculator, FlaskConical, Code2, Cloud, BookOpen, History } from "lucide-react"
```

### Step 5: Restart Services (Only if you modified frontend)

```bash
# If you modified the frontend (dashboard/page.tsx)
cd sources/client
npm run dev  # Or restart your Next.js server

# Backend services auto-discover new skills, no restart needed!
```

### Step 6: Verify

1. Open dashboard: `http://localhost:3000/dashboard`
2. You should see "Cloud Computing" appear automatically
3. Click on it to start learning

---

## Full Example: Adding Kubernetes (From Your test_data.json)

### Step 1: SQL File Already Created ✅

You have `sources/scripts/insert_kubernetes_quiz.sql` ready!

### Step 2: Insert into Database

```bash
PGPASSWORD=postgres psql -h localhost -U postgres -d content_db \
    -f sources/scripts/insert_kubernetes_quiz.sql
```

Expected output:
```
INSERT 0 1
INSERT 0 1
...
 id |                        content_preview                         | correct_answer |       skill_tag        | difficulty_level | is_remedial
----+----------------------------------------------------------------+----------------+------------------------+------------------+-------------
 90 | Hãy mô tả ngắn gọn nhiệm vụ chính của Kubelet...              | Kubelet là...  | kubernetes_architecture|                2 | t
 89 | Một lập trình viên triển khai ứng dụng web...                 | Kiểu Service...| kubernetes_networking  |                3 | f
...
```

### Step 3: Add Frontend Metadata

Edit `sources/client/app/dashboard/page.tsx`:

```typescript
import { Server } from "lucide-react"  // ← Add this import

const SKILL_METADATA: Record<string, { name: string; icon: any; description: string }> = {
  // ... existing skills
  kubernetes_architecture: {
    name: "Kubernetes Architecture",
    icon: Server,
    description: "K8s components, Controllers, and Schedulers"
  },
  kubernetes_workloads: {
    name: "Kubernetes Workloads",
    icon: Server,
    description: "Deployments, StatefulSets, DaemonSets"
  },
  kubernetes_networking: {
    name: "Kubernetes Networking",
    icon: Server,
    description: "Services, Ingress, and Network Policies"
  },
  kubernetes_configuration: {
    name: "Kubernetes Configuration",
    icon: Server,
    description: "ConfigMaps, Secrets, and Resource Management"
  },
  kubernetes_core_concepts: {
    name: "Kubernetes Basics",
    icon: Server,
    description: "Pods, Nodes, and Core Concepts"
  },
  kubectl_commands: {
    name: "kubectl Commands",
    icon: Server,
    description: "kubectl operations and CLI usage"
  },
}
```

**Note:** You have **6 different skill tags** for Kubernetes. They will appear as separate subjects on the dashboard.

**Alternative:** Use a single skill tag like `"kubernetes"` if you want them grouped together.

### Step 4: Restart Frontend (if modified)

```bash
# Only if you edited dashboard/page.tsx
cd sources/client
# Stop the dev server (Ctrl+C) and restart
npm run dev
```

### Step 5: Test

1. Open dashboard: `http://localhost:3000/dashboard`
2. You should see 6 new Kubernetes-related subjects
3. Click any to start learning

---

## Testing Your Changes

### 1. Verify Questions in Database

```bash
PGPASSWORD=postgres psql -h localhost -U postgres -d content_db
```

```sql
-- List all skill tags
SELECT DISTINCT skill_tag FROM questions ORDER BY skill_tag;

-- Count questions per skill
SELECT skill_tag,
       COUNT(*) as total,
       SUM(CASE WHEN is_remedial THEN 1 ELSE 0 END) as remedial,
       SUM(CASE WHEN NOT is_remedial THEN 1 ELSE 0 END) as standard
FROM questions
GROUP BY skill_tag
ORDER BY skill_tag;

-- View specific subject questions
SELECT id, LEFT(content, 50) as preview, correct_answer, is_remedial
FROM questions
WHERE skill_tag = 'kubernetes_architecture'
LIMIT 10;
```

### 2. Test Content Service API

```bash
# Get all available skills
curl http://localhost:8081/api/content/skills | jq .

# Get a specific question
curl http://localhost:8081/api/content/123 | jq .

# Test recommendation (remedial)
curl 'http://localhost:8081/api/content/recommend?skill=kubernetes_architecture&type=remedial&userId=student-123' | jq .

# Test recommendation (standard)
curl 'http://localhost:8081/api/content/recommend?skill=kubernetes_architecture&type=standard&userId=student-123' | jq .
```

### 3. Test Adaptive Engine

```bash
# Get next lesson recommendation
curl -X POST http://localhost:8084/api/adaptive/next-lesson \
  -H 'Content-Type: application/json' \
  -d '{
    "user_id": "student-123",
    "current_skill": "kubernetes_architecture"
  }' | jq .
```

### 4. Test Frontend

1. Login: `http://localhost:3000`
2. Dashboard should show new subjects
3. Click on subject to start learning
4. Answer questions and verify:
   - Different questions each time (randomization works)
   - Mastery score updates after each answer
   - Transitions from remedial → standard when mastery >= 50

---

## Best Practices

### 1. Question Quality

✅ **Do:**
- Create at least 5 remedial + 5 standard questions per subject
- Use clear, unambiguous language
- Provide 4 options for multiple choice (A, B, C, D)
- Test questions yourself before deploying

❌ **Don't:**
- Create only standard OR only remedial (need both!)
- Use options with identical meanings
- Make questions too easy or impossibly hard

### 2. Skill Tag Naming

✅ **Good Examples:**
- `devops` - Simple, clear
- `kubernetes_networking` - Descriptive, underscore-separated
- `python_basics` - Consistent naming

❌ **Bad Examples:**
- `Kubernetes Networking` - Contains spaces
- `k8s-net` - Abbreviations, unclear
- `KUBERNETES_NETWORKING` - All caps

### 3. Options Format

**Multiple Choice Questions:**
```sql
options: '["A. First option", "B. Second option", "C. Third option", "D. Fourth option"]'
correct_answer: 'B'  -- Just the letter
```

**Short Answer Questions:**
```sql
options: NULL  -- No options for short answer
correct_answer: 'Full text of the correct answer here'
```

### 4. Difficulty Levels

- `difficulty_level = 1`: Basic concepts, definitions
- `difficulty_level = 2`: Application, understanding
- `difficulty_level = 3`: Analysis, complex scenarios

### 5. Remedial vs Standard

**Remedial Questions (is_remedial = TRUE):**
- Shown when mastery < 50%
- Very basic concepts
- Should be easier than standard

**Standard Questions (is_remedial = FALSE):**
- Shown when mastery >= 50%
- More challenging
- Real-world scenarios

---

## Common Workflows

### Workflow 1: Quick Subject Addition

```bash
# 1. Create SQL file
cat > sources/scripts/insert_python_quiz.sql << 'EOF'
INSERT INTO questions (content, options, correct_answer, skill_tag, difficulty_level, is_remedial, created_at)
VALUES
('What is Python?', '["A. A snake", "B. A programming language", "C. A framework", "D. A database"]', 'B', 'python', 1, TRUE, NOW()),
('What is a list in Python?', '["A. A data type", "B. An ordered collection", "C. A function", "D. A module"]', 'B', 'python', 1, FALSE, NOW());
EOF

# 2. Insert
PGPASSWORD=postgres psql -h localhost -U postgres -d content_db -f sources/scripts/insert_python_quiz.sql

# 3. Done! Check dashboard
```

### Workflow 2: Converting JSON to SQL

If you have questions in JSON format (like test_data.json):

```bash
# Use the Kubernetes SQL file as a template
# Key conversion:
# - question_type: "mcq" → options: '["A. ...", "B. ..."]'
# - question_type: "short_answer" → options: NULL
# - skill_tag from JSON → skill_tag in SQL (keep as-is)
# - is_remedial from JSON → is_remedial in SQL
```

### Workflow 3: Bulk Import Multiple Subjects

```bash
# Create a master SQL file
cat sources/scripts/insert_devops_quiz.sql \
    sources/scripts/insert_kubernetes_quiz.sql \
    sources/scripts/insert_python_quiz.sql \
    > sources/scripts/insert_all_content.sql

# Import all at once
PGPASSWORD=postgres psql -h localhost -U postgres -d content_db \
    -f sources/scripts/insert_all_content.sql
```

---

## Troubleshooting

### Problem: New subject doesn't appear on dashboard

**Solutions:**
1. Check database:
   ```sql
   SELECT DISTINCT skill_tag FROM questions;
   ```
2. Test Content Service:
   ```bash
   curl http://localhost:8081/api/content/skills | jq .
   ```
3. Clear browser cache and refresh
4. Check browser console for errors

### Problem: Questions not randomizing

**Cause:** Services not restarted after code changes.

**Solution:**
```bash
# Restart Content Service
cd sources/content
./mvnw spring-boot:run &

# Restart Adaptive Engine
cd sources/adaptive-engine
./bin/adaptive-engine &
```

### Problem: "No questions found" error

**Cause:** No questions with matching `skill_tag` and `is_remedial` combination.

**Solution:**
```sql
-- Check what exists
SELECT skill_tag, is_remedial, COUNT(*)
FROM questions
WHERE skill_tag = 'your_skill_tag'
GROUP BY skill_tag, is_remedial;

-- You need BOTH:
-- - Questions with is_remedial = TRUE (for low mastery)
-- - Questions with is_remedial = FALSE (for high mastery)
```

### Problem: Options not displaying correctly

**Check:**
1. Options format is valid JSON: `'["A. Text", "B. Text"]'`
2. Quotes are escaped properly
3. No trailing commas in JSON array

**Fix:**
```sql
-- Bad
options: '["A. Option1", "B. Option2",]'  -- Trailing comma!

-- Good
options: '["A. Option1", "B. Option2"]'
```

---

## Summary Checklist

When adding new content:

- [ ] Create SQL file in `sources/scripts/`
- [ ] Include both remedial (is_remedial=TRUE) and standard (is_remedial=FALSE) questions
- [ ] Use consistent skill_tag naming (lowercase, underscores)
- [ ] Insert questions into database
- [ ] (Optional) Add metadata to `sources/client/app/dashboard/page.tsx`
- [ ] Test via curl commands
- [ ] Test in browser
- [ ] Verify mastery updates correctly
- [ ] Verify question randomization works

---

## Quick Reference Commands

```bash
# Insert questions
PGPASSWORD=postgres psql -h localhost -U postgres -d content_db -f sources/scripts/YOUR_FILE.sql

# View all skills
PGPASSWORD=postgres psql -h localhost -U postgres -d content_db -c "SELECT DISTINCT skill_tag FROM questions ORDER BY skill_tag;"

# Test Content Service
curl http://localhost:8081/api/content/skills | jq .

# Test recommendation
curl 'http://localhost:8081/api/content/recommend?skill=SKILL_NAME&type=remedial&userId=student-123' | jq .

# Test Adaptive Engine
curl -X POST http://localhost:8084/api/adaptive/next-lesson -H 'Content-Type: application/json' -d '{"user_id":"student-123","current_skill":"SKILL_NAME"}' | jq .
```

---

## Need Help?

- Check logs: `sources/content/logs/` and `sources/adaptive-engine/logs/`
- Review database schema: `sources/db/schema.sql`
- Examine existing questions: Check `sources/scripts/insert_*.sql` files
- Test each component individually using curl commands above

---

**Last Updated:** 2025-11-23
**System Version:** ITS v1.0
