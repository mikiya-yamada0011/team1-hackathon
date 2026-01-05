CREATE TABLE IF NOT EXISTS tags (
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(255) UNIQUE NOT NULL,
    is_category BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- nameにインデックスを作成（検索の高速化）
CREATE INDEX idx_tags_name ON tags(name);

-- is_categoryにインデックスを作成（カテゴリ一覧取得の高速化）
CREATE INDEX idx_tags_is_category ON tags(is_category);

-- updated_atの自動更新トリガーを設定
CREATE TRIGGER update_tags_updated_at
BEFORE UPDATE ON tags
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
