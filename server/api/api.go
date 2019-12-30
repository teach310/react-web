package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"todo/auth"
	"todo/domain"
)

func respondDefault(w http.ResponseWriter) {
	respond(w, 200, "ok")
}

func respondJson(w http.ResponseWriter, statusCode int, data interface{}) {
	bytes, _ := json.Marshal(data)
	w.WriteHeader(statusCode)
	w.Write(bytes)
}

func respond(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	fmt.Fprint(w, data)
}

func getRequestData(r *http.Request, data interface{}) error {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, data); err != nil {
		return err
	}
	return err
}

type TodoHandler struct {
	TaskRepository domain.TaskRepository
}

func (api *TodoHandler) HandleLoadTodo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	account, ok := auth.GetAccountFromContext(ctx)
	if !ok {
		respond(w, http.StatusBadRequest, "account not found")
		return
	}

	service := &domain.LoadTodo{
		OwnerID:        account.UserID,
		TaskRepository: api.TaskRepository,
	}
	domainTodoList, err := service.Run(ctx)
	if err != nil {
		respond(w, http.StatusInternalServerError, err)
		return
	}

	todoList := make([]UniqueTodo, 0, len(domainTodoList))
	for _, data := range domainTodoList {
		todoList = append(todoList, UniqueTodo{
			ID:     data.ID,
			IsDone: data.IsDone,
			Name:   data.Name,
		})
	}

	responseData := &LoadTodoResponse{
		TodoList: todoList,
	}
	respondJson(w, 200, responseData)
}

func (api *TodoHandler) HandleSaveTodo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	account, ok := auth.GetAccountFromContext(ctx)
	if !ok {
		respond(w, http.StatusBadRequest, "account not found")
		return
	}

	requestData := &SaveTodoRequest{}
	if err := getRequestData(r, requestData); err != nil {
		respond(w, http.StatusBadRequest, err)
		return
	}
	service := &domain.SaveTodo{
		OwnerID:        account.UserID,
		TaskRepository: api.TaskRepository,
	}

	tasks := make([]domain.Task, 0, len(requestData.TodoList))
	for _, data := range requestData.TodoList {
		tasks = append(tasks, domain.Task{
			IsDone:  data.IsDone,
			Name:    data.Name,
			OwnerID: account.UserID,
		})
	}

	if err := service.Run(ctx, tasks); err != nil {
		respond(w, http.StatusInternalServerError, err)
		return
	}
	respondDefault(w)
}
