package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

type AuthService struct {
	username  string
	password  string
	jwtSecret string
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func NewAuthService(username, password, jwtSecret string) *AuthService {
	return &AuthService{
		username:  username,
		password:  password,
		jwtSecret: jwtSecret,
	}
}

func (s *AuthService) Login(username, password string) (string, error) {
	if username != s.username || password != s.password {
		return "", ErrInvalidCredentials
	}

	now := time.Now()

	claims := Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(s.jwtSecret))
}

func (s *AuthService) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, errors.New("unexpected signing method")
			}

			return []byte(s.jwtSecret), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}