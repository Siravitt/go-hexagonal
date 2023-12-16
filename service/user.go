package service

type UserResponse struct {
	Name string `json:"name"`
}

type UserService interface {
	GetAllUser() ([]UserResponse, error)
	GetUser(int) (*UserResponse, error)
}
