package cache

import (
	"context"
	"encoding/json"

	"github.com/ginger-core/errors"
	"github.com/go-redis/redis/v8"
)

func (db *db) SetItem(ctx context.Context,
	key string, id string, value any) errors.Error {
	c := db.getClient()
	err := c.HSet(ctx, key, id, value).Err()
	if err != nil {
		return errors.New(err)
	}
	return nil
}

func (db *db) MarshalSetItem(ctx context.Context,
	key string, id string, value any) errors.Error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return errors.New(err)
	}
	return db.SetItem(ctx, key, id, bytes)
}

func (db *db) GetItem(ctx context.Context,
	key string, id string, resultRef any) errors.Error {
	c := db.getClient()
	result, err := c.HGet(ctx, key, id).Result()
	if err != nil {
		if err == redis.Nil {
			return errors.NotFound(err)
		}
		return errors.New(err)
	}
	if err := json.Unmarshal([]byte(result), resultRef); err != nil {
		return errors.New(err)
	}
	return nil
}

func (db *db) UnsetItem(ctx context.Context,
	key string, id string) errors.Error {
	c := db.getClient()
	r, err := c.HDel(ctx, key, id).Result()
	if err != nil {
		return errors.New(err)
	}
	if r == 0 {
		return errors.NotFound()
	}
	return nil
}
