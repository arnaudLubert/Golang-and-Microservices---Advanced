package transaction

import (
	"context"
	"os"
	"src/transactions/domain"
)

type Cache struct {
	Data map[string]domain.Transaction
}

func NewCache() *Cache {
	cache := Cache{Data: make(map[string]domain.Transaction)}

	if os.Getenv("ENV") == "dev" {
		cache.Data["1"] = domain.Transaction{
			ID:         "1",
			Status:     "waiting",
			BidPrices:  []float64{3, 4},
			Messages:   []string{"hello", "hi"},
			SenderID:   "2",
			AdID:       "ad1",
			AdSellerID: "1",
		}
		cache.Data["2"] = domain.Transaction{
			ID:         "2",
			Status:     "accepted",
			BidPrices:  []float64{1, 2},
			Messages:   []string{"msg1", "msg2"},
			SenderID:   "1",
			AdID:       "2",
			AdSellerID: "2",
		}
		cache.Data["12"] = domain.Transaction{
			ID:         "12",
			Status:     "refused",
			BidPrices:  []float64{12, 12.1},
			Messages:   []string{"bon", "jour"},
			SenderID:   "2",
			AdID:       "1",
			AdSellerID: "1",
		}
		cache.Data["14"] = domain.Transaction{
			ID:         "14",
			Status:     "refused",
			BidPrices:  []float64{-1},
			Messages:   []string{"yo"},
			SenderID:   "-1",
			AdID:       "-1",
			AdSellerID: "-1",
		}
	}

	return &cache
}

func (mem *Cache) Create(ctx context.Context, tsn domain.Transaction) error {
	sessionInfo, ok := ctx.Value("session").(domain.SessionInfo)
	if !ok {
		return domain.ErrCannotRetrieveSession
	}

	tsn.SenderID = sessionInfo.UserID
	if _, ok := mem.Data[tsn.ID]; ok {
		return domain.ErrAlreadyExists
	}

	mem.Data[tsn.ID] = tsn

	return nil
}

func (mem *Cache) Update(ctx context.Context, tsnID string, tsn domain.Transaction) error {
	sessionInfo, ok := ctx.Value("session").(domain.SessionInfo)
	if !ok {
		return domain.ErrCannotRetrieveSession
	}

	oldTsn, ok := mem.Data[tsnID]
	if !ok {
		return domain.ErrNotFound
	}

	if sessionInfo.Access < 1 {
		if oldTsn.SenderID != sessionInfo.UserID {
			return domain.ErrAccessForbidden
		}
	}

	/*if tsn.Status != "" {
		oldTsn.Status = tsn.Status
	}
	if tsn.SenderID != "" {
		oldTsn.SenderID = tsn.SenderID
	}*/

	mem.Data[tsnID] = tsn

	return nil
}

func (mem *Cache) UpdateStatus(ctx context.Context, tsnID string, status string) error {
	sessionInfo, ok := ctx.Value("session").(domain.SessionInfo)
	if !ok {
		return domain.ErrCannotRetrieveSession
	}

	oldTsn, ok := mem.Data[tsnID]
	if !ok {
		return domain.ErrNotFound
	}

	if sessionInfo.Access != 2 {
		switch status {
		case "accepted", "refused":
			if sessionInfo.UserID != oldTsn.AdSellerID {
				return domain.ErrAccessForbidden
			}
		case "canceled":
			if sessionInfo.UserID != oldTsn.SenderID {
				return domain.ErrAccessForbidden
			}
		}
	}

	oldTsn.Status = status
	mem.Data[tsnID] = oldTsn

	return nil
}

/*func remove(s []domain.Transaction, i int) []domain.Transaction {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}*/

func (mem *Cache) GetAll(ctx context.Context) (*[]domain.Transaction, error) {
	sessionInfo, ok := ctx.Value("session").(domain.SessionInfo)
	if !ok {
		return nil, domain.ErrCannotRetrieveSession
	}

	tsnx := make([]domain.Transaction, 0, len(mem.Data))
	for _, tsn := range mem.Data {
		tsnx = append(tsnx, tsn)
	}

	if sessionInfo.Access < 1 {
		userTsnx := make([]domain.Transaction, 0, len(mem.Data))
		for _, tsn := range mem.Data {
			if sessionInfo.UserID == tsn.SenderID || sessionInfo.UserID == tsn.AdSellerID {
				userTsnx = append(userTsnx, tsn)
			}
		}
		return &userTsnx, nil
	}

	return &tsnx, nil
}

func (mem *Cache) Get(ctx context.Context, tsnID string) (*domain.Transaction, error) {
	sessionInfo, ok := ctx.Value("session").(domain.SessionInfo)
	if !ok {
		return nil, domain.ErrCannotRetrieveSession
	}

	tsn, ok := mem.Data[tsnID]
	if !ok { //|| sessionInfo.UserID != tsn.SenderID || sessionInfo.UserID != tsn.ReceiverID {
		return nil, domain.ErrNotFound
	}

	if sessionInfo.Access < 1 {
		if sessionInfo.UserID != tsn.SenderID && sessionInfo.UserID != tsn.AdSellerID {
			return nil, domain.ErrAccessForbidden
		}
	}

	return &tsn, nil
}

func (mem *Cache) Delete(ctx context.Context, tsnID string) error {
	sessionInfo, ok := ctx.Value("session").(domain.SessionInfo)
	if !ok {
		return domain.ErrCannotRetrieveSession
	}
	if sessionInfo.Access < 1 {
		return domain.ErrAccessForbidden
	}

	if _, ok := mem.Data[tsnID]; !ok {
		return domain.ErrNotFound
	}
	delete(mem.Data, tsnID)

	return nil
}
