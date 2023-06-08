package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/ginger-core/errors"
)

func (db *db) Store(ctx context.Context,
	key string, value any, expiration ...time.Duration) errors.Error {
	c := db.getClient()

	var exp time.Duration = 0
	if len(expiration) > 0 {
		exp = expiration[0]
	}

	_, err := c.Set(ctx, key, value, exp).Result()
	if err != nil {
		return errors.New(err)
	}
	return nil
}

func (db *db) MarshalStore(ctx context.Context,
	key string, value any, expiration ...time.Duration) errors.Error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return errors.New(err)
	}
	return db.Store(ctx, key, bytes, expiration...)
}
