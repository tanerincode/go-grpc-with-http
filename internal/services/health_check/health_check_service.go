package health_check

import (
	"context"
	"log"
)

// Ensure that HealthCheckServiceServer implements all methods.
var _ HealthCheckServiceServer = (*HealthCheckServiceServerImpl)(nil)

// HealthCheckServiceServerImpl provides an implementation of the HealthCheckServiceServer interface.
type HealthCheckServiceServerImpl struct {
	UnimplementedHealthCheckServiceServer
}

// NewHealthCheckServiceServerImpl returns a new HealthCheckServiceServerImpl.
func NewHealthCheckServiceServerImpl() *HealthCheckServiceServerImpl {
	return &HealthCheckServiceServerImpl{}
}

// Check implements the HealthCheckServiceServer interface.
func (srv *HealthCheckServiceServerImpl) Check(ctx context.Context, in *HealthCheckRequest) (*HealthCheckResponse, error) {
	log.Println("Health check request received")
	return &HealthCheckResponse{
		Status:  true,
		Message: "OK",
	}, nil
}
