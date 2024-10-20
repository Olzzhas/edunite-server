package config

import "os"

type Config struct {
	DatabaseURL  string
	KeycloakURL  string
	ClientID     string
	ClientSecret string
}

func LoadConfig() *Config {
	return &Config{
		DatabaseURL:  os.Getenv("DATABASE_URL"),
		KeycloakURL:  os.Getenv("KEYCLOAK_URL"),
		ClientID:     "user-service",
		ClientSecret: "your-client-secret",
	}
}
