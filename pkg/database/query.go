package database

import (
	"context"
	"github.com/go-jet/jet/v2/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"iter"
)

type Queryable interface {
	Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
	Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
}

func QueryRaw(ctx context.Context, db Queryable, stmt postgres.Statement) (pgx.Rows, error) {
	query, args := stmt.Sql()

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func QueryWithReturningField[T any](ctx context.Context, db Queryable, stmt postgres.Statement) (*T, error) {
	rows, err := QueryRaw(ctx, db, stmt)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	v, err := pgx.CollectOneRow(rows, pgx.RowTo[T])

	if err != nil {
		return nil, err
	}

	return &v, nil
}

func QueryRow[T any](ctx context.Context, db Queryable, stmt postgres.Statement) (*T, error) {
	rows, err := QueryRaw(ctx, db, stmt)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	return pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByPos[T])
}

func Query[T any](ctx context.Context, db Queryable, stmt postgres.Statement) ([]T, error) {
	rows, err := QueryRaw(ctx, db, stmt)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	return pgx.CollectRows(rows, pgx.RowToStructByPos[T])
}

func Exec(ctx context.Context, db Queryable, stmt postgres.Statement) error {
	query, args := stmt.Sql()
	_, err := db.Exec(ctx, query, args...)
	return err
}

func Chunk[T any](ctx context.Context, db Queryable, stmt postgres.SelectStatement, chunkSize int) iter.Seq2[T, error] {
	return func(yield func(x T, e error) bool) {
		page := 1
		offset := (page - 1) * chunkSize

		stmt = stmt.OFFSET(
			int64(offset),
		).LIMIT(
			int64(chunkSize),
		)

		items, err := Query[T](ctx, db, stmt)

		if err != nil {
			if !yield(*new(T), err) {
				return
			}
		}

		for _, item := range items {
			if !yield(item, nil) {
				return
			}
		}
	}
}
