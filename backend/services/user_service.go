package services

import (
	"context"
	"errors"

	"github.com/yamada-mikiya/team1-hackathon/models"
	"github.com/yamada-mikiya/team1-hackathon/repositories"
)

type UserService interface {
	GetUserDetail(ctx context.Context, id int) (*models.UserDetailResponse, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

// GetUserDetail は公開用のユーザープロフィールを取得します
func (s *userService) GetUserDetail(ctx context.Context, id int) (*models.UserDetailResponse, error) {
	// 1. Repositoryから取得
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// 2. nilチェック (リポジトリで GORM依存を消したため、ここで判定)
	if user == nil {
		return nil, errors.New("user not found")
	}

	// 3. レスポンスへの変換
	// 注意: ここで Email や PasswordHash を含めないことでセキュリティを守ります
	res := models.UserDetailResponse{
		ID:          user.ID,
		Name:        user.Name,
		Affiliation: user.Affiliation,
		IconURL:     user.IconURL,
		CreatedAt:   user.CreatedAt,
	}

	return &res, nil
}