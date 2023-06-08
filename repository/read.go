package repository

import (
	"encoding/json"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/go-redis/redis/v8"
)

func (repo *repo) Get(q query.Query) (any, errors.Error) {
	key := q.(query.Filter).GetId().(string)
	value, err := repo.client.Get(q.GetContext(), key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, errors.NotFound(err)
		}
		return nil, errors.New(err)
	}

	model := q.(query.ModelQuery).GetModel()
	if err := json.Unmarshal([]byte(value), model); err != nil {
		return nil, errors.New(err)
	}
	return model, nil
}
