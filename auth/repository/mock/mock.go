package mock

import (
	"Portfolio_You/models"
	"context"

	"github.com/stretchr/testify/mock"
)

type userMock struct {
	mock.Mock
}

func (m *userMock) CreateUser(ctx context.Context, user *models.User) error {
	args := m.Called(user)

	return args.Error(0)
}

func (m *userMock) GetUser(ctx context.Context, username, password string) (*models.User, error) {
	args := m.Called(username, password)

	return args.Get(0).(*models.User), args.Error(1)
}
