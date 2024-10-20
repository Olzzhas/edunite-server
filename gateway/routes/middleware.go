package routes

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/olzzhas/edunite-server/gateway/clients"
)

func AuthMiddleware(keycloakClient *clients.KeycloakClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		if _, err := keycloakClient.ValidateToken(token); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		c.Next()
	}
}
