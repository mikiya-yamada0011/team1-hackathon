#!/bin/sh

set -e

echo "Waiting for postgres..."
while ! pg_isready -h db -p 5432 -q; do
  sleep 1
done

echo "PostgreSQL started"

# マイグレーションファイルがあれば実行
if [ -d "/app/db/migrations" ] && [ "$(ls -A /app/db/migrations)" ]; then
  echo "Running database migrations..."
  migrate -database "$DATABASE_URL" -path db/migrations up
fi

# モックデータファイルがあれば読み込み
if [ -f "/app/mockData.sql" ]; then
  echo "Loading mock data..."
  PGPASSWORD="$POSTGRES_PASSWORD" psql -h db -U "$POSTGRES_USER" -d "$POSTGRES_DB" -f /app/mockData.sql
  echo "Mock data loaded successfully"
fi

echo "Starting application with Air..."
air
