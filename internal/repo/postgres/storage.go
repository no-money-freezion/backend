package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
)

const (
	shortUrlsTableName = "urls"
	longUrlsTableName  = "urls_long"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(connStr string) (*Repository, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Repository{
		db: db,
	}, nil
}
