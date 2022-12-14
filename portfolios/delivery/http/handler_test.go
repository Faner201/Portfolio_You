package http

import (
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
	}

	r := gin.Default()

	uc := new(usecase.PortfolioUseCaseMock)

	RegisterHttpEndpoints(r, uc)

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

	inp := &createInputPortf{
		Global: models.Global{
			Name: "lopata",
			View: "very good, polka",
			Bg:   "very interesting",
		},
		Struct: models.Struct{
			StructList: structs,
		},
	}

	body, err := json.Marshal(inp)
	assert.NoError(t, err)

	uc.On("CreatePortolio", user, inp.Global.Name, inp.Global.View, inp.Global.View, inp.Struct.StructList).Return(nil)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "portfolio/create", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
