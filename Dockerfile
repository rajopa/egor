FROM golang:1.25-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o watcher ./cmd/main.go

FROM debian:bookworm-slim 

WORKDIR /app

RUN apt-get update && apt-get install -y postgresql-client && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/watcher /app/watcher

COPY --from=builder /app/schema /app/schema

COPY --from=builder /app/pkg/config ./config

COPY --from=builder /app/pkg/config ./pkg/config

COPY --from=builder /app/.env ./.env

EXPOSE 8000

CMD [ "/app/watcher" ] 

