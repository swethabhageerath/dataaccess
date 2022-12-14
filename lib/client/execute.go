package client

import (
	"context"
	"database/sql"
)

type Execute struct {
	db *sql.DB
}

func NewExecute(db *sql.DB) Execute {
	return Execute{
		db: db,
	}
}

func (e Execute) ExecuteContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return e.db.ExecContext(ctx, query, args)
}

func (e Execute) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return e.db.QueryRowContext(ctx, query, args)
}

func (e Execute) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return e.db.QueryContext(ctx, query, args)
}
