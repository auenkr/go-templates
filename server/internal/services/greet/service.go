// Package greet contains the implementation of the GreetService.
package greet

import (
	"connectrpc.com/connect"
	"connectrpc.com/validate"
	"github.com/auenkr/go-templates/server/gen/proto/greet/v1/greetv1connect"
	"github.com/auenkr/go-templates/server/internal/interceptors"
	"github.com/auenkr/go-templates/server/pkg/server"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

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
