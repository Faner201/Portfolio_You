package usecase

import (
	"Portfolio_You/models"
	"Portfolio_You/portfolios"
	"context"
)

type PortfolioUseCase struct {
	portfolioRepo portfolios.PortfolioRepository
}

func NewPortfolioUseCase(portfolioRepo portfolios.PortfolioRepository) *PortfolioUseCase {
	return &PortfolioUseCase{
		portfolioRepo: portfolioRepo,
	}
}

func CreateURL(name, createrName string) string {
	url := "/portfolio/" + name + "%" + createrName
	return url
}

func (p *PortfolioUseCase) CreatePortfolio(ctx context.Context, user *models.User, portfolio *models.Portfolio) error {
	portfolio.Url = CreateURL(portfolio.Name, user.Username)
	return p.portfolioRepo.CreatePortfolio(ctx, user, portfolio)
}

func (p PortfolioUseCase) CreateMenuPortfolio(ctx context.Context, user *models.User, menuPortfolio *models.Menu) error {
	return p.portfolioRepo.CreateMenuPortfolio(ctx, user, menuPortfolio)
}

func (p PortfolioUseCase) OpenPortfolio(ctx context.Context, user *models.User, portfolioID string) (*models.Portfolio, error) {
	return p.portfolioRepo.GetPortfolioByUserName(ctx, user.Username, portfolioID)
}

func (p PortfolioUseCase) GetListPorfolio(ctx context.Context, user *models.User) (*[]models.Menu, error) {
	return p.portfolioRepo.GetListPortfolioByUserName(ctx, user.Username)
}

func (p PortfolioUseCase) DeletePortfolio(ctx context.Context, user *models.User, portfolioID string) error {
	return p.portfolioRepo.DeletePortfolio(ctx, user, portfolioID)
}
