# Build Stage
FROM golang:1.26 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Run Stage
FROM debian:bookworm-slim

# ←これが重要
RUN apt-get update && \
  apt-get install -y ca-certificates && \
  update-ca-certificates && \
  rm -rf /var/lib/apt/lists/*

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs

EXPOSE 8080

CMD ["./main"]
