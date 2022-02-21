package utils

import (
	"src/ads/domain"
	"src/ads/internal/conf"
	"src/ads/internal/infrastructure/ad"
	"context"
	"net/http"
	"io/ioutil"
	"time"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

type CreateAdCmd func(ctx context.Context, ad domain.Ad) (string, error)

func CreateAd(usersService conf.Service, storage ad.Storage) CreateAdCmd {
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

		ibanRequest, err := http.NewRequest("GET", usersService.Url + "/iban/" + sessionInfo.UserID, nil)
fmt.Println(sessionInfo.UserID)
		if err != nil {
			return "", domain.ErrCannotRetreiveIBAN
		}
		ibanRequest.Header.Set("x-api-key", usersService.ApiKey)
		ibanRequest.Header.Add("Content-Type", "application/json")
		ibanRequest.Header.Add("Host", usersService.Url)

		client := &http.Client{Timeout: time.Second * 20}
		IbanResponse, err := client.Do(ibanRequest)

		if err != nil {
			return "", domain.ErrCannotRetreiveIBAN
		}

		if IbanResponse.StatusCode >= 500 {
			IbanResponse.Body.Close()
			return "", domain.ErrCannotRetreiveIBAN
		} else if IbanResponse.StatusCode < 200 || IbanResponse.StatusCode > 299 {
			IbanResponse.Body.Close()
			return "", domain.ErrCannotRetreiveIBAN
		}
		body, err := ioutil.ReadAll(IbanResponse.Body)

		if err != nil {
			return "", domain.ErrCannotRetreiveIBAN
		}
		IbanResponse.Body.Close()

		if string(body) == "" {
			return "", domain.ErrMissingIBAN
		}

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
