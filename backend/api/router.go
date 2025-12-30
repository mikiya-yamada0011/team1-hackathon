package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yamada-mikiya/team1-hackathon/config"
)

func SetupRouter(cfg *config.Config) *echo.Echo {
	router := echo.New()

	corsConfig := middleware.CORSConfig{
		AllowOrigins:     cfg.CORS.AllowedOrigins,
		AllowMethods:     cfg.CORS.AllowedMethods,
		AllowHeaders:     cfg.CORS.AllowedHeaders,
		AllowCredentials: cfg.CORS.AllowCredentials,
	}

	router.Use(middleware.CORSWithConfig(corsConfig))

	router.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	return router
}
