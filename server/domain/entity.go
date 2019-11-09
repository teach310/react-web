package domain

import "context"

type Todo struct {
	ID     int
	IsDone bool
	Name   string
}

type TodoRepository interface {
	FindAll(ctx context.Context) ([]Todo, error)
	Save(ctx context.Context, todoList []Todo) error
}
