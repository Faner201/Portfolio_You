package mock

import (
	"Portfolio_You/models"
	"context"

	"github.com/stretchr/testify/mock"
)

type PortfolioMock struct {
	mock.Mock
}

func (m *PortfolioMock) CreatePortfolio(ctx context.Context, portfolio *models.Portfolio, user *models.User) error {
	args := m.Called(portfolio, user) // тут падает тест  usecase

	return args.Error(0)
}

func (m *PortfolioMock) CreateMenuPortfolio(ctx context.Context, user *models.User, menu *models.Menu) error {
	args := m.Called(user, menu)

	return args.Error(0)
}

func (m *PortfolioMock) GetPortfolioByUserName(ctx context.Context, userName string, portfolioID string) (*models.Portfolio, error) {
	args := m.Called(userName, portfolioID)

	return args.Get(0).(*models.Portfolio), args.Error(1)
}

func (m *PortfolioMock) GetListPortfolioByUserName(ctx context.Context, userName string) ([]*models.Menu, error) {
	args := m.Called(userName)

	return args.Get(0).([]*models.Menu), args.Error(1)
}

func (m *PortfolioMock) DeletePortfolio(ctx context.Context, user *models.User, portportfolioID string) error {
	args := m.Called(user, portportfolioID)

	return args.Error(0)
}
