package auth

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"sso/internal/domain/models"
	"sso/internal/lib/logger/sl"
	"time"
)

type Auth struct {
	log      *slog.Logger
	getter   Getter
	tokenTTL time.Duration
}

type Getter interface {
	GetToken(ctx context.Context, user models.UserToken) (string, error)
}

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidAppID       = errors.New("invalid app id")
	ErrUserExists         = errors.New("user already exists")
	ErrUserNotFound       = errors.New("user not found")
)

// New returns a new instance of the Auth service
func New(
	log *slog.Logger,
	getter Getter,
	tokenTTL time.Duration,
) *Auth {
	return &Auth{
		getter:   getter,
		log:      log,
		tokenTTL: tokenTTL,
	}
}

// GetToken generates a new JWT token for the given user ID.
func (a *Auth) GetToken(ctx context.Context, user models.UserToken) (string, error) {
	const op = "auth.GetToken"

	log := a.log.With(
		slog.String("op", op),
		slog.Int64("user_id", user.UserID),
	)

	token, err := a.getter.GetToken(ctx, user)
	if err != nil {
		log.Error("failed to get token", sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}

	log.Info("token generated")
	return token, nil
}
