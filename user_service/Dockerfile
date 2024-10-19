# Используем последнюю версию Go
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Копируем файлы для зависимостей
COPY go.mod go.sum ./

# Загружаем все зависимости
RUN go mod download

# Копируем весь исходный код
COPY . .

# Собираем бинарный файл
RUN go build -o user_service ./cmd/main.go

# Используем минимальный образ для выполнения
FROM alpine:latest

WORKDIR /app

# Копируем бинарник из builder
COPY --from=builder /app/user_service .

# Пробрасываем gRPC порт
EXPOSE 50051

# Запускаем сервис
CMD ["./user_service"]
