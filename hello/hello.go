package hello

import "context"
import "time"
import "log"

type Response struct {
	Message string
}

//encore:api public
func World(ctx context.Context) (*Response, error) {
  go waitForMe()
  return &Response{Message: "Hello, world!"}, nil
}

func waitForMe() {
    time.Sleep(1 *time.Second)
    log.Println("you waited!")
}

