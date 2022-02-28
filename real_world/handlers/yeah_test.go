package handlers

import (
	"e2dlabo/real_world/handlers/mocks"
	"e2dlabo/real_world/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestYeah_Me(t *testing.T) {
	ctrl := gomock.NewController(t)
	s := mocks.NewMockUserService(ctrl)
	s.EXPECT().Me().Return(&models.User{
		ID:          "test",
		Name:        "test",
		PhoneNumber: "test",
	}, nil)
	target := &Yeah{UserService: s}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	target.Me(c)
	assert.Equal(t, http.StatusOK, c.Writer.Status())
}
