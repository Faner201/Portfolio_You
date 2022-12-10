package http

import (
	"Portfolio_You/auth/usecase"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSignIn(t *testing.T) {
	r := gin.Default()
	uc := new(usecase.AuthUseCaseMock)

	RegisterHttpEndpoints(r, uc)

	signIn := &SignIn{
		Username: "faner201",
		Password: "polka",
	}

	body, err := json.Marshal(signIn)
	assert.NoError(t, err)

	uc.On("SignIn", signIn.Username, signIn.Password).Return("jwt", nil)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/auth/sign-in", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"token\":\"jwt\"}", w.Body.String())
}

func TestSignUp(t *testing.T) {
	r := gin.Default()
	uc := new(usecase.AuthUseCaseMock)

	RegisterHttpEndpoints(r, uc)

	signUp := &SignUp{
		Username: "faner201",
		Password: "polka",
		Email:    "polka@gmail.com",
	}

	body, err := json.Marshal(signUp)
	assert.NoError(t, err)

	uc.On("SignUp", signUp.Username, signUp.Password, signUp.Email).Return(nil)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/auth/sign-up", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
