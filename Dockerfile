# Используем многоступенчатую сборку для оптимизации размера образа
FROM golang:1.24-alpine AS builder

# Устанавливаем необходимые пакеты для сборки
RUN apk add --no-cache git ca-certificates tzdata

# Создаем рабочую директорию
WORKDIR /app

# Копируем файлы зависимостей
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/server/main.go

# Финальный образ
FROM alpine:latest

# Устанавливаем ca-certificates для HTTPS запросов
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копируем собранное приложение из builder стадии
COPY --from=builder /app/main .

# Копируем конфигурационные файлы если они есть
COPY --from=builder /app/internal/db/migrations ./internal/db/migrations

# Открываем порт
EXPOSE 8080

# Запускаем приложение
CMD ["./main"]


