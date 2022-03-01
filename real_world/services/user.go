package services

import (
	"e2dlabo/real_world/models"
	"fmt"
)

//go:generate mockgen -source=user.go -destination=mocks/user.go -package=mocks

type DB interface {
	GetUserByID(id string) (*models.User, error)
}

type UserService struct {
	DB DB
}

var id = "hoge"

func NewUserService(d DB) (*UserService, error) {
	return &UserService{
		DB: d,
	}, nil
}

func (u *UserService) Me() (*models.User, error) {
	user, err := u.DB.GetUserByID(id)
	if err != nil {
		return nil, fmt.Errorf(`user, err := u.DB.GetUserByID(id): %w`, err)
	}
	return user, nil
}
