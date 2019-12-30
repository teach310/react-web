package domain

import (
	"context"
	"todo/utility/uuid"
)

type LoadTodo struct {
	OwnerID        string
	TaskRepository TaskRepository
}

func (s *LoadTodo) Run(ctx context.Context) ([]Task, error) {
	tasks, err := s.TaskRepository.FindAllByOwnerID(ctx, s.OwnerID)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

type SaveTodo struct {
	OwnerID        string
	TaskRepository TaskRepository
}

func (s *SaveTodo) Run(ctx context.Context, tasks []Task) error {

	// わたされたIDを全てUUIDで上書きしてしまう。
	// 普通ありえないが、更新時に全部書き換えるという仕様の上ならこれでいける
	for i := 0; i < len(tasks); i++ {
		tasks[i].ID = uuid.NewString()
	}

	if err := s.TaskRepository.Save(ctx, s.OwnerID, tasks); err != nil {
		return err
	}
	return nil
}
