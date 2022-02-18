package domain

type SessionInfo struct {
    UserID      string      `json:"user_id"`
    Access      int8        `json:"access"`
}