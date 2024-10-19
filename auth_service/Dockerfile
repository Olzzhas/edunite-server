# Используем последнюю версию Go
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем исходный код в контейнер
COPY . .

# Собираем бинарный файл
RUN go build -o auth_service ./cmd/main.go

# Используем минимальный образ для финального контейнера
FROM alpine:latest

WORKDIR /app

# Копируем бинарный файл из builder-образа
COPY --from=builder /app/auth_service .

# Пробрасываем порт для gRPC
EXPOSE 50052

CMD ["./auth_service"]
