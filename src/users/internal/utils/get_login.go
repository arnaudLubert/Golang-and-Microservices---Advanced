package utils

import (
    "src/users/internal/infrastructure/user"
    "context"
)

type GetLoginCmd func(ctx context.Context, login string, password string) (string, error)

func GetLogin(storage user.Storage) GetLoginCmd {
    return func(ctx context.Context, login string, password string) (string, error) {
        userID, err := storage.GetLogin(ctx, login, password)

        if err != nil {
            return "", err
        }
        return userID, nil
    }
}