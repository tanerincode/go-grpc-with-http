package main

import (
	"log"
	"net"

	"github.com/tanerincode/go-grpc-server/internal/services/health_check"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	log.Println("Server started on port 9000")

	// Create a new gRPC server instance.
	grpcServer := grpc.NewServer()

	// Create an instance of the HealthCheckServiceServerImpl.
	healthServer := health_check.NewHealthCheckServiceServerImpl()

	// Register it with the gRPC server.
	health_check.RegisterHealthCheckServiceServer(grpcServer, healthServer)

	// Start serving incoming connections.
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
