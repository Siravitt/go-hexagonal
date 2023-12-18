// ! Port

package repository

type User struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

type userRepositoryMock struct {
	users []User
}
