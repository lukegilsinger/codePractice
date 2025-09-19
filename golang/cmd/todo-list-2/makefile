.PHONY: start run build docker clean

# Define variables
APP_NAME=my-go-backend
DOCKER_IMAGE=my-go-backend
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
	DATABASE_URL=${DATABASE_URL} go run cmd/server/main.go

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

# Run tests
test:
	go test ./...