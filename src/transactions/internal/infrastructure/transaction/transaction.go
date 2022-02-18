package transaction

import (
	"context"
	"src/transactions/domain"
)

type Storage interface {
	Create(ctx context.Context, tsn domain.Transaction) error
	Update(ctx context.Context, tsnID string, tsn domain.Transaction) error
	GetAll(ctx context.Context) (*[]domain.Transaction, error)
	Get(ctx context.Context, tsnID string) (*domain.Transaction, error)
	Delete(ctx context.Context, tsnID string) error
	UpdateStatus(ctx context.Context, tsnID string, action string) error
}
