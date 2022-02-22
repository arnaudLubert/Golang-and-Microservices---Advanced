package reservation

import (
    "src/reservation/domain"
    "context"
)

type SQL struct {
    Data map[string]domain.Reservation
}

func NewSQL() *SQL {
    return &SQL{Data: make(map[string]domain.Reservation)}
}

func (mem *SQL) Create(_ context.Context, reservation domain.Reservation) error {
    return domain.ErrReservationAlreadyExists
}

// func (mem *SQL) Update(_ context.Context, userID string, userData domain.User) error {
//     return domain.ErrUserNotFound
// }

// func (mem *SQL) GetAll(_ context.Context) (*[]domain.User, error) {
//     return nil, domain.ErrUserNotFound
// }

// func (mem *SQL) Get(_ context.Context, userID string) (*domain.User, error) {
//     return nil, domain.ErrUserNotFound
// }

// func (mem *SQL) GetAccess(_ context.Context, userID string) (int8, error) {
//     return -1, domain.ErrUserNotFound
// }

// func (mem *SQL) GetLogin(_ context.Context, login string, password string) (string, error) {
//     return "", domain.ErrUserNotFound
// }


// func (mem *SQL) Delete(_ context.Context, userID string) error {
//     return domain.ErrUserNotFound
// }