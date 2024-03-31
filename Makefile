build-protos:
	protoc --proto_path=protos --go_out=./ --go-grpc_out=./ pkg/protos/*.proto

build-server:
	GOOS=$(platform) GOARCH=amd64 go build -o bin/server-$(platform) cmd/server/main.go

build-client:
	GOOS=$(platform) GOARCH=amd64 go build -o bin/client-$(platform) cmd/client/main.go

serve-server:
	./bin/server-$(platform)

serve-client:
	./bin/client-$(platform)

build-and-serve-server: build-server serve-server

build-and-serve-client: build-client serve-client

run-server:
	go run cmd/server/main.go

run-client:
	go run cmd/client/main.go
