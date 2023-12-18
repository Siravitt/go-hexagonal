package service

import "github.com/Siravitt/go-hexagonal/repository"

type Service interface {
	// User
	GetAllUser() ([]UserResponse, error)
	GetUser(int) (*UserResponse, error)

	// Todo
	NewTodo(TodoRequest, int) (*TodoResponse, error)
	GetTodo(int) ([]TodoResponse, error)
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return service{repo: repo}
}
