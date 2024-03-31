package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/tanerincode/go-grpc-server/internal/servers"
	"github.com/tanerincode/go-grpc-server/internal/services/health_check"
)

func main() {
	// Start the gRPC connection.
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Could not connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create a new HealthCheckService client from the connection.
	client := health_check.NewHealthCheckServiceClient(conn)

	// Start the HTTP server in a separate goroutine.
	go servers.StartServer(client)

	// Block forever.
	select {}
}
