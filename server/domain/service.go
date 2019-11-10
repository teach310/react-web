package domain

import (
	"context"
)

type LoadTodo struct {
	TodoRepository TodoRepository
}

func (s *LoadTodo) Run(ctx context.Context) ([]Todo, error) {
	todos, err := s.TodoRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

type SaveTodo struct {
	TodoRepository TodoRepository
}

func (s *SaveTodo) Run(ctx context.Context, todoList []Todo) error {
	if err := s.TodoRepository.Save(ctx, todoList); err != nil {
		return err
	}
	return nil
}
