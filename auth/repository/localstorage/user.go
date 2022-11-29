package localstorage

import (
	"Portfolio_You/auth"
	"Portfolio_You/models"
	"context"
	"sync"
)

type userLocalStorage struct {
	users map[int]*models.User
	mutex *sync.Mutex
}

func NewUserLocalStorage() *userLocalStorage {
	return &userLocalStorage{
		users: make(map[int]*models.User),
		mutex: new(sync.Mutex),
	}
}

func (u *userLocalStorage) CreateUser(ctx context.Context, user *models.User) error {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	if user.Email != "" && user.Username != "" && user.Password != "" {
		u.users[user.ID] = user
		return nil
	}

	return auth.ErrCreateUser
}

func (u *userLocalStorage) GetUser(ctx context.Context, username, password string) (*models.User, error) {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	for _, user := range u.users {
		if user.Username == username && user.Password == password {
			return user, nil
		}
	}

	return nil, auth.ErrUserNotFound
}
