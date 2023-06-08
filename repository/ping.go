package repository

import (
	"context"

	"github.com/ginger-core/errors"
)

func (repo *repo) Ping(ctx context.Context) errors.Error {
	r := repo.client.Ping(ctx)
	if err := r.Err(); err != nil {
		return errors.New(err).WithTrace("client.Ping")
	}
	return nil
}
