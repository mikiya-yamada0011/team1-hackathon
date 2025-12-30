FROM golang:1.25.5-alpine

RUN apk add --no-cache git postgresql-client

RUN go install github.com/air-verse/air@v1.62.0
RUN go install github.com/swaggo/swag/cmd/swag@v1.16.4

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]