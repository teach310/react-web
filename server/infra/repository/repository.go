package repository

import (
	"context"
	"todo/cerrors"
	"todo/domain"
	"todo/infra/model"
	"todo/infra/mysql"
)

type TodoRepository struct {
}

func (r *TodoRepository) FindAll(ctx context.Context) ([]domain.Todo, error) {
	db := mysql.DefaultDB()
	var todoList []model.Todo
	if err := db.Find(&todoList).Error; err != nil {
		if mysql.IsErrNotFound(err) {
			return nil, cerrors.ErrNotFound
		}
		return nil, err
	}
	results := make([]domain.Todo, 0, len(todoList))
	for _, data := range todoList {
		results = append(results, domain.Todo{
			ID:     data.ID,
			IsDone: data.IsDone,
			Name:   data.Name,
		})
	}
	return results, nil
}

func (r *TodoRepository) Save(ctx context.Context, todoList []domain.Todo) (rerr error) {
	db := mysql.DefaultDB()
	tx := db.Begin()
	defer func() {
		if rerr != nil {
			tx.Rollback()
		}

		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 全部消して作り直す
	if rerr = tx.Delete(model.Todo{}, "id > 0").Error; rerr != nil {
		return
	}

	for _, task := range todoList {
		if rerr = tx.Create(&task).Error; rerr != nil {
			return
		}
	}

	rerr = tx.Commit().Error
	return
}
