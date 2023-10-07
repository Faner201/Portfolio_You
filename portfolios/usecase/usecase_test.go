package usecase

import (
	"Portfolio_You/models"
	faker "Portfolio_You/portfolios/faker"
	"Portfolio_You/portfolios/repository/mock"
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePortfolio(t *testing.T) {
	rep := new(mock.PortfolioMock)

	pf := NewPortfolioUseCase(rep)

	user, _ := faker.GetUser()

	portfolio, _ := faker.GetPortfolio()

	rep.On("CreatePortfolio", portfolio, user).Return(nil)
	err := pf.CreatePortfolio(context.Background(), user, portfolio)
	assert.NoError(t, err)
}

func TestCreateMenuPortfoli0(t *testing.T) {
	rep := new(mock.PortfolioMock)

	pf := NewPortfolioUseCase(rep)

	user, _ := faker.GetUser()

	menu, _ := faker.GetMenu()

	rep.On("CreateMenuPortfolio", user, menu).Return(nil)
	err := pf.CreateMenuPortfolio(context.Background(), user, menu)
	assert.NoError(t, err)
}

func TestOpenPortfolio(t *testing.T) {
	rep := new(mock.PortfolioMock)

	pf := NewPortfolioUseCase(rep)

	user, _ := faker.GetUser()

	portfolioID := "1"
	portfolio, _ := faker.GetPortfolio()

	log.Println(portfolio.Colors)

	rep.On("GetPortfolioByUserName", user.Username, portfolioID).Return(portfolio, nil)
	portf, err := pf.OpenPortfolio(context.Background(), user, portfolioID)
	assert.NoError(t, err)
	assert.NotEmpty(t, portf)
}

func TestGetListPortfolio(t *testing.T) {
	rep := new(mock.PortfolioMock)

	pf := NewPortfolioUseCase(rep)

	user, _ := faker.GetUser()

	menu, _ := faker.GetMenu()
	menuList := []models.Menu{}
	menuList = append(menuList, *menu)

	rep.On("GetListPortfolioByUserName", user.Username).Return(&menuList, nil)
	list, err := pf.GetListPorfolio(context.Background(), user)
	assert.NoError(t, err)
	assert.NotEmpty(t, list)
}

func TestDeletePortfolio(t *testing.T) {
	rep := new(mock.PortfolioMock)

	pf := NewPortfolioUseCase(rep)

	portfolioID := "1"

	user, _ := faker.GetUser()

	rep.On("DeletePortfolio", user, portfolioID).Return(nil)
	err := pf.DeletePortfolio(context.Background(), user, portfolioID)
	assert.NoError(t, err)
}
