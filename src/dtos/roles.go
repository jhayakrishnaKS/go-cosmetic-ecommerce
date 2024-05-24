package dtos

type RoleReq struct {
	ID string `gorm:"type:uuid;primaryKey"`
	Name string `json:"name"`
}
