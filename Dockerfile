# 1. Билдим Go бинарь
FROM golang:1.24.3-alpine AS builder

WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./
RUN go mod download

# Копируем код сервиса
COPY . .

# Собираем бинарь
RUN go build -o user-service ./cmd/server/main.go

# 2. Минимальный рантайм
FROM alpine:3.20

WORKDIR /app

# Копируем бинарь из builder
COPY --from=builder /app/user-service .

# Запуск
CMD ["./user-service"]
