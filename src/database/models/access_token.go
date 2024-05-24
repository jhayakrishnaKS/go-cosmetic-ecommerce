package models

import "time"

type AccessToken struct {
	Token         string
	RefreshTokens string
	UserID     string `gorm:"type:uuid"`
	ExpiresAt     time.Time
}
