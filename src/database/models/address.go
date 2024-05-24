package models

type Address struct {
	ID      string `gorm:"primaryKey;type:uuid"`
	DoorNo  int
	Street  string
	City    string
	Zipcode int
	UserID  string `gorm:"type:uuid"`
}
