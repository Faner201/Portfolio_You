package http

import (
	"Portfolio_You/portfolios"

	"github.com/gin-gonic/gin"
)

func RegisterHttpEndpoints(router *gin.RouterGroup, uc portfolios.UseCase) {
	h := NewHandler(uc)

	portfolio := router.Group("/portfolio")
	{
		portfolio.POST("/create", h.CreatePortfolio)
		portfolio.POST("/create", h.CreateMenuPortfolio)
		portfolio.GET("/open", h.GetPortfolio)
		portfolio.GET("/menu", h.GetListMenu)
		portfolio.DELETE("", h.DeletePortfolio)
	}
}
