package cache

import (
	"context"

	"github.com/ginger-core/errors"
)

func (db *db) Delete(ctx context.Context, key string) errors.Error {
	c := db.getClient()
	err := c.Del(ctx, key).Err()
	if err != nil {
		return errors.New(err)
	}
	return nil
}
