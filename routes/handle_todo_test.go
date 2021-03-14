package routes_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"egreb.net/todos/database"
	"egreb.net/todos/generated"
	"egreb.net/todos/routes"
	"egreb.net/todos/todo"
	"github.com/docker/go-connections/nat"
	"github.com/go-pg/pg/v10"
	"github.com/matryer/is"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func CreateTestContainer(ctx context.Context, dbname string) (testcontainers.Container, *pg.DB, error) {
	var env = map[string]string{
		"POSTGRES_PASSWORD": "password",
		"POSTGRES_USER":     "postgres",
		"POSTGRES_DB":       dbname,
	}

	port, _ := nat.NewPort("tcp", "5432")
	req := testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "postgres:12.3-alpine",
			ExposedPorts: []string{port.Port()},
			Env:          env,
			WaitingFor:   wait.ForListeningPort(port),
		},
		Started: true,
	}
	container, err := testcontainers.GenericContainer(ctx, req)
	if err != nil {
		return container, nil, fmt.Errorf("failed to start container: %s", err)
	}

	mappedPort, err := container.MappedPort(ctx, nat.Port(port))
	if err != nil {
		return container, nil, fmt.Errorf("failed to get container external port: %s", err)
	}

	url := fmt.Sprintf("localhost:%s", mappedPort.Port())
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "password",
		Database: dbname,
		Addr:     url,
	})
	if err != nil {
		return container, db, fmt.Errorf("failed to establish database connection: %s", err)
	}

	return container, db, nil
}

func TestTodoGet(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()
	container, db, err := CreateTestContainer(ctx, "foo")
	defer container.Terminate(ctx)
	is.NoErr(err)

	err = database.CreateSchema(db)
	is.NoErr(err)
	service := routes.TodoService{DB: db}

	todoID := -1
	res, err := service.Get(ctx, generated.GetTodoRequest{ID: todoID})
	is.True(err != nil)
	is.Equal(res.Error, fmt.Sprintf("Todo by id %v not found", todoID))
}

func TestTodoCreate(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()
	container, db, err := CreateTestContainer(ctx, "createTodoDb")
	defer container.Terminate(ctx)
	is.NoErr(err)

	err = database.CreateSchema(db)
	is.NoErr(err)
	service := routes.TodoService{DB: db}

	res, err := service.Create(ctx, generated.CreateTodoRequest{
		Title:       "Test create todo",
		Description: "Test todo description",
	})
	is.NoErr(err)

	is.True(res != nil)
	is.True(res.Error == "")

	is.True(res.Todo.ID > 0)
}

func TestTodoDelete(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()
	container, db, err := CreateTestContainer(ctx, "createTodoDb")
	defer container.Terminate(ctx)
	is.NoErr(err)

	err = database.CreateSchema(db)
	is.NoErr(err)
	service := routes.TodoService{DB: db}

	createResult, err := service.Create(ctx, generated.CreateTodoRequest{
		Title:       "Test create todo",
		Description: "Test todo description",
	})
	is.NoErr(err)

	deleteResult, err := service.Delete(ctx, generated.DeleteTodoRequest{TodoId: createResult.Todo.ID})
	is.NoErr(err)
	is.True(deleteResult.Success)

	deleteResult2, err := service.Delete(ctx, generated.DeleteTodoRequest{TodoId: -1})
	is.NoErr(err)
	is.True(deleteResult2.Success == false)
}

func TestTodoGetAll(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()
	container, db, err := CreateTestContainer(ctx, "createTodoDb")
	defer container.Terminate(ctx)
	is.NoErr(err)

	err = database.CreateSchema(db)
	is.NoErr(err)
	service := routes.TodoService{DB: db}

	todosToInsert := []todo.Todo{
		{
			Title:       "Test 1",
			Description: "This is 'Test 1' todo",
		},
		{
			Title:       "Test 2",
			Description: "This is 'Test 2' todo",
		},
	}

	for _, tti := range todosToInsert {
		_, err = service.Create(ctx, generated.CreateTodoRequest{
			Title:       tti.Title,
			Description: tti.Description,
		})
		is.NoErr(err)
		time.Sleep(1 * time.Second)
	}

	res, err := service.GetAll(ctx, generated.GetAllTodosRequest{})
	is.NoErr(err)

	is.True(len(res.Todos) == len(todosToInsert))
	is.Equal(res.Todos[0].Title, "Test 2")
}
