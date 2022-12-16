package models

type Menu struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	CreaterName string `json:"createrName"`
	ShortText   string `json:"shortText"`
	Photo       string `json:"photo"`
}
