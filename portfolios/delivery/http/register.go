package http

import (
	"Portfolio_You/portfolios"

	"github.com/gin-gonic/gin"
)

func RegisterHttpEndpoints(router *gin.RouterGroup, uc portfolios.UseCase) {
	h := NewHandler(uc)

	portfolioEndPoints := router.Group("/portfolio")
	{
		portfolioEndPoints.POST("/create", h.CreatePortfolio)
		portfolioEndPoints.POST("/create/menu", h.CreateMenuPortfolio)
		portfolioEndPoints.POST("/open", h.GetPortfolio)
		portfolioEndPoints.GET("/menu", h.GetListMenu)
		portfolioEndPoints.DELETE("", h.DeletePortfolio)
	}
}
