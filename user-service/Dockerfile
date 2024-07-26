FROM golang:1.22-alpine as builder

WORKDIR /app

COPY . .

RUN go mod download

WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o user-service ./cmd

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/user-service .

EXPOSE 8081

HEALTHCHECK --interval=30s --timeout=5s --retries=3 CMD curl --fail http://localhost:8081/healthcheck || exit 1

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

USER appuser

ENTRYPOINT ["./user-service"]