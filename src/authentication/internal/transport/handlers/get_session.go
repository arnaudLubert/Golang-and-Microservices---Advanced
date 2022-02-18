package handlers

import (
    "src/authentication/internal/utils"
    "src/authentication/domain"
    "encoding/json"
    "net/http"
    "strings"
)

func GetSessionHandler(cmd utils.GetSessionCmd) http.HandlerFunc {
    return func(rw http.ResponseWriter, req *http.Request) {
        token := req.Header.Get("Authorization")
        tokenParts := strings.Split(token, "Bearer ")

        if len(tokenParts) != 2 || tokenParts[1] == "" {
            http.Error(rw, "bad token", http.StatusBadRequest)
            return
        }
        sessionInfo, err := cmd(req.Context(), tokenParts[1])

        if err != nil {
            switch err {
            case domain.ErrSessionNotFound, domain.ErrGetUserAccess, domain.ErrSessionExpired:
                http.Error(rw, err.Error(), http.StatusUnauthorized)
            default: http.Error(rw, err.Error(), http.StatusInternalServerError)
            }
        } else {
            rw.WriteHeader(http.StatusCreated)
            rw.Header().Set("Content-type", "application/json")
            json.NewEncoder(rw).Encode(sessionInfo)
        }
    }
}