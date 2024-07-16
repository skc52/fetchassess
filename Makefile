build:
	@go build -o bin/fetchAssess cmd/main.go

test:
	@go test -v ./...

run: build 
	@./bin/fetchAssess