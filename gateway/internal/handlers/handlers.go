package handlers

import (
	"github.com/gin-gonic/gin"
	clients2 "github.com/olzzhas/edunite-server/gateway/clients"
	"github.com/olzzhas/edunite-server/gateway/internal/models"
	"net/http"
)

type Handler struct {
	KeycloakClient *clients2.KeycloakClient
	UserClient     *clients2.UserClient
}

func (h *Handler) Register(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Создаем пользователя в Keycloak
	userID, err := h.KeycloakClient.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user in Keycloak"})
		return
	}

	// Сохраняем пользователя в User Service
	if err := h.UserClient.SaveUser(userID, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user in user service"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (h *Handler) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	token, err := h.KeycloakClient.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
