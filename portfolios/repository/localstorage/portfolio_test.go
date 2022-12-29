package localstorage

import (
	"Portfolio_You/models"
	"Portfolio_You/portfolios"
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func User(id, username, password, email string) *models.User {
	return &models.User{
		ID:       id,
		Username: username,
		Password: password,
		Email:    email,
	}
}

func Portfolio(id, url, createrUser, name string, text *[]models.Text, photo *[]models.Photo, colors *models.Colors,
	structs *[][]models.Block) *models.Portfolio {
	return &models.Portfolio{
		ID:          id,
		Url:         url,
		CreaterUser: createrUser,
		Name:        name,
		Text:        text,
		Photo:       photo,
		Colors:      colors,
		Struct:      structs,
	}
}

func Menu(id, name, createrUser, shortText string) *models.Menu {
	return &models.Menu{
		ID:          id,
		CreaterName: createrUser,
		Name:        name,
		ShortText:   shortText,
	}
}

func TestCreatePortfolio(t *testing.T) {
	p := NewPortfolioLocalStorage()

	id := "1"
	username := "faner201"
	password := "locaut"
	email := "polta@mail.ru"

	user := User(id, username, password, email)

	url := "/portfolios/aboba&faner201"
	createrUser := user.Username
	name := "aboba"
	notName := ""
	text := &[]models.Text{}
	photo := &[]models.Photo{
		{
			Addres: "cd/42342dsfs3",
		},
		{
			Addres: "cd/42342dsfs3",
		},
	}
	colors := &models.Colors{
		Base:      "#fff",
		Text:      "#dsfsdfs",
		Contrast:  "#fdsfsdcxs",
		Primary:   "#fdsfsdf",
		Secondary: "#fdsfxz",
	}

	structs := &[][]models.Block{
		{
			{
				Type:     "text",
				Location: "1",
			},
			{
				Type:     "text",
				Location: "2",
			},
			{
				Type:     "image",
				Location: "1",
			},
		},
		{
			{
				Type:     "image",
				Location: "2",
			},
			{
				Type:     "text",
				Location: "3",
			},
			{
				Type:     "text",
				Location: "4",
			},
		},
	}

	portfolio := Portfolio(id, url, createrUser, name, text, photo, colors, structs)
	portfolioNotName := Portfolio(id, url, createrUser, notName, text, photo, colors, structs)

	err := p.CreatePortfolio(context.Background(), portfolio, user)
	assert.NoError(t, err)

	err = p.CreatePortfolio(context.Background(), portfolioNotName, user)
	assert.Error(t, err)
	assert.Equal(t, err, portfolios.ErrCreatePortfolio)
}

func TestCreateMenuPortfolio(t *testing.T) {
	p := NewPortfolioLocalStorage()

	id := "1"
	username := "faner201"
	password := "locaut"
	email := "polta@mail.ru"

	user := User(id, username, password, email)

	name := "backend"
	notName := ""
	createrName := user.Username
	shortText := "deceided"

	menu := Menu(id, name, createrName, shortText)
	menuNotName := Menu(id, notName, createrName, shortText)

	err := p.CreateMenuPortfolio(context.Background(), user, menu)
	assert.NoError(t, err)

	err = p.CreateMenuPortfolio(context.Background(), user, menuNotName)
	assert.Error(t, err)
	assert.Equal(t, err, portfolios.ErrCreateMenuPortfolio)
}

func TestGetListPortfolioByUserName(t *testing.T) {
	p := NewPortfolioLocalStorage()

	id := "1"
	username := "faner201"
	password := "locaut"
	email := "polta@mail.ru"

	user := User(id, username, password, email)

	shortText := "deceided"

	url := "/portfolios/aboba&faner201"
	createrUser := user.Username
	name := "aboba"
	text := &[]models.Text{
		{
			Sludge: "avtobot",
			Style:  "limitic",
			Size:   "12",
		},
		{
			Sludge: "avtobot",
			Style:  "limitic",
			Size:   "12",
		},
	}
	photo := &[]models.Photo{
		{
			Addres: "cd/42342dsfs3",
		},
		{
			Addres: "cd/42342dsfs3",
		},
	}
	colors := &models.Colors{
		Base:      "#fff",
		Text:      "#dsfsdfs",
		Contrast:  "#fdsfsdcxs",
		Primary:   "#fdsfsdf",
		Secondary: "#fdsfxz",
	}

	structs := &[][]models.Block{
		{
			{
				Type:     "text",
				Location: "1",
			},
			{
				Type:     "text",
				Location: "2",
			},
			{
				Type:     "image",
				Location: "1",
			},
		},
		{
			{
				Type:     "image",
				Location: "2",
			},
			{
				Type:     "text",
				Location: "3",
			},
			{
				Type:     "text",
				Location: "4",
			},
		},
	}

	portfolio := Portfolio(id, url, createrUser, name, text, photo, colors, structs)

	for i := 0; i < 10; i++ {

		menu := Menu(fmt.Sprintf("id%d", i), name, createrUser, shortText)

		p.CreatePortfolio(context.Background(), portfolio, user)
		p.CreateMenuPortfolio(context.Background(), user, menu)
	}

	returnedPortfolio, err := p.GetListPortfolioByUserName(context.Background(), user.Username)
	assert.NoError(t, err)
	assert.Equal(t, 10, len(*returnedPortfolio))

	returnedPortfolio, err = p.GetListPortfolioByUserName(context.Background(), "lodsfsdfs")
	assert.Error(t, err)
	assert.Equal(t, err, portfolios.ErrGetListPortfolio)

}

func TestGetPortfolioByUserName(t *testing.T) {
	p := NewPortfolioLocalStorage()

	id := "1"
	username := "faner201"
	password := "locaut"
	email := "polta@mail.ru"

	user := User(id, username, password, email)

	url := "/portfolios/aboba&faner201"
	createrUser := user.Username
	name := "aboba"
	text := &[]models.Text{
		{
			Sludge: "avtobot",
			Style:  "limitic",
			Size:   "12",
		},
		{
			Sludge: "avtobot",
			Style:  "limitic",
			Size:   "12",
		},
	}
	photo := &[]models.Photo{
		{
			Addres: "cd/42342dsfs3",
		},
		{
			Addres: "cd/42342dsfs3",
		},
	}
	colors := &models.Colors{
		Base:      "#fff",
		Text:      "#dsfsdfs",
		Contrast:  "#fdsfsdcxs",
		Primary:   "#fdsfsdf",
		Secondary: "#fdsfxz",
	}

	structs := &[][]models.Block{
		{
			{
				Type:     "text",
				Location: "1",
			},
			{
				Type:     "text",
				Location: "2",
			},
			{
				Type:     "image",
				Location: "1",
			},
		},
		{
			{
				Type:     "image",
				Location: "2",
			},
			{
				Type:     "text",
				Location: "3",
			},
			{
				Type:     "text",
				Location: "4",
			},
		},
	}

	portfolio := Portfolio(id, url, createrUser, name, text, photo, colors, structs)

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

	id := "1"
	username := "faner201"
	password := "locaut"
	email := "polta@mail.ru"

	user := User(id, username, password, email)

	url := "/portfolios/aboba&faner201"
	createrUser := user.Username
	name := "aboba"
	text := &[]models.Text{
		{
			Sludge: "avtobot",
			Style:  "limitic",
			Size:   "12",
		},
		{
			Sludge: "avtobot",
			Style:  "limitic",
			Size:   "12",
		},
	}
	photo := &[]models.Photo{
		{
			Addres: "cd/42342dsfs3",
		},
		{
			Addres: "cd/42342dsfs3",
		},
	}
	colors := &models.Colors{
		Base:      "#fff",
		Text:      "#dsfsdfs",
		Contrast:  "#fdsfsdcxs",
		Primary:   "#fdsfsdf",
		Secondary: "#fdsfxz",
	}

	structs := &[][]models.Block{
		{
			{
				Type:     "text",
				Location: "1",
			},
			{
				Type:     "text",
				Location: "2",
			},
			{
				Type:     "image",
				Location: "1",
			},
		},
		{
			{
				Type:     "image",
				Location: "2",
			},
			{
				Type:     "text",
				Location: "3",
			},
			{
				Type:     "text",
				Location: "4",
			},
		},
	}

	portfolio := Portfolio(id, url, createrUser, name, text, photo, colors, structs)

	p.CreatePortfolio(context.Background(), portfolio, user)

	err := p.DeletePortfolio(context.Background(), user, portfolio.ID)
	assert.NoError(t, err)

	err = p.DeletePortfolio(context.Background(), user, "5")
	assert.Error(t, err)
	assert.Equal(t, err, portfolios.ErrDeletePortfolio)
}
