package repository

import (
	"encoding/json"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
)

func (repo *repo) Create(q query.Query, entity any) errors.Error {
	key := q.(query.IdGetter).GetId().(string)
	bytes, err := json.Marshal(entity)
	if err != nil {
		return errors.New(err)
	}
	if err := repo.client.HSet(q.GetContext(), key, bytes).Err(); err != nil {
		return errors.New(err)
	}
	return nil
}
