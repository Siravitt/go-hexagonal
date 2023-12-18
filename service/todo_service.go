package service

import (
	"database/sql"

	"github.com/Siravitt/go-hexagonal/errs"
	"github.com/Siravitt/go-hexagonal/logs"
	"github.com/Siravitt/go-hexagonal/repository"
)

// type todoService struct {
// 	todoRepo repository.TodoRepository
// }

// func NewTodoService(todoRepo repository.TodoRepository) todoService {
// 	return todoService{todoRepo: todoRepo}
// }

func (s service) NewTodo(req TodoRequest, userId int) (*TodoResponse, error) {
	todoRequest := repository.Todo{
		Task:      req.Task,
		UserId:    userId,
		Completed: 0,
	}

	newTodo, err := s.repo.Create(todoRequest)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	response := TodoResponse{
		Id:        newTodo.Id,
		Task:      newTodo.Task,
		Completed: newTodo.Completed,
	}
	return &response, nil
}

func (s service) GetTodo(userId int) ([]TodoResponse, error) {
	todos, err := s.repo.GetAllTodo(userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("todo not found")
		}
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	todoResponses := []TodoResponse{}
	for _, todo := range todos {
		todoResponse := TodoResponse{
			Id:        todo.Id,
			Task:      todo.Task,
			Completed: todo.Completed,
		}
		todoResponses = append(todoResponses, todoResponse)
	}
	return todoResponses, nil
}
