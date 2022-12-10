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

func Portfolio(name, view, bg, url, createUser string, structs []interface{}) *models.Portfolio {
	return &models.Portfolio{
		Url:         url,
		CreaterUser: createUser,
		Global: models.Global{
			Name: name,
			View: view,
			Bg:   bg,
		},
		Struct: models.Struct{
			StructList: structs,
		},
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

	url := "/portfolio/lopata%faner201"
	name := "lopata"
	view := "very good, polka"
	bg := "very interesting"
	blockPhoto := "cd/fsdgsdgsd"
	blockText := "fdsfsdfds"
	structs := []interface{}{
		models.BlockPhoto{
			Photo: blockPhoto,
		},
		models.BlockText{
			Text: blockText,
		},
	}

	username := "faner201"
	password := "locaut"
	email := "polta@mail.ru"

	user := User(username, email, password)
	portfolio := Portfolio(name, view, bg, url, user.Username, structs)

	rep.On("CreatePortfolio", portfolio, user).Return(nil)
	err := pf.CreatePortfolio(context.Background(), user, name, view, bg, structs)
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
	email := "polta@mail.ru"
	password := "locaut"

	portfolioID := "1"
	url := "/portfolio/lopata%faner201"
	name := "lopata"
	view := "very good, polka"
	bg := "very interesting"
	blockPhoto := "cd/fsdgsdgsd"
	blockText := "fdsfsdfds"
	structs := []interface{}{
		models.BlockPhoto{
			Photo: blockPhoto,
		},
		models.BlockText{
			Text: blockText,
		},
	}

	user := User(username, email, password)
	portfolio := Portfolio(name, view, bg, url, user.Username, structs)

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

	menuList := []*models.Menu{
		Menu(name, user.Username, shortText, photo),
		Menu(name, user.Username, shortText, photo),
		Menu(name, user.Username, shortText, photo),
	}
	rep.On("GetListPortfolioByUserName", username).Return(menuList, nil)
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
