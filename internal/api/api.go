package api

import (
	"bitexporter/internal/domain"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Executor struct {
	URL    string
	Auth   domain.Auth
	client *http.Client
}

func (e *Executor) Sync() (domain.Sync, error) {
	var sync domain.Sync
	req, err := http.NewRequest("GET", e.URL+"/api/sync", nil)
	if err != nil {
		return sync, err
	}
	req.Header.Add("Authorization", "Bearer "+e.Auth.AccessToken)
	resp, _ := e.client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return sync, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return sync, err
	}
	if err := json.Unmarshal(body, &sync); err != nil {
		return sync, err
	}
	return sync, nil
}

func (b *Executor) authenticate(clientId string, clientSecret string) error {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", clientId)
	data.Set("client_secret", clientSecret)
	data.Set("scope", "api")
	data.Set("scope", "api")
	data.Set("device_type", "cli")
	data.Set("device_identifier", "exporter_bot")
	data.Set("device_name", "exporter")

	r, err := http.NewRequest(
		"POST", b.URL+"/identity/connect/token",
		strings.NewReader(data.Encode()),
	)
	if err != nil {
		return err
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		return err
	}

	defer r.Body.Close()
	client := &http.Client{}
	resp, _ := client.Do(r)
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	json.Unmarshal(body, &b.Auth)

	return nil
}

func New(apiUrl string, clientId string, clientSecret string) (Executor, error) {
	var api Executor
	api.URL = apiUrl
	api.client = &http.Client{
		Timeout: 10 * time.Second,
	}
	err := api.authenticate(clientId, clientSecret)
	return api, err
}
