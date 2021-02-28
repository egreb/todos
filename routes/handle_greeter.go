package routes

import (
	"context"

	"egreb.net/todos/generated"
)

// GreeterService handler
type GreeterService struct{}

// Greet route
func (GreeterService) Greet(ctx context.Context, r generated.GreetRequest) (*generated.GreetResponse, error) {
	resp := &generated.GreetResponse{
		Greeting: "Hello " + r.Name,
	}

	return resp, nil
}
