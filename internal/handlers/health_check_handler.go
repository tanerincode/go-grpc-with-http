package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/tanerincode/go-grpc-server/internal/services/health_check"
)

// HealthCheckHandler returns an HTTP handler for health checks.
func HealthCheckHandler(client health_check.HealthCheckServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Trigger gRPC health check.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := client.Check(ctx, &health_check.HealthCheckRequest{})
		if err != nil {
			http.Error(w, `{"status":"error","message":"Service unhealthy"}`, http.StatusInternalServerError)
			return
		}

		// Check the response and write back to HTTP as JSON.
		w.Header().Set("Content-Type", "application/json")
		if response.GetStatus() {
			json.NewEncoder(w).Encode(map[string]string{"status": "ok", "message": "Service is healthy"})
		} else {
			http.Error(w, `{"status":"error","message":"Service unhealthy"}`, http.StatusInternalServerError)
		}
	}
}
