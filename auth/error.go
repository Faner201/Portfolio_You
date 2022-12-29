package auth

import "errors"

var (
	ErrUserNotFound           = errors.New("User not found")
	ErrCreateUser             = errors.New("User was not created")
	ErrInvalidAccessToken     = errors.New("Session time has expired, please re-enter")
	ErrInvalidLeanguage       = errors.New("Please enter the data in English")
	ErrLenUsername            = errors.New("Your username is long, please make up another one")
	ErrSpecialSymbolsUsername = errors.New("Please remove the special characters from the username field")
)
