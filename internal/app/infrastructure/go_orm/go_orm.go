package go_orm

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewClient),
)
