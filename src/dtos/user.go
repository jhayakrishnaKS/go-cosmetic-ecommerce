package dtos

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	RoleID   string `json:"roleId"`
}
