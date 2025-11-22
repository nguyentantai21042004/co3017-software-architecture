#!/bin/bash

# Reset RabbitMQ Configuration
# This script deletes and recreates the queue/exchange setup

echo "=== RabbitMQ Reset Script ==="
echo ""

# RabbitMQ credentials
RABBITMQ_USER="admintest"
RABBITMQ_PASS="adminTest2025"
RABBITMQ_HOST="localhost"
RABBITMQ_PORT="15672"
RABBITMQ_API="http://${RABBITMQ_HOST}:${RABBITMQ_PORT}/api"

echo "This will DELETE and RECREATE the RabbitMQ configuration"
echo "Press Ctrl+C to cancel, or Enter to continue..."
read

echo ""
echo "1. Deleting queue 'learner.updates'..."
curl -s -u ${RABBITMQ_USER}:${RABBITMQ_PASS} -X DELETE \
  ${RABBITMQ_API}/queues/%2F/learner.updates
echo "Queue deleted (if it existed)"

echo ""
echo "2. Deleting exchange 'its.events'..."
curl -s -u ${RABBITMQ_USER}:${RABBITMQ_PASS} -X DELETE \
  ${RABBITMQ_API}/exchanges/%2F/its.events
echo "Exchange deleted (if it existed)"

echo ""
echo "3. Creating exchange 'its.events'..."
curl -s -u ${RABBITMQ_USER}:${RABBITMQ_PASS} -X PUT \
  ${RABBITMQ_API}/exchanges/%2F/its.events \
  -H "Content-Type: application/json" \
  -d '{
    "type":"topic",
    "durable":true,
    "auto_delete":false,
    "internal":false,
    "arguments":{}
  }'
echo "Exchange created"

echo ""
echo "4. Creating queue 'learner.updates'..."
curl -s -u ${RABBITMQ_USER}:${RABBITMQ_PASS} -X PUT \
  ${RABBITMQ_API}/queues/%2F/learner.updates \
  -H "Content-Type: application/json" \
  -d '{
    "durable":true,
    "auto_delete":false,
    "arguments":{}
  }'
echo "Queue created"

echo ""
echo "5. Binding queue to exchange..."
curl -s -u ${RABBITMQ_USER}:${RABBITMQ_PASS} -X POST \
  ${RABBITMQ_API}/bindings/%2F/e/its.events/q/learner.updates \
  -H "Content-Type: application/json" \
  -d '{
    "routing_key":"event.submission",
    "arguments":{}
  }'
echo "Queue bound to exchange"

echo ""
echo "6. Verifying configuration..."
echo ""
echo "Exchange:"
curl -s -u ${RABBITMQ_USER}:${RABBITMQ_PASS} ${RABBITMQ_API}/exchanges/%2F/its.events | jq '{name:.name, type:.type, durable:.durable}'

echo ""
echo "Queue:"
curl -s -u ${RABBITMQ_USER}:${RABBITMQ_PASS} ${RABBITMQ_API}/queues/%2F/learner.updates | jq '{name:.name, durable:.durable, messages:.messages}'

echo ""
echo "Bindings:"
curl -s -u ${RABBITMQ_USER}:${RABBITMQ_PASS} ${RABBITMQ_API}/queues/%2F/learner.updates/bindings | jq '.[] | {source:.source, routing_key:.routing_key}'

echo ""
echo "=== Reset Complete ==="
echo ""
echo "Next steps:"
echo "1. Restart learner-model service (consumer)"
echo "2. Restart scoring service (publisher)"
echo "3. Test with: ./test_rabbitmq.sh"
