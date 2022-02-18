package ad

import (
	"src/ads/domain"
	"context"
	"os"
	"regexp"
	"strings"
)

type Cache struct {
	Data map[string]domain.Ad
}

func NewCache() *Cache {
	cache := Cache{Data: make(map[string]domain.Ad)}

	// create default ads
	if os.Getenv("ENV") == "dev" {
		cache.Create(nil, domain.Ad{
			ID:          "ad1",
			Title:       "Something",
			Description: "The description",
			Price:       1,
			SellerID:    "1",
		})
	}

	return &cache
}

func (mem *Cache) Create(_ context.Context, newAd domain.Ad) error {

	if _, ok := mem.Data[newAd.ID]; ok {
		return domain.ErrAlreadyExists
	}
	if newAd.Price <= 0 {
		return domain.ErrInvalidAdValues
	}

	// must be unique
	for _, ad := range mem.Data {
		if ad.Title == newAd.Title {
			return domain.ErrAdTitleAlreadyExists
		}
	}
	mem.Data[newAd.ID] = newAd

	return nil
}

func (mem *Cache) Update(ctx context.Context, adID string, adData domain.Ad) error {
	oldAd, ok := mem.Data[adID]

	if !ok {
		return domain.ErrAdNotFound
	}

	sessionInfo, ok := ctx.Value("session").(domain.SessionInfo)
	if !ok {
		return domain.ErrCannotRetreiveSession
	}
	if sessionInfo.UserID != oldAd.SellerID && sessionInfo.Access < 1 {
		return domain.ErrAccessForbidden
	}
	if adData.Title != "" {
		oldAd.Title = adData.Title
	}
	if adData.Description != "" {
		oldAd.Description = adData.Description
	}
	if adData.Price >= 0 {
		oldAd.Price = adData.Price
	}
	if adData.Picture != "" {
		oldAd.Picture = adData.Picture
	}
	mem.Data[adID] = oldAd

	return nil
}

func hasMatchingKeywords(ad domain.Ad, keywords []string) bool {
	if len(keywords) == 0 {
		return true
	}
	keywordsRegexString := strings.Join(keywords, "|")
	matched, err := regexp.MatchString(keywordsRegexString, ad.Title)
	if matched && err == nil {
		return true
	}
	matched, err = regexp.MatchString(keywordsRegexString, ad.Description)
	if matched && err == nil {
		return true
	}
	return false
}

func hasMatchingSeller(ad domain.Ad, sellerID string) bool {
	if sellerID == "" {
		return true
	}
	return ad.SellerID == sellerID
}

func (mem *Cache) Search(ctx context.Context, sellerID string, keywords []string) (*[]domain.Ad, error) {
	ads := make([]domain.Ad, 0, len(mem.Data))
	for _, ad := range mem.Data {
		if hasMatchingKeywords(ad, keywords) && hasMatchingSeller(ad, sellerID) {
			ads = append(ads, ad)
		}
	}
	return &ads, nil
}

func (mem *Cache) SearchBySeller(ctx context.Context, sellerID string) (*[]domain.Ad, error) {
	ads := make([]domain.Ad, 0, len(mem.Data))

	for _, ad := range mem.Data {
		if ad.SellerID == sellerID {
			ads = append(ads, ad)
		}
	}
	return &ads, nil
}

func (mem *Cache) Get(ctx context.Context, adID string) (*domain.Ad, error) {
	sessionInfo, ok := ctx.Value("session").(domain.SessionInfo)

	if !ok {
		return nil, domain.ErrCannotRetreiveSession
	}
	ad, ok := mem.Data[adID]

	if !ok {
		return nil, domain.ErrAdNotFound
	}

	if sessionInfo.Access < 2 {
		ad.SellerID = ""
		return &ad, nil
	}
	return &ad, nil
}

func (mem *Cache) Delete(ctx context.Context, adID string) error {
	sessionInfo, ok := ctx.Value("session").(domain.SessionInfo)
	if !ok {
		return domain.ErrCannotRetreiveSession
	}
	add, err := mem.Get(ctx, adID)

	if err != nil {
		return err
	}
	if sessionInfo.UserID != add.SellerID && sessionInfo.Access < 1 {
		return domain.ErrAccessForbidden
	}
	if _, ok = mem.Data[adID]; !ok {
		return domain.ErrAdNotFound
	}
	delete(mem.Data, adID)

	return nil
}
