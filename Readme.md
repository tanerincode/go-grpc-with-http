# gRPC Health Check Example

This repository demonstrates how to set up a simple gRPC server with a health check service and an HTTP server that triggers the gRPC health check when accessed.

## Overview

The project is structured as follows:

- `cmd/client/main.go`: The main entry point for the client application.
- `internal/services/health_check`: Contains the gRPC service definition and implementation for the health check.
- `internal/servers/httpSrv/httpSrv.go`: Contains the HTTP server setup.
- `internal/handlers/healthCheckHandler.go`: Contains the HTTP handler for triggering the gRPC health check.

## Getting Started

### Prerequisites

- Go (version 1.18 or higher)
- Protocol Buffers Compiler (`protoc`)
- gRPC Go plugins

### Running the Server

```bash
make build-and-serve-server platform=linux
```
OR
```bash
make build-and-serve-client platform=darwin
```