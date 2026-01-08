package controller

import (

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/yamada-mikiya/team1-hackathon/models"
	"github.com/yamada-mikiya/team1-hackathon/repositories"
	"github.com/yamada-mikiya/team1-hackathon/services"
	"gorm.io/gorm"
)

type ArticleController struct {
	service services.ArticleService
}

func NewArticleController(db *gorm.DB) *ArticleController {
	repo := repositories.NewArticleRepository(db)
	service := services.NewArticleService(repo)
	return &ArticleController{service: service}
}

// GetArticles は記事一覧を取得します
func (ac *ArticleController) GetArticles(c echo.Context) error {
	// ページネーションパラメータを取得
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page < 1 {
		page = 1
	}

	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit < 1 || limit > 100 {
		limit = 10
	}

	// フィルタパラメータを取得
	filters := repositories.ArticleFilters{
		Department: c.QueryParam("department"),
		Status:     c.QueryParam("status"),
	}

	// サービスから記事一覧を取得
	response, err := ac.service.GetArticles(filters, page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "記事の取得に失敗しました",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response)
}

//GetArticleBySlugは
func (ac *ArticleController) GetArticleBySlug(c echo.Context) error {
	slug := c.Param("slug")

	response, err := ac.service.GetArticleBySlug(slug)
	if err != nil {
			if err.Error() == "article not found" {
				return c.JSON(http.StatusNotFound, models.ErrorResponse{
				Error: "記事が見つかりません",
			})
		}

		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "記事の取得に失敗しました",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response)
}