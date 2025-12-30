# Backend

## セットアップ

### 開発環境

1. 設定ファイルを作成：

```bash
cp config/config.yaml.example config/config.yaml
```

2. Docker Composeで起動：

```bash
docker compose up --build
```

データベースのマイグレーションは自動実行されます。

### デフォルト設定

- API: http://localhost:8080
- Swagger UI: http://localhost:80
- PostgreSQL: `myuser/mypassword@db:5432/mydb`

### 設定のカスタマイズ

`config/config.yaml`を編集するか、環境変数で上書きできます（環境変数が優先されます）：

```bash
export DATABASE_HOST=db
export DATABASE_PORT=5432
export SERVER_PORT=8080
```

## API仕様

Swagger UI: http://localhost:80

## 本番環境

本番環境では環境変数で設定を管理してください。
