build:
	@go build -o bin/bats

run: build
	@./bin/bats

test:
	@go test -v ./...