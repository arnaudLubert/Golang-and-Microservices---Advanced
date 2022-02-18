package domain

import "errors"

var (
    ErrOperationNotPermitted = errors.New("operation not permitted")
    ErrCannotRetreiveSession = errors.New("cannot retreive session info")
    ErrAccessForbidden       = errors.New("access to this ressource is forbidden")
)

type SessionInfo struct {
    UserID    string      `json:"user_id"`
    Access    int8        `json:"access"`
}