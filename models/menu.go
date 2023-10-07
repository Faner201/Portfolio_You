package models

type Menu struct {
	ID          string `json:"Id"`
	Name        string `json:"Name" bson:"Name"`
	CreaterName string `json:"CreaterName" bson:"CreaterName"`
	ShortText   string `json:"ShortText" bson:"ShortText"`
	Image       string `json:"Image" bson:"Image"`
}
