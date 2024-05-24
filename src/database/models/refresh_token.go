package models

import "time"

type RefreshToken struct {
	Token     string
	UserID string `gorm:"type:uuid"`
	ExpiresAt time.Time
}
