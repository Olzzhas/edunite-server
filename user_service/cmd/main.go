package main

import (
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"user_service/internal/database"
	"user_service/internal/user"
	"user_service/user_service/pb"
)

func main() {
	// Загрузка .env файла
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Подключение к базе данных
	db := database.ConnectDB()
	defer db.Close()

	repo := database.NewUserRepository(db)
	userService := user.NewUserService(repo)
	// Создание gRPC сервера
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userService)

	reflection.Register(grpcServer)

	log.Println("Server is running on port :50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
