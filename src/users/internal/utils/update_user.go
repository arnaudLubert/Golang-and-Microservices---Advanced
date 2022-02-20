package utils

import (
    "src/users/internal/infrastructure/user"
    "src/users/domain"
    "context"
)

type UpdateUserCmd func(ctx context.Context, userID string, userData domain.User) error

func UpdateUser(storage user.Storage) UpdateUserCmd {
    return func(ctx context.Context, userID string, userData domain.User) error {
        sessionInfo, ok := ctx.Value("session").(domain.SessionInfo)

        if !ok {
            return domain.ErrCannotRetreiveSession
        }

        if userID == "" {
            userID = sessionInfo.UserID
        }

        if sessionInfo.UserID != userID && sessionInfo.Access < 1 {
            return domain.ErrOperationNotPermitted
        }

        if userData.IBAN != "" && len(userData.IBAN) < 22 {
            return domain.ErrInvalidIBAN
        }
        err := storage.Update(ctx, userID, userData)

        if err != nil {
            return err
        }
        return nil
    }
}