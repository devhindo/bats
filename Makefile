build:
	@go build -o bin/bats

run: build
	@./bin/bats

test:
	@go test -v ./...

up:
	@docker-compose up -d

adbs: # access db shell

sdbl: # show db logs
	@docker-compose logs -f db