package pgdb

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func New() *sql.DB {
	connStr := "postgresql://postgres:123456@localhost/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
