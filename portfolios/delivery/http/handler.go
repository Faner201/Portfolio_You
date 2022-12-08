package http

import (
	"Portfolio_You/auth"
	"Portfolio_You/models"
	"Portfolio_You/portfolios"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Portfolio struct {
	ID          string        `json:"id"`
	URL         string        `json:"url"`
	CreaterUser string        `json:"createrUser"`
	Global      models.Global `json:"Global"`
	Struct      models.Struct `json:"Struct"`
}

type Menu struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	CreaterName string `json:"createrName"`
	ShortText   string `json:"shortText"`
	Photo       string `json:"photo"`
}

type Handler struct {
	useCase portfolios.UseCase
}

func NewHandler(useCase portfolios.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

type createInputPortf struct {
	URL         string        `json:"url"`
	CreaterUser string        `json:"createrUser"`
	Global      models.Global `json:"Global"`
	Struct      models.Struct `json:"Struct"`
}

type createInputMenu struct {
	Name        string `json:"name"`
	CreaterName string `json:"createrName"`
	ShortText   string `json:"shortText"`
	Photo       string `json:"photo"`
}

type getPortfID struct {
	ID string `json:"id"`
}

type getPortfolio struct {
	Portf *Portfolio `json:"portfolio"`
}

type getMenu struct {
	Menu []*Menu `json:"menu"`
}

func (h *Handler) CreatePortfolio(c *gin.Context) {
	input := new(createInputPortf)

	if err := c.BindJSON(input); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := c.MustGet(auth.CtxUserKey).(*models.User)

	if err := h.useCase.CreatePortfolio(c.Request.Context(), user, input.Global.Name,
		input.Global.View, input.Global.Bg, input.Struct.StructList); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) CreateMenuPortfolio(c *gin.Context) {
	input := new(createInputMenu)

	if err := c.BindJSON(input); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := c.MustGet(auth.CtxUserKey).(*models.User)

	if err := h.useCase.CreateMenuPortfolio(c.Request.Context(), user, input.Name,
		input.ShortText, input.Photo); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) GetPortfolio(c *gin.Context) {
	input := new(getPortfID)

	user := c.MustGet(auth.CtxUserKey).(*models.User)

	portf, err := h.useCase.OpenPortfolio(c.Request.Context(), user, input.ID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, &getPortfolio{
		Portf: toPortfolio(portf),
	})
}

func (h *Handler) GetListMenu(c *gin.Context) {
	user := c.MustGet(auth.CtxUserKey).(*models.User)

	list, err := h.useCase.GetListPorfolio(c.Request.Context(), user)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, &getMenu{
		Menu: toListMenu(list),
	})
}

func toListMenu(ls []*models.Menu) []*Menu {
	out := make([]*Menu, len(ls))

	for i, b := range ls {
		out[i] = toMenu(b)
	}
	return out
}

func toMenu(m *models.Menu) *Menu {
	return &Menu{
		ID:          m.ID,
		CreaterName: m.CreaterName,
		Name:        m.Name,
		ShortText:   m.ShortText,
		Photo:       m.Photo,
	}
}

func toPortfolio(p *models.Portfolio) *Portfolio {
	return &Portfolio{
		ID:          p.ID,
		URL:         p.Url,
		CreaterUser: p.CreaterUser,
		Global:      p.Global,
		Struct:      p.Struct,
	}
}
