package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/olzzhas/edunite-server/gateway/internal/clients"
	"net/http"
)

var userClient *clients.UserClient // Предположим, что клиент уже инициализирован

func RegisterHandler(c *gin.Context) {
	var user struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Логика регистрации может отправлять запрос в User Service
	// Здесь можно добавить проверку на существование пользователя
	_, err := userClient.GetUserByEmail(user.Email)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	// Добавьте вызов регистрации через gRPC, если необходимо
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}
