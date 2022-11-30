package mock

import (
	"Portfolio_You/models"
	"context"

	"github.com/stretchr/testify/mock"
)

type portfolioMock struct {
	mock.Mock
}

func (m *portfolioMock) CreatePortfolio(ctx context.Context, portfolio *models.Portfolio, user *models.User) error {
	args := m.Called(portfolio, user)

	return args.Error(0)
}

func (m *portfolioMock) GetPortfolioByUserName(ctx context.Context, userName string, portfolioID int) (*models.Portfolio, error) {
	args := m.Called(userName, portfolioID)

	return args.Get(0).(*models.Portfolio), args.Error(1)
}

func (m *portfolioMock) GetListPortfolioByUserName(ctx context.Context, userName string) ([]*models.Portfolio, error) {
	args := m.Called(userName)

	return args.Get(0).([]*models.Portfolio), args.Error(1)
}
