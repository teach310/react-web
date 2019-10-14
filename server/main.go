package main

import (
	"fmt"
	"net/http"

	"todo/domain"
	"todo/infra/mysql"
)

func main() {

	mysql.Connect()
	defer mysql.CloseConnect()
	mysql.Migrate()

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
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if r.Method == "OPTIONS" { // 認証用
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func register() {
	http.HandleFunc("/load", allowAccessControl(domain.HandleLoadTodo))
	http.HandleFunc("/save", allowAccessControl(domain.HandleSaveTodo))
}
