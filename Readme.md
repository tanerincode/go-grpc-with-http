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

ps : _Make sure port 9000 and 8080 are available_

### Running the Server

First Generate the gRPC code (if not already generated):
```bash
make build-protos
```


After This you can run project below commands

```bash
make build-and-serve-server platform=linux
```
OR
```bash
make build-and-serve-client platform=darwin
```
You can also run the server directly without build :
```bash
make run-server
```
for client: 
ps : _make sure the 9000_
```bash
make run-client
```

## Accessing the Health Check
To trigger the health check, access the following URL in your browser or using a tool like curl:
```bash
curl http://localhost:8080/health
```

If the gRPC server is running and healthy, you should receive a response:
```json
{"status":"ok","message":"Service is healthy"}
```