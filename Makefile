.PHONY: run build test lint clean dev docker-build docker-up docker-down docker-logs help

# Help
help: ## Show available commands
	@echo "Dev Echo Server - Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-20s %s\n", $$1, $$2}'

# Quick run with script
quick: ## Quick run with interactive script
	@./run.sh

# Development
dev: ## Run in development mode
	@echo "Running in development mode..."
	NODE_ENV=development go run main.go

# Production
run: ## Run in production mode
	@echo "Running in production mode..."
	NODE_ENV=production go run main.go

# Build binary
build: ## Build the binary
	@echo "Building binary..."
	go build -o dev-echo-server .

# Docker commands
docker-build: ## Build Docker image
	@echo "Building Docker image..."
	docker-compose build

docker-up: ## Start Docker containers
	@echo "Starting containers..."
	docker-compose up -d
	@echo "Server running at http://localhost:3000"

docker-down: ## Stop Docker containers
	@echo "Stopping containers..."
	docker-compose down

docker-logs: ## View Docker logs
	@docker-compose logs -f

docker-restart: ## Restart Docker containers
	@make docker-down
	@make docker-up

# Run tests
test: ## Run tests
	@echo "Running tests..."
	go test -v ./...

# Run tests with coverage
test-coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	go test -cover ./...

# Lint code
lint: ## Run linter
	@echo "Running linter..."
	golangci-lint run

# Fix linting issues
lint-fix: ## Fix linting issues
	@echo "Fixing linting issues..."
	golangci-lint run --fix

# Clean build artifacts
clean: ## Clean build artifacts and logs
	@echo "Cleaning..."
	rm -f dev-echo-server server
	rm -rf log/*.log
	go clean

# Download dependencies
deps: ## Download dependencies
	@echo "Downloading dependencies..."
	go mod download

# Tidy dependencies
tidy: ## Tidy dependencies
	@echo "Tidying dependencies..."
	go mod tidy

.DEFAULT_GOAL := help
