//go:build wireinject
// +build wireinject

package wire

import (
	"database/sql"
	"e2dlabo/real_world/handlers"
	"e2dlabo/real_world/infrastructures"
	"e2dlabo/real_world/services"
	wire2 "e2dlabo/real_world/wire/wire"

	"github.com/google/wire"
)

func Initialize(db *sql.DB) (*handlers.Yeah, error) {
	wire.Build(
		wire2.EndSet,
		wire.Bind(new(services.DB), new(*infrastructures.DB)),
		newDB,
		wire.Bind(new(infrastructures.SqlDB), new(*sql.DB)),
	)
	return nil, nil
}

func newDB(db infrastructures.SqlDB) *infrastructures.DB {
	return &infrastructures.DB{
		DB: db,
	}
}
