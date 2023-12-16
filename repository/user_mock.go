package repository

import "errors"

type userRepositoryMock struct {
	users []User
}

func NewUserRepositoryMock() userRepositoryMock {
	users := []User{
		{Id: 1, Name: "Siravit"},
		{Id: 2, Name: "Tanratvijit"},
	}
	return userRepositoryMock{users: users}
}

func (r userRepositoryMock) GetAll() ([]User, error) {
	return r.users, nil
}

func (r userRepositoryMock) GetById(id int) (*User, error) {
	for _, user := range r.users {
		if user.Id == id {
			return &user, nil
		}
	}
	return nil, errors.New("User not found")
}
