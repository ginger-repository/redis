package cache

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/go-redis/redis/v8"
)

func (db *db) Rename(ctx context.Context, key, new string) errors.Error {
	c := db.getClient()
	if err := c.Rename(ctx, key, new).Err(); err != nil {
		if err == redis.Nil {
			return errors.NotFound(err)
		}
		return errors.New(err)
	}
	return nil
}
