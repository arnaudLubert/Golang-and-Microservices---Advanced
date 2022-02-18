package user

import (
	"context"
	"os"
	"src/users/domain"
	"src/users/internal/security"
)

type Cache struct {
	Data map[string]domain.User
}

func NewCache() *Cache {
	cache := Cache{Data: make(map[string]domain.User)}

	// create default users
	if os.Getenv("ENV") == "dev" {
		cache.Create(nil, domain.User{
			ID:       "1",
			Login:    "client",
			Email:    "client@domain.com",
			Password: security.MD5("password"),
			Access:   0,
		})
		cache.Create(nil, domain.User{
			ID:       "2",
			Login:    "client2",
			Email:    "client2@domain.com",
			Password: security.MD5("password"),
			Access:   0,
		})
		cache.Create(nil, domain.User{
			ID:       "3",
			Login:    "admin",
			Email:    "admin@domain.com",
			Password: security.MD5("password"),
			Access:   2,
		})
	}

	return &cache
}

func (mem *Cache) Create(_ context.Context, newUser domain.User) error {

	if _, ok := mem.Data[newUser.ID]; ok {
		return domain.ErrUserAlreadyExists
	}

	// must be unique
	for _, user := range mem.Data {
		if user.Login == newUser.Login {
			return domain.ErrUserLoginAlreadyExists
		} else if user.Email == newUser.Email {
			return domain.ErrUserEmailAlreadyExists
		}
	}
	mem.Data[newUser.ID] = newUser

	return nil
}

func (mem *Cache) Update(_ context.Context, userID string, userData domain.User) error {
	oldUser, ok := mem.Data[userID]

	if !ok {
		return domain.ErrUserNotFound
	}

	if userData.Firstname != "" {
		oldUser.Firstname = userData.Firstname
	}
	if userData.Lastname != "" {
		oldUser.Lastname = userData.Lastname
	}
	if userData.Phone != "" {
		oldUser.Phone = userData.Phone
	}
	if userData.Address.City != "" {
		oldUser.Address.City = userData.Address.City
	}
	if userData.Address.ZipCode != "" {
		oldUser.Address.ZipCode = userData.Address.ZipCode
	}
	if userData.Address.Street != "" {
		oldUser.Address.Street = userData.Address.Street
	}

	// must be unique
	if userData.Email != "" {
		for key, user := range mem.Data {
			if key != oldUser.ID && user.Email == userData.Email {
				return domain.ErrUserEmailAlreadyExists
			}
		}
		oldUser.Email = userData.Email
	}
	mem.Data[userID] = oldUser

	return nil
}

func (mem *Cache) GetAll(ctx context.Context) (*[]domain.User, error) {
	sessionInfo, ok := ctx.Value("session").(domain.SessionInfo)

	if !ok {
		return nil, domain.ErrCannotRetreiveSession
	}
	users := make([]domain.User, 0, len(mem.Data))

	for _, user := range mem.Data {
		if sessionInfo.Access > 0 || sessionInfo.UserID == user.ID {
			users = append(users, user)
		}
	}
	return &users, nil
}

func (mem *Cache) Get(ctx context.Context, userID string) (*domain.User, error) {
	sessionInfo, ok := ctx.Value("session").(domain.SessionInfo)

	if !ok {
		return nil, domain.ErrCannotRetreiveSession
	}
	user, ok := mem.Data[userID]

	if !ok {
		return nil, domain.ErrUserNotFound
	}

	if sessionInfo.UserID != userID && sessionInfo.Access < 1 {
		user.Login = ""
		user.Password = ""

		return &user, nil
	}
	return &user, nil
}

func (mem *Cache) GetAccess(_ context.Context, userID string) (int8, error) {
	user, ok := mem.Data[userID]

	if !ok {
		return -1, domain.ErrUserNotFound
	}
	return user.Access, nil
}

func (mem *Cache) GetLogin(_ context.Context, login string, password string) (string, error) {

	for _, user := range mem.Data {
		if user.Login == login && user.Password == password {
			return user.ID, nil
		}
	}
	return "", domain.ErrUserNotFound
}

func (mem *Cache) Delete(_ context.Context, userID string) error {

	if _, ok := mem.Data[userID]; !ok {
		return domain.ErrUserNotFound
	}
	delete(mem.Data, userID)

	return nil
}
