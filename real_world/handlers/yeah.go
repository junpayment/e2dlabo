package handlers

import (
	"e2dlabo/real_world/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=yeah.go -destination=mocks/yeah.go -package=mocks

type UserService interface {
	Me() (*models.User, error)
}

type Yeah struct {
	UserService UserService
}

func NewYeah(u UserService) (*Yeah, error) {
	return &Yeah{
		UserService: u,
	}, nil
}

func (y *Yeah) Me(c *gin.Context) {
	user, err := y.UserService.Me()
	if err != nil {
		_ = c.AbortWithError(
			http.StatusInternalServerError, fmt.Errorf(`user, err := y.UserService.Me(): %w`, err))
		return
	}
	c.JSON(http.StatusOK, user)
}
