package clients

import (
	"context"
	"github.com/olzzhas/edunite-server/logger_service/pb"
	"google.golang.org/grpc"
	"time"
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

func (uc *LoggerClient) WriteInfoLog(keycloakId, serviceName, message string) error {
	_, err := uc.client.LogEvent(context.Background(), &pb.LogEventRequest{
		KeycloakId:  keycloakId,
		ServiceName: serviceName,
		Message:     message,
		Timestamp:   time.Now().Format(time.RFC3339),
		Level:       "INFO",
	})
	if err != nil {
		return err
	}

	return nil
}
