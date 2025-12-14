package main

import (
	"context"
	"log"
	"net/http"
	"sync"

	"github.com/bufbuild/connect-go"

	todov1 "todoBackend/gen/todo/v1"
	"todoBackend/gen/todo/v1/todo_v1connect"
)

/*
In-memory store.
Mutex is needed because HTTP handlers can be called concurrently.
*/
type todoStore struct {
	mu     sync.Mutex
	todos  []*todov1.Todo
	nextID int32
}

/*
Service implementation.
This struct will implement the methods required by TodoService.
*/
type todoService struct {
	store *todoStore
}

/*
CreateTodo RPC implementation.
*/
func (s *todoService) CreateTodo(
	ctx context.Context,
	req *connect.Request[todov1.CreateTodoRequest],
) (*connect.Response[todov1.CreateTodoResponse], error) {

	s.store.mu.Lock()
	defer s.store.mu.Unlock()

	todo := &todov1.Todo{
		Id:    s.store.nextID,
		Title: req.Msg.Title,
	}
	s.store.nextID++
	s.store.todos = append(s.store.todos, todo)

	resp := connect.NewResponse(&todov1.CreateTodoResponse{
		Todo: todo,
	})

	return resp, nil
}

/*
ListTodos RPC implementation.
*/
func (s *todoService) ListTodos(
	ctx context.Context,
	req *connect.Request[todov1.ListTodosRequest],
) (*connect.Response[todov1.ListTodosResponse], error) {

	s.store.mu.Lock()
	defer s.store.mu.Unlock()

	resp := connect.NewResponse(&todov1.ListTodosResponse{
		Todos: s.store.todos,
	})

	return resp, nil
}

func main() {
	// Initialize in-memory store
	store := &todoStore{
		todos:  make([]*todov1.Todo, 0),
		nextID: 1,
	}

	// Create service implementation
	service := &todoService{
		store: store,
	}

	// Create HTTP mux
	mux := http.NewServeMux()

	/*
		ConnectRPC generates an HTTP handler for the service.
		path  -> URL path where the service is exposed
		handler -> standard http.Handler
	*/
	path, handler := todo_v1connect.NewTodoServiceHandler(service)

	mux.Handle(path, handler)

	log.Println("Todo backend listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
