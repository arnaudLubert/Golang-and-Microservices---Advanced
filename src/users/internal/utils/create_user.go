package utils

import (
	"context"
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
