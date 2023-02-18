package database

import (
	"database/sql"
	"fmt"
	"weehongayden/finance-flow-app/internal/config"

	_ "github.com/lib/pq"
)

func New(c config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", c.Database.User, c.Database.Pass, c.Database.Host, c.Database.Port, c.Database.Name)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
