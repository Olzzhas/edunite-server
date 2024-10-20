package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type KeycloakConfig struct {
	BaseURL      string
	Realm        string
	ClientID     string
	ClientSecret string
}

type UserServiceConfig struct {
	Target string
}

type Config struct {
	Services struct {
		Keycloak    KeycloakConfig
		UserService UserServiceConfig
	}
}

// LoadConfig загружает переменные окружения из .env файла и возвращает конфигурацию
func LoadConfig() *Config {
	// Загружаем переменные из .env файла
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return &Config{
		Services: struct {
			Keycloak    KeycloakConfig
			UserService UserServiceConfig
		}{
			Keycloak: KeycloakConfig{
				BaseURL:      os.Getenv("KEYCLOAK_URL"),
				Realm:        os.Getenv("KEYCLOAK_REALM"),
				ClientID:     os.Getenv("KEYCLOAK_CLIENT_ID"),
				ClientSecret: os.Getenv("KEYCLOAK_CLIENT_SECRET"),
			},
			UserService: UserServiceConfig{
				Target: os.Getenv("USER_SERVICE_URL"),
			},
		},
	}
}
