# Orval への移行完了 🎉

swagger-typescript-api から **Orval** に移行しました!

## ✨ 変更点

### Before (swagger-typescript-api)
```typescript
import { Api } from '@/generated/Api';

const api = new Api({ baseURL: 'http://localhost:8080' });
const response = await api.api.listArticles({ page: 1 });
```

### After (Orval + React Query)
```typescript
import { useGetApiArticles } from '@/generated/api/記事-articles/記事-articles';

function Component() {
  const { data, isLoading } = useGetApiArticles({
    page: 1,
    limit: 10,
  });
  
  return <div>{data?.articles}</div>;
}
```

## 📦 追加されたパッケージ

- **orval** - OpenAPIからTypeScript/React Queryを生成
- **@tanstack/react-query** - データフェッチング・キャッシュライブラリ
- **axios** - HTTPクライアント
- **zod** - バリデーション (将来のために)

## 🚀 使い方

### 1. 型とフックを生成

```bash
# フロントエンドで実行
cd frontend
pnpm generate:api

# またはバックエンドから一括実行
cd backend
./scripts/generate-frontend-types.sh
```

### 2. React Query Providerをセットアップ

```tsx
// app/layout.tsx
import { Providers } from './providers';

export default function RootLayout({ children }) {
  return (
    <html>
      <body>
        <Providers>{children}</Providers>
      </body>
    </html>
  );
}
```

### 3. フックを使う

```tsx
'use client';

import { useGetApiArticles } from '@/generated/api/記事-articles/記事-articles';

export default function ArticlesPage() {
  const { data, isLoading, error } = useGetApiArticles({
    page: 1,
    limit: 10,
    status: 'public',
  });

  if (isLoading) return <div>Loading...</div>;
  if (error) return <div>Error: {error.message}</div>;

  return (
    <div>
      {data?.articles?.map(article => (
        <div key={article.id}>{article.title}</div>
      ))}
    </div>
  );
}
```

## 📁 生成されるファイル

```
frontend/src/
├── generated/
│   ├── api/
│   │   └── 記事-articles/
│   │       └── 記事-articles.ts  # React Queryフック
│   ├── models/
│   │   └── *.ts                  # 型定義
│   └── openapi.json
└── lib/
    ├── api-client.ts             # カスタムAxiosインスタンス
    └── api-hooks-example.tsx     # 使用例
```

## 🎯 Orvalの利点

### 1. React Queryとの完璧な統合
- `useQuery` が自動生成される
- キャッシュ、自動リフェッチが標準搭載
- ローディング状態の管理が簡単

### 2. 型安全性
- リクエスト・レスポンスの完全な型推論
- パラメータのバリデーション
- エラー型も含めて型定義

### 3. 宣言的なコード
```typescript
// 命令的 (Before)
const [data, setData] = useState();
const [loading, setLoading] = useState(true);
useEffect(() => {
  fetch('/api/articles')
    .then(res => res.json())
    .then(setData)
    .finally(() => setLoading(false));
}, []);

// 宣言的 (After)
const { data, isLoading } = useGetApiArticles();
```

### 4. 自動キャッシュ管理
- 同じクエリは自動的にキャッシュ
- バックグラウンドでの自動更新
- 楽観的更新のサポート

## 🔧 設定ファイル

### orval.config.ts
```typescript
import { defineConfig } from 'orval';

export default defineConfig({
  blog: {
    input: {
      target: './src/generated/openapi.json',
    },
    output: {
      mode: 'tags-split',              // タグごとにファイル分割
      target: './src/generated/api',    // 出力先
      schemas: './src/generated/models',
      client: 'react-query',            // React Query使用
      override: {
        mutator: {
          path: './src/lib/api-client.ts',  // カスタムAxios
          name: 'customInstance',
        },
      },
    },
  },
});
```

### src/lib/api-client.ts
```typescript
import Axios from 'axios';

export const AXIOS_INSTANCE = Axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080',
});

export const customInstance = <T>(config, options?) => {
  return AXIOS_INSTANCE({ ...config, ...options }).then(({ data }) => data);
};
```

## 📚 React Queryの便利な機能

### 手動リフェッチ
```typescript
const { data, refetch } = useGetApiArticles({ page: 1 });

<button onClick={() => refetch()}>更新</button>
```

### ローディング状態
```typescript
const { isLoading, isFetching } = useGetApiArticles({ page: 1 });

// isLoading: 初回ロード中
// isFetching: バックグラウンド更新中
```

### エラーハンドリング
```typescript
const { error } = useGetApiArticles({ page: 1 });

if (error) {
  return <div>Error: {error.message}</div>;
}
```

### 依存クエリ
```typescript
const { data: article } = useGetApiArticlesSlug(slug);
const { data: comments } = useGetComments(
  { articleId: article?.id },
  { enabled: !!article?.id }  // articleが取得されるまで待つ
);
```

## 🔄 ワークフロー

1. **バックエンドでSwagger生成**
   ```bash
   cd backend
   swag init -g cmd/api/main.go -o docs
   ```

2. **フロントエンドで型生成**
   ```bash
   cd frontend
   pnpm generate:api
   ```

3. **コンポーネントで使用**
   ```tsx
   const { data } = useGetApiArticles({ page: 1 });
   ```

## 🎨 使用例

詳細な使用例は以下を参照:
- `frontend/src/lib/api-hooks-example.tsx`
- `frontend/src/app/providers.tsx`

## 📖 参考リンク

- [Orval公式](https://orval.dev/)
- [TanStack Query](https://tanstack.com/query/latest)
- [Axios](https://axios-http.com/)

これで型安全でモダンなAPI通信が可能になりました! 🚀
