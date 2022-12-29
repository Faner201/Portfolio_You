package http

import (
	"Portfolio_You/auth"

	"github.com/gin-gonic/gin"
)

func RegisterHttpEndpoints(router *gin.Engine, uc auth.UseCase) {
	h := NewHandler(uc)

	authEndPoints := router.Group("")
	{
		authEndPoints.POST("/sign-up", h.SignUp)
		authEndPoints.POST("/sign-in", h.SignIn)
	}
}
