package jwt

import (
	"github.com/golang-jwt/jwt"
	"sso/internal/domain/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewToken(t *testing.T) {
	// Arrange
	user := models.UserToken{
		UserID: 1,
	}
	duration := time.Minute * 15

	// Act
	token, err := NewToken(user, duration)

	// Assert
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestNewToken_ValidateToken(t *testing.T) {
	// Arrange
	user := models.UserToken{
		UserID: 1,
	}
	duration := time.Minute * 15

	// Act
	token, err := NewToken(user, duration)

	// Assert
	assert.NoError(t, err)

	// Parse and validate token
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	assert.NoError(t, err)
	assert.True(t, parsedToken.Valid)

	// Validate claims
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	assert.True(t, ok)
	assert.Equal(t, float64(user.UserID), claims["uid"])
}

func TestJWT(t *testing.T) {
	testCases := []struct {
		name     string
		user     models.UserToken
		duration time.Duration
		wantErr  bool
	}{
		{
			name:     "Valid Token",
			user:     models.UserToken{UserID: 1},
			duration: time.Minute * 15,
			wantErr:  false,
		},
		//TODO: fix it
		{
			name:     "Zero Duration",
			user:     models.UserToken{UserID: 1},
			duration: 0,
			wantErr:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			token, err := NewToken(tc.user, tc.duration)
			if tc.wantErr {
				assert.Error(t, err)
				assert.Empty(t, token)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, token)
			}
		})
	}
}
