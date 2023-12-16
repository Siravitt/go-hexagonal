// ! Port

package repository

type User struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

type UserRepository interface {
	GetAll() ([]User, error)
	GetById(int) (*User, error)
}
