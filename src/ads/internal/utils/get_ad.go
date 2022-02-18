package utils

import (
	"src/ads/domain"
	"src/ads/internal/infrastructure/ad"
	"context"
)

type GetAdCmd func(ctx context.Context, AdID string) (*domain.Ad, error)

func GetAd(storage ad.Storage) GetAdCmd {
	return func(ctx context.Context, adID string) (*domain.Ad, error) {
		ad, err := storage.Get(ctx, adID)

		if err != nil {
			return nil, err
		}
		return ad, nil
	}
}
