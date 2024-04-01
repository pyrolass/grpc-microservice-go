gate:
	@go build -o bin/gateway ./gateway
	@./bin/gateway

prod:
	@go build -o bin/producer-micro ./producer-micro
	@./bin/producer-micro

con:
	@go build -o bin/consumer-micro ./consumer-micro
	@./bin/consumer-micro

write:
	@go build -o bin/write-db-micro ./write-db-micro
	@./bin/write-db-micro

dist:
	@go build -o bin/distance-micro ./distance-micro
	@./bin/distance-micro


spam:
	@go build -o bin/spam ./spam
	@./bin/spam

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/ptypes.proto

.PHONY: gate proto spam
