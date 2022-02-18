package session

import (
    "src/authentication/domain"
    "context"
    "time"
)

type Cache struct {
    Data map[string]domain.Session
}

func NewCache() *Cache {
    return &Cache{Data: make(map[string]domain.Session)}
}

func (mem *Cache) Create(_ context.Context, session domain.Session) error {

    if _, ok := mem.Data[session.ID]; ok {
        return domain.ErrSessionAlreadyExists
    }
    session.CreatedAt = time.Now()
    session.Open = true
    mem.Data[session.ID] = session

    return nil
}

func (mem *Cache) Get(_ context.Context, sessionID string) (*domain.Session, error) {
    session, ok := mem.Data[sessionID]

    if !ok {
        return nil, domain.ErrSessionNotFound
    }
    return &session, nil
}

func (mem *Cache) GetFromBearer(_ context.Context, bearer string) (*domain.Session, error) {
    now := time.Now()

    for _, session := range(mem.Data) {
        if session.Bearer == bearer {
            if !session.Open || session.ExpiresAt.Before(now) {
                mem.Invalidate(nil, session.ID)
                return nil, domain.ErrSessionExpired
            } else {
                return &session, nil
            }
        }
    }
    return nil, domain.ErrSessionNotFound
}

func (mem *Cache) Invalidate(_ context.Context, sessionID string) error {
    session, ok := mem.Data[sessionID]

    if !ok {
        return domain.ErrSessionNotFound
    }
    session.Open = false
    mem.Data[sessionID] = session

    return nil
}

func (mem *Cache) Delete(_ context.Context, sessionID string) error {

    if _, ok := mem.Data[sessionID]; !ok {
        return domain.ErrSessionNotFound
    }
    delete(mem.Data, sessionID)

    return nil
}