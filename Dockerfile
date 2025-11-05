# Этап 1: сборка
FROM golang:1.23.3-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server .

# Этап 2: запуск
FROM alpine:latest

# Устанавливаем сертификаты, Chromium и зависимости для headless-режима
RUN apk --no-cache add \
    ca-certificates \
    chromium \
    nss \
    freetype \
    harfbuzz \
    ttf-freefont \
    dumb-init

WORKDIR /root/

COPY --from=builder /app/server .

EXPOSE 8080

# Используем dumb-init для корректного завершения процессов Chromium
ENTRYPOINT ["/usr/bin/dumb-init", "--"]
CMD ["./server"]
