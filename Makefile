default:
	sh -c "'scripts/build.sh'"

fmt:
	gofmt -w .

.PHONY: default
