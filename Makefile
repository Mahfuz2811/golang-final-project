.PHONY: build up down logs restart clean help

# Docker Compose file location
COMPOSE_FILE := docker/docker-compose.yaml
OVERRIDE_FILE := docker/docker-compose.override.yaml

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the Docker images
	docker-compose -f $(COMPOSE_FILE) build

up: ## Start the services
	docker-compose -f $(COMPOSE_FILE) up -d

down: ## Stop the services
	docker-compose -f $(COMPOSE_FILE) down

logs: ## Show logs for all services
	docker-compose -f $(COMPOSE_FILE) logs -f

logs-app: ## Show logs for the main application
	docker-compose -f $(COMPOSE_FILE) logs -f golang-final-project

restart: ## Restart the services
	docker-compose -f $(COMPOSE_FILE) restart

clean: ## Remove containers, networks, and images
	docker-compose -f $(COMPOSE_FILE) down --rmi all --volumes --remove-orphans

rebuild: down build up ## Rebuild and restart services

status: ## Show status of services
	docker-compose -f $(COMPOSE_FILE) ps

migrate: ## Run database migration only
	docker-compose -f $(COMPOSE_FILE) up migration

shell: ## Open shell in the running app container
	docker-compose -f $(COMPOSE_FILE) exec golang-final-project sh

dev-up: ## Start services in development mode with hot reload
	docker-compose -f $(COMPOSE_FILE) -f $(OVERRIDE_FILE) up

dev-build: ## Build and start services in development mode
	docker-compose -f $(COMPOSE_FILE) -f $(OVERRIDE_FILE) up --build

dev-down: ## Stop development services
	docker-compose -f $(COMPOSE_FILE) -f $(OVERRIDE_FILE) down
