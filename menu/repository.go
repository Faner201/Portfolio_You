package menu

import (
	"Portfolio_You/models"
	"context"
)

type MenuRepository interface {
	GetListPortfolioByUserName(ctx context.Context, userName string) ([]*models.Portfolio, error)
}
