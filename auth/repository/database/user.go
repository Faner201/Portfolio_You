package database

import (
	"Portfolio_You/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/mongo/driver"
)

type UserRepository struct {
	db *mongo.Collection
}

func NewUserRepository(db *mongo.Database, collection string) *UserRepository {
	return &UserRepository{
		db: db.Collection(collection),
	}
}

func (u UserRepository) CreateUser(ctx context.Context, user *models.User) error {

	valid := u.db.FindOne(ctx, bson.M{
		"username": user.Username,
		"email":    user.Email,
	}).Err()

	if valid == nil {
		return driver.ErrUnacknowledgedWrite
	}

	res, err := u.db.InsertOne(ctx, &models.User{
		ID:       primitive.NewObjectID().Hex(),
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
	})

	if err != nil {
		return err
	}

	user.ID = res.InsertedID.(primitive.ObjectID).Hex()

	return nil
}

func (u UserRepository) GetUser(ctx context.Context, username, password string) (*models.User, error) {
	modelUser := &models.User{}

	err := u.db.FindOne(ctx, bson.M{
		"username": username,
		"password": password,
	}).Decode(modelUser)

	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:       modelUser.ID,
		Username: modelUser.Username,
		Password: modelUser.Password,
	}, nil
}
