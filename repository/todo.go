package repository

type Todo struct {
	Id        int    `db:"id"`
	Task      string `db:"description"`
	Completed int    `db:"completed"`
	UserId    int    `db:"user_id"`
}
