package portfolios

import (
	"Portfolio_You/models"
	"context"
)

type UseCase interface {
	CreatePortfolio(ctx context.Context, user *models.User, name, view, bg, shortText, photo string, structs []interface{}) error
	CreateMenuPortfolio(ctx context.Context, user *models.User, name, shortText, photo string) error
	OpenPortfolio(ctx context.Context, user *models.User, portfolioID string) (*models.Portfolio, error)
	GetListPorfolio(ctx context.Context, user *models.User) ([]*models.Menu, error)
	DeletePortfolio(ctx context.Context, user *models.User, portfolioID string) error
}
