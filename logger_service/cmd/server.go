package main

import (
	"context"
	"github.com/olzzhas/edunite-server/logger_service/pb"

	"log"
)

type LoggerServer struct {
	pb.UnimplementedLoggerServiceServer
}

func (s *LoggerServer) LogEvent(ctx context.Context, req *pb.LogEventRequest) (*pb.LogEventResponse, error) {
	log.Printf("[%s] %s: %s", req.Level, req.ServiceName, req.Message)

	// Сюда можно добавить сохранение в MongoDB или другой стор
	return &pb.LogEventResponse{
		Success: true,
		Message: "Log saved successfully",
	}, nil
}
