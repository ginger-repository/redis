package repository

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/repository"
	"github.com/go-redis/redis/v8"
)

type Repository interface {
	repository.Repository

	Initialize() errors.Error
	GetClient() *redis.Client
}
