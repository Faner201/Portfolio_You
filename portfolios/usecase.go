package portfolios

import (
	"Portfolio_You/models"
	"context"
)

type UseCase interface {
	CreatePortfolio(ctx context.Context, user *models.User, portfolio *models.Portfolio) error
	CreateMenuPortfolio(ctx context.Context, user *models.User, menuPortfolio *models.Menu) error
	OpenPortfolio(ctx context.Context, user *models.User, portfolioID string) (*models.Portfolio, error)
	GetListPorfolio(ctx context.Context, user *models.User) (*[]models.Menu, error)
	DeletePortfolio(ctx context.Context, user *models.User, portfolioID string) error
}
