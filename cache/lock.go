package cache

import (
	"context"
	"time"

	"github.com/ginger-core/errors"
)

func (db *db) Lock(ctx context.Context,
	key string, expiration time.Duration) errors.Error {
	c := db.getClient()
	r, err := c.SetNX(ctx, key, "", expiration).Result()
	if err != nil {
		return errors.New(err)
	}
	if !r {
		return errors.DefaultDuplicateError
	}
	return nil
}
