package hello

import "context"

type Response struct {
	Message string
}

//encore:api public
func World(ctx context.Context) (*Response, error) {
  return &Response{Message: "Hello, world!"}, nil
}
