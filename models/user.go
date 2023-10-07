package models

type User struct {
	ID       string
	Username string `bson:"Username"`
	Password string `bson:"Password"`
	Email    string `bson:"Email"`
}
