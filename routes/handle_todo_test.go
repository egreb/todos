package routes_test

import (
	"context"
	"fmt"
	"testing"

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

func postgresRequest(identifier string) (testcontainers.ContainerRequest, nat.Port) {
	port, _ := nat.NewPort("tcp", "5432")
	dbName := fmt.Sprintf("%s-db", identifier)
	req := testcontainers.ContainerRequest{
		Name:         dbName,
		Image:        "postgres:12.3-alpine",
		ExposedPorts: []string{port.Port()},
		Env: map[string]string{
			"POSTGRES_DB":               "test_db",
			"POSTGRES_HOST_AUTH_METHOD": "trust",
		},
		Networks:   []string{identifier},
		WaitingFor: wait.ForListeningPort(port),
	}

	return req, port
}

func testTodoGet(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()
	req, _ := postgresRequest("foo")
	c, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	defer c.Terminate(ctx)
	is.NoErr(err)

	db := pg.Connect(&pg.Options{
		Database: "test_db",
	})
	database.CreateSchema(db)

	service := routes.TodoService{DB: db}

	res, err := service.Get(ctx, generated.GetTodoRequest{ID: -1})
	is.NoErr(err)

	is.Equal(res.Todo, todo.Todo{ID: -1})
}
