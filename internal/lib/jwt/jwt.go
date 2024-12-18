package jwt

import (
	"sso/internal/domain/models"
	"time"

	"github.com/golang-jwt/jwt"
)

// TODO: add to config
const secretKey = "GdeGeneratorrr"

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
