package http

import (
	"Portfolio_You/models"
	"Portfolio_You/portfolios"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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
	CreaterUser string            `json:"createrUser" form:"createrUser"`
	Name        string            `json:"name" form:"name"`
	Text        *[]models.Text    `json:"texts" form:"texts"`
	Photo       *[]models.Photo   `json:"images" form:"images"`
	Colors      *models.Colors    `json:"colors" form:"colors"`
	Struct      *[][]models.Block `json:"structure" form:"structure"`
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
	Menu *[]models.Menu `json:"menu"`
}

func savePicture(c *gin.Context) *[]models.Photo {
	form, _ := c.MultipartForm()
	photos := form.File["photo"]

	model := new(models.Photo)
	list := []models.Photo{}

	for _, photo := range photos {
		c.SaveUploadedFile(photo, photo.Filename)
		model.Addres = photo.Filename
		list = append(list, *model)
	}

	return &list
}

func (h *Handler) CreatePortfolio(c *gin.Context) {
	input := new(createInputPortf)

	// if err := c.ShouldBind(input); err != nil {
	// 	c.AbortWithStatus(http.StatusBadRequest)
	// 	return
	// }

	// log.Println(input)

	if err := c.BindJSON(input); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// input.Photo = savePicture(c)

	user := c.MustGet(viper.GetString("privileges.user")).(*models.User)

	// user := &models.User{
	// 	Username: "faner201",
	// 	Password: "lopata",
	// }

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

	user := c.MustGet(viper.GetString("privileges.user")).(*models.User)

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

	log.Println(input)

	user := c.MustGet(viper.GetString("privileges.user")).(*models.User)

	// user := &models.User{
	// 	Username: "faner201",
	// 	Password: "lopata",
	// }

	portf, err := h.useCase.OpenPortfolio(c.Request.Context(), user, input.ID) // местная заглушка для решения проблемы получения id с фронта
	// portf = &models.Portfolio{
	// 	Url:         "/portfolios/aboba&faner201",
	// 	CreaterUser: user.Username,
	// 	Name:        "aboba",
	// 	Text: &[]models.Text{
	// 		{
	// 			Sludge: "Хотел бы сказать этому артёму",
	// 			Style:  "bold",
	// 			Size:   "big",
	// 		},
	// 		{
	// 			Sludge: "Был бы ты человеком, а не дотером",
	// 			Style:  "italic",
	// 			Size:   "small",
	// 		},
	// 	},
	// 	Photo: &[]models.Photo{
	// 		{
	// 			Addres: "/Users/fanfurick/Documents/Profile_You/src/photo.jpeg",
	// 		},
	// 		{
	// 			Addres: "/Users/fanfurick/Documents/Profile_You/src/photo.jpeg",
	// 		},
	// 	},
	// 	Colors: &models.Colors{
	// 		Base:      "#4f634b",
	// 		Text:      "#79a5b3",
	// 		Contrast:  "#794e8a",
	// 		Primary:   "#7d8f49",
	// 		Secondary: "#d3f76a",
	// 	},
	// 	Struct: &[][]models.Block{
	// 		{
	// 			{
	// 				Type:     "text",
	// 				Location: "1",
	// 			},
	// 			{
	// 				Type:     "text",
	// 				Location: "1",
	// 			},
	// 		},
	// 		{
	// 			{
	// 				Type:     "image",
	// 				Location: "1",
	// 			},
	// 			{
	// 				Type:     "text",
	// 				Location: "0",
	// 			},
	// 			{
	// 				Type:     "text",
	// 				Location: "1",
	// 			},
	// 		},
	// 	},
	// } // ужасная заглушка для просмотра отображения инфы с бэка на фронт
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, &getPortfolio{
		Portf: portf,
	})
}

func (h *Handler) GetListMenu(c *gin.Context) {
	user := c.MustGet(viper.GetString("privileges.user")).(*models.User)

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

	user := c.MustGet(viper.GetString("privileges.user")).(*models.User)

	if err := h.useCase.DeletePortfolio(c.Request.Context(), user, input.ID); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
