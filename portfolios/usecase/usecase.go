package usecase

import (
	"Portfolio_You/models"
	"Portfolio_You/portfolios"
	"context"
	"net/url"
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
	url := url.URL{
		Scheme: "https",
		Host:   "portfolio_you.com",
		Path:   "/portfolio/" + name + "%" + createrName,
	}

	return url.String()
}

func (p PortfolioUseCase) CreatePortfolio(ctx context.Context, user *models.User, url, tags, name, text, photo, shortText string) error {
	portf := &models.Portfolio{
		Tags:        tags,
		URL:         CreateURL(name, user.Username),
		CreaterName: user.Username,
		Name:        name,
		Photo:       photo,
		Text:        text,
	}

	menu := &models.Menu{
		Name:        name,
		CreaterName: user.Username,
		ShortText:   shortText,
		Photo:       photo,
	}

	return p.portfolioRepo.CreatePortfolio(context.Background(), portf, user, menu)
}

func (p PortfolioUseCase) OpenPortfolio(ctx context.Context, user *models.User, portfolioID int) (*models.Portfolio, error) {
	return p.portfolioRepo.GetPortfolioByUserName(context.Background(), user.Username, portfolioID)
}

func (p PortfolioUseCase) GetListPorfolio(ctx context.Context, user *models.User) ([]*models.Menu, error) {
	return p.portfolioRepo.GetListPortfolioByUserName(context.Background(), user.Username)
}

func (p PortfolioUseCase) DeletePortfolio(ctx context.Context, user *models.User, portfolioID int) error {
	return p.portfolioRepo.DeletePortfolio(context.Background(), user, portfolioID)
}
