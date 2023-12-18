package service

import (
	"database/sql"
	"errors"

	"github.com/Siravitt/go-hexagonal/logs"
	"github.com/Siravitt/go-hexagonal/repository"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) userService {
	return userService{userRepo: userRepo}
}

func (s userService) GetAllUser() ([]UserResponse, error) {
	users, err := s.userRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, err
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

func (s userService) GetUser(id int) (*UserResponse, error) {
	user, err := s.userRepo.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("User not found")
		}
		logs.Error(err)
		return nil, err
	}
	userResponse := UserResponse{
		Name: user.Name,
	}
	return &userResponse, nil
}
