package domain

import "context"

type Task struct {
	ID      string
	IsDone  bool
	Name    string
	OwnerID string
}

type TaskRepository interface {
	FindAllByOwnerID(ctx context.Context, ownerID string) ([]Task, error)
	Save(ctx context.Context, ownerID string, tasks []Task) error
}
