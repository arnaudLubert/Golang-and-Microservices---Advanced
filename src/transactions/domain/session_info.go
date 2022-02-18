package domain

import "errors"

var (
	ErrOperationNotPermitted = errors.New("operation not permitted")
	ErrCannotRetrieveSession = errors.New("cannot retrieve session info")
	ErrAccessForbidden       = errors.New("access to this resource is forbidden")
)

type SessionInfo struct {
	UserID string `json:"user_id"`
	Access int8   `json:"access"`
}
