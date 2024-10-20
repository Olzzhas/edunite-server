package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olzzhas/edunite-server/gateway/clients"
	"net/http"
)

type UserHandler struct {
	UserService *clients.UserClient
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.UserService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error while getting all users: %s", err)})
		return
	}

	c.JSON(http.StatusOK, users)
}
