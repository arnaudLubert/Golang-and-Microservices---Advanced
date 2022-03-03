package utils

import (
	"context"
	"src/reservation/domain"
	//"src/reservation/internal/security"
	//"src/reservation/internal/infrastructure/reservation"

	uuid "github.com/satori/go.uuid"
)

type CreateReservationCmd func(ctx context.Context, reservation domain.Reservation) (string, error)

func CreateReservation(storage reservation.Storage) CreateReservationCmd {
	return func(ctx context.Context, reservation domain.Reservation) (string, error) {
		sessionInfo, _ := ctx.Value("session").(domain.SessionInfo)

		if (sessionInfo.Access < 1) {
			return "", domain.ErrOperationNotPermitted
		}

		var err error
		reservation.ID = uuid.String()

		if err = storage.Create(ctx, reservation); err != nil {
			return "", err
		}
		return reservation.ID, nil
	}
}
