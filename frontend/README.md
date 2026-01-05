This is a [Next.js](https://nextjs.org) project bootstrapped with [`create-next-app`](https://nextjs.org/docs/app/api-reference/cli/create-next-app).

## Getting Started

First, run the development server:

```bash
npm run dev
# or
yarn dev
# or
pnpm dev
# or
bun dev
```

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

You can start editing the page by modifying `app/page.tsx`. The page auto-updates as you edit the file.

This project uses [`next/font`](https://nextjs.org/docs/app/building-your-application/optimizing/fonts) to automatically optimize and load [Geist](https://vercel.com/font), a new font family for Vercel.

## Learn More

To learn more about Next.js, take a look at the following resources:

- [Next.js Documentation](https://nextjs.org/docs) - learn about Next.js features and API.
- [Learn Next.js](https://nextjs.org/learn) - an interactive Next.js tutorial.

You can check out [the Next.js GitHub repository](https://github.com/vercel/next.js) - your feedback and contributions are welcome!

## Deploy on Vercel

The easiest way to deploy your Next.js app is to use the [Vercel Platform](https://vercel.com/new?utm_medium=default-template&filter=next.js&utm_source=create-next-app&utm_campaign=create-next-app-readme) from the creators of Next.js.

Check out our [Next.js deployment documentation](https://nextjs.org/docs/app/building-your-application/deploying) for more details.

## 🔄 API型の自動生成

### バックエンドから型を生成する方法

1. バックエンドディレクトリで型生成スクリプトを実行:
```bash
cd ../backend
./scripts/generate-frontend-types.sh
```

または個別のコマンドで:
```bash
cd ../backend

# Swaggerドキュメントを生成
swag init -g cmd/api/main.go -o docs --parseDependency --parseInternal

# OpenAPIスキーマをコピー
cp docs/swagger.json ../frontend/src/generated/openapi.json

# TypeScript型を生成
cd ../frontend
pnpm exec swagger-typescript-api generate -p src/generated/openapi.json -o src/generated --modular --axios
```

### 生成されるファイル

- `src/generated/data-contracts.ts` - 型定義
- `src/generated/Api.ts` - APIクライアント
- `src/generated/http-client.ts` - HTTPクライアント
- `src/generated/openapi.json` - OpenAPIスキーマ

### APIクライアントの使い方

```typescript
import { Api } from '@/generated/Api';
import type { ArticleListResponse } from '@/generated/data-contracts';

const api = new Api({
  baseURL: 'http://localhost:8080',
});

// 記事一覧を取得
const articles = await api.api.listArticles({ 
  page: 1, 
  limit: 10 
});

// 記事詳細を取得
const article = await api.api.detailArticles('article-slug');
```

詳細な使用例は `src/lib/api-client.example.ts` を参照してください。
