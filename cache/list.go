package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/repository"
	"github.com/go-redis/redis/v8"
)

func (db *db) insertItem(ctx context.Context,
	key string, score uint64, value any) errors.Error {
	c := db.getClient()
	err := c.ZAdd(ctx, key,
		&redis.Z{
			Score:  float64(score),
			Member: value,
		}).Err()
	if err != nil {
		return errors.New(err)
	}
	return nil
}

func (db *db) InsertItem(ctx context.Context,
	key string, value repository.Entity) errors.Error {
	return db.insertItem(ctx, key, value.GetId(), value)
}

func (db *db) MarshalInsertItem(ctx context.Context,
	key string, value repository.Entity) errors.Error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return errors.New(err)
	}
	return db.insertItem(ctx, key, value.GetId(), bytes)
}

func (db *db) Count(ctx context.Context, pattern string) (int, errors.Error) {
	c := db.getClient()
	keys, err := c.Keys(ctx, pattern).Result()
	if err != nil {
		return 0, errors.New(err)
	}
	return len(keys), nil
}

func (db *db) CountItems(ctx context.Context, key string) (int, errors.Error) {
	c := db.getClient()
	result, err := c.ZCount(ctx, key, "-inf", "+inf").Result()
	if err != nil {
		return 0, errors.New(err)
	}
	return int(result), nil
}

func (db *db) LoadItems(ctx context.Context,
	key string, page, pageSize int) ([]string, errors.Error) {
	skip := (page - 1) * pageSize
	c := db.getClient()
	result, err := c.ZRange(ctx, key,
		int64(skip), int64(skip+pageSize-1)).Result()
	if err != nil {
		return nil, errors.New(err)
	}
	return result, nil
}

func (db *db) LoadItem(ctx context.Context,
	key string, id int64, resultRef any) errors.Error {
	c := db.getClient()
	result, err := c.ZRange(ctx, key, id, id).Result()
	if err != nil {
		return errors.New(err)
	}
	if len(result) == 0 {
		return errors.NotFound()
	}
	if err := json.Unmarshal([]byte(result[0]), resultRef); err != nil {
		return errors.New(err)
	}
	return nil
}

func (db *db) DeleteItem(ctx context.Context,
	key string, id int64) errors.Error {
	err := db.getClient().
		ZRemRangeByScore(ctx, key,
			fmt.Sprint(id), fmt.Sprint(id)).Err()
	if err != nil {
		return errors.New(err)
	}
	return nil
}
