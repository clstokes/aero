default:
	go build -o bin/aero

fmt:
	gofmt -w .

.PHONY: default
