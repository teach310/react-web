package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"todo/api"
	"todo/auth/authhttp"
	"todo/infra/firebase"
	"todo/infra/mysql"
	"todo/infra/repository"
)

var authHandler *authhttp.Handler

func main() {

	mysql.Connect()
	defer mysql.CloseConnect()
	mysql.Migrate()

	ctx := context.Background()
	if err := firebase.Init(ctx); err != nil {
		log.Println("error: firebase initialize", err)
		return
	}

	ah, err := authhttp.New(ctx)
	if err != nil {
		log.Println("error: authhttp new", err)
		return
	}
	authHandler = ah

	// files := http.FileServer(http.Dir("./public"))
	// http.Handle("/page/", http.StripPrefix("/page/", files))
	// http.Handle("/admin/", http.StripPrefix("/admin/", http.FileServer(http.Dir("./build"))))

	register()
	fmt.Println("Server Started Port 8080")
	http.ListenAndServe(":8080", nil)
}

func allowAccessControl(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// CORS対策
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Token")

		if r.Method == "OPTIONS" { // 認証用
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next(w, r)
	}
}

func register() {
	todoAPI := &api.TodoHandler{
		TodoRepository: &repository.TodoRepository{},
	}
	http.HandleFunc("/load", allowAccessControl(authHandler.HandleFunc(todoAPI.HandleLoadTodo)))
	http.HandleFunc("/save", allowAccessControl(authHandler.HandleFunc(todoAPI.HandleSaveTodo)))
}
