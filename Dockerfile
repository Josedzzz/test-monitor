# Build stage
FROM golang:1.24 AS builder

WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go binary
RUN go build -o docker-monitor ./cmd

# Final stage (slim runtime)
FROM debian:bullseye-slim

WORKDIR /app

# Copy the binary from the builder
COPY --from=builder /app/docker-monitor .

# Start the application
CMD ["./docker-monitor"]

