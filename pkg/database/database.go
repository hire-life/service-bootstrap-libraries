package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zoondengine/hl/service-bootstrap-libraries/pkg/logger"
	"go.uber.org/zap"
	"sync"
	"time"
)

var (
	connections sync.Map
	mx          sync.Mutex
)

func GetOrNew(conn string) *pgxpool.Pool {
	if v, ok := connections.Load(conn); ok {
		return v.(*pgxpool.Pool)
	}

	mx.Lock()
	defer mx.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, conn)

	if err != nil {
		logger.Get().Fatal("Failed to create new pool for pgsql", zap.Error(err), zap.String("conn", conn))
	}

	connections.Store(conn, pool)
	return pool
}
