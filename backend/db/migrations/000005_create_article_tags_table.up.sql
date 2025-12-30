CREATE TABLE IF NOT EXISTS article_tags (
    article_id UUID NOT NULL REFERENCES articles(id) ON DELETE CASCADE,
    tag_id UUID NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    PRIMARY KEY (article_id, tag_id)
);

-- article_idにインデックスを作成（記事のタグ検索の高速化）
CREATE INDEX idx_article_tags_article_id ON article_tags(article_id);

-- tag_idにインデックスを作成（タグの記事検索の高速化）
CREATE INDEX idx_article_tags_tag_id ON article_tags(tag_id);
