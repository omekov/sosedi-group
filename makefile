run:
	go run ./cmd/api/main.go -port 8080

test:
	go test ./... -v -cover -race

swag:
	swag init -d internal/api -g api.go --parseDependency --parseInternal

.PHONY: run test swag