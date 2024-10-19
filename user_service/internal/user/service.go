package user

import (
	"context"
	"github.com/olzzhas/edunite-server/user_service/internal/database"
	"github.com/olzzhas/edunite-server/user_service/pb"
)

type Service struct {
	repo database.UserRepository
	pb.UnimplementedUserServiceServer
}

func NewUserService(repo database.UserRepository) *Service {
	return &Service{repo: repo}
}

// CreateUser Создать пользователя
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
		Id:      int64(id),
		Name:    user.Name,
		Surname: user.Surname,
		Role:    user.Role,
	}, nil
}

// GetUser Получить пользователя по ID
func (s *Service) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	user, err := s.repo.GetUserByID(ctx, int(req.GetId()))
	if err != nil {
		return nil, err
	}
	return &pb.UserResponse{
		Id:      int64(user.ID),
		Name:    user.Name,
		Surname: user.Surname,
		Role:    user.Role,
	}, nil
}

// GetAllUsers Получить всех пользователей
func (s *Service) GetAllUsers(ctx context.Context, req *pb.EmptyRequest) (*pb.UsersResponse, error) {
	users, err := s.repo.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	var pbUsers []*pb.UserResponse
	for _, user := range users {
		pbUsers = append(pbUsers, &pb.UserResponse{
			Id:      int64(user.ID),
			Name:    user.Name,
			Surname: user.Surname,
			Role:    user.Role,
		})
	}
	return &pb.UsersResponse{Users: pbUsers}, nil
}

// DeleteUser Удалить пользователя
func (s *Service) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.EmptyResponse, error) {
	err := s.repo.DeleteUser(ctx, int(req.GetId()))
	if err != nil {
		return nil, err
	}
	return &pb.EmptyResponse{}, nil
}
