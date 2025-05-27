# Etapa 1: build
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o docker-monitor ./cmd

FROM scratch

WORKDIR /app

COPY --from=builder /app/docker-monitor .

EXPOSE 81

ENTRYPOINT ["./docker-monitor"]

