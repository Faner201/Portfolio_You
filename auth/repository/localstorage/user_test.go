package localstorage

import (
	"Portfolio_You/auth"
	"Portfolio_You/models"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	u := NewUserLocalStorage()

	user := &models.User{
		ID:       1,
		Username: "faner201",
		Password: "locaut",
		Email:    "polta@mail.ru",
	}

	err := u.CreateUser(context.Background(), user)
	assert.NoError(t, err)

	userNotUsername := &models.User{
		ID:       1,
		Username: "",
		Password: "locaut",
		Email:    "polta@mail.ru",
	}

	err = u.CreateUser(context.Background(), userNotUsername)
	assert.Error(t, err)
	assert.Equal(t, err, auth.ErrCreateUser)
}

func TestGetUser(t *testing.T) {
	u := NewUserLocalStorage()

	user := &models.User{
		ID:       1,
		Username: "faner201",
		Password: "locaut",
		Email:    "polta@mail.ru",
	}

	u.CreateUser(context.Background(), user)

	returnedUser, err := u.GetUser(context.Background(), "faner201", "locaut")
	assert.NoError(t, err)
	assert.Equal(t, user, returnedUser)

	returnedUser, err = u.GetUser(context.Background(), "faner201", "")
	assert.Error(t, err)
	assert.Equal(t, err, auth.ErrUserNotFound)
}
