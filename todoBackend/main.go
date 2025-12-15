package main

import (
	"log"
	"net/http"

	todov1connect "todoBackend/gen/todo/v1/todo_v1connect"
)

func withCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Connect-Protocol-Version")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		h.ServeHTTP(w, r)
	})
}

func main() {
	handler := NewTodoHandler()

	mux := http.NewServeMux()
	path, h := todov1connect.NewTodoServiceHandler(handler)
	mux.Handle(path, h)

	log.Println("Backend running on :8080")
	log.Fatal(http.ListenAndServe(":8080", withCORS(mux)))
}
