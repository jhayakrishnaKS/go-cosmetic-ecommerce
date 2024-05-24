package models

type OrderStatus struct {
	ID     string `gorm:"primaryKey;type:uuid"`
	Status string
}
