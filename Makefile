build:
	@go build -o ./bin/calculator

run: build
	@./bin/calculator

test:
	@go test ./... -v
