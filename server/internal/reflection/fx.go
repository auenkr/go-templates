package reflection

import "go.uber.org/fx"

var Module = fx.Module(
	"reflectionService",
	fx.Provide(
		fx.Annotate(
			NewReflectionService,
			fx.ResultTags(`group:"reflection"`),
		),
	),
	fx.Provide(
		fx.Annotate(
			NewReflectionRouteV1Alpha,
			fx.ResultTags(`group:"reflection"`),
		),
	),
)
