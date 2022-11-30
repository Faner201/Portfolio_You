package portfolios

import (
	"Portfolio_You/models"
	"context"
)

type PortfolioRepository interface {
	CreatePortfolio(ctx context.Context, portfolio *models.Portfolio, user *models.User, menu *models.Menu) error
	GetPortfolioByUserName(ctx context.Context, userName string, portfolioID int) (*models.Portfolio, error)
	GetListPortfolioByUserName(ctx context.Context, userName string) ([]*models.Menu, error)
}
