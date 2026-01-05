# Swagger & TypeScript型生成 セットアップ完了 ✨

## 📝 実装内容

ブログサイト用のSwagger実装が完了しました。記事一覧取得APIを例に、フロントエンドと型を共有できる環境を構築しました。

### 実装したもの

1. **Swagger アノテーション** (`backend/cmd/api/main.go`)
   - APIタイトル、バージョン、説明
   - 認証設定 (Bearer Token)

2. **モデル定義** (`backend/models/`)
   - `models.go` - Article, User のデータモデル
   - `response.go` - APIレスポンス型 (ArticleListResponse, ArticleResponse, etc.)

3. **コントローラー** (`backend/controller/article_controller.go`)
   - 記事一覧取得API (`GET /api/articles`)
   - 記事詳細取得API (`GET /api/articles/:slug`)
   - 詳細なSwaggerアノテーション付き

4. **ルーター** (`backend/api/router.go`)
   - Swagger UI のエンドポイント (`/swagger/*`)
   - API ルート設定

5. **自動生成ツール**
   - `backend/scripts/generate-frontend-types.sh` - ワンコマンド型生成スクリプト
   - `backend/Makefile` - make コマンド (現在修正中)

6. **生成されたファイル**
   - `backend/docs/swagger.json` - OpenAPI 2.0 スキーマ
   - `frontend/src/generated/data-contracts.ts` - TypeScript型定義
   - `frontend/src/generated/Api.ts` - APIクライアント
   - `frontend/src/lib/api-client.example.ts` - 使用例

## 🚀 使い方

### 1. Swaggerドキュメントを確認

サーバーを起動して、ブラウザでアクセス:
```bash
cd backend
air  # または go run cmd/api/main.go

# ブラウザで開く
open http://localhost:8080/swagger/index.html
```

### 2. フロントエンド用の型を生成

```bash
cd backend
./scripts/generate-frontend-types.sh
```

このスクリプトは以下を自動実行します:
1. Swaggerドキュメント生成 (`swag init`)
2. OpenAPIスキーマをフロントエンドにコピー
3. TypeScript型とAPIクライアントを生成

### 3. フロントエンドで型を使用

```typescript
import { Api } from '@/generated/Api';
import type { ArticleListResponse } from '@/generated/data-contracts';

const api = new Api({
  baseURL: 'http://localhost:8080',
});

// 記事一覧を取得
const response = await api.api.listArticles({ 
  page: 1, 
  limit: 10,
  department: 'Dev',
  status: 'public'
});

console.log(response.data.articles);
```

## 📚 Swaggerアノテーションの書き方

### メインの設定 (main.go)

```go
// @title        API タイトル
// @version      1.0
// @description  API の説明

// @host         localhost:8080
// @BasePath     /

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description 認証トークンを'Bearer 'に続けて入力
```

### エンドポイントの定義 (controller)

```go
// GetArticles は記事一覧を取得します
// @Summary      記事一覧を取得
// @Description  詳細な説明をここに書く
// @Tags         記事 (Articles)
// @Accept       json
// @Produce      json
// @Param        page query int false "ページ番号 (デフォルト: 1)" default(1)
// @Param        department query string false "部署" Enums(Dev, MKT, Ops)
// @Success      200 {object} models.ArticleListResponse "成功時のレスポンス"
// @Failure      400 {object} models.ErrorResponse "エラー時のレスポンス"
// @Router       /api/articles [get]
func (ac *ArticleController) GetArticles(c echo.Context) error {
    // ...
}
```

### モデルの定義

```go
type ArticleResponse struct {
    ID    string `json:"id" example:"123"`
    Title string `json:"title" example:"タイトル"`
} // @name ArticleResponse
```

## 🔄 ワークフロー

新しいAPIを追加する場合:

1. **models/ にレスポンス型を定義**
   ```go
   type NewResponse struct {
       Field string `json:"field" example:"example value"`
   } // @name NewResponse
   ```

2. **controller/ にハンドラを実装**
   ```go
   // @Summary 新しいエンドポイント
   // @Router /api/new [get]
   func (c *Controller) NewHandler(ctx echo.Context) error {
       // ...
   }
   ```

3. **api/router.go にルートを追加**
   ```go
   api.GET("/new", controller.NewHandler)
   ```

4. **型を再生成**
   ```bash
   cd backend
   ./scripts/generate-frontend-types.sh
   ```

5. **フロントエンドで使用**
   ```typescript
   const data = await api.api.newEndpoint();
   ```

## 📁 ディレクトリ構造

```
backend/
├── cmd/api/main.go          # メインエントリーポイント (Swagger設定)
├── api/router.go            # ルーター (Swagger UI含む)
├── controller/              # コントローラー (Swaggerアノテーション)
│   └── article_controller.go
├── models/                  # モデル定義
│   ├── models.go
│   ├── response.go
│   └── request.go
├── docs/                    # 自動生成されるドキュメント
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
└── scripts/
    └── generate-frontend-types.sh  # 型生成スクリプト

frontend/
├── src/
│   ├── generated/           # 自動生成されるファイル
│   │   ├── data-contracts.ts    # TypeScript型定義
│   │   ├── Api.ts               # APIクライアント
│   │   ├── http-client.ts       # HTTPクライアント
│   │   └── openapi.json         # OpenAPIスキーマ
│   └── lib/
│       └── api-client.example.ts # 使用例
```

## ⚡ コマンド一覧

```bash
# Swagger ドキュメント生成
cd backend
swag init -g cmd/api/main.go -o docs --parseDependency --parseInternal

# フロントエンドに型を生成 (推奨)
cd backend
./scripts/generate-frontend-types.sh

# 手動で型生成
cd frontend
pnpm exec swagger-typescript-api generate -p src/generated/openapi.json -o src/generated --modular --axios
```

## 🎯 実装されたAPI

### 記事関連

- **GET /api/articles** - 記事一覧を取得
  - クエリパラメータ: page, limit, department, status
  - レスポンス: ArticleListResponse

- **GET /api/articles/:slug** - 記事詳細を取得
  - パスパラメータ: slug
  - レスポンス: ArticleResponse

### Swagger UI

- **GET /swagger/*** - Swagger UI
  - http://localhost:8080/swagger/index.html

## 📖 参考リンク

- [Swaggo ドキュメント](https://github.com/swaggo/swag)
- [swagger-typescript-api](https://github.com/acacode/swagger-typescript-api)
- [Echo Swagger](https://github.com/swaggo/echo-swagger)

## ✅ 完了したこと

- ✅ Swagger アノテーションの実装
- ✅ モデル定義 (Article, User, Response型)
- ✅ コントローラー実装 (記事一覧・詳細取得)
- ✅ Swagger UI の設定
- ✅ TypeScript型の自動生成
- ✅ APIクライアントの生成
- ✅ 使用例の作成
- ✅ ドキュメント整備

これで`./scripts/generate-frontend-types.sh`を実行するだけで、
バックエンドの変更がフロントエンドに自動的に反映されます! 🎉
