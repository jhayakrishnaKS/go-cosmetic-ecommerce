package dtos

type ProductReq struct {
	ID            string  `gorm:"type:uuid;primaryKey"`
	Product_title string  `json:"Product_title"`
	Description   string  `json:"Description"`
	Price         float64 `json:"Price"`
	Brand         string  `json:"Brand"`
}

type UpdateProductReq struct{
	Product_title string  `json:"Product_title"`
	Description   string  `json:"Description"`
	Price         float64 `json:"Price"`
	Brand         string  `json:"Brand"`
}