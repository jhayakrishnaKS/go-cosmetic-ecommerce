package models

import "time"

type Users struct {
	ID        string `gorm:"type:uuid;primaryKey"`
	Email     string
	Username  string
	Password  string
	RoleID    string    `gorm:"type:uuid"`
	CreatedAt time.Time `gorm:"column:created_at"`
}
