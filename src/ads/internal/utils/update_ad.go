package utils

import (
	"src/ads/domain"
	"src/ads/internal/infrastructure/ad"
	"context"
)

type UpdateAdCmd func(ctx context.Context, adID string, adData domain.Ad) error

func UpdateAd(storage ad.Storage) UpdateAdCmd {
	return func(ctx context.Context, adID string, adData domain.Ad) error {
		err := storage.Update(ctx, adID, adData)

		if err != nil {
			return err
		}
		return nil
	}
}
