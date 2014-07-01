all: build test run

build:
	@mkdir -p bin/
	go build -o bin/ouija

test:
	go test ./...

run:
	@echo "$(LATENCIES_CSV)"
	./bin/ouija "$(LATENCIES_CSV)"

.PHONY: build test run
