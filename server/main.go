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

const (
	dbType = 2 // 1: mysql 2: firestore
)

func main() {

	if dbType == 1 {
		mysql.Connect()
		defer mysql.CloseConnect()
		mysql.Migrate()
	}

	ctx := context.Background()
	if err := firebase.Init(ctx); err != nil {
		log.Println("error: firebase initialize", err)
		return
	}

	if dbType == 2 {
		firebase.InitDefaultFirestoreClient(ctx)
		defer firebase.CloseDefaultFirestoreClient()
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
		// TaskRepository: &repository.TaskRepository{},
		TaskRepository: &repository.FSTaskRepository{},
	}
	http.HandleFunc("/load", allowAccessControl(authHandler.HandleFunc(todoAPI.HandleLoadTodo)))
	http.HandleFunc("/save", allowAccessControl(authHandler.HandleFunc(todoAPI.HandleSaveTodo)))

	http.HandleFunc("/test", allowAccessControl(test))
	http.HandleFunc("/testread", allowAccessControl(testread))
}

func test(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	client, err := firebase.NewFirestore(ctx)
	if err != nil {
		panic(err)
	}
	defer client.Close()
	// Setで上書き
	_, err = client.Doc("users/1").Set(ctx, map[string]interface{}{
		"first": "Ada",
		"last":  "Lovelace",
		"born":  1815,
	})
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}
	fmt.Fprintln(w, "ok")
}

func testread(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	client, err := firebase.NewFirestore(ctx)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	type User struct {
		First string `firestore:"first"`
		Last  string `firestore:"last"`
		Born  int    `firestore:"born"`
	}

	// Getで取得
	docsnap, err := client.Doc("users/1").Get(ctx)
	if err != nil {
		panic(err)
	}
	user := &User{}
	if err := docsnap.DataTo(user); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, fmt.Sprintf("%v", user))
}
