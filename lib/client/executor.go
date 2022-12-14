package client

import (
	"context"
	"database/sql"
)

type Executor interface {
	ExecuteContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}
