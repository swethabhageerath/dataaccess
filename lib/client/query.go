package client

import (
	"context"
	"database/sql"
)

type Query struct {
	db *sql.DB
}

func NewQuery(db *sql.DB) Query {
	return Query{
		db: db,
	}
}

func (q Query) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return q.db.QueryRowContext(ctx, query, args)
}

func (q Query) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return q.db.QueryContext(ctx, query, args)
}
