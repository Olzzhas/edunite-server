package auth

import (
	"context"
	"github.com/olzzhas/edunite-server/auth_service/pb"
)

type Handler struct {
	pb.UnimplementedAuthServiceServer
}

func NewAuthHandler() *Handler {
	return &Handler{}
}

// Login Заглушка для метода Login
func (h *Handler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	return &pb.AuthResponse{
		Token:   "fake-token-123",
		Message: "This is a stub response",
	}, nil
}
