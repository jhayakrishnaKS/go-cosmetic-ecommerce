package dtos

type AddressReq struct {
	ID      string `gorm:"primaryKey;type:uuid"`
	DoorNo  int    `json:"door_no"`
	Street  string `json:"street"`
	City    string `json:"city"`
	Zipcode int    `json:"zipcode"`
}

type UpdateAddressReq struct {
	DoorNo  int    `json:"door_no"`
	Street  string `json:"street"`
	City    string `json:"city"`
	Zipcode int    `json:"zipcode"`
}
