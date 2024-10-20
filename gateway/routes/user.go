package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/olzzhas/edunite-server/gateway/clients"
	"github.com/olzzhas/edunite-server/gateway/handlers"
)

func SetupUserRoutes(r *gin.Engine, keycloakClient *clients.KeycloakClient, userHandler *handlers.UserHandler) {
	authGroup := r.Group("/user")
	authGroup.Use(AuthMiddleware(keycloakClient))
	{
		authGroup.GET("/users", userHandler.GetAllUsers)
	}

}
