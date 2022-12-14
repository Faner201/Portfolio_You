package http

import (
	"Portfolio_You/portfolios"

	"github.com/gin-gonic/gin"
)

func RegisterHttpEndpoints(router *gin.Engine, uc portfolios.UseCase) {
	h := NewHandler(uc)

	portfolioEndPoints := router.Group("/portfolio")
	{
		portfolioEndPoints.POST("/create", h.CreatePortfolio)
		// portfolioEndPoints.POST("/create", h.CreateMenuPortfolio)
		// portfolioEndPoints.GET("/open", h.GetPortfolio)
		// portfolioEndPoints.GET("/menu", h.GetListMenu)
		// portfolioEndPoints.DELETE("", h.DeletePortfolio)
	}
}
