package models

type Text struct {
	Sludge string `json:"text" bson:"text"`
	Style  string `json:"style" bson:"style"`
	Size   string `json:"size" bson:"size"`
}
