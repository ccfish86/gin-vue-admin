package persist

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"context"
	"errors"
	"time"

	"github.com/casbin/casbin/v2/persist/cache"
	"github.com/go-redis/redis/v8"
)

// CasbinRedisStore the Redis cache for policy storage.
type CasbinRedisStore struct {
	Context context.Context
	PreKey  string
}

// NewCasbinCache initializes an instance of RedisStore.
func NewCasbinCache() *CasbinRedisStore {
	return &CasbinRedisStore{
		PreKey:  "casbin:",
		Context: context.Background(),
	}
}

// Get returns cached value by given key.
func (redisStore *CasbinRedisStore) Get(key string) (bool, error) {
	val, err := global.GVA_REDIS.Get(redisStore.Context, redisStore.PreKey+key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return false, cache.ErrNoSuchKey
		}
		return false, err
	}

	return val == "1", nil
}

// Set sets cached value with key and expire time.
func (redisStore *CasbinRedisStore) Set(key string, value bool, extra ...interface{}) error {
	ttl := time.Hour
	if len(extra) > 0 {
		ttl = extra[0].(time.Duration) * time.Second
	}

	return global.GVA_REDIS.Set(redisStore.Context, redisStore.PreKey+key, "1", ttl).Err()
}

// Delete deletes cached value by given key.
func (redisStore *CasbinRedisStore) Delete(key string) error {
	return global.GVA_REDIS.Del(redisStore.Context, redisStore.PreKey+key).Err()
}

// Clear clears all cached data.
func (redisStore *CasbinRedisStore) Clear() error {
	keys, cursor, err := global.GVA_REDIS.Scan(redisStore.Context, 0, redisStore.PreKey+"*", 1000).Result()
	if err != nil {
		return err
	}
	for cursor != 0 {
		if err := global.GVA_REDIS.Del(redisStore.Context, keys...).Err(); err != nil {
			return err
		}
		keys, cursor, err = global.GVA_REDIS.Scan(redisStore.Context, 0, redisStore.PreKey+"*", 1000).Result()
		if err != nil {
			return err
		}
	}
	return nil
}
