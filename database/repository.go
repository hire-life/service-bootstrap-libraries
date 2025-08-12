package database

import (
	"context"
)

type CrudRepository[T any] interface {
	FindById(ctx context.Context, id int64) (*T, error)
	Create(ctx context.Context, m *T) error
	Update(ctx context.Context, m *T) error
	Delete(ctx context.Context, m *T) error
	DeleteById(ctx context.Context, id int64) error
}
