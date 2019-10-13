package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	ID     int    `json:"id"`
	IsDone bool   `json:"isDone"`
	Name   string `json:"name"`
}

func main() {
	// http.HandleFunc("/", handler)
	files := http.FileServer(http.Dir("./public"))
	http.Handle("/page/", http.StripPrefix("/page/", files))
	http.Handle("/admin/", http.StripPrefix("/admin/", http.FileServer(http.Dir("./build"))))

	// http.Handle("/", http.FileServer(http.Dir("./build")))

	fmt.Println("Server Started Port 8080")
	http.HandleFunc("/load", loadTodo)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func loadTodo(w http.ResponseWriter, r *http.Request) {
	todoModels := []Todo{
		Todo{ID: 1, IsDone: false, Name: "task1"},
		Todo{ID: 2, IsDone: false, Name: "task2"},
	}

	json, _ := json.Marshal(todoModels)
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(json)
}
