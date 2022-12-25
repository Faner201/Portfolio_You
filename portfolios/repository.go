package portfolios

import (
	"Portfolio_You/models"
	"context"
)

type PortfolioRepository interface {
	CreatePortfolio(ctx context.Context, portfolio *models.Portfolio, user *models.User) error
	CreateMenuPortfolio(ctx context.Context, user *models.User, menu *models.Menu) error
	GetPortfolioByUserName(ctx context.Context, userName, portfolioID string) (*models.Portfolio, error)
	GetListPortfolioByUserName(ctx context.Context, userName string) ([]*models.Menu, error)
	DeletePortfolio(ctx context.Context, user *models.User, portfolioID string) error
}
