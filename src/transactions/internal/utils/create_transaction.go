package utils

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"src/transactions/domain"
	"src/transactions/internal/infrastructure/transaction"
)

type CreateCmd func(ctx context.Context, tsn domain.Transaction) (string, error)

func Create(storage transaction.Storage) CreateCmd {
	return func(ctx context.Context, tsn domain.Transaction) (string, error) {
		id, err := uuid.NewV4()

		if err != nil {
			return "", err
		}
		sessionInfo, ok := ctx.Value("session").(domain.SessionInfo)
		if !ok {
			return "", domain.ErrCannotRetrieveSession
		}

		tsn.ID = id.String()
		tsn.Status = "waiting"
		tsn.SenderID = sessionInfo.UserID

		if tsn.AdID == "" {
			return "", domain.ErrRequiredField
		}
		if tsn.AdSellerID == "" {
			return "", domain.ErrCannotRetrieveSeller
		}
		if len(tsn.BidPrices) < 1 {
			return "", domain.ErrRequiredField
		}

		if err = storage.Create(ctx, tsn); err != nil {
			return "", err
		}
		return tsn.ID, nil
	}
}
