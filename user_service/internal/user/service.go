package user

import (
	"context"
	"github.com/olzzhas/edunite-server/user_service/internal/database"
	"github.com/olzzhas/edunite-server/user_service/pb"
	"time"
)

type Service struct {
	repo database.UserRepository
	pb.UnimplementedUserServiceServer
}

func NewUserService(repo database.UserRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	user := database.User{
		Name:    req.GetName(),
		Surname: req.GetSurname(),
		Role:    req.GetRole(),
	}

	id, err := s.repo.CreateUser(ctx, &user)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		Id:        int64(id),
		Name:      user.Name,
		Surname:   user.Surname,
		Role:      user.Role,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
		Version:   int32(user.Version),
	}, nil
}

func (s *Service) GetUser(ctx context.Context, request *pb.GetUserRequest) (*pb.UserResponse, error) {
	//TODO implement me
	return nil, nil
}

func (s *Service) UpdateUser(ctx context.Context, request *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	//TODO implement me
	return nil, nil
}

func (s *Service) DeleteUser(ctx context.Context, request *pb.DeleteUserRequest) (*pb.EmptyResponse, error) {
	//TODO implement me
	return nil, nil
}
