# Compile
FROM golang:1.22 AS builder

WORKDIR /app

# get dependencies 
COPY go.mod go.sum ./
RUN go mod download

# copy the rest of the code
COPY . .

# get the binary
RUN go build -o docker-monitor ./cmd

# final image
FROM debian:bullseye-slim

WORKDIR /app

# copy the binary
COPY --from=builder /app/docker-monitor .

# command to run the server
CMD ["./docker-monitor"]

