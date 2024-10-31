package postgres

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/jmoiron/sqlx"
	"sso/internal/domain/models"
	"sso/internal/lib/jwt"
	"time"
)

const salt = "hjqrhjqw124617ajfhajs"

type AuthStorage struct {
	db *sqlx.DB
}

func NewTokenStorage(db *sqlx.DB) *AuthStorage {
	return &AuthStorage{db: db}
}

func (s *AuthStorage) GetToken(ctx context.Context, user models.UserToken) (string, error) {
	const op = "storage.postgres.GetToken"

	token, err := generateTokenHash(user.UserID)
	if err != nil {
		return "", fmt.Errorf("%s: failed to generate token: %w", op, err)
	}

	// Store token usage in database
	q := `
        INSERT INTO users (user_id, hashToken, created_at, expires_at) 
        VALUES ($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP + INTERVAL '24 hours')
    	`

	if _, err := s.db.ExecContext(ctx, q, user.UserID, token); err != nil {
		return "", fmt.Errorf("%s: failed to store token: %w", op, err)
	}

	return token, nil
}

func generateTokenHash(userID int64) (token string, err error) {
	const op = "storage.postgres.generateTokenHash"

	token, err = jwt.NewToken(models.UserToken{UserID: userID}, 24*time.Hour)
	if err != nil {
		return "", fmt.Errorf("%s: failed to create JWT token: %w", op, err)
	}

	hash := sha256.Sum256([]byte(token + salt))

	return fmt.Sprintf("%x", hash), nil
}
