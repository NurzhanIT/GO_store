FROM golang:1.19-alpine

RUN mkdir /market
WORKDIR /market

COPY /shop /market

RUN go build -o market ./cmd/api

CMD [ "./market" ]


