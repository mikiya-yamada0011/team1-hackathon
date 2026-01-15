package api

import (
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/yamada-mikiya/team1-hackathon/models"
)

// OptionalAuthMiddleware はJWTトークンがあれば検証してユーザー情報をセットし、
// なければゲスト扱いで通すミドルウェア
func OptionalAuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// JWTトークンの取得を試みる
			tokenString := extractToken(c)

			if tokenString != "" {
				// トークンがある場合は検証を試みる
				token, err := jwt.ParseWithClaims(tokenString, &models.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
					return []byte(os.Getenv("SECRET_KEY")), nil
				})

				if err == nil && token.Valid {
					// トークンが有効な場合、ユーザー情報をContextにセット
					if claims, ok := token.Claims.(*models.JwtCustomClaims); ok {
						c.Set("user", claims)
						c.Set("user_id", claims.UserID)
					}
				}
				// エラーがあっても続行（ゲスト扱い）
			}

			// トークンがない、または無効でもハンドラーを実行
			return next(c)
		}
	}
}

// extractToken はリクエストからJWTトークンを抽出する
func extractToken(c echo.Context) string {
	// 1. Cookieから取得を試みる
	if cookie, err := c.Cookie("token"); err == nil {
		return cookie.Value
	}

	// 2. Authorization headerから取得を試みる
	auth := c.Request().Header.Get("Authorization")
	if auth != "" {
		parts := strings.Split(auth, " ")
		if len(parts) == 2 && parts[0] == "Bearer" {
			return parts[1]
		}
	}

	return ""
}
