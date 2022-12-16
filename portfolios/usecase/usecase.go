package usecase

import (
	"Portfolio_You/models"
	"Portfolio_You/portfolios"
	"context"
	"log"
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

func (p *PortfolioUseCase) CreatePortfolio(ctx context.Context, user *models.User, name string, text *[]models.Text,
	photo *[]models.Photo, colors *models.Colors, structs *[][]models.Block) error {
	portf := &models.Portfolio{
		Url:         CreateURL(name, user.Username),
		CreaterUser: user.Username,
		Name:        name,
		Text:        text,
		Photo:       photo,
		Colors:      colors,
		Struct:      structs,
	}

	log.Println(portf)

	return p.portfolioRepo.CreatePortfolio(ctx, portf, user)
}

func (p PortfolioUseCase) CreateMenuPortfolio(ctx context.Context, user *models.User, name, shortText, photo string) error {
	menu := &models.Menu{
		Name:        name,
		CreaterName: user.Username,
		ShortText:   shortText,
		Photo:       photo,
	}

	return p.portfolioRepo.CreateMenuPortfolio(ctx, user, menu)
}

func (p PortfolioUseCase) OpenPortfolio(ctx context.Context, user *models.User, portfolioID string) (*models.Portfolio, error) {
	return p.portfolioRepo.GetPortfolioByUserName(ctx, user.Username, portfolioID)
}

func (p PortfolioUseCase) GetListPorfolio(ctx context.Context, user *models.User) ([]*models.Menu, error) {
	return p.portfolioRepo.GetListPortfolioByUserName(ctx, user.Username)
}

func (p PortfolioUseCase) DeletePortfolio(ctx context.Context, user *models.User, portfolioID string) error {
	return p.portfolioRepo.DeletePortfolio(ctx, user, portfolioID)
}
