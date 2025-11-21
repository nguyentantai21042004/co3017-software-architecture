# Integration Tests

## Overview

Comprehensive end-to-end integration tests for the Intelligent Tutoring System (ITS), covering the complete adaptive learning flow across all four microservices.

## Test Scenarios

### 1. Complete Learning Flow - Low Mastery to High
**File**: `integration_test.go::TestCompleteLearningFlow_LowMasteryToHigh`

**Flow**:
1. Check initial mastery (0% for new user)
2. Get lesson recommendation → Remedial content (mastery < 50%)
3. Fetch question details from Content Service
4. Submit correct answer → Score = 100
5. Wait for async mastery update via RabbitMQ → New mastery = 50%
6. Get next lesson → Standard content (mastery ≥ 50%)
7. Submit another correct answer → Score = 100
8. Wait for mastery update → New mastery = 75%

**Verifies**:
- Adaptive Engine recommendation logic (remedial vs standard)
- Scoring Service grading (correct = 100, incorrect = 0)
- RabbitMQ event publishing and consumption
- Learner Model mastery calculation: `(old + new) / 2`
- Content type switching based on mastery threshold (50%)

### 2. Complete Learning Flow - High Mastery to Low
**File**: `integration_test.go::TestCompleteLearningFlow_HighMasteryToLow`

**Flow**:
1. Build up mastery to ~94% (if starting from 0)
2. Get lesson → Standard content (high mastery)
3. Submit incorrect answer → Score = 0
4. Wait for mastery update → Mastery drops to ~47%
5. Get next lesson → Remedial content (mastery < 50%)

**Verifies**:
- Mastery degradation with incorrect answers
- Content type switching from standard to remedial
- Adaptive logic responds to mastery changes

### 3. Multiple Submissions - Mastery Progression
**File**: `integration_test.go::TestMultipleSubmissions`

**Flow**:
1. Submit 3 correct answers in sequence
2. Track mastery progression: 0% → 50% → 75% → 87.5%
3. Verify each mastery calculation

**Verifies**:
- Consistent mastery formula application
- Mastery progression over multiple submissions
- No race conditions in async updates

## Prerequisites

### Services Running

All 4 microservices must be running locally:

```bash
# Terminal 1: Content Service (Java)
cd src/content
mvn spring-boot:run

# Terminal 2: Scoring Service (Go)
cd src/scoring
go run cmd/api/main.go

# Terminal 3: Learner Model Service (Go)
cd src/learner-model
go run cmd/api/main.go

# Terminal 4: Adaptive Engine (Go)
cd src/adaptive-engine
go run cmd/api/main.go
```

### Dependencies Running

- **PostgreSQL**: Running on `localhost:5432`
- **RabbitMQ**: Running on `localhost:5672`

### Test Data

The tests assume questions exist in the Content Service database with:
- Math questions (both remedial and standard)
- Science questions (both remedial and standard)

If questions don't exist, the tests will fail when trying to fetch question details.

## Running Tests

### Run All Integration Tests

```bash
cd src/tests/integration
go test -v -timeout 60s
```

### Run Specific Test

```bash
go test -v -run TestCompleteLearningFlow_LowMasteryToHigh
go test -v -run TestCompleteLearningFlow_HighMasteryToLow
go test -v -run TestMultipleSubmissions
```

### Run with Detailed Output

```bash
go test -v -timeout 60s 2>&1 | tee test_output.log
```

## Configuration

Edit `config.go` to change:
- Service URLs (default: localhost:8081-8084)
- Timeouts (default: 10s for mastery updates)
- Polling intervals (default: 200ms)

## Helper Functions

### API Calls
- `GetNextLesson(userID, skill)` - Call Adaptive Engine
- `SubmitAnswer(userID, questionID, answer)` - Call Scoring Service
- `GetMastery(userID, skill)` - Call Learner Model
- `GetQuestion(questionID)` - Call Content Service

### Async Handling
- `WaitForMasteryUpdate(userID, skill, expectedScore, timeout)` - Poll Learner Model until mastery reaches expected value or timeout

## Expected Output

```
Test: Complete Learning Flow - Low Mastery to High
User: integration-test-user-1, Skill: math

Step 1: Check initial mastery...
   Initial mastery: 0%

Step 2: Get first lesson recommendation...
   Recommended: Question ID 123, Type: remedial, Reason: Your mastery is 0%. Let's review the basics.

Step 3: Fetch question details...
   Question: What is 2 + 2?
   Correct Answer: 4

Step 4: Submit correct answer...
   Result: Chính xác! Bạn đã trả lời đúng. (Score: 100)

Step 5: Waiting for mastery update (async via RabbitMQ)...
   Mastery updated: 0% → 50%

Step 6: Get second lesson recommendation...
   Recommended: Question ID 456, Type: standard, Reason: Great! Your mastery is 50%. Continue with the next challenge.

Step 7: Submit second correct answer...
   Result: Chính xác! Bạn đã trả lời đúng. (Score: 100)

Step 8: Waiting for second mastery update...
   Mastery updated: 50% → 75%

Test completed successfully!
Final mastery: 75%
```

## Troubleshooting

### Tests Timeout Waiting for Mastery Update

**Cause**: RabbitMQ event not processed or Learner Model not consuming events.

**Solution**:
1. Check RabbitMQ logs: `docker logs rabbitmq`
2. Check Learner Model logs for event consumption
3. Verify RabbitMQ connection in Learner Model `.env`
4. Increase `MasteryUpdateTimeout` in `config.go`

### "Content Service returned 404"

**Cause**: No questions exist in database for the requested skill/type.

**Solution**:
1. Insert test questions into Content Service database
2. Use existing question IDs in tests
3. Check Content Service logs

### "Connection refused"

**Cause**: One or more services not running.

**Solution**:
1. Verify all 4 services are running
2. Check service ports match `config.go`
3. Check service health endpoints

## Cleanup

After tests, you may want to clean up test user data:

```sql
-- Clean up mastery data
DELETE FROM skill_mastery WHERE user_id LIKE 'integration-test-user-%';

-- Clean up submissions
DELETE FROM submissions WHERE user_id LIKE 'integration-test-user-%';
```

## Next Steps

- Add Docker Compose setup for automated test environment
- Add CI/CD pipeline integration
- Add more test scenarios (edge cases, error handling)
- Add performance/load testing
