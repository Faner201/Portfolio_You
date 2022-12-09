package auth

import (
	"Portfolio_You/models"
	"context"
)

const CtxUserKey = "user"

type UseCase interface {
	SignUp(ctx context.Context, username, password, email string) error
	SignIn(ctx context.Context, username, password string) (string, error)
	ParseTokenJWT(ctx context.Context, accessToken string) (*models.User, error)
}
