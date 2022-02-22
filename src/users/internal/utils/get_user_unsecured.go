package utils

import (
    "src/users/internal/infrastructure/user"
    "src/users/domain"
    "context"
)

func GetUserUnsecured(storage user.Storage) GetUserCmd {
    return func(ctx context.Context, userID string) (*domain.User, error) {
        user, err := storage.GetUserUnsecured(ctx, userID)

        if err != nil {
            return nil, err
        }
        return user, nil
    }
}