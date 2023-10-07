package models

type Block struct {
	Type     string `json:"Type" bson:"Type" form:"Type"`
	Position string `json:"Position" bson:"Position" form:"Position"`
}
