package menu

import (
	"Portfolio_You/models"
	"context"
)

type UseCase interface {
	GetListPorfolio(ctx context.Context, portfolio *models.Portfolio) ([]*models.Portfolio, error)
}
