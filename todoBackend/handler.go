package main

import (
	"context"
	"sync"

	todov1 "todoBackend/gen/todo/v1"

	"github.com/bufbuild/connect-go"
)

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
