FROM golang:1.23-alpine

# Установим необходимые пакеты
RUN apk add --no-cache git

# Настроим рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum, скачиваем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем проект
COPY . .

# Открываем порт и запускаем Gateway
EXPOSE 8081
CMD ["go", "run", "./main.go"]
