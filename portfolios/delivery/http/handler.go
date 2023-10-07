package http

import (
	"Portfolio_You/models"
	"Portfolio_You/portfolios"
	"net/http"
	"unicode"

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

type portfolioDTO struct {
	CreaterUser string            `json:"CreaterUser" form:"CreaterUser"`
	Name        string            `json:"Name" form:"Name"`
	Text        *[]models.Text    `json:"Texts" form:"Texts"`
	Images      *[]models.Image   `json:"Images" form:"Images"`
	Colors      *models.Colors    `json:"Colors" form:"Colors"`
	Struct      *[][]models.Block `json:"Structure" form:"Structure"`
}

type portfolioMenuDTO struct {
	Name        string `json:"Name"`
	CreaterName string `json:"CreaterName"`
	ShortText   string `json:"ShortText"`
	Image       string `json:"Image"`
}

type imageDTO struct {
	Image *[]models.Image `json:"Image"`
}

type portfoliofID struct {
	ID string `json:"Id"`
}

type portfolio struct {
	Portfolio *portfolioDTO `json:"Portfolio"`
}

type menu struct {
	Menu *[]portfolioMenuDTO `json:"Menu"`
}

func (h *Handler) validateDatePortfolio(p *portfolioDTO) error {

	for _, letter := range p.Name {
		if unicode.IsSymbol(letter) {
			return portfolios.ErrSpecialSymbolName
		}
	}

	if p.Text == nil {
		return portfolios.ErrNotTextPortfolio
	}

	if p.Struct == nil {
		return portfolios.ErrFullnessPortfolio
	}
	return nil
}

func (h *Handler) validateDateMenu(m *portfolioMenuDTO) error {
	for _, letter := range m.Name {
		if unicode.IsSymbol(letter) {
			return portfolios.ErrSpecialSymbolName
		}
	}
	return nil
}

func (h *Handler) savePicture(c *gin.Context) *imageDTO {
	form, err := c.MultipartForm()
	if err != nil {
		return nil
	}
	photos := form.File["photo"]

	list := []models.Image{}

	for _, photo := range photos {
		c.SaveUploadedFile(photo, photo.Filename)
		list = append(list, models.Image{
			Src: photo.Filename,
		})
	}

	return &imageDTO{
		Image: &list,
	}
}

func (h *Handler) CreatePortfolio(c *gin.Context) {
	input := new(portfolioDTO)

	if err := c.ShouldBind(input); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := c.BindJSON(input); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := h.validateDatePortfolio(input)

	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
		return
	}
	// input.Images = h.savePicture(c).Image

	user := c.MustGet(viper.GetString("privileges.user")).(*models.User)

	if err := h.useCase.CreatePortfolio(c.Request.Context(), user, &models.Portfolio{
		Name:   input.Name,
		Texts:  input.Text,
		Images: input.Images,
		Colors: input.Colors,
		Struct: input.Struct,
	}); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) CreateMenuPortfolio(c *gin.Context) {
	input := new(portfolioMenuDTO)

	if err := c.BindJSON(input); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := h.validateDateMenu(input)

	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
		return
	}

	user := c.MustGet(viper.GetString("privileges.user")).(*models.User)

	if err := h.useCase.CreateMenuPortfolio(c.Request.Context(), user, &models.Menu{
		Name:        input.Name,
		CreaterName: input.CreaterName,
		ShortText:   input.ShortText,
		Image:       input.Image,
	}); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) GetPortfolio(c *gin.Context) {
	input := new(portfoliofID)

	if err := c.BindJSON(input); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := c.MustGet(viper.GetString("privileges.user")).(*models.User)
	portf, err := h.useCase.OpenPortfolio(c.Request.Context(), user, input.ID)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, &portfolio{
		Portfolio: &portfolioDTO{
			CreaterUser: portf.CreaterUser,
			Name:        portf.Name,
			Text:        portf.Texts,
			Images:      portf.Images,
			Colors:      portf.Colors,
			Struct:      portf.Struct,
		},
	})
}

func (h *Handler) GetListMenu(c *gin.Context) {
	user := c.MustGet(viper.GetString("privileges.user")).(*models.User)

	list, err := h.useCase.GetListPorfolio(c.Request.Context(), user)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	menuList := []portfolioMenuDTO{}

	for _, menu := range *list {
		menuList = append(menuList, portfolioMenuDTO{
			Name:        menu.Name,
			CreaterName: menu.CreaterName,
			ShortText:   menu.ShortText,
			Image:       menu.Image,
		})
	}

	c.JSON(http.StatusOK, &menu{
		Menu: &menuList,
	})
}

func (h *Handler) DeletePortfolio(c *gin.Context) {
	input := new(portfoliofID)

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
