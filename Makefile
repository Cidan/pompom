.EXPORT_ALL_VARIABLES:

PUBSUB_EMULATOR_HOST=localhost:8085

build:
	go build ./...
	go install ./...

test:
	go vet ./...; \
	go test -v -test.short -covermode=atomic ./...

run:
	pompom

.PHONY: test