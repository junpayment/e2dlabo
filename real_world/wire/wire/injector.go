package wire

import (
	"e2dlabo/real_world/handlers"
	"e2dlabo/real_world/services"

	"github.com/google/wire"
)

var EndSet = wire.NewSet(
	handlers.NewYeah,
	wire.Bind(new(handlers.UserService), new(*services.UserService)),
	services.NewUserService,
)
