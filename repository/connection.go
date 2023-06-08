package repository

import (
	"github.com/ginger-core/query"
	"github.com/go-redis/redis/v8"
)

func (repo *repo) GetDB(q query.Query) any {
	return repo.GetClient()
}

func (repo *repo) GetClient() *redis.Client {
	return repo.client
}
