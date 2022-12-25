package models

type Block struct {
	Type     string `json:"type" bson:"type" form:"type"`
	Location string `json:"index" bson:"index" form:"index"`
}
