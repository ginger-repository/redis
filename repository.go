package redis

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-repository/redis/repository"
)

func NewRepository(registry registry.Registry) repository.Repository {
	return repository.New(registry)
}
