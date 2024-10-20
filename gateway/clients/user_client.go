package clients

import (
	"github.com/olzzhas/edunite-server/user_service/pb"
	"google.golang.org/grpc"
)

// UserClient представляет клиент для взаимодействия с User Service через gRPC
type UserClient struct {
	client pb.UserServiceClient
}

// NewUserClient создаёт новый экземпляр UserClient
func NewUserClient(conn *grpc.ClientConn) *UserClient {
	return &UserClient{
		client: pb.NewUserServiceClient(conn),
	}
}
