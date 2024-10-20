FROM golang:1.21-alpine

WORKDIR /app

COPY . .

RUN go mod tidy && go mod download

RUN go build -o logger_service ./cmd/main.go

CMD ["./logger_service"]
