package handlers

import (
    "src/authentication/internal/utils"
    "src/authentication/internal/conf"
    "src/authentication/domain"
    "encoding/json"
    "net/http"
    "bytes"
    "time"
)

type LoginRequest struct {
    Login       string      `json:"login"`
    Password    string      `json:"password"`
}

type CreateLoginResponse struct {
    ID          string      `json:"id"`
    UserID      string      `json:"user_id"`
    Bearer      string      `json:"bearer"`
    ExpiresAt   time.Time   `json:"expires_at"`
}

type GetUserLoginResponse struct {
    UserID      string          `json:"user_id"`
}

func LoginHandler(cmd utils.CreateSessionCmd, usersService conf.Service) http.HandlerFunc {
    return func(rw http.ResponseWriter, req *http.Request) {
        var loginReq LoginRequest

        decoder := json.NewDecoder(req.Body)
        err := decoder.Decode(&loginReq)

        if err != nil {
            http.Error(rw, err.Error(), http.StatusBadRequest)
            return
        }

        if loginReq.Login == "" || loginReq.Password == "" {
            http.Error(rw, "missing login nor password", http.StatusBadRequest)
            return
        }
        jsonLoginReq, err := json.Marshal(loginReq)

        if err != nil {
            http.Error(rw, err.Error(), http.StatusInternalServerError)
            return
        }
        userRequest, err := http.NewRequest("POST", usersService.Url + "/auth/login", bytes.NewBuffer(jsonLoginReq))

        if err != nil {
            http.Error(rw, err.Error(), http.StatusInternalServerError)
            return
        }
        userRequest.Header.Set("x-api-key", usersService.ApiKey)
        userRequest.Header.Add("Content-Type", "application/json")
        userRequest.Header.Add("Host", usersService.Url)

        client := &http.Client{Timeout: time.Second * 20}
        userResponse, err := client.Do(userRequest)

        if err != nil {
            http.Error(rw, err.Error(), http.StatusInternalServerError)
            return
        }

        if userResponse.StatusCode >= 500 {
            rw.WriteHeader(http.StatusInternalServerError)
            userResponse.Body.Close()
            return
        } else if userResponse.StatusCode < 200 || userResponse.StatusCode > 299 {
            rw.WriteHeader(http.StatusForbidden)
            userResponse.Body.Close()
            return
        }
        var loginResponse GetUserLoginResponse

        decoder = json.NewDecoder(userResponse.Body)

        if err = decoder.Decode(&loginResponse); err != nil {
            http.Error(rw, err.Error(), http.StatusInternalServerError)
            userResponse.Body.Close()
            return
        }
        userResponse.Body.Close()

        session, err := cmd(req.Context(), loginResponse.UserID)

        if err != nil {
            switch err {
            case domain.ErrSessionAlreadyExists: http.Error(rw, err.Error(), http.StatusConflict)
            default: http.Error(rw, err.Error(), http.StatusInternalServerError)
            }
        } else {
            rw.WriteHeader(http.StatusCreated)
            rw.Header().Set("Content-type", "application/json")
            json.NewEncoder(rw).Encode(CreateLoginResponse{
                ID: session.ID,
                UserID: session.UserID,
                Bearer: session.Bearer,
                ExpiresAt: session.ExpiresAt,
            })
        }
    }
}