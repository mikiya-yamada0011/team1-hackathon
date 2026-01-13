package models

import "time"

// ArticleListResponse は記事一覧取得のレスポンス
type ArticleListResponse struct {
	Articles   []ArticleResponse `json:"articles"`
	TotalCount int               `json:"total_count" example:"100"`
	Page       int               `json:"page" example:"1"`
	Limit      int               `json:"limit" example:"10"`
	TotalPages int               `json:"total_pages" example:"10"`
} // @name ArticleListResponse

// ArticleResponse は記事の詳細レスポンス
type ArticleResponse struct {
	ID           int            `json:"id" example:"1"`
	Title        string         `json:"title" example:"Go言語でのAPI開発入門"`
	ArticleType  string         `json:"article_type" example:"markdown" enums:"markdown,external"`
	Content      *string        `json:"content,omitempty" example:"記事の本文です..."`
	ExternalURL  *string        `json:"external_url,omitempty" example:"https://example.com/article"`
	ThumbnailURL *string        `json:"thumbnail_url,omitempty" example:"https://example.com/thumbnail.jpg"`
	Slug         string         `json:"slug" example:"go-api-development"`
	Department   string         `json:"department" example:"Dev" enums:"Dev,MKT,Ops"`
	Status       string         `json:"status" example:"public" enums:"draft,internal,public"`
	Author       AuthorResponse `json:"author"`
	CreatedAt    time.Time      `json:"created_at" example:"2026-01-06T12:00:00Z"`
	UpdatedAt    time.Time      `json:"updated_at" example:"2026-01-06T12:00:00Z"`
	Tags         []string       `json:"tags" example:"Go,Backend,Echo"`
} // @name ArticleResponse

// AuthorResponse は記事の著者情報
type AuthorResponse struct {
	ID          int     `json:"id" example:"1"`
	Name        string  `json:"name" example:"山田太郎"`
	Affiliation *string `json:"affiliation,omitempty" example:"開発部"`
	IconURL     *string `json:"icon_url,omitempty" example:"https://example.com/icon.jpg"`
} // @name AuthorResponse

// ErrorResponse はエラーレスポンス
type ErrorResponse struct {
	Error   string `json:"error" example:"エラーメッセージ"`
	Message string `json:"message,omitempty" example:"詳細なエラー情報"`
} // @name ErrorResponse
