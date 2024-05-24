package models

type Cart struct {
	ID        string `gorm:"primaryKey;autoIncrement"`
	ProductID string `gorm:"type:uuid"`
	UserID    string `gorm:"type:uuid"`
	Count     int
}
