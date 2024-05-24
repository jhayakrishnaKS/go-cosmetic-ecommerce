package models

import "time"

type Order struct {
	ID            string `gorm:"primaryKey;type:uuid"`
	ProductID     string `gorm:"type:uuid"`
	UserID        string `gorm:"type:uuid"`
	OrderstatusID string `gorm:"type:uuid"`
	Address_id    string `gorm:"type:uuid"`
	CreatedAt     time.Time
}
