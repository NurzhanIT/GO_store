# Dockerfile for the server and client applications

# Build stage
FROM golang:1.17 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./server/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o client ./client/main.go

# Final stage
FROM golang:1.17

WORKDIR /app

COPY --from=builder /app/server .
COPY --from=builder /app/client .

CMD ["./server"]