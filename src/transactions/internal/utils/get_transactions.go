package utils

import (
	"context"
	"src/transactions/domain"
	"src/transactions/internal/infrastructure/transaction"
)

type GetAllCmd func(ctx context.Context) (*[]domain.Transaction, error)

func GetAll(storage transaction.Storage) GetAllCmd {
	return func(ctx context.Context) (*[]domain.Transaction, error) {
		tsnx, err := storage.GetAll(ctx)

		if err != nil {
			return nil, err
		}
		return tsnx, nil
	}
}
