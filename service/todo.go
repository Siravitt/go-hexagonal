package service

type TodoResponse struct {
	Id        int    `json:"id"`
	Task      string `json:"task"`
	Completed int    `json:"completed"`
}

type TodoRequest struct {
	Task string `json:"task"`
}
