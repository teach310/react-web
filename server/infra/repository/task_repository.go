package repository

import (
	"context"
	"todo/domain"
	"todo/infra/model"
	"todo/infra/mysql"
)

type TaskRepository struct {
}

func (r *TaskRepository) FindAllByOwnerID(ctx context.Context, ownerID string) ([]domain.Task, error) {
	db := mysql.DefaultDB()
	var tasks []model.Task
	if err := db.Where("owner_id = ?", ownerID).Find(&tasks).Error; err != nil {
		if mysql.IsErrNotFound(err) {
			emptyList := make([]domain.Task, 0)
			return emptyList, nil
		}
		return nil, err
	}
	results := make([]domain.Task, 0, len(tasks))
	for _, data := range tasks {
		results = append(results, domain.Task{
			ID:      data.ID,
			IsDone:  data.IsDone,
			Name:    data.Name,
			OwnerID: data.OwnerID,
		})
	}
	return results, nil
}

func (r *TaskRepository) Save(ctx context.Context, ownerID string, todoList []domain.Task) (rerr error) {
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
	if rerr = tx.Where("owner_id = ?", ownerID).Delete(model.Task{}).Error; rerr != nil {
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
