package models

type Portfolio struct {
	ID          string     `json:"Id"`
	Url         string     `json:"Url" bson:"Url"`
	CreaterUser string     `json:"CreaterUser" bson:"CreaterUser"`
	Name        string     `json:"Name" bson:"Name"`
	Texts       *[]Text    `json:"Texts" bson:"Texts"`
	Images      *[]Image   `json:"Images" bson:"Images"`
	Colors      *Colors    `json:"Colors" bson:"Colors"`
	Struct      *[][]Block `json:"Structures" bson:"Structures"`
}
