build:
	@go build -o bin/Go-Rest cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/Go-Rest