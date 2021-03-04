package routes

import (
	"context"
	"fmt"

	"egreb.net/todos/generated"
)

// GreeterService handler
type GreeterService struct{}

// Greet route
func (GreeterService) Greet(ctx context.Context, r generated.GreetRequest) (*generated.GreetResponse, error) {
	if r.Name == "" {
		return nil, fmt.Errorf("Name cannot be empty")
	}

	resp := &generated.GreetResponse{
		Greeting: "Hello " + r.Name,
	}

	return resp, nil
}
