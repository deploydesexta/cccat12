package repository

import (
	"context"
	"database/sql"
)

// DbConn ISP - Interface Segregation Principle
type DbConn interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryRow(query string, args ...any) *sql.Row
}
