package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
	"src/transactions/domain"
	"src/transactions/internal/conf"
)

func handleErrors(rw http.ResponseWriter, err error) {
	switch err {
	case domain.ErrNotFound, domain.ErrAdNotFound, domain.ErrCannotRetrieveSeller:
		http.Error(rw, err.Error(), http.StatusNotFound)
	case domain.ErrAlreadyExists:
		http.Error(rw, err.Error(), http.StatusConflict)
	case domain.ErrAccessForbidden, domain.ErrOperationNotPermitted:
		http.Error(rw, err.Error(), http.StatusForbidden)
	default:
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

type GetAdCmd func(adID string) (*domain.Ad, error)

func GetAd(adsService conf.Service, authService conf.Service) GetAdCmd {
	return func(adID string) (*domain.Ad, error) {
		req, err := http.NewRequest("GET", adsService.Url+"/ad/"+adID, nil)
		if err != nil {
			return nil, err
		}

		adminToken, err := getAdminToken(authService)
		if err != nil {
			return nil, err
		}

		req.Header.Set("x-api-key", adsService.ApiKey)
		req.Header.Add("Authorization", "Bearer "+adminToken)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Host", adsService.Url)

		client := &http.Client{Timeout: time.Second * 20}
		res, err := client.Do(req)
		if err != nil {
			return nil, err
		}

		var ad domain.Ad
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		if err = json.Unmarshal(body, &ad); err != nil {
			return nil, err
		}
		/*decoder := json.NewDecoder(res.Body)
		decoder.Decode(&ad)
		if err = decoder.Decode(&ad); err != nil {
			return nil, domain.ErrAdNotFound
		}*/
		return &ad, nil
	}
}

type CreateLoginResponse struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Bearer    string    `json:"bearer"`
	ExpiresAt time.Time `json:"expires_at"`
}

func getAdminToken(authService conf.Service) (string, error) {
	var creds = []byte(`{
		"login": "admin",
		"password": "password"
	}`)
	req, err := http.NewRequest("POST", authService.Url+"/login", bytes.NewBuffer(creds))
	if err != nil {
		return "", err
	}
	req.Header.Set("x-api-key", authService.ApiKey)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Host", authService.Url)

	client := &http.Client{Timeout: time.Second * 20}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	if res.StatusCode < 200 || res.StatusCode > 299 {
		return "", domain.ErrCannotRetrieveSeller
	}

	var loginRes CreateLoginResponse
	decoder := json.NewDecoder(res.Body)
	if err = decoder.Decode(&loginRes); err != nil {
		return "", domain.ErrCannotRetrieveSeller
	}
	return loginRes.Bearer, nil
}
