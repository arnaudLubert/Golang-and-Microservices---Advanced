package handlers

import (
    "src/users/internal/utils"
    "src/users/domain"
    "encoding/json"
    "net/http"
    "fmt"
)

func GetUsersHandler(cmd utils.GetUsersCmd) http.HandlerFunc {
    return func(rw http.ResponseWriter, req *http.Request) {
        users, err := cmd(req.Context())

        if err != nil {
            switch err {
            case domain.ErrUserNotFound: http.Error(rw, err.Error(), http.StatusNotFound)
            case domain.ErrUserAlreadyExists: http.Error(rw, err.Error(), http.StatusConflict)
            case domain.ErrAccessForbidden, domain.ErrOperationNotPermitted:
                http.Error(rw, err.Error(), http.StatusForbidden)
            default: http.Error(rw, err.Error(), http.StatusInternalServerError)
            }
        } else {
            rw.Header().Set("Content-type", "application/json")
            json.NewEncoder(rw).Encode(users)
        }
    }
}