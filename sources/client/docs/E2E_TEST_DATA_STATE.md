# E2E Test Data State Documentation

This document describes the expected state of test data for E2E testing.

## Overview

E2E tests require specific test data to be present in the databases. This data is set up using the `setup-test-data.sh` script and can be cleaned up using `cleanup-test-data.sh`.

## Content Service Test Data (content_db)

### Database: `content_db`
### Table: `questions`

#### Math Questions

**Remedial Questions (5 questions, difficulty_level = 1):**
1. "What is 2 + 2?" - Options: ["2", "3", "4", "5"] - Correct: "4"
2. "What is 5 - 3?" - Options: ["1", "2", "3", "4"] - Correct: "2"
3. "What is 3 × 4?" - Options: ["9", "10", "12", "15"] - Correct: "12"
4. "What is 10 ÷ 2?" - Options: ["2", "5", "10", "20"] - Correct: "5"
5. "What is 7 + 8?" - Options: ["14", "15", "16", "17"] - Correct: "15"

**Standard Questions (5 questions, difficulty_level = 2):**
1. "Solve: 2x + 5 = 13. What is x?" - Options: ["2", "4", "6", "8"] - Correct: "4"
2. "What is the square root of 144?" - Options: ["10", "11", "12", "13"] - Correct: "12"
3. "Calculate: (3 + 5) × 2 - 4" - Options: ["8", "10", "12", "16"] - Correct: "12"
4. "What is 15% of 200? (Enter number only)" - Free text - Correct: "30"
5. "If a triangle has angles 60° and 70°, what is the third angle? (Enter number only)" - Free text - Correct: "50"

#### Science Questions

**Remedial Questions (5 questions, difficulty_level = 1):**
1. "What is the chemical symbol for water?" - Options: ["H2O", "CO2", "O2", "H2"] - Correct: "H2O"
2. "How many planets are in our solar system?" - Options: ["6", "7", "8", "9"] - Correct: "8"
3. "What gas do plants absorb from the atmosphere?" - Options: ["O2", "N2", "CO2", "H2"] - Correct: "CO2"
4. "What is the freezing point of water in Celsius?" - Options: ["-10", "0", "10", "100"] - Correct: "0"
5. "What force pulls objects toward Earth?" - Options: ["magnetism", "gravity", "friction", "tension"] - Correct: "gravity"

**Standard Questions (5 questions, difficulty_level = 2):**
1. "What is the powerhouse of the cell?" - Options: ["nucleus", "mitochondria", "ribosome", "chloroplast"] - Correct: "mitochondria"
2. "What is the atomic number of Carbon?" - Options: ["4", "6", "8", "12"] - Correct: "6"
3. "What is the smallest unit of life?" - Options: ["atom", "molecule", "cell", "tissue"] - Correct: "cell"
4. "What is the speed of light in vacuum (in km/s)? (Enter number only)" - Free text - Correct: "300000"
5. "What is the process by which plants make food?" - Free text - Correct: "photosynthesis"

### Total Questions
- Math: 10 questions (5 remedial + 5 standard)
- Science: 10 questions (5 remedial + 5 standard)
- **Total: 20 questions**

## Learner Model Service Test Data (learner_db)

### Database: `learner_db`
### Table: `skill_mastery`

#### Test User: `test-user-123`

This user ID matches the test user configured in E2E tests (set via localStorage in test setup).

**Skill Mastery Records:**

| user_id | skill_tag | current_score | last_updated |
|---------|-----------|---------------|--------------|
| test-user-123 | math | 50 | Current timestamp |
| test-user-123 | science | 60 | Current timestamp |

### Expected Behavior

- **Math skill**: 50% mastery
  - Should receive standard questions (not remedial, as mastery is above threshold)
  - Adaptive Engine should recommend standard difficulty content

- **Science skill**: 60% mastery
  - Should receive standard questions
  - Adaptive Engine should recommend standard difficulty content

## Test Data Verification

### Verify Content Service Data

```bash
# Connect to PostgreSQL
docker exec -it its-postgres psql -U postgres -d content_db

# Count questions by skill
SELECT skill_tag, is_remedial, COUNT(*) as count 
FROM questions 
WHERE skill_tag IN ('math', 'science')
GROUP BY skill_tag, is_remedial 
ORDER BY skill_tag, is_remedial;

# View all test questions
SELECT id, skill_tag, is_remedial, difficulty_level, content 
FROM questions 
WHERE skill_tag IN ('math', 'science')
ORDER BY skill_tag, is_remedial, difficulty_level;
```

### Verify Learner Model Data

```bash
# Connect to PostgreSQL
docker exec -it its-postgres psql -U postgres -d learner_db

# View test user mastery
SELECT * FROM skill_mastery WHERE user_id = 'test-user-123';

# Verify mastery scores
SELECT user_id, skill_tag, current_score, last_updated 
FROM skill_mastery 
WHERE user_id = 'test-user-123'
ORDER BY skill_tag;
```

## Test Data Setup Script

The `setup-test-data.sh` script performs the following:

1. **Content Service Setup**:
   - Connects to `content_db`
   - Inserts test questions using `insert_e2e_test_data.sql`
   - Verifies questions were inserted

2. **Learner Model Service Setup**:
   - Connects to `learner_db`
   - Creates/updates test user mastery records
   - Sets math skill to 50%
   - Sets science skill to 60%
   - Verifies mastery records exist

## Test Data Cleanup

The `cleanup-test-data.sh` script removes:

1. **Content Service**: All questions with `skill_tag IN ('math', 'science')`
2. **Learner Model Service**: All mastery records for `user_id = 'test-user-123'`

## Notes

- Test data uses `ON CONFLICT DO NOTHING` to avoid duplicates
- Test user ID `test-user-123` is hardcoded in E2E tests
- Mastery scores are set to specific values to test adaptive behavior
- Questions include both multiple-choice and free-text formats
- Questions cover both remedial and standard difficulty levels

