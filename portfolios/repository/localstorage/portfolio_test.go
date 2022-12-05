package localstorage

import (
	"Portfolio_You/models"
	"Portfolio_You/portfolios"
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePortfolio(t *testing.T) {
	p := NewPortfolioLocalStorage()

	user := &models.User{
		ID:       "1",
		Username: "faner201",
		Password: "locaut",
		Email:    "polta@mail.ru",
	}

	portfolio := &models.Portfolio{
		ID:          "1",
		Url:         "https://portfolio_you/&lopata/6",
		CreaterUser: user.Username,
		Global: models.Global{
			Name: "Very best",
			View: "nice, popit",
			Bg:   "not nice, impossible",
		},
		Struct: models.Struct{
			StructList: []interface{}{
				models.BlockPhoto{
					Photo: "cd:/fdsfsdfsd",
				}, models.BlockPhoto{
					Photo: "cd:/fdsfwerwefds",
				}, models.BlockText{
					Text: "Very impoaible",
				},
			},
		},
	}

	portfolioNotName := &models.Portfolio{
		ID:          "1",
		Url:         "https://portfolio_you/&lopata/6",
		CreaterUser: user.Username,
		Global: models.Global{
			Name: "",
			View: "nice, popit",
			Bg:   "not nice, impossible",
		},
		Struct: models.Struct{
			StructList: []interface{}{
				models.BlockPhoto{
					Photo: "cd:/fdsfsdfsd",
				}, models.BlockPhoto{
					Photo: "cd:/fdsfwerwefds",
				}, models.BlockText{
					Text: "Very impoaible",
				},
			},
		},
	}

	err := p.CreatePortfolio(context.Background(), portfolio, user)
	assert.NoError(t, err)

	err = p.CreatePortfolio(context.Background(), portfolioNotName, user)
	assert.Error(t, err)
	assert.Equal(t, err, portfolios.ErrCreatePortfolio)
}

func TestCreateMenuPortfolio(t *testing.T) {
	p := NewPortfolioLocalStorage()

	user := &models.User{
		ID:       "1",
		Username: "faner201",
		Password: "locaut",
		Email:    "polta@mail.ru",
	}

	menu := &models.Menu{
		ID:          "1",
		Name:        "backend",
		CreaterName: user.Username,
		ShortText:   "deceided",
	}

	menuNotName := &models.Menu{
		ID:          "1",
		Name:        "",
		CreaterName: user.Username,
		ShortText:   "deceided",
	}

	err := p.CreateMenuPortfolio(context.Background(), user, menu)
	assert.NoError(t, err)

	err = p.CreateMenuPortfolio(context.Background(), user, menuNotName)
	assert.Error(t, err)
	assert.Equal(t, err, portfolios.ErrCreateMenuPortfolio)
}

func TestGetListPortfolioByUserName(t *testing.T) {
	p := NewPortfolioLocalStorage()

	user := &models.User{
		ID:       "1",
		Username: "faner201",
		Password: "locaut",
		Email:    "polta@mail.ru",
	}

	portfolio := &models.Portfolio{
		ID:          "1",
		Url:         "https://portfolio_you/&lopata/6",
		CreaterUser: user.Username,
		Global: models.Global{
			Name: "Very best",
			View: "nice, popit",
			Bg:   "not nice, impossible",
		},
		Struct: models.Struct{
			StructList: []interface{}{
				models.BlockPhoto{
					Photo: "cd:/fdsfsdfsd",
				}, models.BlockPhoto{
					Photo: "cd:/fdsfwerwefds",
				}, models.BlockText{
					Text: "Very impoaible",
				},
			},
		},
	}

	for i := 0; i < 10; i++ {

		menu := &models.Menu{
			ID:          fmt.Sprintf("id%d", i),
			Name:        "backend",
			CreaterName: user.Username,
			ShortText:   "deceided",
		}

		p.CreatePortfolio(context.Background(), portfolio, user)
		p.CreateMenuPortfolio(context.Background(), user, menu)
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

	user := &models.User{
		ID:       "1",
		Username: "faner201",
		Password: "locaut",
		Email:    "polta@mail.ru",
	}

	portfolio := &models.Portfolio{
		ID:          "1",
		Url:         "https://portfolio_you/&lopata/6",
		CreaterUser: user.Username,
		Global: models.Global{
			Name: "Very best",
			View: "nice, popit",
			Bg:   "not nice, impossible",
		},
		Struct: models.Struct{
			StructList: []interface{}{
				models.BlockPhoto{
					Photo: "cd:/fdsfsdfsd",
				}, models.BlockPhoto{
					Photo: "cd:/fdsfwerwefds",
				}, models.BlockText{
					Text: "Very impoaible",
				},
			},
		},
	}

	p.CreatePortfolio(context.Background(), portfolio, user)

	returnedPortfolio, err := p.GetPortfolioByUserName(context.Background(), "faner201", "1")
	assert.NoError(t, err)
	assert.Equal(t, portfolio, returnedPortfolio)

	returnedPortfolio, err = p.GetPortfolioByUserName(context.Background(), "", "1")
	assert.Error(t, err)
	assert.Equal(t, err, portfolios.ErrGetPortfolioByUserName)
}

func TestDeletePortfolio(t *testing.T) {
	p := NewPortfolioLocalStorage()

	user := &models.User{
		ID:       "1",
		Username: "faner201",
		Password: "locaut",
		Email:    "polta@mail.ru",
	}

	portfolio := &models.Portfolio{
		ID:          "1",
		Url:         "https://portfolio_you/&lopata/6",
		CreaterUser: user.Username,
		Global: models.Global{
			Name: "Very best",
			View: "nice, popit",
			Bg:   "not nice, impossible",
		},
		Struct: models.Struct{
			StructList: []interface{}{
				models.BlockPhoto{
					Photo: "cd:/fdsfsdfsd",
				}, models.BlockPhoto{
					Photo: "cd:/fdsfwerwefds",
				}, models.BlockText{
					Text: "Very impoaible",
				},
			},
		},
	}

	p.CreatePortfolio(context.Background(), portfolio, user)

	err := p.DeletePortfolio(context.Background(), user, portfolio.ID)
	assert.NoError(t, err)

	err = p.DeletePortfolio(context.Background(), user, "5")
	assert.Error(t, err)
	assert.Equal(t, err, portfolios.ErrDeletePortfolio)
}
