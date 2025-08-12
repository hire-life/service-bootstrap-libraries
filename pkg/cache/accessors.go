package cache

import (
	"context"
	"github.com/goccy/go-json"
	"github.com/hire-life/service-bootstrap-libraries/pkg/logger"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"time"
)

func Has(ctx context.Context, key string, r *redis.Client) bool {
	return r.Exists(ctx, key).Val() == 1
}

func Delete(ctx context.Context, key string, r *redis.Client) bool {
	return r.Del(ctx, key).Val() == 1
}

func Write(ctx context.Context, key string, val any, ttl time.Duration, r *redis.Client) {
	bytes, err := json.Marshal(val)

	if err != nil {
		logger.Get().Fatal("redis: Failed to marshal value", zap.Any("value", val), zap.Error(err))
	}

	WriteRaw(ctx, key, string(bytes), ttl, r)
}

func WriteRaw(ctx context.Context, key string, val string, ttl time.Duration, r *redis.Client) {
	err := r.Set(ctx, key, val, ttl).Err()

	if err != nil {
		logger.Get().Fatal("redis: Failed to set value", zap.Any("value", val), zap.Error(err))
	}
}

func Read[T any](ctx context.Context, key string, r *redis.Client) *T {
	val := ReadRaw(ctx, key, r)

	res := new(T)
	err := json.Unmarshal([]byte(val), res)

	if err != nil {
		logger.Get().Error("redis: Failed to unmarshal value", zap.Any("value", val), zap.Error(err))
		return nil
	}

	return res
}

func ReadRaw(ctx context.Context, key string, r *redis.Client) string {
	val, err := r.Get(ctx, key).Result()
	if err != nil {
		logger.Get().Fatal("redis: Failed to get value", zap.Any("value", val), zap.Error(err))
	}

	return val
}
