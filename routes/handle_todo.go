package routes

import (
	"context"
	"log"

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
	entity := &todo.Entity{ID: req.ID}
	err := t.DB.Model(entity).WherePK().Select()
	if err != nil {
		log.Println("TodoService.Get:", err)
		return nil, err
	}
	todo := entity.ToModel()

	return &generated.GetTodoResponse{
		Todo: todo,
	}, nil
}
