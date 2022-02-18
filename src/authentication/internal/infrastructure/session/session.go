package session

import (
    "src/authentication/domain"
    "context"
)

type Storage interface {
    Create(ctx context.Context, session domain.Session) error
    Get(ctx context.Context, sessionID string) (*domain.Session, error)
    GetFromBearer(ctx context.Context, bearer string) (*domain.Session, error)
    Invalidate(ctx context.Context, sessionID string) error
    Delete(ctx context.Context, sessionID string) error
}