# Backend

## セットアップ

### 開発環境

Docker Composeを使用して起動します。

```bash
docker compose up --build
```

`.env`ファイルに開発環境用のPostgreSQL設定が含まれています。

デフォルトで以下の設定が使用されます：
- Database Host: `db`
- Database Port: `5432`
- Database User: `myuser`
- Database Password: `mypassword`
- Database Name: `mydb`
- Server Port: `8080`

### 設定のカスタマイズ

アプリケーションの設定を変更したい場合は、環境変数で上書きできます：

```bash
export DATABASE_HOST=localhost
export DATABASE_PORT=5432
export DATABASE_USER=custom_user
export DATABASE_PASSWORD=custom_password
export DATABASE_NAME=custom_db
export SERVER_PORT=8080
```

PostgreSQLの設定を変更する場合は、`.env`ファイルを編集してください。

### マイグレーション

アプリケーション起動時に自動的にマイグレーションが実行されます。

## API仕様

Swagger UIは以下のURLで確認できます：
- http://localhost:80

## 本番環境

本番環境では、環境変数や機密情報管理システムを使用して設定を管理してください。
`.env`ファイルに記載されている認証情報は開発環境専用です。
