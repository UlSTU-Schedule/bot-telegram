package postgres

import (
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

// NewStudentDB ...
func NewStudentDB(databaseUrl string) (*sqlx.DB, error) {
	db, err := sqlx.Open("pgx", databaseUrl)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
