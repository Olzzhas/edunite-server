package handlers

import (
	"encoding/json"
	"github.com/olzzhas/edunite-server/gateway/clients"
	"net/http"
)

type UserHandler struct {
	userClient *clients.UserClient
}

func NewUserHandler(userClient *clients.UserClient) *UserHandler {
	return &UserHandler{userClient: userClient}
}

func (h *UserHandler) GetUserByEmailHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	user, err := h.userClient.GetUserByEmail(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}
