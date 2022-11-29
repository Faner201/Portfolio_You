package auth

import "errors"

var ErrUserNotFound = errors.New("User not found")
var ErrCreateUser = errors.New("User was not created")
