package cache

import (
	"context"
	"github.com/hire-life/service-bootstrap-libraries/logger"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"sync"
)

var (
	connections sync.Map
	mx          sync.Mutex
)

func GetOrNew(conn string) *redis.Client {
	if val, ok := connections.Load(conn); ok {
		return val.(*redis.Client)
	}

	mx.Lock()
	defer mx.Unlock()

	opt, err := redis.ParseURL(conn)

	if err != nil {
		logger.Get().Fatal("Failed to extract redis connection url", zap.Error(err))
	}

	client := redis.NewClient(opt)

	if err = client.Ping(context.Background()).Err(); err != nil {
		logger.Get().Fatal("Failed to connect to redis", zap.Error(err))
	}

	connections.Store(conn, client)
	return client
}
