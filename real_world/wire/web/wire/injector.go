//go:build wireinject
// +build wireinject

package wire

import (
	"e2dlabo/real_world/handlers"
	"e2dlabo/real_world/infrastructures"
	"e2dlabo/real_world/services"
	wire2 "e2dlabo/real_world/wire/wire"

	"github.com/google/wire"
)

func Initialize(dsn string) (*handlers.Yeah, error) {
	wire.Build(
		wire2.EndSet,
		wire.Bind(new(services.DB), new(*infrastructures.DB)),
		infrastructures.NewDB,
	)
	return nil, nil
}
