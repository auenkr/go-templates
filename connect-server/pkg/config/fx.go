// Package config contains the configuration for the application.
package config

import "go.uber.org/fx"

var Module = fx.Module(
	"config",
	fx.Provide(NewConfig),
	fx.Provide(NewLogger),
)
