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

type UserSaver interface {
	SaveUser(ctx context.Context, email string, passHash []byte) (uid int64, err error)
}

type UserProvider interface {
	User(ctx context.Context, email string) (models.User, error)
	IsAdmin(ctx context.Context, userID int64) (bool, error)
}

type AppProvider interface {
	App(ctx context.Context, appID int) (models.App, error)
}

type Storage struct {
	UserSaver
	UserProvider
	AppProvider
}

func NewStorage(db *sqlx.DB, logg *slog.Logger) (*Storage, error) {
	return &Storage{
		UserSaver:    postgres.NewAuthStorage(db),
		UserProvider: postgres.NewAuthStorage(db),
		AppProvider:  postgres.NewAuthStorage(db),
	}, nil
}
