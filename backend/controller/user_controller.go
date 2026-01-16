package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/yamada-mikiya/team1-hackathon/models"
	"github.com/yamada-mikiya/team1-hackathon/services"
)

type UserController interface {
	GetUserDetailHandler(c echo.Context) error
}

type userController struct {
	s services.UserService
}

func NewUserController(s services.UserService) UserController {
	return &userController{s: s}
}

// GetUserDetailHandler はユーザー詳細を取得します
// @Summary      ユーザー詳細取得
// @Description  指定されたIDのユーザー公開プロフィール（名前、所属、アイコン）を取得します。メールアドレスは含まれません。
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  models.UserDetailResponse
// @Failure      400  {object}  models.ErrorResponse
// @Failure      404  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Router       /users/{id} [get]
func (uc *userController) GetUserDetailHandler(c echo.Context) error {
	// 1. URLパラメータからIDを取得
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "IDの形式が不正です",
		})
	}

	// 2. Service呼び出し (Contextを渡す)
	res, err := uc.s.GetUserDetail(c.Request().Context(), id)

	if err != nil {
		// Serviceで "user not found" というエラーを作ったので、それを検知します
		if err.Error() == "user not found" {
			return c.JSON(http.StatusNotFound, models.ErrorResponse{
				Error: "ユーザーが見つかりません",
			})
		}

		// その他のエラー (DB接続エラーなど)
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "ユーザー情報の取得に失敗しました",
			Message: err.Error(),
		})
	}

	// 3. レスポンスを返す
	return c.JSON(http.StatusOK, res)
}