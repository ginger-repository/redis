package repository

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/go-redis/redis/v8"
)

func (repo *repo) Delete(q query.Query) errors.Error {
	key := q.(query.IdGetter).GetId().(string)
	count, err := repo.client.Del(q.GetContext(), key).Result()
	if err != nil {
		if err == redis.Nil {
			return errors.NotFound(err)
		}
		return errors.New(err)
	}
	if count == 0 {
		return errors.NotFound()
	}
	return nil
}
