// Package greet contains the implementation of the GreetService.
package greet

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"connectrpc.com/validate"
	greetv1 "github.com/auenkr/go-templates/server/gen/proto/greet/v1"
	"github.com/auenkr/go-templates/server/gen/proto/greet/v1/greetv1connect"
	"github.com/auenkr/go-templates/server/internal/interceptors"
	"github.com/auenkr/go-templates/server/pkg/server"
	"go.uber.org/fx"
	"go.uber.org/zap"
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

type ServiceParams struct {
	fx.In

	Logger *zap.Logger
}

func NewService(in ServiceParams) server.ServiceHandler {
	svc := &Service{}

	path, handler := greetv1connect.NewGreetServiceHandler(
		svc,
		connect.WithInterceptors(
			validate.NewInterceptor(),
			interceptors.NewLatencyLoggerInterceptor(in.Logger),
		),
	)
	return server.ServiceHandler{
		ServiceName: greetv1connect.GreetServiceName,
		Path:        path,
		Handler:     handler,
	}
}
