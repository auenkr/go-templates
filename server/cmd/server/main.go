package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	internalfx "github.com/auenkr/go-templates/server/internal"
	pkgfx "github.com/auenkr/go-templates/server/pkg"
	"github.com/auenkr/go-templates/server/pkg/config"
	"github.com/auenkr/go-templates/server/pkg/server"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	app := fx.New(
		internalfx.Module,
		pkgfx.Module,

		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),

		fx.Invoke(StartServer),
	)
	app.Run()
}

type StartServerParams struct {
	fx.In
	Lifecycle fx.Lifecycle

	Logger *zap.Logger
	Config *config.Config

	ServiceHandler    []server.ServiceHandler `group:"service-handlers"`
	ReflectionHandler []server.ServiceHandler `group:"reflection"`
}

func StartServer(in StartServerParams) {
	mux := http.NewServeMux()
	for _, handler := range in.ServiceHandler {
		in.Logger.Info("Registering service", zap.String("service", handler.ServiceName))
		mux.Handle(handler.Path, handler.Handler)
	}

	if in.Config.Environment != config.PROD_ENVIRONMENT {
		for _, handler := range in.ReflectionHandler {
			in.Logger.Info("Registering reflection service", zap.String("service", handler.ServiceName))
			mux.Handle(handler.Path, handler.Handler)
		}
	}

	protocol := &http.Protocols{}
	protocol.SetHTTP1(true)
	// Use h2c so we can serve HTTP/2 without TLS.
	protocol.SetUnencryptedHTTP2(true)

	server := http.Server{
		Addr:      fmt.Sprintf(":%s", in.Config.Port),
		Handler:   mux,
		Protocols: protocol,
	}

	in.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					panic(err)
				}
			}()
			in.Logger.Info("Server started", zap.String("port", in.Config.Port))

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
}
