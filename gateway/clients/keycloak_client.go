package clients

import (
	"context"
	"fmt"
	"github.com/Nerzal/gocloak/v13"
	"log"
)

type KeycloakClient struct {
	client       *gocloak.GoCloak
	token        *gocloak.JWT
	ctx          context.Context
	realm        string
	clientID     string
	clientSecret string
}

// NewKeycloakClient создает и инициализирует новый KeycloakClient
func NewKeycloakClient(baseURL, realm, clientID, clientSecret string) *KeycloakClient {
	client := gocloak.NewClient(baseURL)
	ctx := context.Background()
	return &KeycloakClient{
		client:       client,
		ctx:          ctx,
		realm:        realm,
		clientID:     clientID,
		clientSecret: clientSecret,
	}
}

// Login логинит пользователя и возвращает JWT токен
func (kc *KeycloakClient) Login(username, password string) (*gocloak.JWT, error) {
	token, err := kc.client.Login(kc.ctx, kc.clientID, kc.clientSecret, kc.realm, username, password)
	if err != nil {
		return nil, fmt.Errorf("failed to login user: %w", err)
	}
	log.Println("User logged in successfully")
	kc.token = token
	return token, nil
}

// GetUserInfo получает информацию о пользователе по токену
func (kc *KeycloakClient) GetUserInfo(accessToken string) (*gocloak.UserInfo, error) {
	userInfo, err := kc.client.GetUserInfo(kc.ctx, accessToken, kc.realm)
	if err != nil {
		log.Printf("Error fetching user info: %v", err)
		return nil, fmt.Errorf("failed to get user info: %w", err)
	}
	log.Printf("User info: %+v", userInfo)
	return userInfo, nil
}

// RefreshToken обновляет access и refresh токены
func (kc *KeycloakClient) RefreshToken(refreshToken string) (*gocloak.JWT, error) {
	newToken, err := kc.client.RefreshToken(kc.ctx, refreshToken, kc.clientID, kc.clientSecret, kc.realm)
	if err != nil {
		return nil, fmt.Errorf("failed to refresh token: %w", err)
	}
	kc.token = newToken
	return newToken, nil
}

// ValidateToken проверяет действительность токена
func (kc *KeycloakClient) ValidateToken(accessToken string) (bool, error) {
	result, err := kc.client.RetrospectToken(kc.ctx, accessToken, kc.clientID, kc.clientSecret, kc.realm)
	if err != nil {
		log.Printf("Error validating token: %v", err)
		return false, fmt.Errorf("failed to validate token: %w", err)
	}
	log.Printf("Token valid: %v", result.Active)
	return *result.Active, nil
}

// RegisterUser создает нового пользователя в Keycloak и возвращает его ID
func (kc *KeycloakClient) RegisterUser(username, password, email, firstName, lastName string) (string, error) {
	// Получаем административный токен
	token, err := kc.client.LoginClient(context.Background(), kc.clientID, kc.clientSecret, kc.realm)
	if err != nil {
		return "", fmt.Errorf("failed to obtain access token: %w", err)
	}

	user := gocloak.User{
		Username:  gocloak.StringP(username),
		Email:     gocloak.StringP(email),
		FirstName: gocloak.StringP(firstName),
		LastName:  gocloak.StringP(lastName),
		Enabled:   gocloak.BoolP(true),
	}

	// Создаем пользователя в Keycloak
	userID, err := kc.client.CreateUser(kc.ctx, token.AccessToken, kc.realm, user)
	if err != nil {
		return "", fmt.Errorf("failed to register user in Keycloak: %w", err)
	}
	log.Printf("User registered successfully with ID: %s", userID)

	// Устанавливаем пароль для пользователя
	err = kc.client.SetPassword(kc.ctx, token.AccessToken, userID, kc.realm, password, false)
	if err != nil {
		return "", fmt.Errorf("failed to set user password: %w", err)
	}
	log.Println("Password set successfully")

	return userID, nil
}
