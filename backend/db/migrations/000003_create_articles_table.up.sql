CREATE TABLE IF NOT EXISTS articles (
    id SERIAL PRIMARY KEY NOT NULL,
    author_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    article_type VARCHAR(50) NOT NULL CHECK (article_type IN ('markdown', 'external')),
    title VARCHAR(255) NOT NULL,
    content TEXT,
    external_url TEXT,
    thumbnail_url TEXT,
    slug VARCHAR(255) UNIQUE NOT NULL,
    department VARCHAR(50) NOT NULL CHECK (department IN ('Dev', 'MKT', 'Ops')),
    status VARCHAR(50) NOT NULL CHECK (status IN ('draft', 'internal', 'public')) DEFAULT 'draft',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- slugにインデックスを作成（検索の高速化）
CREATE INDEX idx_articles_slug ON articles(slug);

-- author_idにインデックスを作成（ユーザーの記事検索の高速化）
CREATE INDEX idx_articles_author_id ON articles(author_id);

-- departmentとstatusにインデックスを作成（フィルタリングの高速化）
CREATE INDEX idx_articles_department ON articles(department);
CREATE INDEX idx_articles_status ON articles(status);

-- updated_atの自動更新トリガーを設定
CREATE TRIGGER update_articles_updated_at
BEFORE UPDATE ON articles
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
