// Package servicesfx contains the fx.Module for the services package.
package servicesfx

import (
	"github.com/auenkr/go-templates/server/internal/services/greet"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"services",

	// GreetService
	fx.Provide(
		fx.Annotate(
			greet.NewService,
			fx.ResultTags(`group:"service-handlers"`),
		),
	),
)
