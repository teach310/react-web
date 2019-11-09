package api

// RequestとかResponseとか　いずれは自動生成したい

type LoadTodoResponse struct {
	TodoList []Todo `json:"todoList"`
}

type Todo struct {
	ID     int    `json:"id"`
	IsDone bool   `json:"isDone"`
	Name   string `json:"name"`
}

type SaveTodoRequest struct {
	TodoList []Todo `json:"todoList"`
}
