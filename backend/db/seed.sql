-- 開発環境専用のテストデータ

-- テストユーザーの挿入
INSERT INTO users (id, name, affiliation, password_hash, icon_url) VALUES
(1, '田中 太郎', 'Dev部門', '$2a$10$dummyhash1111111111111111111111111111111111111111', 'https://i.pravatar.cc/150?img=1'),
(2, '佐藤 花子', 'MKT部門', '$2a$10$dummyhash2222222222222222222222222222222222222222', 'https://i.pravatar.cc/150?img=2'),
(3, '鈴木 一郎', 'Ops部門', '$2a$10$dummyhash3333333333333333333333333333333333333333', 'https://i.pravatar.cc/150?img=3'),
(4, '高橋 美咲', 'Dev部門', '$2a$10$dummyhash4444444444444444444444444444444444444444', 'https://i.pravatar.cc/150?img=4')
ON CONFLICT (id) DO NOTHING;

-- タグの挿入
INSERT INTO tags (name, is_category) VALUES
('React', false),
('Go', false),
('Docker', false),
('マーケティング', false),
('インフラ', false),
('チュートリアル', true),
('ベストプラクティス', true),
('トラブルシューティング', true)
ON CONFLICT (name) DO NOTHING;

-- テスト記事の挿入
INSERT INTO articles (author_id, article_type, title, content, slug, department, status, thumbnail_url) VALUES
(
    1,
    'markdown',
    'React Hooks完全ガイド',
    E'# React Hooks完全ガイド\n\nReact Hooksの基本から応用まで、実践的な使い方を解説します。\n\n## useState\n\nコンポーネントの状態管理に使用します。\n\n```javascript\nconst [count, setCount] = useState(0);\n```\n\n## useEffect\n\n副作用を扱うためのHookです。\n\n```javascript\nuseEffect(() => {\n  console.log(\'Component mounted\');\n}, []);\n```\n\n## カスタムフック\n\n独自のフックを作成して、ロジックを再利用できます。',
    'react-hooks-guide',
    'Dev',
    'public',
    'https://images.unsplash.com/photo-1633356122544-f134324a6cee?w=800'
),
(
    2,
    'markdown',
    'マーケティング戦略2026',
    E'# マーケティング戦略2026\n\n2026年のマーケティングトレンドと戦略について解説します。\n\n## SNSマーケティング\n\n最新のSNSマーケティング手法を紹介します。\n\n- Instagram\n- TikTok\n- X (Twitter)\n\n## データ分析\n\nマーケティングデータの分析手法について説明します。',
    'marketing-strategy-2026',
    'MKT',
    'public',
    'https://images.unsplash.com/photo-1460925895917-afdab827c52f?w=800'
),
(
    3,
    'markdown',
    'Kubernetes運用ガイド',
    E'# Kubernetes運用ガイド\n\nKubernetesの運用に関するベストプラクティスを紹介します。\n\n## デプロイ戦略\n\n- Rolling Update\n- Blue-Green Deployment\n- Canary Deployment',
    'kubernetes-operations',
    'Ops',
    'internal',
    'https://images.unsplash.com/photo-1666875753105-c63a6f3bdc86?w=800'
),
(
    1,
    'markdown',
    'Go言語入門',
    E'# Go言語入門\n\nGo言語の基本的な文法と特徴を学びます。\n\n## Go言語の特徴\n\n- シンプルな文法\n- 高速なコンパイル\n- 並行処理のサポート\n\n## サンプルコード\n\n```go\npackage main\n\nimport \"fmt\"\n\nfunc main() {\n    fmt.Println(\"Hello, Go!\")\n}\n```',
    'go-introduction',
    'Dev',
    'internal',
    'https://images.unsplash.com/photo-1617854818583-09e7f077a156?w=800'
),
(
    1,
    'external',
    'TypeScript公式ドキュメント',
    NULL,
    'typescript-official-docs',
    'Dev',
    'public',
    'https://images.unsplash.com/photo-1587620962725-abab7fe55159?w=800'
),
(
    2,
    'markdown',
    'コンテンツマーケティングの基礎',
    E'# コンテンツマーケティングの基礎\n\n効果的なコンテンツマーケティングの手法を紹介します。\n\n## コンテンツの企画\n\nターゲットオーディエンスを明確にしましょう。\n\n## 配信チャネル\n\n- ブログ\n- メールマガジン\n- SNS',
    'content-marketing-basics',
    'MKT',
    'draft',
    'https://images.unsplash.com/photo-1542744094-3a31f272c490?w=800'
)
ON CONFLICT (slug) DO NOTHING;

-- 外部記事のURLを設定
UPDATE articles 
SET external_url = 'https://www.typescriptlang.org/docs/'
WHERE slug = 'typescript-official-docs';

-- 記事とタグの関連付け
INSERT INTO article_tags (article_id, tag_id) VALUES
-- React Hooks完全ガイド
(1, 1),
(1, 6),
(1, 7),
-- マーケティング戦略2026
(2, 4),
-- Kubernetes運用ガイド
(3, 5),
(3, 8),
-- Go言語入門
(4, 2),
(4, 6),
-- TypeScript公式ドキュメント
(5, 1),
-- コンテンツマーケティングの基礎
(6, 4)
ON CONFLICT (article_id, tag_id) DO NOTHING;
