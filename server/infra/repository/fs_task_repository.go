package repository

import (
	"context"
	"fmt"
	"todo/domain"
	"todo/infra/firebase"
	"todo/infra/model"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type FSTaskRepository struct {
}

func (r *FSTaskRepository) FindAllByOwnerID(ctx context.Context, ownerID string) ([]domain.Task, error) {
	db := firebase.DefaultFirestoreClient()
	// CollectionRefを取得
	collectionPath := fmt.Sprintf("users/%s/tasks", ownerID)
	iter := db.Collection(collectionPath).Documents(ctx)
	var tasks []*model.FSTask
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		fmt.Println(doc.Data())
		task := &model.FSTask{}
		if err := doc.DataTo(task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	rtn := make([]domain.Task, 0, len(tasks))
	for i := 0; i < len(tasks); i++ {
		task := tasks[i]
		rtn = append(rtn, domain.Task{
			ID:      task.ID,
			IsDone:  task.IsDone,
			Name:    task.Name,
			OwnerID: task.OwnerID,
		})
	}
	return rtn, nil
}

func (r *FSTaskRepository) Save(ctx context.Context, ownerID string, todoList []domain.Task) error {
	db := firebase.DefaultFirestoreClient()
	// CollectionRefを取得
	collectionPath := fmt.Sprintf("users/%s/tasks", ownerID)
	collectionRef := db.Collection(collectionPath)
	err := db.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		// read
		documentRefs, err := tx.DocumentRefs(collectionRef).GetAll()
		if err != nil {
			return err
		}

		// write
		for _, doc := range documentRefs {
			tx.Delete(doc)
		}

		for i := 0; i < len(todoList); i++ {
			domainTask := todoList[i]
			fsTask := model.FSTask{
				ID:      domainTask.ID,
				IsDone:  domainTask.IsDone,
				Name:    domainTask.Name,
				OwnerID: domainTask.OwnerID,
			}
			err = tx.Set(collectionRef.Doc(fsTask.ID), fsTask)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
	// batchSize := 100 // 最大100個まとめて処理
	// // コレクションの中のDocを全部Delete
	// for {
	// 	iter := collectionRef.Limit(batchSize).Documents(ctx)
	// 	numDeleted := 0
	// 	// batchを作成
	// 	batch := db.Batch()
	// 	for {
	// 		doc, err := iter.Next()
	// 		if err == iterator.Done {
	// 			break
	// 		}
	// 		if err != nil {
	// 			return err
	// 		}

	// 		batch.Delete(doc.Ref)
	// 		numDeleted++
	// 	}
	// 	if numDeleted == 0 {
	// 		break
	// 	}

	// 	_, err := batch.Commit(ctx)
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	// Docを作成
	// batchをコミット

	// コレクションの中のドキュメントを全て取得
	// - users  col
	// - userID doc
	//     - tasks col
	//         - taskID
	//             - FSTask
}
