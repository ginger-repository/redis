package cache

import (
	"context"
	"time"

	"github.com/ginger-core/errors"
)

func (db *db) Expire(ctx context.Context,
	key string, expiration time.Duration) errors.Error {
	c := db.getClient()
	err := c.Expire(ctx, key, expiration).Err()
	if err != nil {
		return errors.New(err)
	}
	return nil
}
