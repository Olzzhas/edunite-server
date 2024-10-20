package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type KeycloakClient struct {
	BaseURL string
}

func NewKeycloakClient(baseURL string) *KeycloakClient {
	return &KeycloakClient{BaseURL: baseURL}
}

func (kc *KeycloakClient) Login(email, password string) (string, error) {
	data := map[string]string{
		"username":   email,
		"password":   password,
		"grant_type": "password",
		"client_id":  "auth-client",
	}

	body, _ := json.Marshal(data)
	resp, err := http.Post(fmt.Sprintf("%s/realms/edunite/protocol/openid-connect/token", kc.BaseURL),
		"application/x-www-form-urlencoded", bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	token, ok := result["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("failed to retrieve access token")
	}

	return token, nil
}
