package jwt

import (
	"sso/internal/domain/models"
	"time"

	"github.com/golang-jwt/jwt"
)

// TODO: add to config
const secretKey = "xH7cT9pK2vN4mX8qL3jY5nR6bA1wE0iS"

func NewToken(user models.UserToken, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.UserID
	claims["exp"] = time.Now().Add(duration).Unix()

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
	// TODO: test
}
