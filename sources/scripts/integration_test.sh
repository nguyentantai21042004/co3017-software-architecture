#!/bin/bash

echo "=== Integration Test: RabbitMQ Event Flow ==="
echo ""

USER_ID="test_user_$(date +%s)"

echo "Testing with user: $USER_ID"
echo ""

echo "Step 1: Submit WRONG answer (score=0)"
echo "----------------------------------------"
RESPONSE1=$(curl -s -X POST "http://localhost:8082/api/scoring/submit" \
  -H "Content-Type: application/json" \
  -d "{\"user_id\":\"$USER_ID\",\"question_id\":1,\"answer\":\"C\"}")
echo "$RESPONSE1" | jq
echo ""

echo "Waiting 2 seconds for event processing..."
sleep 2

echo ""
echo "Step 2: Check mastery after wrong answer"
echo "----------------------------------------"
MASTERY1=$(curl -s "http://localhost:8080/internal/learner/$USER_ID/mastery?skill=math_algebra")
echo "$MASTERY1" | jq
SCORE1=$(echo "$MASTERY1" | jq -r '.data.mastery_score')
echo "Current mastery: $SCORE1"
echo ""

echo "Step 3: Submit CORRECT answer (score=100)"
echo "----------------------------------------"
RESPONSE2=$(curl -s -X POST "http://localhost:8082/api/scoring/submit" \
  -H "Content-Type: application/json" \
  -d "{\"user_id\":\"$USER_ID\",\"question_id\":1,\"answer\":\"A\"}")
echo "$RESPONSE2" | jq
echo ""

echo "Waiting 2 seconds for event processing..."
sleep 2

echo ""
echo "Step 4: Check mastery after correct answer"
echo "----------------------------------------"
MASTERY2=$(curl -s "http://localhost:8080/internal/learner/$USER_ID/mastery?skill=math_algebra")
echo "$MASTERY2" | jq
SCORE2=$(echo "$MASTERY2" | jq -r '.data.mastery_score')
echo "Updated mastery: $SCORE2"
echo ""

echo "=== Test Summary ==="
echo "User: $USER_ID"
echo "Initial mastery: $SCORE1 (after wrong answer, formula: (0+0)/2 = 0)"
echo "Final mastery: $SCORE2 (after correct answer, formula: ($SCORE1+100)/2 = $SCORE2)"
echo ""

if [ "$SCORE2" -gt "$SCORE1" ]; then
    echo "✅ SUCCESS: Mastery increased after correct answer!"
    echo "✅ RabbitMQ event flow is working correctly!"
else
    echo "❌ FAIL: Mastery did not increase"
    echo "Check consumer logs for errors"
fi
