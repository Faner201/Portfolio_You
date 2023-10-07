package models

type Colors struct {
	Main           string `json:"Main" bson:"Main"`
	Text           string `json:"Text" bson:"Text"`
	Contrast       string `json:"Contrast" bson:"Contrast"`
	PrimaryBlock   string `json:"PrimaryBlock" bson:"PrimaryBlock"`
	SecondaryBlock string `json:"SecondaryBlock" bson:"SecondaryBlock"`
}
