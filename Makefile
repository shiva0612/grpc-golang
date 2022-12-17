.PHONY: domain server

domain:
	protoc --proto_path=proto --go_out=. --go-grpc_out=. --go_opt=module=shiva --go-grpc_opt=module=shiva  proto/*.proto

server: domain
	cd server && go run main.go

unary:
	cd cobra-grpc-client && go run main.go unary
cstream:
	cd cobra-grpc-client && go run main.go clientStream
sstream:
	cd cobra-grpc-client && go run main.go serverStream
bistream:
	cd cobra-grpc-client && go run main.go biStream
