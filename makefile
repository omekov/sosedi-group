run:
	go run main.go -port 8080

test:
	go test ./... -v -cover -race

.PHONY: run test