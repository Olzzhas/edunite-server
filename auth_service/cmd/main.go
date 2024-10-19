package main

import (
	"log"
	"net"

	"github.com/olzzhas/edunite-server/auth_service/internal/auth"
	"github.com/olzzhas/edunite-server/auth_service/pb"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen on port 50052: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, auth.NewAuthHandler())

	log.Println("Auth service is running on port 50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
