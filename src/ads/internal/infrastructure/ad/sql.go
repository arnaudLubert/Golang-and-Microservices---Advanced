package ad

import (
	"src/ads/domain"
	"context"
)

type SQL struct {
	Data map[string]domain.Ad
}

func NewSQL() *SQL {
	return &SQL{Data: make(map[string]domain.Ad)}
}

func (mem *SQL) Create(_ context.Context, user domain.Ad) error {
	return domain.ErrAdNotFound
}

func (mem *SQL) Update(_ context.Context, userID string, userData domain.Ad) error {
	return domain.ErrAdNotFound
}

func (mem *SQL) UpdateBalance(_ context.Context, userID string, balanceInc float64) error {
	return domain.ErrAdNotFound
}

func (mem *SQL) Search(_ context.Context, coordinates []float64, distance float64) (*[]domain.Ad, error) {
	return nil, domain.ErrAdNotFound
}

func (mem *SQL) Get(_ context.Context, userID string) (*domain.Ad, error) {
	return nil, domain.ErrAdNotFound
}

func (mem *SQL) GetAccess(_ context.Context, userID string) (int8, error) {
	return -1, domain.ErrAdNotFound
}

func (mem *SQL) GetLogin(_ context.Context, login string, password string) (string, error) {
	return "", domain.ErrAdNotFound
}

func (mem *SQL) Delete(_ context.Context, userID string) error {
	return domain.ErrAdNotFound
}
