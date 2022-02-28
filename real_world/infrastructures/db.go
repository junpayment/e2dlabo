package infrastructures

import (
	"database/sql"
	"e2dlabo/real_world/models"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//go:generate mockgen -source=db.go -destination=mocks/db.go -package=mocks

type SqlDB interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

type DB struct {
	DB SqlDB
}

func NewDB(dsn string) (*DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf(`db, err := sql.Open("mysql", dsn): %w`, err)
	}
	return &DB{
		DB: db,
	}, nil
}

func (d *DB) GetUserByID(id string) (*models.User, error) {
	rows, err := d.DB.Query("SELECT * FROM users WHERE id = ? LIMIT 1", id)
	if err != nil {
		return nil, fmt.Errorf(
			`rows, err := d.DB.Query("SELECT * FROM users WHERE id = ? LIMIT 1", id): %w`, err)
	}
	if rows.Next() == false {
		return nil, errors.New(`no user`)
	}
	user := &models.User{}
	err = rows.Scan(&user.ID, &user.Name, &user.PhoneNumber)
	if err != nil {
		return nil, fmt.Errorf(
			`err = rows.Scan(&user.ID, &user.Name, &user.PhoneNumber): %w`, err)
	}
	return user, nil
}
