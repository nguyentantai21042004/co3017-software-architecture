# =============================================================================
# Makefile for ITS Microservices - Docker Management
# =============================================================================

.PHONY: help build up down restart logs clean test rebuild status health db-init infra services

# Default target
.DEFAULT_GOAL := help

# Colors for output
GREEN  := \033[0;32m
YELLOW := \033[0;33m
RED    := \033[0;31m
NC     := \033[0m # No Color

# Docker Compose Files
COMPOSE_INFRA := -f docker-compose.infra.yml
COMPOSE_APP   := -f docker-compose.yml
COMPOSE_ALL   := $(COMPOSE_INFRA) $(COMPOSE_APP)

# =============================================================================
# Help
# =============================================================================

help: ## Show this help message
	@echo "$(GREEN)ITS Microservices - Docker Management$(NC)"
	@echo ""
	@echo "$(YELLOW)Available targets:$(NC)"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  $(GREEN)%-20s$(NC) %s\n", $$1, $$2}'

# =============================================================================
# Docker Compose Management
# =============================================================================

build: ## Build application Docker images
	@echo "$(GREEN)Building application Docker images...$(NC)"
	docker-compose $(COMPOSE_APP) build

build-no-cache: ## Build application images without cache
	@echo "$(GREEN)Building application images without cache...$(NC)"
	docker-compose $(COMPOSE_APP) build --no-cache

infra: ## Start infrastructure only (DBs, Queue)
	@echo "$(GREEN)Starting infrastructure...$(NC)"
	docker-compose $(COMPOSE_INFRA) up -d
	@echo "$(GREEN)Infrastructure started!$(NC)"

services: ## Start application services only (requires infrastructure)
	@echo "$(GREEN)Starting application services...$(NC)"
	docker-compose $(COMPOSE_APP) up -d
	@echo "$(GREEN)Application services started!$(NC)"

up: infra services ## Start all services (Infrastructure + Applications)
	@echo "$(GREEN)All systems go!$(NC)"

down: ## Stop all services
	@echo "$(YELLOW)Stopping all services...$(NC)"
	docker-compose $(COMPOSE_APP) down
	docker-compose $(COMPOSE_INFRA) down

down-volumes: ## Stop all services and remove volumes (WARNING: deletes data)
	@echo "$(RED)WARNING: This will delete all data!$(NC)"
	@read -p "Are you sure? [y/N] " -n 1 -r; \
	echo; \
	if [[ $$REPLY =~ ^[Yy]$$ ]]; then \
		docker-compose $(COMPOSE_APP) down -v; \
		docker-compose $(COMPOSE_INFRA) down -v; \
		echo "$(RED)All services stopped and volumes removed.$(NC)"; \
	fi

restart: ## Restart all services
	@echo "$(YELLOW)Restarting all services...$(NC)"
	docker-compose $(COMPOSE_APP) restart
	docker-compose $(COMPOSE_INFRA) restart

restart-infra: ## Restart infrastructure only
	@echo "$(YELLOW)Restarting infrastructure...$(NC)"
	docker-compose $(COMPOSE_INFRA) restart

restart-services: ## Restart application services only
	@echo "$(YELLOW)Restarting application services...$(NC)"
	docker-compose $(COMPOSE_APP) restart

rebuild: down build up ## Rebuild and restart all services

# =============================================================================
# Monitoring and Debugging
# =============================================================================

status: ## Show status of all services
	@echo "$(GREEN)Infrastructure Status:$(NC)"
	@docker-compose $(COMPOSE_INFRA) ps
	@echo ""
	@echo "$(GREEN)Application Status:$(NC)"
	@docker-compose $(COMPOSE_APP) ps

logs: ## Show logs from all services
	docker-compose $(COMPOSE_ALL) logs -f

logs-infra: ## Show logs from infrastructure
	docker-compose $(COMPOSE_INFRA) logs -f

logs-app: ## Show logs from application services
	docker-compose $(COMPOSE_APP) logs -f

health: ## Check health of all services
	@echo "$(GREEN)Health Check Results:$(NC)"
	@echo ""
	@echo "$(YELLOW)Frontend:$(NC)"
	@curl -sf http://localhost:3000 > /dev/null && echo "  $(GREEN)✓ Client (3000)$(NC)" || echo "  $(RED)✗ Client (3000)$(NC)"
	@echo ""
	@echo "$(YELLOW)Backend Services:$(NC)"
	@curl -sf http://localhost:8081/health > /dev/null && echo "  $(GREEN)✓ Content Service (8081)$(NC)" || echo "  $(RED)✗ Content Service (8081)$(NC)"
	@curl -sf http://localhost:8082/health > /dev/null && echo "  $(GREEN)✓ Scoring Service (8082)$(NC)" || echo "  $(RED)✗ Scoring Service (8082)$(NC)"
	@curl -sf http://localhost:8083/health > /dev/null && echo "  $(GREEN)✓ Learner Model API (8083)$(NC)" || echo "  $(RED)✗ Learner Model API (8083)$(NC)"
	@curl -sf http://localhost:8084/health > /dev/null && echo "  $(GREEN)✓ Adaptive Engine (8084)$(NC)" || echo "  $(RED)✗ Adaptive Engine (8084)$(NC)"
	@echo ""
	@echo "$(YELLOW)Infrastructure:$(NC)"
	@curl -sf http://localhost:15672 > /dev/null && echo "  $(GREEN)✓ RabbitMQ (15672)$(NC)" || echo "  $(RED)✗ RabbitMQ (15672)$(NC)"

ps: status ## Alias for status

# =============================================================================
# Database Operations
# =============================================================================

db-create: ## Create databases if they don't exist
	@echo "$(GREEN)Creating databases if they don't exist...$(NC)"
	@docker-compose $(COMPOSE_INFRA) exec -T postgres psql -U postgres -c "CREATE DATABASE content_db;" postgres 2>/dev/null || echo "  Database 'content_db' already exists or error occurred"
	@docker-compose $(COMPOSE_INFRA) exec -T postgres psql -U postgres -c "CREATE DATABASE scoring_db;" postgres 2>/dev/null || echo "  Database 'scoring_db' already exists or error occurred"
	@docker-compose $(COMPOSE_INFRA) exec -T postgres psql -U postgres -c "CREATE DATABASE learner_db;" postgres 2>/dev/null || echo "  Database 'learner_db' already exists or error occurred"
	@echo "$(GREEN)Database creation check complete!$(NC)"

db-init: db-create ## Initialize all databases with test data
	@echo "$(GREEN)Initializing databases with test data...$(NC)"
	@# We execute commands inside the single 'postgres' container, targeting different DBs
	@docker-compose $(COMPOSE_INFRA) exec -T postgres psql -U postgres -d content_db < scripts/01-init-content-db.sql || true
	@docker-compose $(COMPOSE_INFRA) exec -T postgres psql -U postgres -d scoring_db < scripts/02-init-scoring-db.sql || true
	@docker-compose $(COMPOSE_INFRA) exec -T postgres psql -U postgres -d learner_db < scripts/03-init-learner-db.sql || true
	@echo "$(GREEN)Database initialization complete!$(NC)"

db-content: ## Connect to Content database
	docker-compose $(COMPOSE_INFRA) exec postgres psql -U postgres -d content_db

db-scoring: ## Connect to Scoring database
	docker-compose $(COMPOSE_INFRA) exec postgres psql -U postgres -d scoring_db

db-learner: ## Connect to Learner database
	docker-compose $(COMPOSE_INFRA) exec postgres psql -U postgres -d learner_db

db-backup: ## Backup all databases
	@echo "$(GREEN)Backing up databases...$(NC)"
	@mkdir -p backups
	@docker-compose $(COMPOSE_INFRA) exec -T postgres pg_dump -U postgres content_db | gzip > backups/content_db_$$(date +%Y%m%d_%H%M%S).sql.gz
	@docker-compose $(COMPOSE_INFRA) exec -T postgres pg_dump -U postgres scoring_db | gzip > backups/scoring_db_$$(date +%Y%m%d_%H%M%S).sql.gz
	@docker-compose $(COMPOSE_INFRA) exec -T postgres pg_dump -U postgres learner_db | gzip > backups/learner_db_$$(date +%Y%m%d_%H%M%S).sql.gz
	@echo "$(GREEN)Backup complete! Check backups/ directory.$(NC)"

# =============================================================================
# Service-Specific Commands
# =============================================================================

content: ## Start only Content Service and dependencies
	docker-compose $(COMPOSE_INFRA) up -d postgres
	docker-compose $(COMPOSE_APP) up -d content-service

scoring: ## Start only Scoring Service and dependencies
	docker-compose $(COMPOSE_INFRA) up -d postgres rabbitmq
	docker-compose $(COMPOSE_APP) up -d scoring-service

learner: ## Start only Learner Model Service and dependencies
	docker-compose $(COMPOSE_INFRA) up -d postgres rabbitmq
	docker-compose $(COMPOSE_APP) up -d learner-model-api learner-model-consumer

adaptive: ## Start only Adaptive Engine and dependencies
	docker-compose $(COMPOSE_APP) up -d adaptive-engine

client: ## Start only Client and dependencies
	docker-compose $(COMPOSE_APP) up -d client

# =============================================================================
# Testing
# =============================================================================

test: ## Run end-to-end test
	@echo "$(GREEN)Running end-to-end test...$(NC)"
	@echo ""
	@echo "$(YELLOW)1. Checking initial mastery...$(NC)"
	@curl -s "http://localhost:8083/internal/learner/user_01/mastery?skill=math_algebra" | jq . || echo "  $(RED)Failed$(NC)"
	@echo ""
	@echo "$(YELLOW)2. Submitting wrong answer...$(NC)"
	@curl -s -X POST http://localhost:8082/api/scoring/submit \
		-H "Content-Type: application/json" \
		-d '{"user_id": "user_01", "question_id": 1, "answer": "C"}' | jq . || echo "  $(RED)Failed$(NC)"
	@sleep 2
	@echo ""
	@echo "$(YELLOW)3. Requesting next lesson...$(NC)"
	@curl -s -X POST http://localhost:8084/api/adaptive/next-lesson \
		-H "Content-Type: application/json" \
		-d '{"user_id": "user_01", "current_skill": "math_algebra"}' | jq . || echo "  $(RED)Failed$(NC)"

test-integration: ## Run integration tests
	@echo "$(GREEN)Running integration tests...$(NC)"
	cd tests/integration && ./run_tests.sh

# =============================================================================
# Cleanup
# =============================================================================

clean: ## Remove stopped containers and unused images
	@echo "$(YELLOW)Cleaning up Docker resources...$(NC)"
	docker-compose $(COMPOSE_APP) down
	docker-compose $(COMPOSE_INFRA) down
	docker system prune -f
	@echo "$(GREEN)Cleanup complete!$(NC)"

clean-all: ## Remove all containers, images, volumes (WARNING: deletes everything)
	@echo "$(RED)WARNING: This will remove ALL Docker resources!$(NC)"
	@read -p "Are you sure? [y/N] " -n 1 -r; \
	echo; \
	if [[ $$REPLY =~ ^[Yy]$$ ]]; then \
		docker-compose $(COMPOSE_APP) down -v; \
		docker-compose $(COMPOSE_INFRA) down -v; \
		docker system prune -af --volumes; \
		echo "$(RED)All Docker resources removed.$(NC)"; \
	fi

# =============================================================================
# Scaling
# =============================================================================

scale-consumer-%: ## Scale consumer (e.g., make scale-consumer-3)
	@echo "$(GREEN)Scaling learner-model-consumer to $* instances...$(NC)"
	docker-compose $(COMPOSE_APP) up -d --scale learner-model-consumer=$*

# =============================================================================
# Development Helpers
# =============================================================================

shell-%: ## Open shell in service container (e.g., make shell-scoring-service)
	docker-compose $(COMPOSE_APP) exec $* sh || docker-compose $(COMPOSE_INFRA) exec $* sh

open-urls: ## Open all service URLs in browser
	@echo "$(GREEN)Opening service URLs...$(NC)"
	@open http://localhost:3000          # Client
	@open http://localhost:8081/swagger-ui.html # Content API
	@open http://localhost:15672         # RabbitMQ

dev: ## Start services for local development
	@echo "$(GREEN)Starting infrastructure services only...$(NC)"
	docker-compose $(COMPOSE_INFRA) up -d
	@echo "$(GREEN)Infrastructure services started!$(NC)"
	@echo "$(YELLOW)You can now run services locally with:$(NC)"
	@echo "  - Content: cd content && mvn spring-boot:run"
	@echo "  - Scoring: cd scoring && go run cmd/api/main.go"
	@echo "  - Learner: cd learner-model && go run cmd/api/main.go"
	@echo "  - Adaptive: cd adaptive-engine && go run cmd/api/main.go"
	@echo "  - Client: cd client && npm run dev"

# =============================================================================
# Quick Setup
# =============================================================================

setup: build infra services db-init health ## Complete setup: build, start infra, start services, initialize DBs, check health
	@echo ""
	@echo "$(GREEN)════════════════════════════════════════════════════════$(NC)"
	@echo "$(GREEN)  Setup Complete! ITS Microservices are ready to use.$(NC)"
	@echo "$(GREEN)════════════════════════════════════════════════════════$(NC)"
	@echo ""
