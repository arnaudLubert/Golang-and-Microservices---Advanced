package utils

import (
	"context"
	"src/transactions/domain"
	"src/transactions/internal/infrastructure/transaction"
)

type GetCmd func(ctx context.Context, ID string) (*domain.Transaction, error)

func Get(storage transaction.Storage) GetCmd {
	return func(ctx context.Context, ID string) (*domain.Transaction, error) {
		tsn, err := storage.Get(ctx, ID)

		if err != nil {
			return nil, err
		}
		return tsn, nil
	}
}
