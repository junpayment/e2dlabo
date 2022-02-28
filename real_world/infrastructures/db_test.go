package infrastructures

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestDB_GetUserByID(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.MonitorPingsOption(true))
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM users WHERE id = ? LIMIT 1`)).WillReturnRows(
		sqlmock.NewRows([]string{"id", "name", "phone_number"}).AddRow("hoge", "hoge", "09012345678"))
	target := &DB{
		DB: db,
	}
	user, err := target.GetUserByID("hoge")
	assert.Nil(t, err)
	assert.NotNil(t, user)
}
