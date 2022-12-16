package http

import (
	"Portfolio_You/auth"
	"Portfolio_You/models"
	"Portfolio_You/portfolios/usecase"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreatePortfolio(t *testing.T) {
	user := &models.User{
		Username: "faner201",
		Password: "lopata",
		Email:    "polta@mail.ru",
	}

	r := gin.Default()
	group := r.Group("/api", func(c *gin.Context) {
		c.Set(auth.CtxUserKey, user)
	})

	uc := new(usecase.PortfolioUseCaseMock)

	RegisterHttpEndpoints(group, uc)

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

	inp := &createInputPortf{
		CreaterUser: user.Username,
		Name:        name,
		Text:        text,
		Photo:       photo,
		Colors:      colors,
		Struct:      structs,
	}

	body, err := json.Marshal(inp)
	assert.NoError(t, err)

	uc.On("CreatePortfolio", user, inp.Name, inp.Text, inp.Photo, inp.Colors, inp.Struct).Return(nil)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/api/portfolio/create", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestCreateMenuPortfolio(t *testing.T) {
	user := &models.User{
		Username: "faner201",
		Password: "lopata",
		Email:    "polta@mail.ru",
	}

	r := gin.Default()
	group := r.Group("/api", func(c *gin.Context) {
		c.Set(auth.CtxUserKey, user)
	})

	uc := new(usecase.PortfolioUseCaseMock)

	RegisterHttpEndpoints(group, uc)

	name := "aboba"
	createrName := user.Username
	shortText := "lopol"
	photo := "cd/34242e3"

	inp := &createInputMenu{
		Name:        name,
		CreaterName: createrName,
		ShortText:   shortText,
		Photo:       photo,
	}

	body, err := json.Marshal(inp)
	assert.NoError(t, err)

	uc.On("CreateMenuPortfolio", user, inp.Name, inp.ShortText, inp.Photo).Return(nil)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/api/portfolio/create/menu", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetPortfolio(t *testing.T) {

	user := &models.User{
		Username: "faner201",
		Password: "lopata",
		Email:    "polta@mail.ru",
	}

	r := gin.Default()
	group := r.Group("/api", func(c *gin.Context) {
		c.Set(auth.CtxUserKey, user)
	})

	uc := new(usecase.PortfolioUseCaseMock)

	RegisterHttpEndpoints(group, uc)

	id := "1"
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

	inp := &models.Portfolio{
		ID:          id,
		CreaterUser: user.Username,
		Name:        name,
		Text:        text,
		Photo:       photo,
		Colors:      colors,
		Struct:      structs,
	}

	portfId := &getPortfID{
		ID: id,
	}

	body, err := json.Marshal(portfId)
	assert.NoError(t, err)

	uc.On("OpenPortfolio", user, id).Return(inp, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/portfolio/open", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	expectedOut := &getPortfolio{Portf: inp}
	expectedOutBody, err := json.Marshal(expectedOut)
	assert.NoError(t, err)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(expectedOutBody), w.Body.String())
}

func TestGetListMenu(t *testing.T) {
	user := &models.User{
		Username: "faner201",
		Password: "lopata",
		Email:    "polta@mail.ru",
	}

	r := gin.Default()
	group := r.Group("/api", func(c *gin.Context) {
		c.Set(auth.CtxUserKey, user)
	})

	uc := new(usecase.PortfolioUseCaseMock)

	RegisterHttpEndpoints(group, uc)

	list := make([]*models.Menu, 5)

	for i := 0; i < 5; i++ {
		list[i] = &models.Menu{
			ID:          "1",
			Name:        "aboba",
			CreaterName: user.Username,
			ShortText:   "lopata",
			Photo:       "cd/fsdfsdfs",
		}
	}

	uc.On("GetListPorfolio", user).Return(list, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/portfolio/menu", nil)
	r.ServeHTTP(w, req)

	expectedOut := &getMenu{Menu: list}

	expectedOutBody, err := json.Marshal(expectedOut)
	assert.NoError(t, err)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(expectedOutBody), w.Body.String())
}

func TestDeletePortfolio(t *testing.T) {
	user := &models.User{
		Username: "faner201",
		Password: "lopata",
		Email:    "polta@mail.ru",
	}

	r := gin.Default()
	group := r.Group("/api", func(c *gin.Context) {
		c.Set(auth.CtxUserKey, user)
	})

	uc := new(usecase.PortfolioUseCaseMock)

	RegisterHttpEndpoints(group, uc)

	inp := &getPortfID{
		ID: "1",
	}

	body, err := json.Marshal(inp)
	assert.NoError(t, err)

	uc.On("DeletePortfolio", user, inp.ID).Return(nil)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("DELETE", "/api/portfolio", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

}
