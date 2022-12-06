package usecase

import (
	"Portfolio_You/models"
	"Portfolio_You/portfolios/repository/mock"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePortfolio(t *testing.T) {
	rep := new(mock.PortfolioMock)

	pf := NewPortfolioUseCase(rep)

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

	user := &models.User{
		ID:       "1",
		Username: "faner201",
		Password: "locaut",
		Email:    "polta@mail.ru",
	}

	portfolio := &models.Portfolio{
		Url:         "https://portfolio_you.com/portfolio/lopata%25faner201",
		CreaterUser: user.Username,
		Global: models.Global{
			Name: name,
			View: view,
			Bg:   bg,
		},
		Struct: models.Struct{
			StructList: structs,
		},
	}

	rep.On("CreatePortfolio", portfolio, user).Return(nil)
	err := pf.CreatePortfolio(context.Background(), user, name, view, bg, structs)
	assert.NoError(t, err)
}

func TestCreateMenuPortfoli0(t *testing.T) {
	rep := new(mock.PortfolioMock)

	pf := NewPortfolioUseCase(rep)

	user := &models.User{
		Username: "faner201",
		Password: "locaut",
		Email:    "polta@mail.ru",
	}

	name := "Lopata"
	shortText := "lol"
	photo := "cd/fdsfewtrwfsd"

	menu := &models.Menu{
		Name:        name,
		CreaterName: user.Username,
		ShortText:   shortText,
		Photo:       photo,
	}

	rep.On("CreateMenuPortfolio", user, menu).Return(nil)
	err := pf.CreateMenuPortfolio(context.Background(), user, name, shortText, photo)
	assert.NoError(t, err)
}

func TestOpenPortfolio(t *testing.T) {
	rep := new(mock.PortfolioMock)

	pf := NewPortfolioUseCase(rep)

	userName := "faner201"
	email := "polta@mail.ru"
	password := "locaut"

	portfolioID := "1"

	user := &models.User{
		Username: userName,
		Email:    email,
		Password: password,
	}

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

	portfolio := &models.Portfolio{
		ID:          portfolioID,
		Url:         "https://portfolio_you.com/portfolio/lopata%25faner201",
		CreaterUser: user.Username,
		Global: models.Global{
			Name: name,
			View: view,
			Bg:   bg,
		},
		Struct: models.Struct{
			StructList: structs,
		},
	}

	rep.On("GetPortfolioByUserName", userName, portfolioID).Return(portfolio, nil)
	portf, err := pf.OpenPortfolio(context.Background(), user, portfolioID)
	assert.NoError(t, err)
	assert.NotEmpty(t, portf)
}

func TestGetListPortfolio(t *testing.T) {
	rep := new(mock.PortfolioMock)

	pf := NewPortfolioUseCase(rep)

	userName := "faner201"
	email := "polta@mail.ru"
	password := "locaut"

	user := &models.User{
		Username: userName,
		Email:    email,
		Password: password,
	}

	name := "Lopata"
	shortText := "lol"
	photo := "cd/fdsfewtrwfsd"

	menuList := []*models.Menu{
		&models.Menu{
			Name:      name,
			ShortText: shortText,
			Photo:     photo,
		},
		&models.Menu{
			Name:      name,
			ShortText: shortText,
			Photo:     photo,
		},
		&models.Menu{
			Name:      name,
			ShortText: shortText,
			Photo:     photo,
		},
	}
	rep.On("GetListPortfolioByUserName", userName).Return(menuList, nil)
	list, err := pf.GetListPorfolio(context.Background(), user)
	assert.NoError(t, err)
	assert.NotEmpty(t, list)
}

func TestDeletePortfolio(t *testing.T) {
	rep := new(mock.PortfolioMock)

	pf := NewPortfolioUseCase(rep)

	userName := "faner201"
	email := "polta@mail.ru"
	password := "locaut"

	portfolioID := "1"

	user := &models.User{
		Username: userName,
		Email:    email,
		Password: password,
	}

	rep.On("DeletePortfolio", user, portfolioID).Return(nil)
	err := pf.DeletePortfolio(context.Background(), user, portfolioID)
	assert.NoError(t, err)
}
