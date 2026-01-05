package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/yamada-mikiya/team1-hackathon/config"
	"github.com/yamada-mikiya/team1-hackathon/controller"
	"gorm.io/gorm"
)

func SetupRouter(cfg *config.Config, db *gorm.DB) *echo.Echo {
	router := echo.New()

	corsConfig := middleware.CORSConfig{
		AllowOrigins:     cfg.CORS.AllowedOrigins,
		AllowMethods:     cfg.CORS.AllowedMethods,
		AllowHeaders:     cfg.CORS.AllowedHeaders,
		AllowCredentials: cfg.CORS.AllowCredentials,
	}

	router.Use(middleware.CORSWithConfig(corsConfig))
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	// ヘルスチェック
	router.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	// Swagger UI
	router.GET("/swagger/*", echoSwagger.WrapHandler)

	// コントローラー初期化
	articleController := controller.NewArticleController(db)

	// APIルート
	api := router.Group("/api")
	{
		// 記事関連
		articles := api.Group("/articles")
		{
			articles.GET("", articleController.GetArticles)
			articles.GET("/:slug", articleController.GetArticleBySlug)
		}
	}

	return router
}
