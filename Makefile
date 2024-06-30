generate-proto:
	protoc --go_out=./src/pkg/grpc --go_opt=paths=source_relative \
		--go-grpc_out=./src/pkg/grpc --go-grpc_opt=paths=source_relative \
		./src/api/hello.proto

startup-server:
	go run ./src/cmd/server/main.go

try-grpcurl:
	 grpcurl -plaintext -d '{"name": "kazu"}' localhost:8080 hello.GreetingService.Hello