package utils

import (
	"context"
	"src/transactions/domain"
	"src/transactions/internal/infrastructure/transaction"
)

type ActionCmd func(ctx context.Context, ID string) error

func Action(storage transaction.Storage, status string) ActionCmd {
	return func(ctx context.Context, ID string) error {

		if ID == "" {
			return domain.ErrNotFound
		}

		return storage.UpdateStatus(ctx, ID, status)
	}
}
