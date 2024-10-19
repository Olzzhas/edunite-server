package config

import (
	"encoding/json"
	"log"
	"os"
)

// Config структура для хранения конфигурации
type Config struct {
	Database struct {
		URL string `json:"url"`
	} `json:"database"`
	Server struct {
		Port int `json:"port"`
	} `json:"server"`
}

// LoadConfig загружает конфигурацию из файла config.json
func LoadConfig() (*Config, error) {
	file, err := os.Open("config/config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &Config{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	log.Println("Config loaded successfully.")
	return config, nil
}
