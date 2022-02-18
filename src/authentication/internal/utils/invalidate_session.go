package utils

import (
    "src/authentication/internal/infrastructure/session"
    "context"
)

type InvalidateSessionCmd func(ctx context.Context, bearer string) error

func InvalidateSession(storage session.Storage) InvalidateSessionCmd {
    return func(ctx context.Context, bearer string) error {
        session, err := storage.GetFromBearer(ctx, bearer)

        if err != nil {
            return err
        }

        if err = storage.Delete(ctx, session.ID); err != nil {
            return err
        }
        return nil
    }
}