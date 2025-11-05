# Этап 1: сборка
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server .

# Этап 2: запуск
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/server .

ENV CHROME_PATH=/usr/bin/chromium-browser

EXPOSE 8080

CMD ["./server"]
