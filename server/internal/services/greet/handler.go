package greet

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	greetv1 "github.com/auenkr/go-templates/server/gen/proto/greet/v1"
	"github.com/auenkr/go-templates/server/gen/proto/greet/v1/greetv1connect"
)

type Service struct{}

var _ greetv1connect.GreetServiceHandler = (*Service)(nil)

func (s *Service) Greet(ctx context.Context, req *connect.Request[greetv1.GreetRequest]) (*connect.Response[greetv1.GreetResponse], error) {
	return &connect.Response[greetv1.GreetResponse]{
		Msg: &greetv1.GreetResponse{
			Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
		},
	}, nil
}
