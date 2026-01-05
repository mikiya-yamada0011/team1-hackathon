-- 開発環境専用のテストデータ

-- テストユーザーの挿入
INSERT INTO users (id, name, affiliation, password_hash, icon_url) VALUES
('11111111-1111-1111-1111-111111111111', '田中 太郎', 'Dev部門', '$2a$10$dummyhash1111111111111111111111111111111111111111', 'https://i.pravatar.cc/150?img=1'),
('22222222-2222-2222-2222-222222222222', '佐藤 花子', 'MKT部門', '$2a$10$dummyhash2222222222222222222222222222222222222222', 'https://i.pravatar.cc/150?img=2'),
('33333333-3333-3333-3333-333333333333', '鈴木 一郎', 'Ops部門', '$2a$10$dummyhash3333333333333333333333333333333333333333', 'https://i.pravatar.cc/150?img=3'),
('44444444-4444-4444-4444-444444444444', '高橋 美咲', 'Dev部門', '$2a$10$dummyhash4444444444444444444444444444444444444444', 'https://i.pravatar.cc/150?img=4')
ON CONFLICT (id) DO NOTHING;

-- タグの挿入
INSERT INTO tags (id, name, is_category) VALUES
('aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', 'React', false),
('bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', 'Go', false),
('cccccccc-cccc-cccc-cccc-cccccccccccc', 'Docker', false),
('dddddddd-dddd-dddd-dddd-dddddddddddd', 'マーケティング', false),
('eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeee', 'インフラ', false),
('ffffffff-ffff-ffff-ffff-ffffffffffff', 'チュートリアル', true),
('00000001-0001-0001-0001-000000000001', 'ベストプラクティス', true),
('00000002-0002-0002-0002-000000000002', 'トラブルシューティング', true)
ON CONFLICT (name) DO NOTHING;

-- テスト記事の挿入
INSERT INTO articles (id, author_id, article_type, title, content, slug, department, status, thumbnail_url) VALUES
(
    'a1111111-1111-1111-1111-111111111111',
    '11111111-1111-1111-1111-111111111111',
    'markdown',
    'React Hooks完全ガイド',
    E'# React Hooks完全ガイド\n\nReact Hooksの基本から応用まで解説します。\n\n## useState\n\n最も基本的なHookです。\n\n```jsx\nconst [count, setCount] = useState(0);\n```\n\n## useEffect\n\n副作用を扱うためのHookです。\n\n```jsx\nuseEffect(() => {\n  document.title = `You clicked ${count} times`;\n}, [count]);\n```',
    'react-hooks-guide',
    'Dev',
    'public',
    'https://images.unsplash.com/photo-1633356122544-f134324a6cee?w=800'
),
(
    'a2222222-2222-2222-2222-222222222222',
    '22222222-2222-2222-2222-222222222222',
    'markdown',
    'マーケティング戦略2026',
    E'# マーケティング戦略2026\n\n2026年のマーケティングトレンドと戦略について解説します。\n\n## SNSマーケティング\n\n最新のSNSマーケティング手法を紹介します。\n\n- Instagram\n- TikTok\n- X (Twitter)\n\n## データ分析\n\nマーケティングデータの分析手法について説明します。',
    'marketing-strategy-2026',
    'MKT',
    'public',
    'https://images.unsplash.com/photo-1460925895917-afdab827c52f?w=800'
),
(
    'a3333333-3333-3333-3333-333333333333',
    '33333333-3333-3333-3333-333333333333',
    'markdown',
    'Docker環境構築ガイド',
    E'# Docker環境構築ガイド\n\nDockerを使った開発環境の構築方法を解説します。\n\n## Dockerのインストール\n\nまずはDockerをインストールしましょう。\n\n## docker-composeの使い方\n\n複数のコンテナを管理するためのツールです。\n\n```yaml\nversion: ''3.8''\nservices:\n  app:\n    build: .\n    ports:\n      - "8080:8080"\n```',
    'docker-environment-guide',
    'Ops',
    'public',
    'https://images.unsplash.com/photo-1605745341112-85968b19335b?w=800'
),
(
    'a4444444-4444-4444-4444-444444444444',
    '44444444-4444-4444-4444-444444444444',
    'markdown',
    'Go言語入門',
    E'# Go言語入門\n\nGo言語の基礎から学びます。\n\n## Goとは\n\nGoogleが開発したプログラミング言語です。\n\n## 環境構築\n\n```bash\n# Goのインストール\nbrew install go\n\n# バージョン確認\ngo version\n```\n\n## Hello World\n\n```go\npackage main\n\nimport "fmt"\n\nfunc main() {\n    fmt.Println("Hello, World!")\n}\n```',
    'golang-introduction',
    'Dev',
    'internal',
    'https://images.unsplash.com/photo-1617854818583-09e7f077a156?w=800'
),
(
    'a5555555-5555-5555-5555-555555555555',
    '11111111-1111-1111-1111-111111111111',
    'external',
    'TypeScript公式ドキュメント',
    NULL,
    'typescript-official-docs',
    'Dev',
    'public',
    'https://images.unsplash.com/photo-1587620962725-abab7fe55159?w=800'
),
(
    'a6666666-6666-6666-6666-666666666666',
    '22222222-2222-2222-2222-222222222222',
    'markdown',
    'コンテンツマーケティングの基礎',
    E'# コンテンツマーケティングの基礎\n\n効果的なコンテンツマーケティングの手法を紹介します。\n\n## コンテンツの企画\n\nターゲットオーディエンスを明確にしましょう。\n\n## 配信チャネル\n\n- ブログ\n- メールマガジン\n- SNS',
    'content-marketing-basics',
    'MKT',
    'draft',
    'https://images.unsplash.com/photo-1542744094-3a31f272c490?w=800'
)
ON CONFLICT (id) DO NOTHING;

-- 外部記事のURLを設定
UPDATE articles 
SET external_url = 'https://www.typescriptlang.org/docs/'
WHERE id = 'a5555555-5555-5555-5555-555555555555';

-- 記事とタグの関連付け
INSERT INTO article_tags (article_id, tag_id) VALUES
-- React Hooks完全ガイド
('a1111111-1111-1111-1111-111111111111', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa'),
('a1111111-1111-1111-1111-111111111111', 'ffffffff-ffff-ffff-ffff-ffffffffffff'),
('a1111111-1111-1111-1111-111111111111', '00000001-0001-0001-0001-000000000001'),
-- マーケティング戦略2026
('a2222222-2222-2222-2222-222222222222', 'dddddddd-dddd-dddd-dddd-dddddddddddd'),
('a2222222-2222-2222-2222-222222222222', '00000001-0001-0001-0001-000000000001'),
-- Docker環境構築ガイド
('a3333333-3333-3333-3333-333333333333', 'cccccccc-cccc-cccc-cccc-cccccccccccc'),
('a3333333-3333-3333-3333-333333333333', 'eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeee'),
('a3333333-3333-3333-3333-333333333333', 'ffffffff-ffff-ffff-ffff-ffffffffffff'),
-- Go言語入門
('a4444444-4444-4444-4444-444444444444', 'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb'),
('a4444444-4444-4444-4444-444444444444', 'ffffffff-ffff-ffff-ffff-ffffffffffff'),
-- TypeScript公式ドキュメント
('a5555555-5555-5555-5555-555555555555', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa'),
-- コンテンツマーケティングの基礎
('a6666666-6666-6666-6666-666666666666', 'dddddddd-dddd-dddd-dddd-dddddddddddd')
ON CONFLICT DO NOTHING;
