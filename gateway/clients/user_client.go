package clients

import (
	"context"
	"fmt"
	"log"

	"github.com/olzzhas/edunite-server/user_service/pb"
	"google.golang.org/grpc"
)

// UserClient представляет клиент для взаимодействия с User Service через gRPC
type UserClient struct {
	client pb.UserServiceClient
}

// NewUserClient создает новый экземпляр UserClient с подключением к gRPC
func NewUserClient(conn *grpc.ClientConn) *UserClient {
	return &UserClient{
		client: pb.NewUserServiceClient(conn),
	}
}

// CreateUser сохраняет пользователя в базе данных через User Service
func (uc *UserClient) CreateUser(keycloakID, name, surname, email string) error {
	_, err := uc.client.CreateUser(context.Background(), &pb.CreateUserRequest{
		KeycloakID: keycloakID,
		Name:       name,
		Surname:    surname,
		Email:      email,
	})
	if err != nil {
		log.Printf("Failed to save user in User Service: %v", err)
		return fmt.Errorf("failed to create user: %w", err)
	}
	log.Printf("User saved in User Service with Keycloak ID: %s", keycloakID)
	return nil
}

func (uc *UserClient) GetAllUsers() (*pb.UsersResponse, error) {
	users, err := uc.client.GetAllUsers(context.Background(), &pb.EmptyRequest{})
	if err != nil {
		return nil, err
	}

	return users, nil
}
