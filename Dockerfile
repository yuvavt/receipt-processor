# syntax=docker/dockerfile:1
FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o receipt-processor .

CMD ["./receipt-processor"]
