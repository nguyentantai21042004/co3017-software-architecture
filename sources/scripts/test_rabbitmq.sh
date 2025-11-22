#!/bin/bash

# Test RabbitMQ Configuration
# This script verifies the RabbitMQ setup for the ITS system

echo "=== RabbitMQ Configuration Test ==="
echo ""

# RabbitMQ credentials
RABBITMQ_USER="admintest"
RABBITMQ_PASS="adminTest2025"
RABBITMQ_HOST="localhost"
RABBITMQ_PORT="15672"
RABBITMQ_API="http://${RABBITMQ_HOST}:${RABBITMQ_PORT}/api"

echo "1. Checking RabbitMQ connection..."
curl -s -u ${RABBITMQ_USER}:${RABBITMQ_PASS} ${RABBITMQ_API}/overview > /dev/null
if [ $? -eq 0 ]; then
    echo "RabbitMQ is running and accessible"
else
    echo "Cannot connect to RabbitMQ"
    exit 1
fi

echo ""
echo "2. Checking Exchange 'its.events'..."
EXCHANGE_EXISTS=$(curl -s -u ${RABBITMQ_USER}:${RABBITMQ_PASS} ${RABBITMQ_API}/exchanges/%2F/its.events | grep -c "its.events")
if [ $EXCHANGE_EXISTS -gt 0 ]; then
    echo "Exchange 'its.events' exists"
    curl -s -u ${RABBITMQ_USER}:${RABBITMQ_PASS} ${RABBITMQ_API}/exchanges/%2F/its.events | jq '{name:.name, type:.type, durable:.durable}'
else
    echo "Exchange 'its.events' does not exist"
fi

echo ""
echo "3. Checking Queue 'learner.updates'..."
QUEUE_EXISTS=$(curl -s -u ${RABBITMQ_USER}:${RABBITMQ_PASS} ${RABBITMQ_API}/queues/%2F/learner.updates | grep -c "learner.updates")
if [ $QUEUE_EXISTS -gt 0 ]; then
    echo "Queue 'learner.updates' exists"
    curl -s -u ${RABBITMQ_USER}:${RABBITMQ_PASS} ${RABBITMQ_API}/queues/%2F/learner.updates | jq '{name:.name, messages:.messages, consumers:.consumers, state:.state}'
else
    echo "Queue 'learner.updates' does not exist"
fi

echo ""
echo "4. Checking Queue Bindings..."
BINDINGS=$(curl -s -u ${RABBITMQ_USER}:${RABBITMQ_PASS} ${RABBITMQ_API}/queues/%2F/learner.updates/bindings)
echo "$BINDINGS" | jq '.[] | {source:.source, routing_key:.routing_key, destination:.destination}'

BINDING_COUNT=$(echo "$BINDINGS" | jq '. | length')
if [ $BINDING_COUNT -gt 0 ]; then
    echo "Queue has $BINDING_COUNT binding(s)"
else
    echo "Queue has no bindings"
fi

echo ""
echo "5. Publishing Test Message..."
curl -s -u ${RABBITMQ_USER}:${RABBITMQ_PASS} -X POST \
  ${RABBITMQ_API}/exchanges/%2F/its.events/publish \
  -H "Content-Type: application/json" \
  -d '{
    "properties":{},
    "routing_key":"event.submission",
    "payload":"{\"event\":\"SubmissionCompleted\",\"user_id\":\"test_user\",\"skill_tag\":\"test_skill\",\"score_obtained\":100,\"timestamp\":\"2025-11-22T12:00:00Z\"}",
    "payload_encoding":"string"
  }' | jq '{routed:.routed}'

echo ""
echo "6. Checking Queue Messages..."
sleep 1
curl -s -u ${RABBITMQ_USER}:${RABBITMQ_PASS} ${RABBITMQ_API}/queues/%2F/learner.updates | jq '{messages:.messages, messages_ready:.messages_ready, messages_unacknowledged:.messages_unacknowledged}'

echo ""
echo "=== Test Complete ==="
echo ""
echo "Summary:"
echo "- If 'routed: true' → Message was routed to queue"
echo "- If 'messages > 0' → Messages are waiting in queue"
echo "- If 'consumers > 0' → Consumer is connected"
echo ""
echo "To manually check RabbitMQ Management UI:"
echo "http://localhost:15672"
echo "Username: admintest"
echo "Password: adminTest2025"
