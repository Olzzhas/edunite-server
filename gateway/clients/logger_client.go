package clients

import (
	"context"
	"fmt"
	"github.com/olzzhas/edunite-server/logger_service/pb"
	"google.golang.org/grpc"
)

// LoggerClient представляет клиент для взаимодействия с Logger Service через gRPC
type LoggerClient struct {
	client pb.LoggerServiceClient
}

// NewLoggerClient создает новый экземпляр LoggerClient с подключением к gRPC
func NewLoggerClient(conn *grpc.ClientConn) *LoggerClient {
	return &LoggerClient{
		client: pb.NewLoggerServiceClient(conn),
	}
}

func (uc *LoggerClient) WriteLog(level, message, serviceName string, data map[string]string) error {
	response, err := uc.client.LogEvent(context.Background(), &pb.LogEventRequest{
		Level:       level,
		Message:     message,
		ServiceName: serviceName,
		Data:        data,
	})
	if err != nil {
		fmt.Println(fmt.Sprintf("error is occured while writing log. response form logger service:%v", response))
		return err
	}

	return nil
}
