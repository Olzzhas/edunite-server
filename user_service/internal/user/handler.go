package user

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"net/http"

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

func ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Здесь используем публичный ключ Keycloak для проверки токена
		return []byte("your-public-key"), nil
	})
	return token, err
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		token, err := ValidateToken(tokenString)
		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
