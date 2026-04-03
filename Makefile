.PHONY: all build test test-race bench lint clean

MODULE := github.com/KyumKyum/lionet/v1

all: lint test build

build:
	go build ./...

test:
	go test ./... -v -count=1

test-race:
	go test ./... -v -race -count=1

bench:
	go test -bench=. -benchmem ./...

lint:
	golangci-lint run ./...

test-integration:
	go test ./test/ -v -run TestMillionKeyWriteRead -count=1 -timeout=10m

test-crash:
	go test ./test/ -v -run TestSIGKILLRecovery -count=1 -timeout=5m

test-stress:
	go test ./test/ -v -race -run TestConcurrentStress -count=1 -timeout=10m

clean:
	rm -rf /tmp/lionet-test-*
