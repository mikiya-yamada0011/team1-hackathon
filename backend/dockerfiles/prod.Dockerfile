# ビルドステージ
FROM golang:1.25.5-alpine AS builder

# 必要なツールをインストール
RUN apk add --no-cache git ca-certificates tzdata curl tar

# マイグレーションツールをインストール
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz && \
    mv migrate /usr/local/bin/migrate

WORKDIR /app

# 依存関係をキャッシュ
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピー
COPY . .

# Swaggerドキュメントを生成
RUN go install github.com/swaggo/swag/cmd/swag@v1.16.4
RUN swag init -g cmd/api/main.go

# 静的バイナリをビルド
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo \
    -o main ./cmd/api

# 実行ステージ
FROM alpine:latest

# 必要なランタイム依存関係をインストール
RUN apk add --no-cache ca-certificates tzdata

# 非rootユーザーを作成
RUN adduser -D -s /bin/sh appuser

WORKDIR /app

# ビルドステージから成果物をコピー
COPY --from=builder /app/main /app/main
COPY --from=builder /app/db/migrations /app/db/migrations
COPY --from=builder /app/config/config.yaml.example /app/config/config.yaml

# 所有権を変更
RUN chown -R appuser:appuser /app

# 非rootユーザーで実行
USER appuser

EXPOSE 8080

# ヘルスチェック
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/swagger/index.html || exit 1

CMD ["/app/main"]
