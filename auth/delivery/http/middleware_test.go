package http

import (
	"Portfolio_You/auth"
	"Portfolio_You/auth/usecase"
	"Portfolio_You/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNoAuthHeader(t *testing.T) {
	r := gin.Default()

	uc := new(usecase.AuthUseCaseMock)

	r.POST("/api/endpoint", NewAuthMiddleware(uc), func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	nr := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/api/endpoint", nil)
	r.ServeHTTP(nr, req)
	assert.Equal(t, http.StatusUnauthorized, nr.Code)
}

func TestEmptyAuthHeader(t *testing.T) {
	r := gin.Default()

	uc := new(usecase.AuthUseCaseMock)

	r.POST("/api/endpoint", NewAuthMiddleware(uc), func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	nr := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/endpoint", nil)

	req.Header.Set("Authorization", "")
	r.ServeHTTP(nr, req)
	assert.Equal(t, http.StatusUnauthorized, nr.Code)
}

func TestBearerNoToken(t *testing.T) {
	r := gin.Default()

	uc := new(usecase.AuthUseCaseMock)

	r.POST("/api/endpoint", NewAuthMiddleware(uc), func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	nr := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/endpoint", nil)

	uc.On("ParseTokenJWT", "").Return(&models.User{}, auth.ErrInvalidAccessToken)

	req.Header.Set("Authorization", "Bearer")
	r.ServeHTTP(nr, req)
	assert.Equal(t, http.StatusUnauthorized, nr.Code)
}

func TestValidAuth(t *testing.T) {
	r := gin.Default()

	uc := new(usecase.AuthUseCaseMock)

	r.POST("/api/endpoint", NewAuthMiddleware(uc), func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	nr := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/endpoint", nil)

	uc.On("ParseTokenJWT", "token").Return(&models.User{}, nil)
	req.Header.Set("Authorization", "Bearer token")
	r.ServeHTTP(nr, req)
	assert.Equal(t, http.StatusOK, nr.Code)
}
