package repository

import (
	"Profile_You/models"
	"context"
)

type Repository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, id int) (*models.User, error)
	CreatePortfolio(ctx context.Context, portfolio *models.Portfolio) error
	GetPortfolioByUserName(ctx context.Context, userName string, positionPortfolio int) (*models.Portfolio, error)
	GetListPortfolioByUserName(ctx context.Context, userName string) ([]*models.Portfolio, error)
}
