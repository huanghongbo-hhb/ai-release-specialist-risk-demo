.PHONY: build test junit

build:
	go build ./...

test:
	go test ./...

junit:
	mkdir -p reports
	go test -json ./... > reports/go-test.json || true
