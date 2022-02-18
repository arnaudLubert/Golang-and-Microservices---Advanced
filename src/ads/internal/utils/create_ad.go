package utils

import (
	"src/ads/domain"
	"src/ads/internal/infrastructure/ad"
	"context"

	uuid "github.com/satori/go.uuid"
)

type CreateAdCmd func(ctx context.Context, ad domain.Ad) (string, error)

func CreateAd(storage ad.Storage) CreateAdCmd {
	return func(ctx context.Context, ad domain.Ad) (string, error) {
		uuid, err := uuid.NewV4()

		if err != nil {
			return "", err
		}
		sessionInfo, ok := ctx.Value("session").(domain.SessionInfo)
		if !ok {
			return "", domain.ErrCannotRetreiveSession
		}
		ad.ID = uuid.String()
		ad.SellerID = sessionInfo.UserID

		if ad.Title == "" {
			return "", domain.ErrAdNoTitle
		}

		if ad.Description == "" {
			return "", domain.ErrAdNoDescription
		}

		if ad.Price < 0 {
			return "", domain.ErrAdNoPrice
		}

		if ad.Capacity <= 0 {
			return "", domain.ErrAdNoCapacity
		}

		if len(ad.Pictures) == 0 {
			return "", domain.ErrAdNoPicture
		}

		if ad.Location.Street == "" || ad.Location.City == "" || ad.Location.ZipCode == "" {
			return "", domain.ErrAdNoLocation
		}

		if err = storage.Create(ctx, ad); err != nil {
			return "", err
		}
		return ad.ID, nil
	}
}
