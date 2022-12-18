package models

type User struct {
	ID       string
	Username string `bson:"username"`
	Password string `bson:"password"`
	Email    string `bson:"email"`
}
