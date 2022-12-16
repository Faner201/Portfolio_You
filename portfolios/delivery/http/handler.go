package http

import (
	"Portfolio_You/auth"
	"Portfolio_You/models"
	"Portfolio_You/portfolios"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase portfolios.UseCase
}

func NewHandler(useCase portfolios.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

type createInputPortf struct {
	CreaterUser string            `json:"createrUser"`
	Name        string            `json:"name"`
	Text        *[]models.Text    `json:"text"`
	Photo       *[]models.Photo   `json:"photo"`
	Colors      *models.Colors    `json:"colors"`
	Struct      *[][]models.Block `json:"blocks"`
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
	Portf *models.Portfolio `json:"portfolio"`
}

type getMenu struct {
	Menu []*models.Menu `json:"menu"`
}

func (h *Handler) CreatePortfolio(c *gin.Context) {
	input := new(createInputPortf)

	if err := c.BindJSON(input); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := c.MustGet(auth.CtxUserKey).(*models.User)

	if err := h.useCase.CreatePortfolio(c.Request.Context(), user, input.Name, input.Text, input.Photo, input.Colors, input.Struct); err != nil {
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

	if err := c.BindJSON(input); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// user := c.MustGet(auth.CtxUserKey).(*models.User)

	user := &models.User{
		Username: "faner201",
		Password: "lopata",
	}

	portf, err := h.useCase.OpenPortfolio(c.Request.Context(), user, input.ID) // местная заглушка для решения проблемы получения id с фронта
	portf = &models.Portfolio{
		Url:         "/portfolios/aboba&faner201",
		CreaterUser: user.Username,
		Name:        "aboba",
		Text: &[]models.Text{
			{
				Sludge: "Хотел бы сказать этому артёму",
				Style:  "bold",
				Size:   "big",
			},
			{
				Sludge: "Был бы ты человеком, а не дотером",
				Style:  "italic",
				Size:   "small",
			},
		},
		Photo: &[]models.Photo{
			{
				Addres: "",
			},
			{
				Addres: "",
			},
		},
		Colors: &models.Colors{
			Base:      "#4f634b",
			Text:      "#79a5b3",
			Contrast:  "#794e8a",
			Primary:   "#7d8f49",
			Secondary: "#d3f76a",
		},
		Struct: &[][]models.Block{
			{
				{
					Type:     "text",
					Location: "1",
				},
				{
					Type:     "text",
					Location: "1",
				},
				{
					Type:     "image",
					Location: "0",
				},
			},
			{
				{
					Type:     "image",
					Location: "1",
				},
				{
					Type:     "text",
					Location: "0",
				},
				{
					Type:     "text",
					Location: "1",
				},
			},
		},
	} // ужасная заглушка для просмотра отображения инфы с бэка на фронт
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, &getPortfolio{
		Portf: portf,
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
		Menu: list,
	})
}

func (h *Handler) DeletePortfolio(c *gin.Context) {
	input := new(getPortfID)

	if err := c.BindJSON(input); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := c.MustGet(auth.CtxUserKey).(*models.User)

	if err := h.useCase.DeletePortfolio(c.Request.Context(), user, input.ID); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
