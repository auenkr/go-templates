// Package greet contains the implementation of the GreetService.
package greet

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"connectrpc.com/validate"
	greetv1 "github.com/auenkr/go-templates/connect-server/gen/proto/greet/v1"
	"github.com/auenkr/go-templates/connect-server/gen/proto/greet/v1/greetv1connect"
	"github.com/auenkr/go-templates/connect-server/pkg/server"
	"go.uber.org/fx"
)

type GreetService struct{}

var _ greetv1connect.GreetServiceHandler = (*GreetService)(nil)

func (s *GreetService) Greet(ctx context.Context, req *connect.Request[greetv1.GreetRequest]) (*connect.Response[greetv1.GreetResponse], error) {
	return &connect.Response[greetv1.GreetResponse]{
		Msg: &greetv1.GreetResponse{
			Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
		},
	}, nil
}

type GreetServiceParams struct {
	fx.In
}

func NewGreetService(in GreetServiceParams) server.ServiceHandler {
	svc := &GreetService{}

	path, handler := greetv1connect.NewGreetServiceHandler(
		svc,
		connect.WithInterceptors(
			validate.NewInterceptor(),
		),
	)
	return server.ServiceHandler{
		ServiceName: greetv1connect.GreetServiceName,
		Path:        path,
		Handler:     handler,
	}
}
