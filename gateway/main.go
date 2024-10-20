package main

import (
	"github.com/gin-gonic/gin"
	"github.com/olzzhas/edunite-server/gateway/clients"
	"github.com/olzzhas/edunite-server/gateway/handlers"
	"github.com/olzzhas/edunite-server/gateway/routes"
	"google.golang.org/grpc"
	"log"
)

func main() {
	r := gin.Default()

	// Keycloak Client
	keycloakClient := clients.NewKeycloakClient(
		"http://keycloak:8080",             // URL Keycloak сервера
		"edunite",                          // Имя вашего realm
		"auth",                             // ID клиента в Keycloak
		"Jd55KeL2zwSiZiEUlewrBKY8lwfV05NJ", // Секрет клиента
	)

	// Устанавливаем gRPC соединение с User Service
	conn, err := grpc.Dial("user_service:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Не удалось подключиться к User Service: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	// User Service Client
	userServiceClient := clients.NewUserClient(conn)

	// Инициализация обработчиков
	authHandler := &handlers.AuthHandler{
		KeycloakClient: keycloakClient,
		UserService:    userServiceClient,
	}

	// Настройка маршрутов
	routes.SetupAuthRoutes(r, authHandler)

	// Middleware для авторизации
	r.Use(routes.AuthMiddleware(keycloakClient))

	// Пример защищённого маршрута
	r.GET("/userinfo", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "User info"})
	})

	// Запуск сервера
	r.Run(":8081")
}
