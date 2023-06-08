package repository

import (
	"context"

	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/errors"
	"github.com/ginger-core/repository"
	"github.com/go-redis/redis/v8"
)

type repo struct {
	repository.Repository

	config config

	client *redis.Client
}

func New(registry registry.Registry) Repository {
	repo := &repo{}
	if err := registry.Unmarshal(&repo.config); err != nil {
		panic(err)
	}
	repo.config.initialize()

	opts := new(redis.Options)
	if err := registry.Unmarshal(opts); err != nil {
		panic(err)
	}
	if repo.config.Timeout > 0 {
		opts.DialTimeout = repo.config.Timeout
		opts.ReadTimeout = repo.config.Timeout
		opts.WriteTimeout = repo.config.Timeout
	}
	repo.client = redis.NewClient(opts)
	return repo
}

func (repo *repo) Initialize() errors.Error {
	ctx, cancel := context.WithTimeout(context.Background(), repo.config.Timeout)
	defer cancel()

	_, errConnect := repo.client.Ping(ctx).Result()
	if errConnect != nil {
		return errors.New(errConnect)
	}
	return nil
}
