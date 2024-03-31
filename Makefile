gate:
	@go build -o bin/gateway ./gateway
	@./bin/gateway

.PHONY: gate
