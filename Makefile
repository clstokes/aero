default:
	sh -c "'scripts/build.sh'"

release: fmt
	sh -c "'scripts/release.sh'"

fmt:
	gofmt -w .

.PHONY: default
