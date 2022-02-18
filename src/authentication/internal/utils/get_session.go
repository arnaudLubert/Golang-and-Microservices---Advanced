package utils

import (
    "src/authentication/internal/infrastructure/session"
    "src/authentication/internal/conf"
    "src/authentication/domain"
    "encoding/json"
    "net/http"
    "context"
    "time"
)

type UserAccessResponse struct {
    Access       int8      `json:"access"`
}

type GetSessionCmd func(ctx context.Context, bearer string) (*domain.SessionInfo, error)

func GetSession(storage session.Storage, usersService conf.Service) GetSessionCmd {
    return func(ctx context.Context, bearer string) (*domain.SessionInfo, error) {
        var sessionInfo domain.SessionInfo
        session, err := storage.GetFromBearer(ctx, bearer)

        if err != nil {
            return nil, err
        }
        sessionInfo.UserID = session.UserID
        userRequest, err := http.NewRequest("GET", usersService.Url + "/auth/access/" + session.UserID, nil)

        if err != nil {
            return nil, err
        }
        userRequest.Header.Set("x-api-key", usersService.ApiKey)
        userRequest.Header.Add("Content-Type", "application/json")
        userRequest.Header.Add("Host", usersService.Url)

        client := &http.Client{Timeout: time.Second * 20}
        userResponse, err := client.Do(userRequest)

        if err != nil {
            return nil, err
        }

        if userResponse.StatusCode < 200 || userResponse.StatusCode > 299 {
            userResponse.Body.Close()
            return nil, domain.ErrGetUserAccess
        }
        var userAccessResponse UserAccessResponse

        decoder := json.NewDecoder(userResponse.Body)

        if err = decoder.Decode(&userAccessResponse); err != nil {
            userResponse.Body.Close()
            return nil, err
        }
        userResponse.Body.Close()
        sessionInfo.Access = userAccessResponse.Access

        return &sessionInfo, nil
    }
}