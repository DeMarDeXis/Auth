package storage

import (
	"context"
	"errors"
	"github.com/jmoiron/sqlx"
	"log/slog"
	"sso/internal/domain/models"
	"sso/internal/storage/postgres"
)

var (
	ErrUserExists   = errors.New("user already exists")
	ErrUserNotFound = errors.New("user not found")
	ErrAppNotFound  = errors.New("app not found")
)

type Token interface {
	GetToken(ctx context.Context, user models.UserToken) (string, error)
}

type Storage struct {
	Token
}

func NewStorage(db *sqlx.DB, logg *slog.Logger) (*Storage, error) {
	return &Storage{
		Token: postgres.NewTokenStorage(db),
	}, nil
}
