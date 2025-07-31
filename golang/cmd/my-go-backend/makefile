# Define variables
APP_NAME=my-go-backend
DOCKER_IMAGE=my-go-backend
DOCKER_PORT=8080
ENV_FILE=dev.env # Change this based on your environment

# Build the Go application
build:
	@echo "Building the Go application..."
	go build -o $(APP_NAME) .

# Build and Run the application
buildrun: build
	@echo "Building and Running the application..."
	APP_ENV=dev ./$(APP_NAME)

# Run the application
run:
	@echo "Running the application..."
	APP_ENV=dev go run main.go

# Run tests
test:
	@echo "Running tests..."
	go test ./...

# Docker targets
docker-build:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE) .

docker-run:
	@echo "Running Docker container..."
	docker run -p $(DOCKER_PORT):$(DOCKER_PORT) --env-file=$(ENV_FILE) $(DOCKER_IMAGE)