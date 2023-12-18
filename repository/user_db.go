// ! Adapter

package repository

// type userRepositoryDB struct {
// 	db *sqlx.DB
// }

// func NewUserRepositoryDB(db *sqlx.DB) UserRepository {
// 	return userRepositoryDB{db: db}
// }

func (r repository) GetAllUser() ([]User, error) {
	users := []User{}
	query := "select id, name from users"
	err := r.db.Select(&users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r repository) GetById(id int) (*User, error) {
	user := User{}
	query := "select id, name from users where id = ?"
	err := r.db.Get(&user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
