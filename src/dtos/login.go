package dtos

type LoginReq struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	RoleID   string `json:"role"`
}
