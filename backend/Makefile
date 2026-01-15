.PHONY: help swagger openapi-gen frontend-types all

help:
	@echo "使用可能なコマンド:"
	@echo "  make swagger         - Swaggerドキュメントを生成"
	@echo "  make openapi-gen     - OpenAPIスキーマをフロントエンドにコピー"
	@echo "  make frontend-types  - フロントエンド用の型を生成"
	@echo "  make all             - すべてを実行"

swagger:
	swag init -g cmd/api/main.go -o docs --parseDependency --parseInternal

openapi-gen: swagger
	@echo "OpenAPI スキーマをフロントエンドにコピー中..."
	@mkdir -p ../frontend/src/generated
	@cp docs/swagger.json ../frontend/src/generated/openapi.json
	@echo "✅ OpenAPI スキーマをコピーしました"

frontend-types: openapi-gen
	@echo "フロントエンド用のTypeScript型を生成中..."
	@cd ../frontend && pnpm add -D openapi-typescript && pnpm exec openapi-typescript src/generated/openapi.json -o src/generated/api-types.ts
	@echo "✅ TypeScript型を生成しました"

all: frontend-types
	@echo ""
	@echo "✅ すべて完了しました!"
	@echo "Swagger UI: http://localhost:8080/swagger/index.html"
	@echo "OpenAPI JSON: ../frontend/src/generated/openapi.json"
	@echo "TypeScript types: ../frontend/src/generated/api-types.ts"
