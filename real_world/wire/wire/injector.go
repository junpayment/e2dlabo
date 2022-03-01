//go:build wireinject
// +build wireinject

package wire

import (
	"e2dlabo/real_world/handlers"
	"e2dlabo/real_world/infrastructures"
	"e2dlabo/real_world/services"

	"github.com/google/wire"
)

func Initialize(dsn string) (*handlers.Yeah, error) {
	wire.Build(
		handlers.NewYeah,
		wire.Bind(new(handlers.UserService), new(*services.UserService)),
		services.NewUserService,
		wire.Bind(new(services.DB), new(*infrastructures.DB)),
		infrastructures.NewDB,
	)
	return nil, nil
}
