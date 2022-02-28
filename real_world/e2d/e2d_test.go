package e2d

import (
	"e2dlabo/real_world/handlers"
	"e2dlabo/real_world/infrastructures"
	"e2dlabo/real_world/services"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.MonitorPingsOption(true))
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM users WHERE id = ? LIMIT 1`)).WillReturnRows(
		sqlmock.NewRows([]string{"id", "name", "phone_number"}).AddRow("hoge", "hoge", "09012345678"))
	d := &infrastructures.DB{
		DB: db,
	}
	target := &handlers.Yeah{
		UserService: &services.UserService{
			DB: d,
		},
	}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	target.Me(c)
	assert.Equal(t, http.StatusOK, c.Writer.Status())
}
