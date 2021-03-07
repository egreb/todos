package definitions

import "egreb.net/todos/todo"

// TodoService create, read, update or delete.
type TodoService interface {
	// Greet makes a greeting.
	Get(GetTodoRequest) GetTodoResponse
}

// GetTodoRequest based by id
type GetTodoRequest struct {
	ID int
}

// GetTodoResponse returns the todo.
type GetTodoResponse struct {
	Todo todo.Todo
}
