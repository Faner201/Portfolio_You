package usecase

import (
	"Portfolio_You/models"
	"context"

	"github.com/stretchr/testify/mock"
)

type AuthUseCaseMock struct {
	mock.Mock
}

func (m *AuthUseCaseMock) SignUp(ctx context.Context, username, password, email string) error {
	args := m.Called(username, password, email)

	return args.Error(0)
}

func (m *AuthUseCaseMock) ParseToketJWT(ctx context.Context, accessToken string) (*models.User, error) {
	args := m.Called(accessToken)

	return args.Get(0).(*models.User), args.Error(1)
}

func (m *AuthUseCaseMock) SignIn(ctx context.Context, username, password string) (string, error) {
	args := m.Called(username, password)

	return args.Get(0).(string), args.Error(1)
}
