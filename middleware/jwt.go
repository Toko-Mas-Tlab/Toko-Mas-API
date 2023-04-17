package middleware

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type IService interface {
	GenerateTokenJWT(id int, username string, minute int, secretToken string) (string, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateTokenJWT(id int, username string, minute int, secretToken string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "nil", err
	}

	claim := jwt.MapClaims{}
	claim["id"] = id
	claim["fullname"] = username
	claim["exp"] = time.Now().Add(time.Minute * time.Duration(minute)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString([]byte(os.Getenv(secretToken)))
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
