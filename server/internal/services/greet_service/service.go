// Package greet_service contains the implementation of the GreetService.
package greet_service

import (
	"connectrpc.com/connect"
	"connectrpc.com/validate"
	"github.com/auenkr/go-templates/server/gen/proto/greet_service/v1/greet_servicev1connect"
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

	path, handler := greet_servicev1connect.NewGreetServiceHandler(
		svc,
		connect.WithInterceptors(
			validate.NewInterceptor(),
			interceptors.NewLatencyLoggerInterceptor(in.Logger),
		),
	)
	return server.ServiceHandler{
		ServiceName: greet_servicev1connect.GreetServiceName,
		Path:        path,
		Handler:     handler,
	}
}
