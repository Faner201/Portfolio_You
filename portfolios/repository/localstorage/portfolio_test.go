package localstorage

import (
	"Portfolio_You/models"
	"Portfolio_You/portfolios"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePortfolio(t *testing.T) {
	p := NewPortfolioLocalStorage()

	user := &models.User{
		ID:       1,
		Username: "faner201",
		Password: "locaut",
		Email:    "polta@mail.ru",
	}

	portfolio := &models.Portfolio{
		ID:          1,
		Tags:        "lopatka, kolpa, nikola",
		URL:         "https://portfolio_you/&lopata/6",
		CreaterName: "faner201",
		Name:        "backend",
		Photo:       "D/photo/lop.png",
		Text:        "hahahahhahah text",
	}

	portfolioNotCreateName := &models.Portfolio{
		ID:          1,
		Tags:        "lopatka, kolpa, nikola",
		URL:         "https://portfolio_you/&lopata/6",
		CreaterName: "faner201",
		Name:        "",
		Photo:       "D/photo/lop.png",
		Text:        "hahahahhahah text",
	}

	err := p.CreatePortfolio(context.Background(), portfolio, user)
	assert.NoError(t, err)

	err = p.CreatePortfolio(context.Background(), portfolioNotCreateName, user)
	assert.Error(t, err)
	assert.Equal(t, err, portfolios.ErrCreatePortfolio)
}

func TestGetListPortfolioByUserName(t *testing.T) {
	p := NewPortfolioLocalStorage()

	user := &models.User{
		ID:       1,
		Username: "faner201",
		Password: "locaut",
		Email:    "polta@mail.ru",
	}

	for i := 0; i < 10; i++ {
		portfolio := &models.Portfolio{
			ID:          i + 1,
			Tags:        "lopatka, kolpa, nikola",
			URL:         "https://portfolio_you/&lopata/6",
			CreaterName: "faner201",
			Name:        "backend",
			Photo:       "D/photo/lop.png",
			Text:        "hahahahhahah text",
		}

		p.CreatePortfolio(context.Background(), portfolio, user)
	}

	returnedPortfolio, err := p.GetListPortfolioByUserName(context.Background(), "faner201")
	assert.NoError(t, err)
	assert.Equal(t, 10, len(returnedPortfolio))

	returnedPortfolio, err = p.GetListPortfolioByUserName(context.Background(), "lodsfsdfs")
	assert.Error(t, err)
	assert.Equal(t, err, portfolios.ErrGetListPortfolio)

}

func TestGetPortfolioByUserName(t *testing.T) {
	p := NewPortfolioLocalStorage()

	portfolio := &models.Portfolio{
		ID:          1,
		Tags:        "lopatka, kolpa, nikola",
		URL:         "https://portfolio_you/&lopata/6",
		CreaterName: "faner201",
		Name:        "backend",
		Photo:       "D/photo/lop.png",
		Text:        "hahahahhahah text",
	}

	user := &models.User{
		ID:       1,
		Username: "faner201",
		Password: "locaut",
		Email:    "polta@mail.ru",
	}

	p.CreatePortfolio(context.Background(), portfolio, user)

	returnedPortfolio, err := p.GetPortfolioByUserName(context.Background(), "faner201", 1)
	assert.NoError(t, err)
	assert.Equal(t, portfolio, returnedPortfolio)

	returnedPortfolio, err = p.GetPortfolioByUserName(context.Background(), "", 1)
	assert.Error(t, err)
	assert.Equal(t, err, portfolios.ErrGetPortfolioByUserName)

}
