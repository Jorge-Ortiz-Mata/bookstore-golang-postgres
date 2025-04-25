package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

var JwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))

func GenerateAuthToken(userId uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Second * 30).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(JwtSecret)

	if err != nil {
		return "", err
	}

	return signedToken, nil
}
