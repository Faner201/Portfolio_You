package database

import (
	"Portfolio_You/models"
	"context"
)

type PortfolioRepository struct{}

func NewPortfolioRepository() *PortfolioRepository {
	return &PortfolioRepository{}
}

func (p PortfolioRepository) CreatePortfolio(ctx context.Context, portfolio *models.Portfolio, user *models.User) error {
	return nil
}

func (p PortfolioRepository) CreateMenuPortfolio(ctx context.Context, user *models.User, menu *models.Menu) error {
	return nil
}

func (p PortfolioRepository) GetPortfolioByUserName(ctx context.Context, userName string, portfolioID string) (*models.Portfolio, error) {
	return &models.Portfolio{}, nil
}

func (p PortfolioRepository) GetListPortfolioByUserName(ctx context.Context, userName string) ([]*models.Menu, error) {
	return []*models.Menu{}, nil
}
func (p PortfolioRepository) DeletePortfolio(ctx context.Context, user *models.User, portportfolioID string) error {
	return nil
}
