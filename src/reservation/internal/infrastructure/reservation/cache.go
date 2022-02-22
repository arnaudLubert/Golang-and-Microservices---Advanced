package reservation

import (
	"context"
	"os"
	//"time"
	//"fmt"
	"src/reservation/domain"
	//"src/reservation/internal/security"
)

type Cache struct {
	Data map[string]domain.Reservation
}




func NewCache() *Cache {
	// layout := "1999-02-21"
	cache := Cache{Data: make(map[string]domain.Reservation)}

	// create default users
	if os.Getenv("ENV") == "dev" {

		// d1, err1 := time.Parse(layout, "2022-11-05")
		// d2, err2 := time.Parse(layout, "2022-11-07")
		// fmt.Println(d1)
		// if err1 != nil || err2 != nil {
		// 	fmt.Println("Error")
		// }

		cache.Create(nil, domain.Reservation{
			ID:				"1",
			Date_of_arrival:"05-11-2022",
			Number_night:	3,
			Number_people:	2,
			Status:			0,
			Ad_id:			"1",
			Owner_id:		"2",
			Customer_id:	"3",
		})
	}
	// if os.Getenv("ENV") == "dev" {
	// 	cache.Create(nil, domain.User{
	// 		ID:       "1",
	// 		Pseudo:    "client",
	// 		Email:    "client@domain.com",
	// 		Password: security.MD5("password"),
	// 		Access:   0,
	// 	})
	// 	cache.Create(nil, domain.User{
	// 		ID:       "2",
	// 		Pseudo:    "client2",
	// 		Email:    "client2@domain.com",
	// 		Password: security.MD5("password"),
	// 		Access:   0,
	// 	})
	// 	cache.Create(nil, domain.User{
	// 		ID:       "3",
	// 		Pseudo:    "admin",
	// 		Email:    "admin@domain.com",
	// 		Password: security.MD5("password"),
	// 		Access:   2,
	// 	})
	// }

	return &cache
}

func (mem *Cache) Create(_ context.Context, newReservation domain.Reservation) error {

	if _, ok := mem.Data[newReservation.ID]; ok {
		return domain.ErrReservationAlreadyExists
	}

	// // must be unique
	// for _, newReservation := range mem.Data {
	// 	if reservation.Pseudo == newUser.Pseudo {
	// 		return domain.ErrUserPseudoAlreadyExists
	// 	} else if user.Email == newUser.Email {
	// 		return domain.ErrUserEmailAlreadyExists
	// 	}
	// }
	mem.Data[newReservation.ID] = newReservation

	return nil
}

// func (mem *Cache) Update(_ context.Context, userID string, userData domain.User) error {
// 	oldUser, ok := mem.Data[userID]

// 	// if !ok {
// 	// 	return domain.ErrUserNotFound
// 	// }

// 	// if userData.Firstname != "" {
// 	// 	oldUser.Firstname = userData.Firstname
// 	// }
// 	// if userData.Lastname != "" {
// 	// 	oldUser.Lastname = userData.Lastname
// 	// }
// 	// if userData.Phone != "" {
// 	// 	oldUser.Phone = userData.Phone
// 	// }
// 	// if userData.Address.City != "" {
// 	// 	oldUser.Address.City = userData.Address.City
// 	// }
// 	// if userData.Address.ZipCode != "" {
// 	// 	oldUser.Address.ZipCode = userData.Address.ZipCode
// 	// }
// 	// if userData.Address.Street != "" {
// 	// 	oldUser.Address.Street = userData.Address.Street
// 	// }

// 	// // must be unique
// 	// if userData.Email != "" {
// 	// 	for key, user := range mem.Data {
// 	// 		if key != oldUser.ID && user.Email == userData.Email {
// 	// 			return domain.ErrUserEmailAlreadyExists
// 	// 		}
// 	// 	}
// 	// 	oldUser.Email = userData.Email
// 	// }
// 	// mem.Data[userID] = oldUser

// 	return nil
// }

// func (mem *Cache) GetAll(ctx context.Context) (*[]domain.User, error) {
// 	sessionInfo, ok := ctx.Value("session").(domain.SessionInfo)

// 	// if !ok {
// 	// 	return nil, domain.ErrCannotRetreiveSession
// 	// }
// 	// users := make([]domain.User, 0, len(mem.Data))

// 	// for _, user := range mem.Data {
// 	// 	if sessionInfo.Access > 0 || sessionInfo.UserID == user.ID {
// 	// 		users = append(users, user)
// 	// 	}
// 	// }
// 	//return &users, nil
// }

// func (mem *Cache) Get(ctx context.Context, userID string) (*domain.User, error) {
// 	sessionInfo, ok := ctx.Value("session").(domain.SessionInfo)

// 	// if !ok {
// 	// 	return nil, domain.ErrCannotRetreiveSession
// 	// }
// 	// user, ok := mem.Data[userID]

// 	// if !ok {
// 	// 	return nil, domain.ErrUserNotFound
// 	// }

// 	// if sessionInfo.UserID != userID && sessionInfo.Access < 1 {
// 	// 	user.Pseudo = ""
// 	// 	user.Password = ""

// 	// 	return &user, nil
// 	// }
// 	// return &user, nil
// }

// func (mem *Cache) GetAccess(_ context.Context, userID string) (int8, error) {
// 	user, ok := mem.Data[userID]

// 	// if !ok {
// 	// 	return -1, domain.ErrUserNotFound
// 	// }
// 	// return user.Access, nil
// }

// func (mem *Cache) GetLogin(_ context.Context, pseudo string, password string) (string, error) {

// 	// for _, user := range mem.Data {
// 	// 	if user.Pseudo == pseudo && user.Password == password {
// 	// 		return user.ID, nil
// 	// 	}
// 	// }
// 	// return "", domain.ErrUserNotFound
// }

// func (mem *Cache) Delete(_ context.Context, userID string) error {

// 	// if _, ok := mem.Data[userID]; !ok {
// 	// 	return domain.ErrUserNotFound
// 	// }
// 	// delete(mem.Data, userID)

// 	// return nil
// }
