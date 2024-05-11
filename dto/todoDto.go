package dto

type CreateTodoDto struct {
	Title string `json:"title"`
}

type UpdateTodoDto struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
