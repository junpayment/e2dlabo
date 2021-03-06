// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"e2dlabo/real_world/handlers"
	"e2dlabo/real_world/infrastructures"
	"e2dlabo/real_world/services"
)

// Injectors from injector.go:

func Initialize(dsn string) (*handlers.Yeah, error) {
	db, err := infrastructures.NewDB(dsn)
	if err != nil {
		return nil, err
	}
	userService, err := services.NewUserService(db)
	if err != nil {
		return nil, err
	}
	yeah, err := handlers.NewYeah(userService)
	if err != nil {
		return nil, err
	}
	return yeah, nil
}
