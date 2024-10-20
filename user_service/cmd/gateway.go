package main

import (
	"context"
	"flag"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/olzzhas/edunite-server/user_service/pb"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func main() {
	// Флаг для указания адреса gRPC сервера
	grpcServerEndpoint := flag.String("grpc-server-endpoint", "localhost:50051", "gRPC server endpoint")
	flag.Parse()

	// Создаём HTTP маршрутизатор для REST запросов
	mux := runtime.NewServeMux()

	// Подключаемся к gRPC серверу
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterUserServiceHandlerFromEndpoint(context.Background(), mux, *grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("Failed to register gateway: %v", err)
	}

	// Запускаем HTTP сервер на порту 8080
	log.Println("API Gateway is running on port :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
