package models

type Portfolio struct {
	ID          string     `json:"id"`
	Url         string     `json:"url" bson:"url"`
	CreaterUser string     `json:"createrUser" bson:"createrUser"`
	Name        string     `json:"name" bson:"name"`
	Text        *[]Text    `json:"texts" bson:"texts"`
	Photo       *[]Photo   `json:"images" bson:"images"`
	Colors      *Colors    `json:"colors" bson:"colors"`
	Struct      *[][]Block `json:"structure" bson:"structure"`
}
