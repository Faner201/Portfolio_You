package portfolio

import (
	"Portfolio_You/models"
	"context"
)

type UseCase interface {
	CreatePortfolio(ctx context.Context, user *models.User, url, tags, name string) error
	OpenPortfolio(ctx context.Context, user *models.User, portfolioID int) (*models.Portfolio, error)
}
