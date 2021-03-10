package definitions

import "egreb.net/todos/todo"

// TodoService create, read, update or delete.
type TodoService interface {
	// Greet makes a greeting.
	Get(GetTodoRequest) GetTodoResponse
	Create(CreateTodoRequest) CreateTodoResponse
	Delete(DeleteTodoRequest) DeleteTodoResponse
}

// GetTodoRequest based by id
type GetTodoRequest struct {
	ID int
}

// GetTodoResponse returns the todo.
type GetTodoResponse struct {
	Todo todo.Todo
}

type CreateTodoRequest struct {
	Title       string
	Description string
}

type CreateTodoResponse struct {
	Todo todo.Todo
}
type DeleteTodoRequest struct {
	TodoId int
}

type DeleteTodoResponse struct {
	Success bool
}
