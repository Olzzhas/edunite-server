package clients

import (
	"context"
	"github.com/olzzhas/edunite-server/user_service/pb"

	"google.golang.org/grpc"
	"time"
)

type UserClient struct {
	client pb.UserServiceClient
}

func NewUserClient(conn *grpc.ClientConn) *UserClient {
	return &UserClient{
		client: pb.NewUserServiceClient(conn),
	}
}

func (uc *UserClient) GetUserByEmail(email string) (*pb.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := uc.client.GetUserByEmail(ctx, &pb.GetUserByEmailRequest{Email: email})
	if err != nil {
		return nil, err
	}
	return resp.User, nil
}
