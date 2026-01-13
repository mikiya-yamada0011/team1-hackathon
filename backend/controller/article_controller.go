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

// @Summary      記事一覧を取得
// @Description  公開されているブログ記事の一覧を取得します。ページネーション、部署フィルタ、ステータスフィルタをサポートしています。
// @Tags         記事 (Articles)
// @Accept       json
// @Produce      json
// @Param        page query int false "ページ番号 (デフォルト: 1)" default(1)
// @Param        limit query int false "1ページあたりの件数 (デフォルト: 10, 最大: 100)" default(10)
// @Param        department query string false "部署でフィルタ (Dev, MKT, Ops)" Enums(Dev, MKT, Ops)
// @Param        status query string false "ステータスでフィルタ (draft, internal, public)" Enums(draft, internal, public)
// @Success      200 {object} models.ArticleListResponse "記事一覧"
// @Failure      400 {object} models.ErrorResponse "リクエストパラメータが不正です"
// @Failure      500 {object} models.ErrorResponse "サーバー内部でエラーが発生しました"
// @Router       /api/articles [get]
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
	filters := services.ArticleFilters{
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

// GetArticleBySlug はslugを指定して記事を取得します
// @Summary      記事詳細を取得
// @Description  指定されたslugのブログ記事の詳細を取得します。
// @Tags         記事 (Articles)
// @Accept       json
// @Produce      json
// @Param        slug path string true "記事のスラグ" example("go-api-development")
// @Success      200 {object} models.ArticleResponse "記事詳細"
// @Failure      404 {object} models.ErrorResponse "記事が見つかりません"
// @Failure      500 {object} models.ErrorResponse "サーバー内部でエラーが発生しました"
// @Router       /api/articles/{slug} [get]
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