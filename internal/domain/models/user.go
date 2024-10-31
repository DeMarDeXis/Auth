package models

import "time"

type UserToken struct {
	ID        int
	UserID    int64
	CreatedAt time.Time
	ExpiresAt time.Time
}
