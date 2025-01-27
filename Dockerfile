# Stage 1: Build the application
FROM golang:1.23 AS builder

# Set environment variables
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the application binary
RUN go build -o gin-app main.go

# Stage 2: Create a lightweight final image
FROM alpine:latest

# Install necessary dependencies (e.g., for TLS support)
RUN apk --no-cache add ca-certificates

# Set working directory inside the final image
WORKDIR /root/

# Copy the compiled binary from the build stage
COPY --from=builder /app/gin-app .

# Set environment variables
ENV GIN_MODE=release

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./gin-app"]
