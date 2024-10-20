package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olzzhas/edunite-server/gateway/clients"
	"log"
	"net/http"
	"strings"
)

type AuthHandler struct {
	keycloakClient *clients.KeycloakClient
}

func NewAuthHandler(keycloakClient *clients.KeycloakClient) *AuthHandler {
	return &AuthHandler{keycloakClient: keycloakClient}
}
func (h *AuthHandler) UserInfoHandler(c *gin.Context) {
	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing access token"})
		return
	}

	// Удалить префикс "Bearer "
	accessToken = strings.TrimPrefix(accessToken, "Bearer ")

	userInfo, err := h.keycloakClient.GetUserInfo(accessToken)
	if err != nil {
		log.Printf("Failed to get user info: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to get user info: %v", err)})
		return
	}

	c.JSON(http.StatusOK, userInfo)
}
