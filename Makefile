build:
	@go build -o bin/api-go

run: build
	@./bin/api-go

test:
	@go test -v ./...