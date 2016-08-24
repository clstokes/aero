default: test

deps:
	go get ./...

dev: deps
	sh -c "'scripts/build.sh'"

fmt:
	gofmt -w .

test:
	go test $(shell go list ./...)

.PHONY: default
