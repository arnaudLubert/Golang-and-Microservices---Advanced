package utils

import (
	"context"
	"src/transactions/domain"
	"src/transactions/internal/infrastructure/transaction"
)

type UpdateCmd func(ctx context.Context, ID string, invoice domain.Invoice) error

func Update(storage transaction.Storage) UpdateCmd {
	return func(ctx context.Context, ID string, invoice domain.Invoice) error {
		tsn, err := storage.Get(ctx, ID)
		if tsn == nil {
			return domain.ErrNotFound
		}

		if invoice.Message != "" {
			tsn.Messages = append(tsn.Messages, invoice.Message)
		}
		if invoice.BidPrice != 0 {
			tsn.BidPrices = append(tsn.BidPrices, invoice.BidPrice)
		}

		if err = storage.Update(ctx, ID, *tsn); err != nil {
			return err
		}
		return nil
	}
}
