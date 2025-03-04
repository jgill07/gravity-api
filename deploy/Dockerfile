FROM golang:1.24.0-alpine3.21 AS builder
WORKDIR /app
ADD . .
RUN go build -o /app/service ./cmd/api/main.go
RUN chmod +x /app/service

FROM alpine:3.21
COPY --from=builder /app/service /app/service
EXPOSE 8080
ENTRYPOINT [ "/app/service" ]