package dtos

type RegisterReq struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	RoleID   string `json:"roleId"`
}

