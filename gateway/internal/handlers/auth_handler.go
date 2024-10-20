package handlers

import (
	"github.com/olzzhas/edunite-server/gateway/internal/clients"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Предположим, что у нас есть глобальная переменная для клиентов
var keycloakClient = clients.NewKeycloakClient("http://localhost:8080")

func LoginHandler(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	token, err := keycloakClient.Login(credentials.Email, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": token})
}
