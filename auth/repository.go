package auth

import (
	"Portfolio_You/models"
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, id int) (*models.User, error)
}
