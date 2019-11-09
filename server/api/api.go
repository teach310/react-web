package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

type TodoAPI struct {
	TodoRepository domain.TodoRepository
}

func (api *TodoAPI) HandleLoadTodo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	service := &domain.LoadTodo{
		TodoRepository: api.TodoRepository,
	}
	domainTodoList, err := service.Process(ctx)
	if err != nil {
		respond(w, http.StatusInternalServerError, err)
		return
	}

	todoList := make([]Todo, 0, len(domainTodoList))
	for _, data := range domainTodoList {
		todoList = append(todoList, Todo{
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

func (api *TodoAPI) HandleSaveTodo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestData := &SaveTodoRequest{}
	if err := getRequestData(r, requestData); err != nil {
		respond(w, http.StatusBadRequest, err)
		return
	}
	service := &domain.SaveTodo{
		TodoRepository: api.TodoRepository,
	}

	todoList := make([]domain.Todo, 0, len(requestData.TodoList))
	for _, data := range requestData.TodoList {
		todoList = append(todoList, domain.Todo{
			ID:     data.ID,
			IsDone: data.IsDone,
			Name:   data.Name,
		})
	}

	if err := service.Process(ctx, todoList); err != nil {
		respond(w, http.StatusInternalServerError, err)
		return
	}
	respondDefault(w)
}
