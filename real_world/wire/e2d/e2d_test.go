package e2d

import (
	"e2dlabo/real_world/wire/e2d/wire"
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
	target, _ := wire.Initialize(db)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	target.Me(c)
	assert.Equal(t, http.StatusOK, c.Writer.Status())
}
