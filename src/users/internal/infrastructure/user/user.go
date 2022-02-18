package user

import (
    "src/users/domain"
    "context"
)

type Storage interface {
    Create(ctx context.Context, user domain.User) error
    Update(ctx context.Context, userID string, userData domain.User) error
    GetAll(ctx context.Context) (*[]domain.User, error)
    Get(ctx context.Context, ID string) (*domain.User, error)
    GetAccess(ctx context.Context, userID string) (int8, error)
    GetLogin(ctx context.Context, login string, password string) (string, error)
    Delete(ctx context.Context, ID string) error
}