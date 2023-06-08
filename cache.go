package redis

import (
	"github.com/ginger-core/repository"
	"github.com/ginger-repository/redis/cache"
	r "github.com/ginger-repository/redis/repository"
)

func NewCache(redis r.Repository) repository.Cache {
	return cache.New(redis)
}
