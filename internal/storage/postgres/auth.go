package postgres

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"sso/internal/domain/models"
)

type AuthStorage struct {
	db *sqlx.DB
}

func NewAuthStorage(db *sqlx.DB) *AuthStorage {
	return &AuthStorage{db: db}
}

func (s *AuthStorage) SaveUser(ctx context.Context, email string, passHash []byte) (int64, error) {
	const op = "storage.postgres.SaveUser"

	query := `INSERT INTO users(email, pass_hash) VALUES($1, $2) RETURNING id`

	var id int64
	err := s.db.QueryRowContext(ctx, query, email, passHash).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (s *AuthStorage) User(ctx context.Context, email string) (models.User, error) {
	const op = "storage.postgres.User"

	stmt, err := s.db.Prepare("SELECT id, email, pass_hash FROM users WHERE email = $1")
	if err != nil {
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, email)

	var user models.User

	err = row.Scan(&user.ID, &user.Email, &user.PassHash)
	if err != nil {
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}
	// TODO: add errors
	return user, nil
}

func (s *AuthStorage) IsAdmin(ctx context.Context, userID int64) (bool, error) {
	const op = "storage.postgres.IsAdmin"

	stmt, err := s.db.Prepare("SELECT is_admin FROM users WHERE id = $1")
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, userID)
	var isAdmin bool

	err = row.Scan(&isAdmin)
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	return isAdmin, nil
}

func (s *AuthStorage) App(ctx context.Context, appID int) (models.App, error) {
	const op = "storage.postgres.App"

	stmt, err := s.db.Prepare("SELECT id, name, description, user_id FROM apps WHERE id = $1")
	if err != nil {
		return models.App{}, fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, appID)

	var app models.App
	err = row.Scan(&app.ID, &app.Name, &app.Secret)
	if err != nil {
		return models.App{}, fmt.Errorf("%s: %w", op, err)
	}

	return app, nil
}
