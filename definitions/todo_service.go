package definitions

import "egreb.net/todos/todo"

// TodoService create, read, update or delete.
type TodoService interface {
	GetAll(GetAllTodosRequest) GetAllTodosResponse
	Get(GetTodoRequest) GetTodoResponse
	Create(CreateTodoRequest) CreateTodoResponse
	Delete(DeleteTodoRequest) DeleteTodoResponse
	Update(UpdateTodoRequest) UpdateTodoResponse
}

// GetAllTodosRequest - needs pagination
type GetAllTodosRequest struct {
	Completed bool
}

// GetAllTodosResponse - needs pagination
type GetAllTodosResponse struct {
	Todos []todo.Todo
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
	Completed   bool
}
type CreateTodoResponse struct {
	Todo todo.Todo
}

type UpdateTodoRequest struct {
	Todo todo.Todo
}
type UpdateTodoResponse struct {
	Todo todo.Todo
}

type DeleteTodoRequest struct {
	TodoId int
}
type DeleteTodoResponse struct {
	Success bool
}
