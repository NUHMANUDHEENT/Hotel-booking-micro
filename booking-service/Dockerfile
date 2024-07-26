FROM golang:1.22-alpine as builder

WORKDIR /app

COPY . .

RUN go mod download

WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o booking-service ./cmd

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/booking-service .

EXPOSE 8082

HEALTHCHECK --interval=30s --timeout=5s --retries=3 CMD curl --fail http://localhost:8082/healthcheck || exit 1

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

USER appuser

ENTRYPOINT ["./booking-service"]