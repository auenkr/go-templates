// Package internalfx contains the fx.Module for the internal package.
package internalfx

import (
	clientsfx "github.com/auenkr/go-templates/server/internal/clients"
	interceptorsfx "github.com/auenkr/go-templates/server/internal/interceptors"
	reflectionfx "github.com/auenkr/go-templates/server/internal/reflection"
	servicesfx "github.com/auenkr/go-templates/server/internal/services"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"internal",
	reflectionfx.Module,
	servicesfx.Module,
	clientsfx.Module,
	interceptorsfx.Module,
)
