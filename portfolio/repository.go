package portfolio

import (
	"Profile_You/models"
	"context"
)

type PortfolioRepository interface {
	CreatePortfolio(ctx context.Context, portfolio *models.Portfolio) error
	GetPortfolioByUserName(ctx context.Context, userName string, positionPortfolio int) (*models.Portfolio, error)
}
