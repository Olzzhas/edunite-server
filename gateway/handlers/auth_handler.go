package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olzzhas/edunite-server/gateway/clients"
	"github.com/olzzhas/edunite-server/gateway/models"
	"net/http"
)

type AuthHandler struct {
	KeycloakClient *clients.KeycloakClient
	UserService    *clients.UserClient
	LoggerService  *clients.LoggerClient
}

// RegisterHandler обрабатывает регистрацию пользователя
func (h *AuthHandler) RegisterHandler(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Создание пользователя в Keycloak
	userID, err := h.KeycloakClient.RegisterUser(req.Username, req.Password, req.Email, req.Name, req.Surname)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error while registering user: %s", err)})
		return
	}

	logData := map[string]string{
		"user_keycloak_id": userID,
		"user_name":        req.Name,
		"user_surname":     req.Surname,
		"user_email":       req.Email,
	}

	_ = h.LoggerService.WriteLog("INFO", "user created in keycloak successfully", "keycloak", logData)

	// Сохранение данных пользователя в User Service
	if err := h.UserService.CreateUser(userID, req.Name, req.Surname, req.Email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user in database"})
		return
	}

	_ = h.LoggerService.WriteLog("INFO", "user created in user database successfully", "user", logData)

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "user_id": userID})
}

// LoginHandler обрабатывает авторизацию пользователя
func (h *AuthHandler) LoginHandler(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Логин через Keycloak
	token, err := h.KeycloakClient.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": token.AccessToken})
}
