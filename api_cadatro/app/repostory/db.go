package repostory

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type DB interface {
	Exec(ctx context.Context, sql string, arg ...any) (pgconn.CommandTag, error)

	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)

	QueryRow(ctx context.Context, sql string, args ...any) pgx.Rows
}
