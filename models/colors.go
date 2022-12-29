package models

type Colors struct {
	Base      string `json:"base" bson:"base"`
	Text      string `json:"text" bson:"text"`
	Contrast  string `json:"contrast" bson:"contrast"`
	Primary   string `json:"primary" bson:"primary"`
	Secondary string `json:"secondary" bson:"secondary"`
}
