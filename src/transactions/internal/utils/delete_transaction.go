package utils

import (
	"context"
	"src/transactions/domain"
	"src/transactions/internal/infrastructure/transaction"
)

type DeleteCmd func(ctx context.Context, ID string) error

func Delete(storage transaction.Storage) DeleteCmd {
	return func(ctx context.Context, ID string) error {

		if ID == "" {
			return domain.ErrNotFound
		}
		return storage.Delete(ctx, ID)
	}
}
