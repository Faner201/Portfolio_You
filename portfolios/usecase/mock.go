package usecase

import (
	"Portfolio_You/models"
	"context"

	"github.com/stretchr/testify/mock"
)

type PortfolioUseCaseMock struct {
	mock.Mock
}

func (m *PortfolioUseCaseMock) CreatePortfolio(ctx context.Context, user *models.User, portfolio *models.Portfolio) error {
	args := m.Called(user, portfolio)
	return args.Error(0)
}

func (m *PortfolioUseCaseMock) CreateMenuPortfolio(ctx context.Context, user *models.User, menuPortfolio *models.Menu) error {
	args := m.Called(user, menuPortfolio)
	return args.Error(0)
}

func (m *PortfolioUseCaseMock) OpenPortfolio(ctx context.Context, user *models.User, portfolioID string) (*models.Portfolio, error) {
	args := m.Called(user, portfolioID)
	return args.Get(0).(*models.Portfolio), args.Error(1)
}

func (m *PortfolioUseCaseMock) GetListPorfolio(ctx context.Context, user *models.User) (*[]models.Menu, error) {
	args := m.Called(user)
	return args.Get(0).(*[]models.Menu), args.Error(1)
}

func (m *PortfolioUseCaseMock) DeletePortfolio(ctx context.Context, user *models.User, portfolioID string) error {
	args := m.Called(user, portfolioID)
	return args.Error(0)
}
