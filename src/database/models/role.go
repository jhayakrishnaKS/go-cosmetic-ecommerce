package models

type Roles struct {
	ID   string `gorm:"primaryKey;type:uuid"`
	Name string
}
