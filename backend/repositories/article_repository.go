package repositories

import (
	"github.com/yamada-mikiya/team1-hackathon/models"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	FindAll(filters ArticleFilters, page, limit int) ([]models.Article, int64, error)
	FindBySlug(slug string) (*models.Article, error)
}

type ArticleFilters struct {
	Department string
	Status     string
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{db: db}
}

// FindAll はフィルタとページネーションを適用して記事一覧を取得します
func (r *articleRepository) FindAll(filters ArticleFilters, page, limit int) ([]models.Article, int64, error) {
	var articles []models.Article
	var totalCount int64

	// クエリを構築
	query := r.db.Model(&models.Article{}).Preload("Author")

	// フィルタを適用
	if filters.Department != "" {
		query = query.Where("department = ?", filters.Department)
	}

	if filters.Status != "" {
		query = query.Where("status = ?", filters.Status)
	} else {
		// デフォルトでは公開記事のみ
		query = query.Where("status = ?", "public")
	}

	// 総件数を取得
	if err := query.Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	// ページネーション適用
	offset := (page - 1) * limit
	if err := query.Order("created_at DESC").Limit(limit).Offset(offset).Find(&articles).Error; err != nil {
		return nil, 0, err
	}

	return articles, totalCount, nil
}

// FindBySlug はslugを指定して記事を取得します
func (r *articleRepository) FindBySlug(slug string) (*models.Article, error) {
	var article models.Article
	result := r.db.Preload("Author").Where("slug = ? AND status = ?", slug, "public").Find(&article)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &article, nil
}
