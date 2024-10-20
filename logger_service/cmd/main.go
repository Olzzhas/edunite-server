package main

import (
	"context"
	"github.com/streadway/amqp"
	"log"
	"net"
	"time"

	"github.com/olzzhas/edunite-server/logger_service/internal"
	"github.com/olzzhas/edunite-server/logger_service/pb"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

type LoggerServer struct {
	pb.UnimplementedLoggerServiceServer
	logger     *internal.Logger
	mongoDB    *mongo.Database
	rabbitConn *amqp.Connection
}

func NewLoggerServer(logger *internal.Logger, db *mongo.Database, rabbitConn *amqp.Connection) *LoggerServer {
	return &LoggerServer{
		logger:     logger,
		mongoDB:    db,
		rabbitConn: rabbitConn,
	}
}

func (s *LoggerServer) LogEvent(ctx context.Context, req *pb.LogEventRequest) (*pb.LogEventResponse, error) {
	logData := map[string]interface{}{
		"level":    req.Level,
		"message":  req.Message,
		"service":  req.ServiceName,
		"data":     req.Data,
		"datetime": time.Now().UTC(),
	}

	if err := internal.SaveLogToMongo(s.mongoDB, logData); err != nil {
		return nil, err
	}

	if err := internal.PublishToRabbitMQ(s.rabbitConn, logData); err != nil {
		return nil, err
	}

	s.logger.LogEvent(req.Level, req.Message, req.ServiceName, req.Data)

	return &pb.LogEventResponse{
		Success: true,
		Message: "Log saved successfully",
	}, nil
}

func main() {
	// Подключение к MongoDB
	db, err := internal.ConnectMongoDB("mongodb://mongo:27017")
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer db.Client().Disconnect(context.Background())

	// Подключение к RabbitMQ
	rabbitConn, err := internal.ConnectRabbitMQ("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer rabbitConn.Close()

	// Инициализация логгера
	logger := internal.NewLogger()

	// Запуск gRPC сервера
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterLoggerServiceServer(grpcServer, NewLoggerServer(logger, db, rabbitConn))

	log.Println("Logger Service is running on port 50052...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
