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
	entity := &todo.Entity{
		Title:       req.Title,
		Description: req.Description,
		Completed:   false,
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
