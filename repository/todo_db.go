package repository

// type todoRepositoryDB struct {
// 	db *sqlx.DB
// }

// func NewTodoRepositoryDB(db *sqlx.DB) TodoRepository {
// 	return todoRepositoryDB{db: db}
// }

func (r repository) Create(t Todo) (*Todo, error) {
	query := "insert into todos (task, completed, userId) values (?, ?, ?)"
	result, err := r.db.Exec(
		query,
		t.Task,
		t.Completed,
		t.UserId,
	)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	t.UserId = int(id)
	return &t, nil
}

func (r repository) GetAllTodo(userId int) ([]Todo, error) {
	query := "select id, task, completed from todos where user_id = ?"
	todos := []Todo{}
	err := r.db.Select(&todos, query, userId)
	if err != nil {
		return nil, err
	}
	return todos, nil
}
