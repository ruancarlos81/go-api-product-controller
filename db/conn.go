package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	connStr := "postgres://postgres:1234@go_db:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr) // Usa o driver pq para "postgres"
	if err != nil {
		return nil, err
	}
	return db, nil
}
