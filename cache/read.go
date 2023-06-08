package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/ginger-core/errors"
	"github.com/go-redis/redis/v8"
)

func (db *db) ListKeys(ctx context.Context,
	pattern string) ([]string, errors.Error) {
	c := db.getClient()
	cmd := c.Keys(ctx, pattern)
	if err := cmd.Err(); err != nil {
		return nil, errors.New(err)
	}
	r, err := cmd.Result()
	if err != nil {
		return nil, errors.New(err)
	}
	return r, nil
}

func (db *db) Fetch(ctx context.Context,
	key string) ([]byte, errors.Error) {
	c := db.getClient()
	value, err := c.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, errors.NotFound(err)
		}
		return nil, errors.New(err)
	}
	return []byte(value), nil
}

func (db *db) Load(ctx context.Context,
	key string, resultRef any) errors.Error {
	c := db.getClient()
	value, err := c.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return errors.NotFound(err)
		}
		return errors.New(err)
	}
	if err := json.Unmarshal([]byte(value), resultRef); err != nil {
		return errors.New(err)
	}
	return nil
}

func (db *db) GetExpiration(ctx context.Context,
	key string) (time.Duration, errors.Error) {
	cmd := db.redis.GetClient().TTL(ctx, key)
	if err := cmd.Err(); err != nil {
		return 0, errors.New(err)
	}
	return cmd.Val(), nil
}
