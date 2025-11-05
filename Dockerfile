# Этап 1: сборка Go-приложения
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server .

# Этап 2: запуск на headless Chromium
FROM chromedp/headless-shell:latest

WORKDIR /root/

COPY --from=builder /app/server .

# Устанавливаем сертификаты (если нужно)
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

EXPOSE 8080

# Удаляем дефолтный entrypoint и запускаем только сервер
ENTRYPOINT []
CMD ["./server"]

