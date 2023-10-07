package models

type Text struct {
	Sludge string `json:"Text" bson:"Text"`
	Style  string `json:"Style" bson:"Style"`
	Size   string `json:"Size" bson:"Size"`
}
