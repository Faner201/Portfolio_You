package database

import (
	"Portfolio_You/models"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID primitive.ObjectID `bson:"_id, omitempy"`
}

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	return nil
}

func (u UserRepository) GetUser(ctx context.Context, username, password string) (*models.User, error) {
	return &models.User{
		Username: username,
		Password: password,
	}, nil
}

// type UserRepository struct {
// 	db *mongo.Collection
// }

// func NewUserRepository(db *mongo.Database, collection string) *UserRepository {
// 	return &UserRepository{
// 		db: db.Collection(collection),
// 	}
// }

// func (u UserRepository) CreateUser(ctx context.Context, user *models.User) error {
// 	res, err := u.db.InsertOne(ctx, user)
// 	if err != nil {
// 		return err
// 	}

// 	user.ID = res.InsertedID.(primitive.ObjectID).Hex()
// 	return nil
// }

// func (u UserRepository) GetUser(ctx context.Context, username, password string) (*models.User, error) {
// 	user := new(User)
// 	modelUser := &models.User{}
// 	err := u.db.FindOne(ctx, bson.M{
// 		"username": username,
// 		"password": password,
// 	}).Decode(modelUser)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &models.User{
// 		ID:       user.ID.Hex(),
// 		Username: modelUser.Username,
// 		Password: modelUser.Password,
// 	}, nil
// }
