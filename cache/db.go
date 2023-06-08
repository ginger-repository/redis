package cache

import (
	"github.com/ginger-core/repository"
	r "github.com/ginger-repository/redis/repository"
)

type db struct {
	redis r.Repository
}

func New(redis r.Repository) repository.Cache {
	db := &db{
		redis: redis,
	}
	return db
}
