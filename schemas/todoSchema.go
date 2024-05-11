package schemas

type Todo struct {
	BaseSchema
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
