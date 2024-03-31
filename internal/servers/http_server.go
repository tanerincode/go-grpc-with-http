package servers

import (
	"log"
	"net/http"

	"github.com/tanerincode/go-grpc-server/internal/handlers"
	"github.com/tanerincode/go-grpc-server/internal/services/health_check"
)

// StartServer starts the HTTP server.
func StartServer(client health_check.HealthCheckServiceClient) {
	// Register the health check handler.
	http.HandleFunc("/health", handlers.HealthCheckHandler(client))

	log.Println("HTTP server started on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
