package service

import (
	"errors"
	"pokedex/shared/config"

	"github.com/dgrijalva/jwt-go"
)

var (
	key = config.GetConfig().JwtSecret
)

//go:generate mockery --name AuthService --case snake --output ../mocks/mocksServices --disable-version-string
type AuthService interface {
	ValidateToken(encodedToken string) (*jwt.Token, error)
	GenerateToken(userID int) (string, error)
}

type jwtService struct {
}

func NewAuthService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(encodedToken *jwt.Token) (interface{}, error) {
		_, ok := encodedToken.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(key), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{
		"id": userID,
	}

	// generate token using HS256 with claim
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString([]byte(key))

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
