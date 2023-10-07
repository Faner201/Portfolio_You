package localstorage

import (
	"Portfolio_You/portfolios"
	faker "Portfolio_You/portfolios/faker"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePortfolio(t *testing.T) {
	p := NewPortfolioLocalStorage()

	user, _ := faker.GetUser()

	portfolio, _ := faker.GetPortfolio()
	portfolioNotName := *portfolio
	portfolioNotName.Name = ""

	err := p.CreatePortfolio(context.Background(), user, portfolio)
	assert.NoError(t, err)

	err = p.CreatePortfolio(context.Background(), user, &portfolioNotName)
	assert.Error(t, err)
	assert.Equal(t, err, portfolios.ErrCreatePortfolio)
}

func TestCreateMenuPortfolio(t *testing.T) {
	p := NewPortfolioLocalStorage()

	user, _ := faker.GetUser()

	menu, _ := faker.GetMenu()
	menuNotName := *menu
	menuNotName.Name = ""

	err := p.CreateMenuPortfolio(context.Background(), user, menu)
	assert.NoError(t, err)

	err = p.CreateMenuPortfolio(context.Background(), user, &menuNotName)
	assert.Error(t, err)
	assert.Equal(t, err, portfolios.ErrCreateMenuPortfolio)
}

func TestGetListPortfolioByUserName(t *testing.T) {
	p := NewPortfolioLocalStorage()

	user, _ := faker.GetUser()

	portfolio, _ := faker.GetPortfolio()

	for i := 0; i < 10; i++ {

		menu, _ := faker.GetMenu()

		p.CreatePortfolio(context.Background(), user, portfolio)
		p.CreateMenuPortfolio(context.Background(), user, menu)
	}

	returnedPortfolio, err := p.GetListPortfolioByUserName(context.Background(), user.Username)
	assert.NoError(t, err)
	assert.Equal(t, 10, len(*returnedPortfolio))

	_, err = p.GetListPortfolioByUserName(context.Background(), "lodsfsdfs")
	assert.Error(t, err)
	assert.Equal(t, err, portfolios.ErrGetListPortfolio)

}

func TestGetPortfolioByUserName(t *testing.T) {
	p := NewPortfolioLocalStorage()

	user, _ := faker.GetUser()

	portfolio, _ := faker.GetPortfolio()

	p.CreatePortfolio(context.Background(), user, portfolio)

	returnedPortfolio, err := p.GetPortfolioByUserName(context.Background(), portfolio.CreaterUser, portfolio.ID)
	assert.NoError(t, err)
	assert.Equal(t, portfolio, returnedPortfolio)

	_, err = p.GetPortfolioByUserName(context.Background(), "", portfolio.ID)
	assert.Error(t, err)
	assert.Equal(t, err, portfolios.ErrGetPortfolioByUserName)
}

func TestDeletePortfolio(t *testing.T) {
	p := NewPortfolioLocalStorage()

	user, _ := faker.GetUser()

	portfolio, _ := faker.GetPortfolio()

	p.CreatePortfolio(context.Background(), user, portfolio)

	err := p.DeletePortfolio(context.Background(), user, portfolio.ID)
	assert.NoError(t, err)

	err = p.DeletePortfolio(context.Background(), user, "5")
	assert.Error(t, err)
	assert.Equal(t, err, portfolios.ErrDeletePortfolio)
}
