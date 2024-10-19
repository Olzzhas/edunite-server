package user

import (
	"context"

	"github.com/olzzhas/edunite-server/user_service/pb"
)

type Handler struct {
	service *Service
	pb.UnimplementedUserServiceServer
}

// NewUserHandler Конструктор для UserHandler
func NewUserHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// CreateUser Обработчик для создания пользователя
func (h *Handler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	return h.service.CreateUser(ctx, req)
}

// GetUser Обработчик для получения пользователя по ID
func (h *Handler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	return h.service.GetUser(ctx, req)
}

// GetAllUsers Обработчик для получения всех пользователей
func (h *Handler) GetAllUsers(ctx context.Context, req *pb.EmptyRequest) (*pb.UsersResponse, error) {
	return h.service.GetAllUsers(ctx, &pb.EmptyRequest{})
}

// DeleteUser Обработчик для удаления пользователя
func (h *Handler) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.EmptyResponse, error) {
	return h.service.DeleteUser(ctx, req)
}
