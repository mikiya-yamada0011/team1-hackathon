#!/bin/bash

set -e

echo "🔄 Swagger ドキュメントを生成中..."
swag init -g cmd/api/main.go -o docs --parseDependency --parseInternal

echo ""
echo "📋 OpenAPI スキーマをフロントエンドにコピー中..."
mkdir -p ../frontend/src/generated
cp docs/swagger.json ../frontend/src/generated/openapi.json

echo ""
echo "⚡ Orval でTypeScript型とReact Queryフックを生成中..."
cd ../frontend
pnpm generate:api

echo ""
echo "✅ すべて完了しました!"
echo ""
echo "生成されたファイル:"
echo "  - backend/docs/swagger.json"
echo "  - backend/docs/swagger.yaml"
echo "  - frontend/src/generated/openapi.json"
echo "  - frontend/src/generated/api/ (React Queryフック)"
echo "  - frontend/src/generated/models/ (型定義)"
echo ""
echo "Swagger UI: http://localhost:8080/swagger/index.html"
echo ""
echo "使い方:"
echo "  import { useListArticles } from '@/generated/api/articles';"
echo "  const { data, isLoading } = useListArticles({ page: 1, limit: 10 });"
