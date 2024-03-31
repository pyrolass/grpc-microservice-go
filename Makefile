gate:
	@go build -o bin/gateway ./gateway
	@./bin/gateway

prod:
	@go build -o bin/producer-micro ./producer-micro
	@./bin/producer-micro

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/ptypes.proto

.PHONY: gate proto
