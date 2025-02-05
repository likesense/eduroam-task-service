FROM golang:1.23.4-alpine3.20 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./cmd/main.go

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .
CMD ["./main"]