package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/olzzhas/edunite-server/gateway/handlers"
)

func SetupAuthRoutes(r *gin.Engine, authHandler *handlers.AuthHandler) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
	}
}
