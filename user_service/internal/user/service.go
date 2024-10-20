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
		ID:         req.GetId(),
		KeycloakID: req.GetKeycloakID(),
		Name:       req.GetName(),
		Surname:    req.GetSurname(),
		Email:      req.GetEmail(),
		Role:       req.GetRole(),
	}
	err := s.repo.CreateUser(ctx, &user)
	if err != nil {
		return nil, err
	}
	return &pb.UserResponse{
		Id:         user.ID,
		KeycloakID: user.KeycloakID,
		Name:       user.Name,
		Surname:    user.Surname,
		Role:       user.Role,
		Email:      user.Email,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
		Version:    int32(user.Version),
	}, nil
}

// GetUser Получить пользователя по ID
func (s *Service) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	user, err := s.repo.GetUserByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.UserResponse{
		Id:         user.ID,
		KeycloakID: user.KeycloakID,
		Name:       user.Name,
		Surname:    user.Surname,
		Role:       user.Role,
		Email:      user.Email,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
		Version:    int32(user.Version),
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
			Id:         user.ID,
			KeycloakID: user.KeycloakID,
			Name:       user.Name,
			Surname:    user.Surname,
			Role:       user.Role,
			Email:      user.Email,
			CreatedAt:  user.CreatedAt,
			UpdatedAt:  user.UpdatedAt,
			Version:    int32(user.Version),
		})
	}
	return &pb.UsersResponse{Users: pbUsers}, nil
}

// DeleteUser Удалить пользователя
func (s *Service) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.EmptyResponse, error) {
	err := s.repo.DeleteUser(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.EmptyResponse{}, nil
}

func (s *Service) GetUserByEmail(ctx context.Context, req *pb.GetUserByEmailRequest) (*pb.UserResponse, error) {
	user, err := s.repo.GetUserByEmail(ctx, req.GetEmail())
	if err != nil {
		return nil, err
	}

	userResponse := &pb.UserResponse{
		Id:         user.ID,
		KeycloakID: user.KeycloakID,
		Name:       user.Name,
		Surname:    user.Surname,
		Role:       user.Role,
		Email:      user.Email,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
		Version:    int32(user.Version),
	}

	return userResponse, nil
}
