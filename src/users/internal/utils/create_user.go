package utils

import (
	"context"
	"unicode"
	"src/users/domain"
	"src/users/internal/infrastructure/user"

	uuid "github.com/satori/go.uuid"
)

type CreateUserCmd func(ctx context.Context, user domain.User) (string, error)

func CreateUser(storage user.Storage) CreateUserCmd {
	return func(ctx context.Context, user domain.User) (string, error) {
		sessionInfo, _ := ctx.Value("session").(domain.SessionInfo)

		if (sessionInfo.Access < 1 || (sessionInfo.Access != 2 && sessionInfo.Access <= user.Access)) && user.Access > 0 {
			return "", domain.ErrOperationNotPermitted
		} else if !IsValidPassword(user.Password) {
			return "", domain.ErrUnsecuredPassword
		}
		uuid := uuid.NewV4()
/*
		uuid, err := uuid.NewV4()

		if err != nil {
			return "", err
		}*/
		var err error
		user.ID = uuid.String()

		if err = storage.Create(ctx, user); err != nil {
			return "", err
		}
		return user.ID, nil
	}
}

// check if password contains at least one digit, one uppercase and one lowercase
func IsValidPassword(str string) bool {
	var uppercase bool
	var lowercase bool
	var digit bool

	if len(str) < 8 {
		return false
	}

	for _, r := range str {
		if unicode.IsLetter(r) && unicode.IsUpper(r) {
			uppercase = true
		} else if unicode.IsLetter(r) && unicode.IsLower(r) {
			lowercase = true
		} else if unicode.IsDigit(r) {
			digit = true
		}
	}
	return (uppercase && lowercase && digit)
}