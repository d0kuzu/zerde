# Этап 1: сборка
FROM golang:1.23.3-alpine AS builder

WORKDIR /app

# Устанавливаем зависимости для сборки
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o server .

# Этап 2: запуск
FROM alpine:latest

# Устанавливаем сертификаты и Chromium
RUN apk --no-cache add \
    ca-certificates \
    chromium \
    nss \
    freetype \
    harfbuzz \
    ttf-freefont \
    font-noto-emoji

WORKDIR /root/

COPY --from=builder /app/server .

# Указываем путь до chromium (на Alpine он вот такой)
ENV CHROME_PATH=/usr/bin/chromium-browser
# Можно также явно отключить sandbox, чтобы chromedp не падал под root
ENV CHROMEDP_NO_SANDBOX=true

EXPOSE 8080

CMD ["./server"]
