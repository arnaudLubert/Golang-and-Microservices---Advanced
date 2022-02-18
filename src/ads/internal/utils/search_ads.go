package utils

import (
	"src/ads/domain"
	"src/ads/internal/infrastructure/ad"
	"context"
)

type SearchAdsCmd func(ctx context.Context, sellerID string, keywords []string) (*[]domain.Ad, error)

func SearchAds(storage ad.Storage) SearchAdsCmd {
	return func(ctx context.Context, sellerID string, keywords []string) (*[]domain.Ad, error) {
		ads, err := storage.Search(ctx, sellerID, keywords)

		if err != nil {
			return nil, err
		}
		return ads, nil
	}
}
