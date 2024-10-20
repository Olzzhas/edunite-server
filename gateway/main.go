package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"github.com/olzzhas/edunite-server/gateway/clients"
	"github.com/olzzhas/edunite-server/gateway/handlers"
	"github.com/olzzhas/edunite-server/gateway/routes"
)

func main() {
	r := gin.Default()

	// Создаем клиента Keycloak
	keycloakClient := clients.NewKeycloakClient(
		"http://keycloak:8080",             // URL Keycloak
		"edunite",                          // Realm
		"auth",                             // Client ID
		"GRV1KxE8BYV45WVxP7s4d4WCm7cKZVOm", // Client Secret
	)

	// Устанавливаем соединение с gRPC User Service
	conn, err := grpc.Dial("user_service:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to User Service: %v", err)
	}
	defer conn.Close()

	// Создаем клиента User Service
	userServiceClient := clients.NewUserClient(conn)

	// Инициализируем обработчик аутентификации
	authHandler := &handlers.AuthHandler{
		KeycloakClient: keycloakClient,
		UserService:    userServiceClient,
	}

	// Настраиваем маршруты
	routes.SetupAuthRoutes(r, authHandler)

	// Запускаем сервер
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
