package utils

import (
	"src/ads/domain"
	"src/ads/internal/infrastructure/ad"
	"context"
)

type SearchAdsCmd func(ctx context.Context, coordinates []float64, distance float64) (*[]domain.Ad, error)

func SearchAds(storage ad.Storage) SearchAdsCmd {
	return func(ctx context.Context, coordinates []float64, distance float64) (*[]domain.Ad, error) {
		ads, err := storage.Search(ctx, coordinates, distance)

		if err != nil {
			return nil, err
		}
		return ads, nil
	}
}
