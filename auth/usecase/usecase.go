package usecase

import (
	"Portfolio_You/auth"
	"Portfolio_You/models"
	"context"
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

type AuthClaims struct {
	jwt.StandardClaims
	User *models.User `json:"user"`
}

type AuthUseCase struct {
	userRepo      auth.UserRepository
	hashSalt      string
	signingKey    []byte
	residenceTime time.Duration
}

func NewAuthUseCase(userRepo auth.UserRepository, hashSalt string, signingKey []byte, toketTTlSeconds time.Duration) *AuthUseCase {
	return &AuthUseCase{
		userRepo:      userRepo,
		hashSalt:      hashSalt,
		signingKey:    signingKey,
		residenceTime: time.Second * toketTTlSeconds,
	}
}

func (a *AuthUseCase) SignUp(ctx context.Context, username, password, email string) error {
	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(a.hashSalt))

	user := &models.User{
		Username: username,
		Email:    email,
		Password: fmt.Sprintf("%x", pwd.Sum(nil)),
	}

	return a.userRepo.CreateUser(ctx, user)
}

func (a *AuthUseCase) ParseTokenJWT(ctx context.Context, accessToken string) (*models.User, error) {
	token, err := jwt.ParseWithClaims(accessToken, &AuthClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return a.signingKey, nil
		},
	)

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.User, nil
	}

	return nil, auth.ErrInvalidAccessToken
}

func (a *AuthUseCase) SignIn(ctx context.Context, username, password string) (string, error) {
	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(a.hashSalt))
	password = fmt.Sprintf("%x", pwd.Sum(nil))

	user, err := a.userRepo.GetUser(ctx, username, password)
	if err != nil {
		return "", auth.ErrUserNotFound
	}

	claims := AuthClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(a.residenceTime)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(a.signingKey)
}
