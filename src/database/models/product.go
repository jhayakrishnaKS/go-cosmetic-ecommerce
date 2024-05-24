package models

type Products struct {
	ID           string `gorm:"primaryKey;type:uuid"`
	Product_title string
	Description   string
	Price         float64
	Brand         string
}
