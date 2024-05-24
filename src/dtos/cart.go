package dtos

type CartReq struct {
	ID        string `gorm:"type:uuid;primaryKey"`
	ProductID string `json:"product_id"`
	UserID    string `json:"user_id"`
	Count     int    `json:"count"`
}
