FROM golang:1.25.5-alpine AS builder

RUN apk add --no-cache git ca-certificates tzdata curl tar

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz && \
    mv migrate /usr/local/bin/migrate

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo \
    -o main ./cmd/api

FROM alpine:latest

RUN apk add --no-cache ca-certificates tzdata netcat-openbsd postgresql-client

RUN adduser -D -s /bin/sh appuser

COPY --from=builder /app/main /app/main
COPY --from=builder /usr/local/bin/migrate /usr/local/bin/migrate
COPY --from=builder /app/db/migrations /app/db/migrations
COPY --from=builder /app/entrypoint-prod.sh /app/entrypoint-prod.sh
COPY --from=builder /app/mockData.sql /app/mockData.sql
COPY --from=builder /app/images /app/images

RUN chmod +x /app/entrypoint-prod.sh

USER appuser

WORKDIR /app

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD ["/app/main", "-health-check"] || exit 1

ENTRYPOINT ["/app/entrypoint-prod.sh"]
