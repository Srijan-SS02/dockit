FROM golang:latest AS builder

WORKDIR /app

COPY .  .

RUN go build -o dockit


FROM debian:stable-slim

WORKDIR /app

COPY --from=builder /app/dockit /app/
COPY --from=builder /app/templates /app/templates

CMD ["./dockit"]