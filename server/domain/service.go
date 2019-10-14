package domain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"todo/infra/model"
	"todo/infra/mysql"
)

func HandleLoadTodo(w http.ResponseWriter, r *http.Request) {
	var todoList []model.Todo
	db := mysql.DefaultDB()
	db.Find(&todoList)
	json, _ := json.Marshal(todoList)
	w.WriteHeader(200)
	w.Write(json)
}

func HandleSaveTodo(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var todoList []model.Todo
	if err := json.Unmarshal(body, &todoList); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := setTodoList(todoList); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func setTodoList(todoList []model.Todo) (rerr error) {
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
	rerr = tx.Delete(model.Todo{}, "id > 0").Error
	if rerr != nil {
		return
	}

	for _, task := range todoList {
		rerr = tx.Create(&task).Error
		if rerr != nil {
			return
		}
	}

	return tx.Commit().Error
}
