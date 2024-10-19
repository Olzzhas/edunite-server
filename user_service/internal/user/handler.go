package user

import (
	"context"
	"github.com/olzzhas/edunite-server/user_service/internal/database"
	"github.com/olzzhas/edunite-server/user_service/pb"

	"time"
)

type ServiceServer struct {
	pb.UnimplementedUserServiceServer
}

func (s *ServiceServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	user := database.User{
		ID:        4,
		Name:      req.GetName(),
		Surname:   req.GetSurname(),
		Role:      req.GetRole(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Version:   1,
	}

	return &pb.UserResponse{
		Id:        int64(user.ID),
		Name:      user.Name,
		Surname:   user.Surname,
		Role:      user.Role,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *ServiceServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	// TODO
	return nil, nil
}

func (s *ServiceServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	// TODO
	return nil, nil
}

func (s *ServiceServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.EmptyResponse, error) {
	// TODO
	return nil, nil
}
