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
