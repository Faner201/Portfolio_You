package portfolios

import (
	"Portfolio_You/models"
	"context"
)

type UseCase interface {
	CreatePortfolio(ctx context.Context, user *models.User, url, tags, name, text, photo string) error
	OpenPortfolio(ctx context.Context, user *models.User, portfolioID int) (*models.Portfolio, error)
	GetListPorfolio(ctx context.Context, user *models.User) ([]*models.Portfolio, error)
}
