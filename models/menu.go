package models

type Menu struct {
	ID          string `json:"id"`
	Name        string `json:"name" bson:"name"`
	CreaterName string `json:"createrName" bson:"createrName"`
	ShortText   string `json:"shortText" bson:"shortText"`
	Photo       string `json:"photo" bson:"photo"`
}
