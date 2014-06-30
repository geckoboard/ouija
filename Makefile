all:
	@mkdir -p bin/
	go build -o bin/ouija

test:
	go test ./...
