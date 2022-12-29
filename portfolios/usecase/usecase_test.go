package usecase

import (
	"Portfolio_You/models"
	"Portfolio_You/portfolios/repository/mock"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func User(name, email, password string) *models.User {
	return &models.User{
		Username: name,
		Password: password,
		Email:    email,
	}
}

func Portfolio(url, createrUser, name string, text *[]models.Text, photo *[]models.Photo, colors *models.Colors,
	structs *[][]models.Block) *models.Portfolio {
	return &models.Portfolio{
		Url:         url,
		CreaterUser: createrUser,
		Name:        name,
		Text:        text,
		Photo:       photo,
		Colors:      colors,
		Struct:      structs,
	}
}

func Menu(name, createrName, shortText, photo string) *models.Menu {
	return &models.Menu{
		Name:        name,
		CreaterName: createrName,
		ShortText:   shortText,
		Photo:       photo,
	}
}

func TestCreatePortfolio(t *testing.T) {
	rep := new(mock.PortfolioMock)

	pf := NewPortfolioUseCase(rep)

	username := "faner201"
	password := "locaut"
	email := "polta@mail.ru"

	user := User(username, email, password)

	url := "/portfolio/aboba%faner201"
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

	portfolio := Portfolio(url, user.Username, name, text, photo, colors, structs)

	rep.On("CreatePortfolio", portfolio, user).Return(nil)
	err := pf.CreatePortfolio(context.Background(), user, name, text, photo, colors, structs)
	assert.NoError(t, err)
}

func TestCreateMenuPortfoli0(t *testing.T) {
	rep := new(mock.PortfolioMock)

	pf := NewPortfolioUseCase(rep)

	username := "faner201"
	password := "locaut"
	email := "polta@mail.ru"

	name := "Lopata"
	shortText := "lol"
	photo := "cd/fdsfewtrwfsd"

	user := User(username, email, password)
	menu := Menu(name, user.Username, shortText, photo)
	rep.On("CreateMenuPortfolio", user, menu).Return(nil)
	err := pf.CreateMenuPortfolio(context.Background(), user, menu.Name, menu.ShortText, menu.Photo)
	assert.NoError(t, err)
}

func TestOpenPortfolio(t *testing.T) {
	rep := new(mock.PortfolioMock)

	pf := NewPortfolioUseCase(rep)

	username := "faner201"
	password := "locaut"
	email := "polta@mail.ru"

	user := User(username, email, password)

	portfolioID := "1"
	url := "/portfolio/aboba%faner201"
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

	portfolio := Portfolio(url, user.Username, name, text, photo, colors, structs)

	rep.On("GetPortfolioByUserName", username, portfolioID).Return(portfolio, nil)
	portf, err := pf.OpenPortfolio(context.Background(), user, portfolioID)
	assert.NoError(t, err)
	assert.NotEmpty(t, portf)
}

func TestGetListPortfolio(t *testing.T) {
	rep := new(mock.PortfolioMock)

	pf := NewPortfolioUseCase(rep)

	username := "faner201"
	email := "polta@mail.ru"
	password := "locaut"

	name := "Lopata"
	shortText := "lol"
	photo := "cd/fdsfewtrwfsd"

	user := User(username, email, password)

	menu := Menu(name, username, shortText, photo)
	menuList := []models.Menu{}
	menuList = append(menuList, *menu)

	rep.On("GetListPortfolioByUserName", username).Return(&menuList, nil)
	list, err := pf.GetListPorfolio(context.Background(), user)
	assert.NoError(t, err)
	assert.NotEmpty(t, list)
}

func TestDeletePortfolio(t *testing.T) {
	rep := new(mock.PortfolioMock)

	pf := NewPortfolioUseCase(rep)

	username := "faner201"
	email := "polta@mail.ru"
	password := "locaut"

	portfolioID := "1"

	user := User(username, email, password)

	rep.On("DeletePortfolio", user, portfolioID).Return(nil)
	err := pf.DeletePortfolio(context.Background(), user, portfolioID)
	assert.NoError(t, err)
}
