package routes

import (
	"context"
	"fmt"

	"egreb.net/todos/generated"
	"egreb.net/todos/todo"
	"github.com/go-pg/pg/v10"
)

// TodoService handler
type TodoService struct {
	DB *pg.DB
}

// Get Todo by ID
func (t TodoService) Get(ctx context.Context, req generated.GetTodoRequest) (*generated.GetTodoResponse, error) {
	entity := new(todo.Entity)
	err := t.DB.Model(entity).Where("id = ?", req.ID).Select()
	if err != nil {
		return &generated.GetTodoResponse{Error: fmt.Sprintf("Todo by id %v not found", req.ID)}, err
	}
	todo := entity.ToModel()

	return &generated.GetTodoResponse{
		Todo: todo,
	}, nil
}

// Create Todo with Title and optional Description
func (t TodoService) Create(ctx context.Context, req generated.CreateTodoRequest) (*generated.CreateTodoResponse, error) {
	var (
		description = ""
		completed   = false
	)
	if req.Description != "" {
		description = req.Description
	}

	completed = req.Completed

	entity := &todo.Entity{
		Title:       req.Title,
		Description: description,
		Completed:   completed,
	}

	_, err := t.DB.Model(entity).Returning("*").Insert()
	if err != nil {
		return nil, err
	}

	todo := entity.ToModel()

	return &generated.CreateTodoResponse{
		Todo: todo,
	}, nil
}

// Create Todo with Title and optional Description
func (t TodoService) Update(ctx context.Context, req generated.UpdateTodoRequest) (*generated.UpdateTodoResponse, error) {
	entity := req.Todo.ToEntity()
	_, err := t.DB.Model(&entity).Set("title = ?title, description = ?description, completed = ?completed").Where("id = ?id").Update()
	if err != nil {
		return nil, err
	}

	todo := entity.ToModel()
	return &generated.UpdateTodoResponse{
		Todo: todo,
	}, nil
}

// Delete Todo by Id
func (t TodoService) Delete(ctx context.Context, req generated.DeleteTodoRequest) (*generated.DeleteTodoResponse, error) {
	entity := &todo.Entity{ID: req.TodoId}
	res, err := t.DB.Model(entity).Where("id = ?", req.TodoId).Delete()
	if err != nil || res.RowsAffected() == 0 {
		return &generated.DeleteTodoResponse{
			Success: false,
			Error:   fmt.Sprintf("Could not delete Todo by id %d", req.TodoId),
		}, err
	}

	return &generated.DeleteTodoResponse{Success: true}, nil
}

// GetAll todos
func (t TodoService) GetAll(ctx context.Context, req generated.GetAllTodosRequest) (*generated.GetAllTodosResponse, error) {
	var (
		entities []todo.Entity
		err      error
	)
	err = t.DB.Model(&entities).Order("created_at desc").Select()
	if err != nil {
		return &generated.GetAllTodosResponse{
				Error: "Could not fetch todos",
			},
			err
	}

	var todos []todo.Todo
	for _, e := range entities {
		todos = append(todos, e.ToModel())
	}

	return &generated.GetAllTodosResponse{
		Todos: todos,
	}, nil
}
