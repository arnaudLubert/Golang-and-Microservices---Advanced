package utils

import (
	"src/authentication/domain"
	"src/authentication/internal/infrastructure/session"
	"context"
	"time"

	uuid "github.com/satori/go.uuid"
)

type CreateSessionCmd func(ctx context.Context, userID string) (*domain.Session, error)

func CreateSession(storage session.Storage) CreateSessionCmd {
	return func(ctx context.Context, userID string) (*domain.Session, error) {
		var session domain.Session

		sessionUuid := uuid.NewV4()
		bearerUuid := uuid.NewV4()
		session.ID = sessionUuid.String()
		session.UserID = userID
		session.Bearer = bearerUuid.String()
		session.ExpiresAt = time.Now().Add(time.Hour * 2)

		if err := storage.Create(ctx, session); err != nil {
			return nil, err
		}
		return &session, nil
	}
}
