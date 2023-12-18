package service

import (
	"database/sql"

	"github.com/Siravitt/go-hexagonal/errs"
	"github.com/Siravitt/go-hexagonal/logs"
)

// type userService struct {
// 	userRepo repository.UserRepository
// }

// func NewUserService(userRepo repository.UserRepository) userService {
// 	return userService{userRepo: userRepo}
// }

func (s service) GetAllUser() ([]UserResponse, error) {
	users, err := s.repo.GetAllUser()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	userResponses := []UserResponse{}
	for _, user := range users {
		userResponse := UserResponse{
			Name: user.Name,
		}
		userResponses = append(userResponses, userResponse)
	}
	return userResponses, nil
}

func (s service) GetUser(id int) (*UserResponse, error) {
	user, err := s.repo.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("user not found")
		}
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	userResponse := UserResponse{
		Name: user.Name,
	}
	return &userResponse, nil
}
