# Team1 Blog API Backend

チーム1のテックブログシステムのバックエンドAPI

## 🚀 セットアップ

### 1. 依存関係のインストール
```bash
go mod download
```

### 2. 環境変数の設定
```bash
cp .env.example .env
# .envファイルを編集して、データベース接続情報などを設定
```

### 3. データベースのマイグレーション
```bash
# Dockerコンテナを起動している場合は自動でマイグレーションが実行されます
# 手動で実行する場合:
go run cmd/api/main.go
```

### 4. サーバーの起動
```bash
# 開発環境 (ホットリロード)
air

# 本番環境
go run cmd/api/main.go
```

## 📚 Swagger ドキュメント

### Swagger UIで確認
サーバーを起動後、以下のURLにアクセス:
```
http://localhost:8080/swagger/index.html
```

### Swaggerドキュメントの生成
```bash
# Swaggerドキュメントのみ生成
make swagger

# OpenAPIスキーマをフロントエンドにコピー
make openapi-gen

# フロントエンド用のTypeScript型を生成
make frontend-types

# すべて実行
make all
```

#### 生成されるファイル
- `docs/swagger.json` - OpenAPI 2.0形式のスキーマ
- `docs/swagger.yaml` - YAML形式のスキーマ
- `docs/docs.go` - Go用の埋め込みドキュメント
- `../frontend/src/generated/openapi.json` - フロントエンド用のOpenAPIスキーマ
- `../frontend/src/generated/api-types.ts` - TypeScript型定義

## 🔧 開発

### Swaggerアノテーションの書き方

#### メインの設定 (cmd/api/main.go)
```go
// @title        API Title
// @version      1.0
// @description  API Description

// @host         localhost:8080
// @BasePath     /

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
```

#### エンドポイントの定義 (controller/*.go)
```go
// GetArticles は記事一覧を取得します
// @Summary      記事一覧を取得
// @Description  詳細な説明
// @Tags         記事 (Articles)
// @Accept       json
// @Produce      json
// @Param        page query int false "ページ番号"
// @Success      200 {object} models.ArticleListResponse
// @Failure      400 {object} models.ErrorResponse
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

## 📖 API エンドポイント

### 記事関連
- `GET /api/articles` - 記事一覧を取得
- `GET /api/articles/:slug` - 記事詳細を取得

## 🛠️ 使用技術

- **Go** 1.25.5
- **Echo** - Webフレームワーク
- **GORM** - ORM
- **PostgreSQL** - データベース
- **Swaggo** - Swagger生成
- **Air** - ホットリロード

## 📝 ディレクトリ構造

```
backend/
├── api/           # ルーター定義
├── cmd/api/       # エントリーポイント
├── config/        # 設定管理
├── controller/    # コントローラー層
├── database/      # データベース接続
├── db/migrations/ # マイグレーションファイル
├── docs/          # Swaggerドキュメント (自動生成)
├── models/        # データモデル
├── repository/    # リポジトリ層
├── service/       # サービス層
└── Makefile       # タスク管理
```

## 🔄 ワークフロー

1. **モデルを定義** (`models/`)
   - リクエスト/レスポンスの型を定義
   - Swaggerアノテーションを追加

2. **コントローラーを実装** (`controller/`)
   - ハンドラー関数を実装
   - Swaggerアノテーションでドキュメント化

3. **ルーターに登録** (`api/router.go`)
   - エンドポイントを追加

4. **ドキュメント生成**
   ```bash
   make all
   ```

5. **フロントエンドで型を使用**
   ```typescript
   import type { components } from '@/generated/api-types'
   
   type Article = components['schemas']['ArticleResponse']
   ```

## ⚡ よく使うコマンド

```bash
# ヘルプを表示
make help

# Swagger生成のみ
make swagger

# フロントエンド用の型まで生成
make all

# 開発サーバー起動 (ホットリロード)
air

# フォーマット
go fmt ./...

# リント
go vet ./...
```
