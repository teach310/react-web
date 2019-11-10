package authhttp

import (
	"context"
	"fmt"
	"net/http"
	"todo/auth"
	"todo/infra/firebase"
)

type Handler struct {
	firebaseAuth *firebase.FirebaseAuth
}

func New(ctx context.Context) (*Handler, error) {
	client, err := firebase.NewAuth(ctx)
	if err != nil {
		return nil, err
	}
	handler := &Handler{
		firebaseAuth: client,
	}
	return handler, nil
}

func (h *Handler) Handle(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Token")
		if token == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "no token")
			return
		}
		ctx := r.Context()
		account, err := h.firebaseAuth.VerifyIDToken(ctx, token)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "error: verify id token %v", err)
			return
		}
		ctx = auth.SetAccountOnContect(r.Context(), account)
		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h *Handler) HandleFunc(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Token")
		if token == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "no token")
			return
		}
		ctx := r.Context()
		account, err := h.firebaseAuth.VerifyIDToken(ctx, token)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "error: verify id token %v", err)
			return
		}
		ctx = auth.SetAccountOnContect(r.Context(), account)
		handler(w, r.WithContext(ctx))
	}
}
