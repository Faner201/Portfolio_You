package faker

import (
	"Portfolio_You/models"

	"github.com/go-faker/faker/v4"
)

type PortfolioFaker struct {
	ID          string
	Url         string
	CreaterUser string
	Name        string
	Texts       *[]models.Text
	Images      *[]models.Image
	Colors      *models.Colors
	Struct      *[][]models.Block
}

type MenuFaker struct {
	ID          string
	Name        string
	CreaterName string
	ShortText   string
	Image       string
}

type UserFaker struct {
	Username string
	Password string
	Email    string
}

func generatorPortfolio() (PortfolioFaker, error) {
	portfolio := PortfolioFaker{
		ID:          faker.UUIDDigit(),
		Url:         faker.URL(),
		CreaterUser: faker.Username(),
		Name:        faker.Word(),
		Texts: &[]models.Text{
			{
				Sludge: faker.Word(),
				Style:  "limitic",
				Size:   "12",
			},
			{
				Sludge: faker.Word(),
				Style:  "italuc",
				Size:   "15",
			},
		},
		Images: &[]models.Image{
			{
				Src: "cd/" + faker.Word(),
			},
			{
				Src: "cd/" + faker.Word(),
			},
		},
		Colors: &models.Colors{
			Main:           "#FFFFFF",
			Text:           "#8884FF",
			Contrast:       "#5D576B",
			PrimaryBlock:   "#E8CEE4",
			SecondaryBlock: "#E8CEE4",
		},
		Struct: &[][]models.Block{
			{
				{
					Type:     "text",
					Position: "1",
				},
				{
					Type:     "text",
					Position: "2",
				},
				{
					Type:     "image",
					Position: "1",
				},
			},
			{
				{
					Type:     "image",
					Position: "2",
				},
				{
					Type:     "text",
					Position: "3",
				},
				{
					Type:     "text",
					Position: "4",
				},
			},
		},
	}

	err := faker.FakeData(&portfolio)
	if err != nil {
		return portfolio, err
	}

	return portfolio, nil
}

func GetPortfolio() (*models.Portfolio, error) {
	fakerDate, err := generatorPortfolio()
	if err != nil {
		return &models.Portfolio{}, err
	}

	return &models.Portfolio{
		ID:          fakerDate.ID,
		Url:         fakerDate.Url,
		CreaterUser: fakerDate.CreaterUser,
		Name:        fakerDate.Name,
		Texts:       fakerDate.Texts,
		Images:      fakerDate.Images,
		Colors:      fakerDate.Colors,
		Struct:      fakerDate.Struct,
	}, nil
}

func generateMenu() (MenuFaker, error) {
	menu := MenuFaker{
		ID:          faker.UUIDDigit(),
		Name:        faker.Word(),
		CreaterName: faker.Username(),
		ShortText:   faker.Word(),
		Image:       "cd/" + faker.Word(),
	}

	err := faker.FakeData(&menu)
	if err != nil {
		return menu, err
	}

	return menu, nil
}

func GetMenu() (*models.Menu, error) {
	fakerDate, err := generateMenu()
	if err != nil {
		return &models.Menu{}, err
	}

	return &models.Menu{
		ID:          fakerDate.ID,
		Name:        fakerDate.Name,
		CreaterName: fakerDate.CreaterName,
		ShortText:   fakerDate.ShortText,
		Image:       fakerDate.Image,
	}, nil
}

func generateUser() (UserFaker, error) {
	user := UserFaker{
		Username: faker.Username(),
		Password: faker.Password(),
		Email:    faker.Email(),
	}

	err := faker.FakeData(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func GetUser() (*models.User, error) {
	fakerDate, err := generateUser()
	if err != nil {
		return &models.User{}, err
	}

	return &models.User{
		Username: fakerDate.Username,
		Password: fakerDate.Password,
		Email:    fakerDate.Email,
	}, nil
}
