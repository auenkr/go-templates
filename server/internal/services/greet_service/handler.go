package greet_service

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	greet_servicev1 "github.com/auenkr/go-templates/server/gen/proto/greet_service/v1"
	"github.com/auenkr/go-templates/server/gen/proto/greet_service/v1/greet_servicev1connect"
)

type Service struct{}

var _ greet_servicev1connect.GreetServiceHandler = (*Service)(nil)

func (s *Service) Greet(ctx context.Context, req *connect.Request[greet_servicev1.GreetRequest]) (*connect.Response[greet_servicev1.GreetResponse], error) {
	return &connect.Response[greet_servicev1.GreetResponse]{
		Msg: &greet_servicev1.GreetResponse{
			Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
		},
	}, nil
}
