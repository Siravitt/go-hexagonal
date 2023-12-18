package repository

import (
	"errors"
)

func (r userRepositoryMock) GetAllUser() ([]User, error) {
	users := []User{
		{Id: 1, Name: "Siravit"},
		{Id: 2, Name: "Tanratvijit"},
	}
	return users, nil
}

func (r userRepositoryMock) GetById(id int) (*User, error) {
	for _, user := range r.users {
		if user.Id == id {
			return &user, nil
		}
	}
	return nil, errors.New("User not found")
}
