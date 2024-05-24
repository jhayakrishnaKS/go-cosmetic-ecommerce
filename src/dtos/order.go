package dtos

type OrderReq struct {
	ID            string `gorm:"type:uuid;primaryKey"`
	ProductID     string `json:"product_id"`
	UserID        string `json:"user_id"`
	OrderstatusID string `json:"orderstatus_id"`
	AddressId     string `json:"address_id"`
}
