package services

import (
	"e2dlabo/real_world/models"
	"e2dlabo/real_world/services/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserService_Me(t *testing.T) {
	ctrl := gomock.NewController(t)
	db := mocks.NewMockDB(ctrl)
	db.EXPECT().GetUserByID(gomock.Any()).Return(&models.User{
		ID:          "test",
		Name:        "test",
		PhoneNumber: "test",
	}, nil)
	target := &UserService{DB: db}
	user, err := target.Me()
	assert.Nil(t, err)
	assert.NotNil(t, user)
}
