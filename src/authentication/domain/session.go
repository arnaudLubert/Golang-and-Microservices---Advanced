package domain

import (
    "errors"
    "time"
)

var (
    ErrSessionExpired       = errors.New("session has expired")
    ErrSessionNotFound      = errors.New("session not found")
    ErrSessionAlreadyExists = errors.New("session already exists")
    ErrGetUserAccess        = errors.New("an error occured while retreiving the user's access level")
)

type Session struct {
    ID          string      `json:"id"`
    UserID      string      `json:"user_id"`
    Open        bool        `json:"open"`
    Bearer      string      `json:"bearer"`
    CreatedAt   time.Time   `json:"created_at"`
    ExpiresAt   time.Time   `json:"expires_at"`
}