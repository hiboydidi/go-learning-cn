.PHONY: fmt test vet check examples

fmt:
	gofmt -w .

test:
	go test ./...

vet:
	go vet ./...

check: fmt test vet

examples:
	go run ./examples/hello
	go run ./examples/control-flow
	go run ./examples/collections
	go run ./examples/functions
	go run ./examples/structs
	go run ./examples/errors
	go run ./examples/files
	go run ./examples/json
	go run ./examples/concurrency
	go run ./examples/worker-pool
	go run ./examples/logging
