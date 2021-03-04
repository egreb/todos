package routes

import (
	"context"
	"fmt"
	"testing"

	"egreb.net/todos/generated"
	"github.com/matryer/is"
)

func TestHandleGreeting(t *testing.T) {
	is := is.New(t)
	gs := GreeterService{}

	ctx := context.Background()
	req := generated.GreetRequest{
		Name: "Toby",
	}
	resp, err := gs.Greet(ctx, req)
	is.NoErr(err)
	is.Equal(resp.Greeting, "Hello "+"Toby")
}

func TestEmptyHandleGreeting(t *testing.T) {
	is := is.New(t)
	gs := GreeterService{}

	ctx := context.Background()
	req := generated.GreetRequest{
		Name: "",
	}
	resp, err := gs.Greet(ctx, req)
	is.Equal(fmt.Errorf("Name cannot be empty"), err)
	is.Equal(resp, nil)
}
