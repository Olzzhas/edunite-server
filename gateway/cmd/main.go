package main

import (
	"github.com/gin-gonic/gin"
	"github.com/olzzhas/edunite-server/gateway/internal/handlers"
	"log"
)

func main() {
	r := gin.Default()

	// Роуты для авторизации и регистрации
	r.POST("/login", handlers.LoginHandler)
	r.POST("/register", handlers.RegisterHandler)

	log.Println("API Gateway is running on port 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
