package models

import "time"

// Article はブログ記事のモデル
type Article struct {
	ID           int       `json:"id" gorm:"primaryKey;autoIncrement"`
	AuthorID     int       `json:"author_id" gorm:"not null"`
	ArticleType  string    `json:"article_type" gorm:"type:varchar(50);not null"`
	Title        string    `json:"title" gorm:"type:varchar(255);not null"`
	Content      *string   `json:"content" gorm:"type:text"`
	ExternalURL  *string   `json:"external_url" gorm:"type:text"`
	ThumbnailURL *string   `json:"thumbnail_url" gorm:"type:text"`
	Slug         string    `json:"slug" gorm:"type:varchar(255);unique;not null"`
	Department   string    `json:"department" gorm:"type:varchar(50);not null"`
	Status       string    `json:"status" gorm:"type:varchar(50);not null;default:draft"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Author       *User     `json:"author,omitempty" gorm:"foreignKey:AuthorID"`
}

// User はユーザーのモデル
type User struct {
	ID           int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name         string    `json:"name" gorm:"type:varchar(255);not null"`
	Affiliation  *string   `json:"affiliation" gorm:"type:varchar(255)"`
	PasswordHash string    `json:"-" gorm:"type:varchar(255);not null"`
	IconURL      *string   `json:"icon_url" gorm:"type:text"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
