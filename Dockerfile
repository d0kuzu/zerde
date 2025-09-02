# Стадия сборки
FROM golang:1.22-alpine AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum отдельно (чтобы кэшировать зависимости)
COPY go.mod go.sum ./
RUN go mod download

# Копируем всё остальное и собираем бинарь
COPY . .
RUN go build -o server .

# Стадия запуска
FROM alpine:latest

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем бинарь из стадии сборки
COPY --from=builder /app/server .

# Открываем порт (например, 8080)
EXPOSE 8080

# Команда запуска
CMD ["./server"]
