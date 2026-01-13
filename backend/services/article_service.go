package services

import (
	"github.com/yamada-mikiya/team1-hackathon/models"
	"github.com/yamada-mikiya/team1-hackathon/repositories"
	"errors"
)

type ArticleService interface {
<<<<<<< HEAD
	GetArticles(filters repositories.ArticleFilters, page, limit int) (*models.ArticleListResponse, error)
	GetArticleBySlug(slug string, isAuthenticated bool) (*models.ArticleResponse, error)
=======
	GetArticles(filters ArticleFilters, page, limit int) (*models.ArticleListResponse, error)
	GetArticleBySlug(slug string) (*models.ArticleResponse, error)
>>>>>>> d20fcd9 (リファクタできた)
}

type ArticleFilters struct {
	Department string
	Status     string
}

type articleService struct {
	repo repositories.ArticleRepository
}

func NewArticleService(repo repositories.ArticleRepository) ArticleService {
	return &articleService{repo: repo}
}

// GetArticles は記事一覧を取得します
func (s *articleService) GetArticles(filters ArticleFilters, page, limit int) (*models.ArticleListResponse, error) {
	// リポジトリから記事を取得
	filtersInRepository := repositories.ArticleFilters{
		Department: filters.Department,
		Status:     filters.Status,
	}
	articles, totalCount, err := s.repo.FindAll(filtersInRepository, page, limit)
	if err != nil {
		return nil, err
	}

	// レスポンスを構築
	articleResponses := make([]models.ArticleResponse, len(articles))
	for i, article := range articles {
		articleResponses[i] = s.convertArticleToResponse(&article)
	}

	totalPages := int((totalCount + int64(limit) - 1) / int64(limit))

	return &models.ArticleListResponse{
		Articles:   articleResponses,
		TotalCount: int(totalCount),
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
	}, nil
}


// GetArticleBySlug はslugを指定して記事を取得します
func (s *articleService) GetArticleBySlug(slug string, isAuthenticated bool) (*models.ArticleResponse, error) {
	article, err := s.repo.FindBySlug(slug, isAuthenticated)
	if err != nil {
		return nil, err
	}

	//データがnilだった場合のチェック
	if article == nil {
		return nil, errors.New("article not found")
	}

	res := s.convertArticleToResponse(article)
	return &res, nil
}

// 共通の変換ロジック (Helper Method)
// DBモデル(*models.Article)を受け取り、レスポンスモデル(models.ArticleResponse)を返す
func (s *articleService) convertArticleToResponse(article *models.Article) models.ArticleResponse {
	authorResponse := models.AuthorResponse{}

	// Authorのnilチェックと詰め替え
	if article.Author != nil {
		authorResponse = models.AuthorResponse{
			ID:          article.Author.ID,
			Name:        article.Author.Name,
			Affiliation: article.Author.Affiliation,
			IconURL:     article.Author.IconURL,
		}
	}

	return models.ArticleResponse{
		ID:           article.ID,
		Title:        article.Title,
		ArticleType:  article.ArticleType,
		Content:      article.Content,
		ExternalURL:  article.ExternalURL,
		ThumbnailURL: article.ThumbnailURL,
		Slug:         article.Slug,
		Department:   article.Department,
		Status:       article.Status,
		Author:       authorResponse,
		CreatedAt:    article.CreatedAt,
		UpdatedAt:    article.UpdatedAt,
	}
}
