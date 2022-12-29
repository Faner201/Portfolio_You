package http

import (
	"Portfolio_You/auth"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase auth.UseCase
}

func NewHandler(useCase auth.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

type SignUp struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type SignIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type signInputToken struct {
	Token string `json:"token"`
}

func (h *Handler) SignUp(c *gin.Context) {
	input := new(SignUp)

	log.Println(input)

	if err := c.BindJSON(input); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.useCase.SignUp(c.Request.Context(), input.Username, input.Password, input.Email); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) SignIn(c *gin.Context) {
	input := new(SignIn)

	if err := c.BindJSON(input); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	log.Println()

	token, err := h.useCase.SignIn(c.Request.Context(), input.Username, input.Password)
	if err != nil {
		if err == auth.ErrUserNotFound || err == auth.ErrInvalidAccessToken {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, signInputToken{Token: token})
}
