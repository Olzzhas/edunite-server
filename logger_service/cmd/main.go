package main

import (
	"log"
	"net"

	"github.com/olzzhas/edunite-server/logger_service/pb"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterLoggerServiceServer(grpcServer, &LoggerServer{})

	log.Println("Logger Service is running on port 50052...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
