FROM golang:latest

WORKDIR /cmd/chat

COPY server /cmd/chat/server
COPY client /cmd/chat/client

RUN go mod download

RUN go build -o chatapp ./server/main.go

EXPOSE 50051

CMD ["./chatapp"]