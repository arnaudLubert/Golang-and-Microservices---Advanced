package ad

import (
	"src/ads/domain"
	"strings"
	"context"
	"regexp"
	"math"
	"os"
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
			Title:       "Apartment",
			Capacity:    2,
			Description: "The description",
			Price:       70,
			SellerID:    "1",
			Pictures:    []string{"https://media.istockphoto.com/photos/business-district-with-tall-modern-office-buildings-at-amsterdam-zuid-picture-id1313787429?b=1&k=20&m=1313787429&s=170667a&w=0&h=_AK9s-kRHxiw3RG4ToobxwVI4bzjgHaq3TeOtTsRq6s=", "https://media.istockphoto.com/photos/modern-elegant-kitchen-stock-photo-picture-id1297586166?b=1&k=20&m=1297586166&s=170667a&w=0&h=Ka-3OYiTlbCiwCJhoXeTqRewh3DI4qfSh1B0baJMcCk="},
			Location: 	 domain.Location{"Nice", "06000", "Epitech, Masséna", 43.6958037, 7.2697099},
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
	if adData.Capacity >= 0 {
		oldAd.Capacity = adData.Capacity
	}

	if adData.Location.City != "" || adData.Location.ZipCode != "" || adData.Location.Street != "" {
		if adData.Location.City != "" {
			oldAd.Location.City = oldAd.Location.City
		}
		if adData.Location.ZipCode != "" {
			oldAd.Location.ZipCode = oldAd.Location.ZipCode
		}
		if adData.Location.Street != "" {
			oldAd.Location.Street = oldAd.Location.Street
		}
		// update coorinates
	}
	if len(adData.Pictures) != 0 {
		oldAd.Pictures = adData.Pictures
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

func (mem *Cache) Search(ctx context.Context, coordinates []float64, distance float64) (*[]domain.Ad, error) {
	ads := make([]domain.Ad, 0, len(mem.Data))
	for _, ad := range mem.Data {
		if withinLocation(ad, coordinates, distance) {
			ads = append(ads, ad)
		}
	}
	return &ads, nil
}

func withinLocation(ad domain.Ad, coordinates []float64, targetDistance float64) bool {
	a1 := ad.Location.Latitude * math.Pi / 180
	a2 := coordinates[0] * math.Pi / 180
	a3 := (coordinates[0] - ad.Location.Latitude) * math.Pi / 180
	a4 := (coordinates[1] - ad.Location.Longitude) * math.Pi / 180
	a := math.Sin(a3/2) * math.Sin(a3/2) + math.Cos(a1) * math.Cos(a2) * math.Sin(a4/2) * math.Sin(a4/2);
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance := 6371000 * c // in metres

	return (distance < targetDistance)
}

/*
func (mem *Cache) Search(ctx context.Context, sellerID string, keywords []string) (*[]domain.Ad, error) {
	ads := make([]domain.Ad, 0, len(mem.Data))
	for _, ad := range mem.Data {
		if hasMatchingKeywords(ad, keywords) && hasMatchingSeller(ad, sellerID) {
			ads = append(ads, ad)
		}
	}
	return &ads, nil
}
*/

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
