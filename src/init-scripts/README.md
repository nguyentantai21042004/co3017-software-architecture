# Database Initialization Scripts

This folder contains SQL scripts to initialize the 3 PostgreSQL databases required for the ITS Microservices.

## 📊 Databases Overview

| Database | Service | Port | Purpose |
|----------|---------|------|---------|
| `content_db` | Content Service (Java) | 8081 | Questions, skills, learning content |
| `scoring_db` | Scoring Service (Golang) | 8082 | Submission records and scores |
| `learner_db` | Learner Model Service (Golang) | 8083 | User skill mastery (AI state) |

**Note:** Adaptive Engine (Port 8084) is stateless and has no database.

---

## 🚀 Quick Setup

### Option 1: Run All Scripts at Once

```bash
# From project root
psql -U postgres -h localhost -p 5432 -f src/init-scripts/01-init-content-db.sql
psql -U postgres -h localhost -p 5432 -f src/init-scripts/02-init-scoring-db.sql
psql -U postgres -h localhost -p 5432 -f src/init-scripts/03-init-learner-db.sql
```

### Option 2: Run One by One

```bash
# 1. Content DB
psql -U postgres -h localhost -p 5432 -f src/init-scripts/01-init-content-db.sql

# 2. Scoring DB
psql -U postgres -h localhost -p 5432 -f src/init-scripts/02-init-scoring-db.sql

# 3. Learner DB
psql -U postgres -h localhost -p 5432 -f src/init-scripts/03-init-learner-db.sql
```

### Option 3: Interactive psql

```bash
psql -U postgres -h localhost -p 5432

# Then run each script:
\i src/init-scripts/01-init-content-db.sql
\i src/init-scripts/02-init-scoring-db.sql
\i src/init-scripts/03-init-learner-db.sql
```

---

## ✅ Verification

After running the scripts, verify the databases:

```bash
# List all databases
psql -U postgres -h localhost -p 5432 -c "\l"

# Check content_db
psql -U postgres -h localhost -p 5432 -d content_db -c "SELECT COUNT(*) FROM questions;"

# Check scoring_db
psql -U postgres -h localhost -p 5432 -d scoring_db -c "SELECT COUNT(*) FROM submissions;"

# Check learner_db
psql -U postgres -h localhost -p 5432 -d learner_db -c "SELECT * FROM skill_mastery;"
```

---

## 📋 Seed Data Summary

### content_db (5 questions)
- **Question 1**: Hard algebra (ID=1, difficulty=3, remedial=false)
- **Question 2**: Easy algebra remedial (ID=2, difficulty=1, remedial=true) ⭐
- **Question 3**: Medium algebra (ID=3, difficulty=2, remedial=false)
- **Question 4**: Easy geometry remedial (ID=4, difficulty=1, remedial=true)
- **Question 5**: Hard geometry (ID=5, difficulty=3, remedial=false)

### scoring_db (2 submissions)
- User `user_01` answered question 3 correctly (score=100)
- User `user_01` answered question 2 correctly (score=100)

### learner_db (4 skill mastery records)
- `user_01` + `math_algebra` = 10 (Low - will trigger remedial) ⭐
- `user_01` + `math_geometry` = 0
- `user_02` + `math_algebra` = 80 (High)
- `user_02` + `math_geometry` = 55 (Medium)

---

## 🎯 Test Scenario (For Demo)

**Initial State:**
- User `user_01` has low algebra mastery (score=10)

**Test Flow:**
1. User answers Question 1 (hard) incorrectly
2. Scoring Service gives score=0
3. RabbitMQ event updates mastery: (10 + 0) / 2 = 5
4. Adaptive Engine detects score=5 < 50
5. **Recommends Question 2 (remedial=true)** ✨

This demonstrates the adaptive learning capability!

---

## 🔄 Reset Databases

To reset all data and start fresh:

```bash
# Method 1: Re-run all scripts (they include DROP DATABASE)
psql -U postgres -h localhost -p 5432 -f src/init-scripts/01-init-content-db.sql
psql -U postgres -h localhost -p 5432 -f src/init-scripts/02-init-scoring-db.sql
psql -U postgres -h localhost -p 5432 -f src/init-scripts/03-init-learner-db.sql

# Method 2: Manual drop
psql -U postgres -h localhost -p 5432 -c "DROP DATABASE IF EXISTS content_db;"
psql -U postgres -h localhost -p 5432 -c "DROP DATABASE IF EXISTS scoring_db;"
psql -U postgres -h localhost -p 5432 -c "DROP DATABASE IF EXISTS learner_db;"
```

---

## 🔗 Connection Details

```yaml
PostgreSQL:
  Host: localhost
  Port: 5432
  Username: postgres
  Password: postgres

Databases:
  - content_db
  - scoring_db
  - learner_db
```
