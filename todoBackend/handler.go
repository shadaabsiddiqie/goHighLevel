package main

import (
	"context"
	"sync"

	todov1 "todoBackend/gen/todo/v1"

	"github.com/bufbuild/connect-go"
)

type UrlHandler struct {
	sortToOriginal map[string]string
	cunter         int
	// urls *todov1.ShortenURLRequest
}

func NewUrlHandler() *UrlHandler {
	return &UrlHandler{}
}

func (h *UrlHandler) ShortenURL(
	ctx context.Context,
	req *connect.Request[todov1.ShortenURLRequest],
) (*connect.Response[todov1.ShortenURLResponse], error) {
	// h.urls = req.Msg
	// rand (7)
	// url1 _ rand1

	// url1 _ rand2

	// MySQL (rand set) (100000) (101-200 given)
	// Redis (rand set) (1-100) all were not used
	// Redis (rand set) (1-100) all were not used

	h.sortToOriginal[req.Msg.OriginalUrl] = "short.ly/" + h.cunter
	h.cunter++
	return connect.NewResponse(&todov1.ShortenURLResponse{
		ShortenedUrl: "short.ly/" + req.Msg.OriginalUrl,
	}), nil
}

func (h *UrlHandler) ExpandURL(
	ctx context.Context,
	req *connect.Request[todov1.ExpandURLRequest],
) (*connect.Response[todov1.ExpandURLResponse], error) {
	originalUrl := h.sortToOriginal[req.Msg.ShortenedUrl]
	return connect.NewResponse(&todov1.ExpandURLResponse{
		OriginalUrl: originalUrl,
	}), nil
}

type TodoHandler struct {
	mu    sync.Mutex
	todos []*todov1.Todo
	next  int32
}

func NewTodoHandler() *TodoHandler {
	return &TodoHandler{next: 1}
}

func (h *TodoHandler) CreateTodo(
	ctx context.Context,
	req *connect.Request[todov1.CreateTodoRequest],
) (*connect.Response[todov1.CreateTodoResponse], error) {

	h.mu.Lock()
	defer h.mu.Unlock()

	todo := &todov1.Todo{
		Id:    h.next,
		Title: req.Msg.Title,
	}
	h.next++
	h.todos = append(h.todos, todo)

	return connect.NewResponse(&todov1.CreateTodoResponse{
		Todo: todo,
	}), nil
}

func (h *TodoHandler) ListTodos(
	ctx context.Context,
	req *connect.Request[todov1.ListTodosRequest],
) (*connect.Response[todov1.ListTodosResponse], error) {

	h.mu.Lock()
	defer h.mu.Unlock()

	return connect.NewResponse(&todov1.ListTodosResponse{
		Todos: h.todos,
	}), nil
}
