package utils

import (
	"src/ads/domain"
	"src/ads/internal/infrastructure/ad"
	"context"
)

type DeleteAdCmd func(ctx context.Context, adID string) error

func DeleteAd(storage ad.Storage) DeleteAdCmd {
	return func(ctx context.Context, adID string) error {

		if adID == "" {
			return domain.ErrAdNotFound
		}
		return storage.Delete(ctx, adID)
	}
}
