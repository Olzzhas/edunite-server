package main

import (
	"github.com/olzzhas/edunite-server/user_service/internal/config"
	"github.com/olzzhas/edunite-server/user_service/internal/database"
	"github.com/olzzhas/edunite-server/user_service/internal/user"
	"github.com/olzzhas/edunite-server/user_service/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	// Загружаем конфигурацию
	cfg := config.LoadConfig()

	// Подключаемся к базе данных
	db := database.ConnectDB(cfg)
	defer db.Close()

	// Запуск gRPC-сервера
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	userService := user.NewUserService(database.NewUserRepository(db))
	pb.RegisterUserServiceServer(grpcServer, userService)

	// Включение рефлексии для дебага
	reflection.Register(grpcServer)

	log.Println("User service is running on port 50051")

	// Блокируем основной поток, пока сервер не завершится
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
