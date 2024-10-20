package main

import (
	"github.com/gin-gonic/gin"
	"github.com/olzzhas/edunite-server/gateway/clients"
	"github.com/olzzhas/edunite-server/gateway/config"
	"github.com/olzzhas/edunite-server/gateway/handlers"
	"github.com/olzzhas/edunite-server/gateway/routes"
	"google.golang.org/grpc"
	"log"
)

func main() {
	r := gin.Default()
	cfg := config.LoadConfig()

	// Создаем клиента Keycloak
	keycloakClient := clients.NewKeycloakClient(
		cfg.Services.Keycloak.BaseURL,      // URL Keycloak
		cfg.Services.Keycloak.Realm,        // Realm
		cfg.Services.Keycloak.ClientID,     // Client ID
		cfg.Services.Keycloak.ClientSecret, // Client Secret
	)

	// Устанавливаем соединение с gRPC User Service
	conn, err := grpc.Dial(cfg.Services.UserService.Target, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to User Service: %v", err)
	}
	defer conn.Close()
	// Создаем клиента User Service
	userServiceClient := clients.NewUserClient(conn)

	conn, err = grpc.Dial(cfg.Services.LoggerService.Target, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to Logger Service: %v", err)
	}

	loggerServiceClient := clients.NewLoggerClient(conn)

	// Инициализируем обработчик аутентификации
	authHandler := &handlers.AuthHandler{
		KeycloakClient: keycloakClient,
		UserService:    userServiceClient,
		LoggerService:  loggerServiceClient,
	}

	userHandler := &handlers.UserHandler{
		UserService:   userServiceClient,
		LoggerService: loggerServiceClient,
	}

	// Настраиваем маршруты
	routes.SetupAuthRoutes(r, authHandler)
	routes.SetupUserRoutes(r, keycloakClient, userHandler)

	// Запускаем сервер
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
