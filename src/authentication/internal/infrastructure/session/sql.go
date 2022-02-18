package session

import (
    "src/authentication/domain"
    "context"
)

type SQL struct {
    Data map[string]domain.Session
}

func NewSQL() *SQL {
    return &SQL{Data: make(map[string]domain.Session)}
}

func (mem *SQL) Create(_ context.Context, session domain.Session) error {
    return domain.ErrSessionNotFound
}

func (mem *SQL) Get(_ context.Context, sessionID string) (*domain.Session, error) {
    return nil, domain.ErrSessionNotFound
}

func (mem *SQL) GetFromBearer(_ context.Context, bearer string) (*domain.Session, error) {
    return nil, domain.ErrSessionNotFound
}

func (mem *SQL) Invalidate(_ context.Context, sessionID string) error {
    return domain.ErrSessionNotFound
}

func (mem *SQL) Delete(_ context.Context, sessionID string) error {
    return domain.ErrSessionNotFound
}