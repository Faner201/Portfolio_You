package usecase

import (
	"Portfolio_You/auth/repository/mock"
	"Portfolio_You/models"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func User(username, password, email string) *models.User {
	return &models.User{
		Username: username,
		Password: password,
		Email:    email,
	}
}

func TestSignUp(t *testing.T) {
	rep := new(mock.UserMock)

	uc := NewAuthUseCase(rep, "lopata", []byte("secret"), 3600)

	password := "pass"
	username := "faner201"
	email := "lopata@mail.ru"
	passwordHash := "eecc7f29701ecd00f00b87e97f0748c1ec0df53d" //sha1 of pass + lopata

	user := User(username, passwordHash, email)

	rep.On("CreateUser", user).Return(nil)
	err := uc.SignUp(context.Background(), username, password, email)
	assert.NoError(t, err)
}

func TestSignIn(t *testing.T) {
	rep := new(mock.UserMock)

	uc := NewAuthUseCase(rep, "lopata", []byte("secret"), 3600)

	password := "pass"
	username := "faner201"
	email := "lopata@mail.ru"
	passwordHash := "eecc7f29701ecd00f00b87e97f0748c1ec0df53d"

	user := User(username, passwordHash, email)

	rep.On("GetUser", user.Username, user.Password).Return(user, nil)
	token, err := uc.SignIn(context.Background(), username, password)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestParseToken(t *testing.T) {
	rep := new(mock.UserMock)

	uc := NewAuthUseCase(rep, "lopata", []byte("secret"), 3600)

	password := "pass"
	username := "faner201"
	email := "lopata@mail.ru"
	passwordHash := "eecc7f29701ecd00f00b87e97f0748c1ec0df53d"

	user := User(username, passwordHash, email)

	rep.On("GetUser", user.Username, user.Password).Return(user, nil)
	token, _ := uc.SignIn(context.Background(), username, password)
	parsedUser, err := uc.ParseToketJWT(context.Background(), token)
	assert.NoError(t, err)
	assert.Equal(t, user, parsedUser)
}
