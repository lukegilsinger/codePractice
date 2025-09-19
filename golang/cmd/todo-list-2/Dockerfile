
# Dockerfile
FROM golang:1.24-alpine

# Install sqlite and gcc (needed for CGO)
RUN apk add --no-cache sqlite gcc musl-dev

# Set working directory
WORKDIR /app

# Copy go files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the app with CGO enabled
RUN CGO_ENABLED=1 go build -o main ./cmd/server

# Expose port
EXPOSE 8080

# Run the app
CMD ["./main"]