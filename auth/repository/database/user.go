package database

import (
	"Portfolio_You/models"
	"context"
)

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
