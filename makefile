run:
	go run ./cmd/api/main.go -port 8080

test:
	go test ./... -v -cover -race

swag:
	swag init -g ./cmd/api/main.go --parseDependency --parseInternal

.PHONY: run test