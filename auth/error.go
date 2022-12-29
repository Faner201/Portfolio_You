package auth

import "errors"

var ErrUserNotFound = errors.New("User not found")
var ErrCreateUser = errors.New("User was not created")
var ErrInvalidAccessToken = errors.New("Session time has expired, please re-enter")
