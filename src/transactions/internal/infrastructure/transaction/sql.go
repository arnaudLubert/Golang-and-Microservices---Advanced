package transaction

import (
	"context"
	"src/transactions/domain"
)

type SQL struct {
	Data map[string]domain.Transaction
}

func NewSQL() *SQL {
	return &SQL{Data: make(map[string]domain.Transaction)}
}

func (mem *SQL) Create(_ context.Context, tsn domain.Transaction) error {
	return domain.ErrNotFound
}

func (mem *SQL) Update(_ context.Context, tsnID string, tsn domain.Transaction) error {
	return domain.ErrNotFound
}

func (mem *SQL) GetAll(_ context.Context) (*[]domain.Transaction, error) {
	return nil, domain.ErrNotFound
}

func (mem *SQL) Get(_ context.Context, tsnID string) (*domain.Transaction, error) {
	return nil, domain.ErrNotFound
}

func (mem *SQL) Delete(_ context.Context, tsnID string) error {
	return domain.ErrNotFound
}
