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
		CreaterName: user.Username,
		Name:        "backend",
		Photo:       "D/photo/lop.png",
		Text:        "hahahahhahah text",
	}

	portfolioNotName := &models.Portfolio{
		ID:          1,
		Tags:        "lopatka, kolpa, nikola",
		URL:         "https://portfolio_you/&lopata/6",
		CreaterName: user.Username,
		Name:        "",
		Photo:       "D/photo/lop.png",
		Text:        "hahahahhahah text",
	}

	menu := &models.Menu{
		ID:          1,
		Name:        "backend",
		CreaterName: user.Username,
		ShortText:   "deceided",
	}

	err := p.CreatePortfolio(context.Background(), portfolio, user, menu)
	assert.NoError(t, err)

	err = p.CreatePortfolio(context.Background(), portfolioNotName, user, menu)
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

	portfolio := &models.Portfolio{
		ID:          1,
		Tags:        "lopatka, kolpa, nikola",
		URL:         "https://portfolio_you/&lopata/6",
		CreaterName: user.Username,
		Name:        "backend",
		Photo:       "D/photo/lop.png",
		Text:        "hahahahhahah text",
	}

	for i := 0; i < 10; i++ {

		menu := &models.Menu{
			ID:          i + 1,
			Name:        "backend",
			CreaterName: user.Username,
			ShortText:   "deceided",
		}

		p.CreatePortfolio(context.Background(), portfolio, user, menu)
	}

	returnedPortfolio, err := p.GetListPortfolioByUserName(context.Background(), user.Username)
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

	menu := &models.Menu{
		ID:          1,
		Name:        "backend",
		CreaterName: user.Username,
		ShortText:   "deceided",
	}

	p.CreatePortfolio(context.Background(), portfolio, user, menu)

	returnedPortfolio, err := p.GetPortfolioByUserName(context.Background(), "faner201", 1)
	assert.NoError(t, err)
	assert.Equal(t, portfolio, returnedPortfolio)

	returnedPortfolio, err = p.GetPortfolioByUserName(context.Background(), "", 1)
	assert.Error(t, err)
	assert.Equal(t, err, portfolios.ErrGetPortfolioByUserName)

}
