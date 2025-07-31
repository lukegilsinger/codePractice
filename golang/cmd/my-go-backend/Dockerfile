# Use the official Golang image as the build stage
FROM golang:1.24 AS builder

# Set the working directory inside the container
WORKDIR /app

# # Install CA certificates
# RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# # Set the GOPROXY environment variable
# ENV GOPROXY=https://proxy.golang.org,direct

# Copy the go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the rest of the application files to the working directory
COPY . .

# Build the Go application
# RUN go build -o mygobackend .
RUN go build -o /mygobackend

# # Use a smaller image for the final stage
# FROM alpine:latest

# # Set the working directory in the final image
# WORKDIR /app

# # Copy the compiled binary from the builder stage
# COPY --from=builder /app/mygobackend .

RUN ls 

RUN pwd

# Expose the port that your application listens on
EXPOSE 8080

# Command to run the application
CMD ["/mygobackend"]