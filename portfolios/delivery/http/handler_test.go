package http

import (
	"Portfolio_You/models"
	"Portfolio_You/portfolios/faker"
	"Portfolio_You/portfolios/usecase"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestCreatePortfolio(t *testing.T) {
	user, _ := faker.GetUser()

	r := gin.Default()
	group := r.Group("/api", func(c *gin.Context) {
		c.Set(viper.GetString("privileges.user"), user)
	})

	uc := new(usecase.PortfolioUseCaseMock)

	RegisterHttpEndpoints(group, uc)

	portf, _ := faker.GetPortfolio()
	portf.ID = ""
	portf.Url = ""
	portf.CreaterUser = ""

	inp := &portfolioDTO{
		Name:   portf.Name,
		Text:   portf.Texts,
		Images: portf.Images,
		Colors: portf.Colors,
		Struct: portf.Struct,
	}

	body, err := json.Marshal(inp)
	assert.NoError(t, err)

	log.Println(inp)

	uc.On("CreatePortfolio", user, portf).Return(nil)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/api/portfolio/create", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestCreateMenuPortfolio(t *testing.T) {
	user, _ := faker.GetUser()

	r := gin.Default()
	group := r.Group("/api", func(c *gin.Context) {
		c.Set(viper.GetString("privileges.user"), user)
	})

	uc := new(usecase.PortfolioUseCaseMock)

	RegisterHttpEndpoints(group, uc)

	inp, _ := faker.GetMenu()
	inp.ID = ""

	body, err := json.Marshal(inp)
	assert.NoError(t, err)

	uc.On("CreateMenuPortfolio", user, inp).Return(nil)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/api/portfolio/create/menu", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetPortfolio(t *testing.T) {

	user, _ := faker.GetUser()

	r := gin.Default()
	group := r.Group("/api", func(c *gin.Context) {
		c.Set(viper.GetString("privileges.user"), user)
	})

	uc := new(usecase.PortfolioUseCaseMock)

	RegisterHttpEndpoints(group, uc)

	inp, _ := faker.GetPortfolio()

	portfId := &portfoliofID{
		ID: inp.ID,
	}

	body, err := json.Marshal(portfId)
	assert.NoError(t, err)

	uc.On("OpenPortfolio", user, inp.ID).Return(inp, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/portfolio/open", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	expectedOut := &portfolio{
		Portfolio: &portfolioDTO{
			CreaterUser: inp.CreaterUser,
			Name:        inp.Name,
			Text:        inp.Texts,
			Images:      inp.Images,
			Colors:      inp.Colors,
			Struct:      inp.Struct,
		}}
	expectedOutBody, err := json.Marshal(expectedOut)
	assert.NoError(t, err)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(expectedOutBody), w.Body.String())
}

func TestGetListMenu(t *testing.T) {
	user, _ := faker.GetUser()

	r := gin.Default()
	group := r.Group("/api", func(c *gin.Context) {
		c.Set(viper.GetString("privileges.user"), user)
	})

	uc := new(usecase.PortfolioUseCaseMock)

	RegisterHttpEndpoints(group, uc)

	list := []models.Menu{}

	for i := 0; i < 5; i++ {
		menu, _ := faker.GetMenu()
		list = append(list, *menu)
	}

	uc.On("GetListPorfolio", user).Return(&list, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/portfolio/menu", nil)
	r.ServeHTTP(w, req)

	menuDTO := []portfolioMenuDTO{}

	for _, menu := range list {
		input := portfolioMenuDTO{
			Name:        menu.Name,
			CreaterName: menu.CreaterName,
			ShortText:   menu.ShortText,
			Image:       menu.Image,
		}
		menuDTO = append(menuDTO, input)
	}

	expectedOut := menu{&menuDTO}

	expectedOutBody, err := json.Marshal(expectedOut)
	assert.NoError(t, err)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(expectedOutBody), w.Body.String())
}

func TestDeletePortfolio(t *testing.T) {
	user, _ := faker.GetUser()

	r := gin.Default()
	group := r.Group("/api", func(c *gin.Context) {
		c.Set(viper.GetString("privileges.user"), user)
	})

	uc := new(usecase.PortfolioUseCaseMock)

	RegisterHttpEndpoints(group, uc)

	portfolio, _ := faker.GetPortfolio()

	inp := &portfoliofID{
		ID: portfolio.ID,
	}

	body, err := json.Marshal(inp)
	assert.NoError(t, err)

	uc.On("DeletePortfolio", user, inp.ID).Return(nil)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("DELETE", "/api/portfolio/menu", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

}
