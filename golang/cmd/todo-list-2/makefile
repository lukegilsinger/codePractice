.PHONY: start run build docker clean

# Define variables
APP_NAME=todo-list-v2
DOCKER_IMAGE=todo-list-v2
DOCKER_PORT=8080

# Server
PORT=8080
DATABASE_URL=data/todos.db

# Logging Configuration  
LOG_LEVEL=info          # debug, info, warn, error
LOG_FORMAT=text         # text or json

# Security
JWT_SECRET=your-super-secret-jwt-key-change-this-in-production

# Run the application
start:
	@echo "Running the application..."
	DATABASE_URL=${DATABASE_URL} APP_ENV=deva go run cmd/server/main.go

# Run locally (without Docker)
run:
	go run cmd/server/main.go

# Build the binary
build:
	go build -o bin/todo-app cmd/server/main.go

# Run with Docker
docker:
	docker-compose up --build

# Stop Docker containers
docker-stop:
	docker-compose down

# Clean up
clean:
	rm -rf bin/
	go clean

# Install dependencies
deps:
	go mod tidy
	go mod download

# Testing commands
test:
	go test ./...

test-verbose:
	go test -v ./...

test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

test-integration:
	go test -v -tags=integration ./...

test-unit:
	go test -v -short ./...

test-watch:
	@echo "Watching for changes... (requires 'entr' - install with: brew install entr)"
	find . -name "*.go" | entr -c make test

# Database migrations
migrate-up:
	go run cmd/migrate/main.go up

migrate-down:
	go run cmd/migrate/main.go down

migrate-status:
	go run cmd/migrate/main.go status

# Rollback multiple migrations
migrate-rollback:
	go run cmd/migrate/main.go down 2
