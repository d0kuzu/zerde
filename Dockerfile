# Этап 1: сборка
FROM golang:1.23.3-alpine AS builder

RUN apk --no-cache add ca-certificates git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy && go mod download

COPY . .

RUN go build -o server .

# Этап 2: запуск
FROM alpine:latest

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

ENV CHROME_PATH=/usr/bin/chromium-browser

EXPOSE 8080

ENTRYPOINT ["/usr/bin/dumb-init", "--"]
CMD ["./server"]
