FROM golang:1.25.5-alpine

RUN apk add --no-cache git postgresql-client && \
    apk add --no-cache --virtual .build-deps curl tar && \
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz && \
    mv migrate /usr/local/bin/ && \
    apk del .build-deps

RUN go install github.com/air-verse/air@v1.62.0
RUN go install github.com/swaggo/swag/cmd/swag@v1.16.4

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN chmod +x /app/entrypoint.sh
ENTRYPOINT ["/app/entrypoint.sh"]

EXPOSE 8080