package database

import (
	"context"
	"errors"
	"github.com/go-jet/jet/v2/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"math"
)

type Page[T any] struct {
	Items       []T
	TotalCount  int
	TotalPages  int
	CurrentPage int
	PageSize    int
}

func Paginate[T any](
	ctx context.Context,
	db *pgxpool.Pool,
	table postgres.Table,
	countColumn postgres.Column,
	baseSelect postgres.SelectStatement,
	page int,
	pageSize int,
) (Page[T], error) {
	var result Page[T]

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		return result, errors.New("invalid page size")
	}

	result.CurrentPage = page
	result.PageSize = pageSize

	countQuery := postgres.
		SELECT(postgres.COUNT(countColumn)).
		FROM(table)

	sqlCount, argsCount := countQuery.Sql()
	err := db.QueryRow(ctx, sqlCount, argsCount...).Scan(&result.TotalCount)
	if err != nil {
		return result, err
	}

	result.TotalPages = int(math.Ceil(float64(result.TotalCount) / float64(pageSize)))

	offset := (page - 1) * pageSize
	pagedQuery := baseSelect.
		LIMIT(int64(pageSize)).
		OFFSET(int64(offset))

	sqlSelect, argsSelect := pagedQuery.Sql()

	rows, err := db.Query(ctx, sqlSelect, argsSelect...)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	items, err := pgx.CollectRows(rows, pgx.RowToStructByPos[T])
	if err != nil {
		return result, err
	}

	result.Items = items
	return result, nil
}
