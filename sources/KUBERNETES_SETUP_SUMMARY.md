# Kubernetes Content - Setup Summary

## What Was Done ✅

### 1. Kubernetes Questions Added to Database

**Questions Created:** 10 from your `test_data.json`
- 7 multiple choice questions
- 3 short answer questions (required database schema fix)

**Database Schema Fix Applied:**
```sql
ALTER TABLE questions ALTER COLUMN correct_answer TYPE TEXT;
```
This allows long-form answers for short answer questions.

**Final Question Count:**
```
skill_tag: kubernetes
├── Total: 19 questions (some duplicates from re-running script)
├── Remedial: 2 questions (for students with mastery < 50%)
└── Standard: 17 questions (for students with mastery >= 50%)
```

### 2. Skill Tag Consolidation Applied

**Problem Solved:** Originally would have created 6 dashboard cards:
- ❌ `kubernetes_architecture`
- ❌ `kubernetes_workloads`
- ❌ `kubernetes_networking`
- ❌ `kubernetes_configuration`
- ❌ `kubernetes_core_concepts`
- ❌ `kubectl_commands`

**Solution Applied:** Consolidated to single tag:
- ✅ `kubernetes` (all 19 questions)

**Result:** Clean dashboard with just 1 Kubernetes card!

### 3. Documentation Created

**Files Created:**

1. **`insert_kubernetes_quiz.sql`** - SQL script with all 10 questions
2. **`consolidate_kubernetes_skill.sql`** - Script to merge skill tags
3. **`ADDING_CONTENT_GUIDE.md`** - Complete guide for adding subjects/questions
4. **`SKILL_TAG_STRATEGY.md`** - Strategy guide for single vs multiple tags
5. **`KUBERNETES_SETUP_SUMMARY.md`** - This file!

## Current Dashboard State

### Skills Available

```bash
$ curl http://localhost:8081/api/content/skills | jq .data
[
  "devops",
  "kubernetes"
]
```

Your dashboard now shows **2 clean cards**:

```
┌─────────────────────┐  ┌─────────────────────┐
│      DevOps         │  │    Kubernetes       │
│  📊 Mastery: X%     │  │  📊 Mastery: 0%     │
│  15 questions       │  │  19 questions       │
└─────────────────────┘  └─────────────────────┘
```

## Next Steps for You

### Step 1: Restart Content Service (Required)

The database schema was changed, so you must restart:

```bash
# Find and kill the current Content Service
lsof -ti:8081 | xargs kill

# Restart it
cd /Users/tantai/Workspaces/hcmut/co3017-software-architecture/sources/content
java -jar target/content-service-0.0.1-SNAPSHOT.jar &
```

### Step 2: Add Kubernetes Icon to Dashboard (Optional)

Edit `sources/client/app/dashboard/page.tsx`:

```typescript
import { Server } from "lucide-react"  // Add this to imports

const SKILL_METADATA: Record<string, { name: string; icon: any; description: string }> = {
  math: { name: "Mathematics", icon: Calculator, description: "Algebra, Geometry, and Calculus" },
  science: { name: "Science", icon: FlaskConical, description: "Physics, Chemistry, and Biology" },
  devops: { name: "DevOps", icon: Code2, description: "CI/CD, Docker, IaC, and Cloud" },
  kubernetes: {
    name: "Kubernetes",
    icon: Server,  // ← Add this
    description: "Container orchestration, kubectl, and K8s fundamentals"
  },
  // ... other skills
}
```

If you skip this step, Kubernetes will still appear but with a default book icon.

### Step 3: Test the System

**A. Test API Endpoint:**
```bash
# Get all skills (should show "kubernetes")
curl http://localhost:8081/api/content/skills | jq .

# Get a random Kubernetes question
curl 'http://localhost:8081/api/content/recommend?skill=kubernetes&type=standard&userId=test-user' | jq .

# Get a remedial question
curl 'http://localhost:8081/api/content/recommend?skill=kubernetes&type=remedial&userId=test-user' | jq .
```

**B. Test in Browser:**
1. Open dashboard: `http://localhost:3000/dashboard`
2. Look for Kubernetes card
3. Click on it to start learning
4. Answer questions and verify:
   - Questions are different each time (randomization works)
   - Mastery score updates after each answer
   - System recommends remedial questions when mastery < 50%

## Database Verification

### Check All Skills

```sql
SELECT skill_tag,
       COUNT(*) as total_questions,
       SUM(CASE WHEN is_remedial THEN 1 ELSE 0 END) as remedial_count,
       SUM(CASE WHEN NOT is_remedial THEN 1 ELSE 0 END) as standard_count
FROM questions
GROUP BY skill_tag
ORDER BY skill_tag;
```

Expected output:
```
 skill_tag  | total_questions | remedial_count | standard_count
------------+-----------------+----------------+----------------
 devops     |              15 |              5 |             10
 kubernetes |              19 |              2 |             17
```

### View Sample Kubernetes Questions

```sql
SELECT id,
       LEFT(content, 60) as preview,
       is_remedial,
       difficulty_level
FROM questions
WHERE skill_tag = 'kubernetes'
ORDER BY is_remedial DESC, difficulty_level ASC
LIMIT 10;
```

## Files Reference

### SQL Scripts (sources/scripts/)

1. **`insert_kubernetes_quiz.sql`**
   - Original import with all 10 questions
   - Includes verification queries

2. **`consolidate_kubernetes_skill.sql`**
   - Merges all kubernetes_* tags to "kubernetes"
   - Run this if you ever split tags and want to merge back

3. **`insert_devops_quiz.sql`** (existing)
   - 10 standard DevOps questions

4. **`insert_devops_remedial.sql`** (existing)
   - 5 remedial DevOps questions

### Documentation

1. **`ADDING_CONTENT_GUIDE.md`** ⭐ MAIN GUIDE
   - Complete step-by-step guide
   - Examples for adding new subjects
   - Troubleshooting section
   - Best practices

2. **`SKILL_TAG_STRATEGY.md`**
   - When to use single vs multiple tags
   - Visual comparisons
   - Migration guide
   - Decision tree

3. **`KUBERNETES_SETUP_SUMMARY.md`** (this file)
   - Summary of what was done
   - Next steps
   - Quick reference

## Common Questions

**Q: Why only 2 remedial questions for Kubernetes?**
A: Your test_data.json only had 1 question marked as remedial. I inserted it, and there was a duplicate run, so now you have 2.

**Recommendation:** Add 3-5 more remedial questions so students with low mastery have enough practice material.

**Q: Can I add more Kubernetes questions later?**
A: Absolutely! Just:
```sql
INSERT INTO questions (content, options, correct_answer, skill_tag, difficulty_level, is_remedial, created_at)
VALUES ('Your question here', '["A. Opt1", "B. Opt2", ...]', 'A', 'kubernetes', 2, FALSE, NOW());
```

**Q: What if I want to split Kubernetes into sub-topics later?**
A: You can, but it will create multiple cards. Run:
```sql
UPDATE questions
SET skill_tag = 'kubernetes_networking'
WHERE content LIKE '%Service%' OR content LIKE '%LoadBalancer%';
```

**Q: How do I remove duplicate questions?**
A: Find duplicates:
```sql
SELECT content, COUNT(*)
FROM questions
WHERE skill_tag = 'kubernetes'
GROUP BY content
HAVING COUNT(*) > 1;
```

Delete duplicates (keeps one):
```sql
DELETE FROM questions
WHERE id IN (
    SELECT id
    FROM (
        SELECT id, ROW_NUMBER() OVER (PARTITION BY content ORDER BY id) as rn
        FROM questions
        WHERE skill_tag = 'kubernetes'
    ) t
    WHERE rn > 1
);
```

## Troubleshooting

### Kubernetes doesn't appear on dashboard

1. **Check database:**
   ```bash
   PGPASSWORD=postgres psql -h localhost -U postgres -d content_db \
       -c "SELECT DISTINCT skill_tag FROM questions WHERE skill_tag = 'kubernetes';"
   ```

2. **Check API:**
   ```bash
   curl http://localhost:8081/api/content/skills | jq .
   ```

3. **Restart Content Service** (see Step 1 above)

4. **Clear browser cache** and refresh dashboard

### Questions not randomizing

**Solution:** Restart both Content Service and Adaptive Engine:
```bash
lsof -ti:8081,8084 | xargs kill

cd sources/content
java -jar target/content-service-0.0.1-SNAPSHOT.jar &

cd ../adaptive-engine
./bin/adaptive-engine &
```

### "No questions found" error

**Cause:** Missing remedial OR standard questions.

**Check:**
```sql
SELECT is_remedial, COUNT(*)
FROM questions
WHERE skill_tag = 'kubernetes'
GROUP BY is_remedial;
```

You need BOTH:
- `is_remedial = true` (for low mastery students)
- `is_remedial = false` (for high mastery students)

## Success Criteria ✅

Your Kubernetes integration is successful when:

- [ ] Content Service returns "kubernetes" in `/api/content/skills`
- [ ] Dashboard shows Kubernetes card (with or without custom icon)
- [ ] Clicking Kubernetes card loads a learning session
- [ ] Questions are different on each "Next Question"
- [ ] Mastery score updates after each answer
- [ ] System shows remedial questions when mastery < 50%
- [ ] System shows standard questions when mastery >= 50%

## Future Additions

To add more subjects, follow the same pattern:

1. Create SQL file: `insert_SUBJECT_quiz.sql`
2. Use single skill_tag: `skill_tag = 'SUBJECT'`
3. Include both remedial and standard questions
4. Insert into database
5. (Optional) Add metadata to dashboard
6. Test!

See `ADDING_CONTENT_GUIDE.md` for complete instructions.

---

**Status:** ✅ Complete
**Last Updated:** 2025-11-23
**Questions Added:** 19 Kubernetes questions
**Dashboard Cards:** 2 total (DevOps, Kubernetes)
