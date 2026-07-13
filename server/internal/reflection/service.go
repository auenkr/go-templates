// Package reflection: Add reflection routes to the server
package reflection

import (
	"connectrpc.com/grpcreflect"
	"github.com/auenkr/go-templates/server/pkg/server"
	"go.uber.org/fx"
)

type ReflectionRouteParams struct {
	fx.In
	Routes []server.ServiceHandler `group:"service-handlers"`
}

func NewReflectionService(in ReflectionRouteParams) server.ServiceHandler {
	serviceNames := make([]string, len(in.Routes))
	for i, r := range in.Routes {
		serviceNames[i] = r.ServiceName
	}

	p, h := grpcreflect.NewHandlerV1(
		grpcreflect.NewStaticReflector(
			serviceNames...,
		),
	)
	return server.ServiceHandler{
		ServiceName: grpcreflect.ReflectV1ServiceName,
		Path:        p,
		Handler:     h,
	}
}

type ReflectionRouteV1AlphaParams struct {
	fx.In
	Routes []server.ServiceHandler `group:"service-handlers"`
}

func NewReflectionRouteV1Alpha(in ReflectionRouteV1AlphaParams) server.ServiceHandler {
	serviceNames := make([]string, len(in.Routes))
	for i, r := range in.Routes {
		serviceNames[i] = r.ServiceName
	}

	p, h := grpcreflect.NewHandlerV1Alpha(
		grpcreflect.NewStaticReflector(
			serviceNames...,
		),
	)
	return server.ServiceHandler{
		ServiceName: grpcreflect.ReflectV1AlphaServiceName,
		Path:        p,
		Handler:     h,
	}
}
