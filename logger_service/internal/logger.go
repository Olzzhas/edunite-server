package internal

import (
	"context"
	"log"

	"github.com/olzzhas/edunite-server/logger_service/pb"
)

// Реализация сервера логгера
type LoggerServer struct {
	pb.UnimplementedLoggerServiceServer
}

// Метод для логирования событий
func (s *LoggerServer) LogEvent(ctx context.Context, req *pb.LogEventRequest) (*pb.LogEventResponse, error) {
	log.Printf("[%s] %s: %s", req.Level, req.ServiceName, req.Message)

	// Здесь можно добавить логику для сохранения в MongoDB
	return &pb.LogEventResponse{
		Success: true,
		Message: "Log saved successfully",
	}, nil
}
