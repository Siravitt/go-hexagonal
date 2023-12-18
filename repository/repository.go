package repository

import "github.com/jmoiron/sqlx"

type Repository interface {
	// User
	GetAllUser() ([]User, error)
	GetById(int) (*User, error)

	// Todo
	Create(Todo) (*Todo, error)
	GetAllTodo(int) ([]Todo, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return repository{db: db}
}

func NewUserRepositoryMock(db *sqlx.DB) Repository {
	return repository{db: nil}
}
