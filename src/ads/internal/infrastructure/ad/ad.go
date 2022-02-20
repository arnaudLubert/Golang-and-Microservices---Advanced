package ad

import (
	"src/ads/domain"
	"context"
)

type Storage interface {
	Create(ctx context.Context, ad domain.Ad) error
	Update(ctx context.Context, adID string, adData domain.Ad) error
	Search(ctx context.Context, location []float64, distance float64) (*[]domain.Ad, error)
	Get(ctx context.Context, ID string) (*domain.Ad, error)
	Delete(ctx context.Context, ID string) error
}
