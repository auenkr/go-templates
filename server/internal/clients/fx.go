// Package clients contains constructor for creating clients to different services.
package clients

import "go.uber.org/fx"

var Module = fx.Module(
	"clients",
	fx.Provide(NewDBQueries),
)
