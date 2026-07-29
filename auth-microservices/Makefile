run:
	go run ./cmd/main.go

build:
	go build -o ./build/auth-microservices ./cmd/main.go

protoc:	
	protoc --go_out=./proto --go-grpc_out=./proto --proto_path=./proto ./proto/*.proto